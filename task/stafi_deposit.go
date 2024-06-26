package task

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/credential"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/utils"
)

func (task *Task) getUsablePoolBalance() (*big.Int, error) {
	poolBalance, err := task.userDepositContract.GetBalance(nil)
	if err != nil {
		return nil, err
	}
	if poolBalance.Cmp(task.poolReservedBalance) > 0 {
		return new(big.Int).Sub(poolBalance, task.poolReservedBalance), nil
	} else {
		return big.NewInt(0), nil
	}
}

func (task *Task) checkAndDeposit() (retErr error) {
	if !task.offchainStateIsLatest() {
		return nil
	}

	poolBalance, err := task.getUsablePoolBalance()
	if err != nil {
		return err
	}
	poolBalanceDeci := decimal.NewFromBigInt(poolBalance, 0)
	logrus.WithFields(logrus.Fields{
		"balance": poolBalanceDeci.String(),
	}).Debug("deposit-poolBalance")

	if poolBalanceDeci.LessThan(minAmountNeedDeposit) {
		return nil
	}

	// wait if some validators wait stake
	for i := 0; i < len(task.validatorsByKeyIndex); i++ {
		val := task.validatorsByKeyIndex[i]
		if val.statusOnStafi == valStatusDeposited || val.statusOnStafi == valStatusMatch {
			return nil
		}
	}

	// validators need deposit
	depositLen := poolBalanceDeci.Div(minAmountNeedDeposit).IntPart()
	maxAvailableDepositLen := uint64(0)
	if task.ValidatorsPerTrustNodeLimit > uint64(len(task.validatorsByKeyIndex)) {
		maxAvailableDepositLen = task.ValidatorsPerTrustNodeLimit - uint64(len(task.validatorsByKeyIndex))
	}
	if depositLen > int64(maxAvailableDepositLen) {
		depositLen = int64(maxAvailableDepositLen)
	}
	if depositLen > int64(task.ValidatorsLimitByGas) {
		depositLen = int64(task.ValidatorsLimitByGas)
	}

	if depositLen == 0 {
		return nil
	}

	validatorPubkeys := make([][]byte, depositLen)
	sigs := make([][]byte, depositLen)
	dataRoots := make([][32]byte, depositLen)
	oldKeyIndex := task.nextKeyIndex

	logrus.WithFields(logrus.Fields{
		"depositLen":   depositLen,
		"nextKeyIndex": task.nextKeyIndex,
	}).Debug("deposit-info")

	defer func() {
		if retErr != nil {
			// got err, should reset validator info
			task.nextKeyIndex = oldKeyIndex
			for i := task.nextKeyIndex; i < task.nextKeyIndex+int(depositLen); i++ {
				pubkey := hex.EncodeToString(task.validatorsByKeyIndex[i].privateKey.PublicKey().Marshal())

				delete(task.validatorsByKeyIndex, i)
				delete(task.validatorsByPubkey, pubkey)
			}
		}
	}()

	for i := 0; i < int(depositLen); i++ {
		credential, err := credential.NewCredential(task.seed, task.nextKeyIndex,
			trustNodeDepositAmount.Div(utils.GweiDeci).BigInt(), task.chain, task.eth1WithdrawalAddress)
		if err != nil {
			return err
		}

		pubkeyBts := credential.SigningSk.PublicKey().Marshal()
		pubkeyStatus, err := task.mustGetTrustNodePubkeyStatus(pubkeyBts)
		if err != nil {
			return fmt.Errorf("mustGetTrustNodePubkeyStatus err: %s", err.Error())
		}
		if pubkeyStatus != utils.ValidatorStatusUnInitial {
			return fmt.Errorf("pubkey %s at index %d already on chain", hex.EncodeToString(pubkeyBts), task.nextKeyIndex)
		}

		validatorPubkeys[i] = pubkeyBts
		depositData, err := credential.SigningDepositData()
		if err != nil {
			return err
		}
		sigBts, err := hex.DecodeString(depositData.Signature)
		if err != nil {
			return err
		}
		sigs[i] = sigBts
		dataRootBts, err := hex.DecodeString(depositData.DepositDataRoot)
		if err != nil {
			return err
		}
		if len(dataRootBts) != 32 {
			return fmt.Errorf("dataRoot length %d not match", len(dataRootBts))
		}
		dataRoots[i] = [32]byte(dataRootBts)

		val := &Validator{
			privateKey:    credential.SigningSk,
			statusOnStafi: valStatusUnInitiated,
			keyIndex:      task.nextKeyIndex,
		}
		task.validatorsByKeyIndex[task.nextKeyIndex] = val
		task.validatorsByPubkey[hex.EncodeToString(pubkeyBts)] = val

		logrus.WithFields(logrus.Fields{
			"keyIndex":            task.nextKeyIndex,
			"pubkey":              hex.EncodeToString(pubkeyBts),
			"dataRoot":            depositData.DepositDataRoot,
			"messageRoot":         depositData.DepositMessageRoot,
			"sig":                 depositData.Signature,
			"withdrawCredentials": depositData.WithdrawalCredentials,
			"amount":              depositData.Amount,
			"forkVersion":         depositData.ForkVersion,
			"networkName":         depositData.NetworkName,
		}).Info("deposit-params")

		// increase nextKeyIndex
		task.nextKeyIndex++

	}

	err = task.connectionOfTrustNodeAccount.LockAndUpdateTxOpts()
	if err != nil {
		return fmt.Errorf("LockAndUpdateTxOpts err: %s", err)
	}
	defer task.connectionOfTrustNodeAccount.UnlockTxOpts()

	depositTx, err := task.nodeDepositContract.Deposit(task.connectionOfTrustNodeAccount.TxOpts(), validatorPubkeys, sigs, dataRoots)
	if err != nil {
		return errors.Wrap(err, "trustNodeContract.Deposit failed")
	}

	logrus.WithFields(logrus.Fields{
		"txHash":           depositTx.Hash(),
		"validatorPubkeys": validatorPubkeys,
		"sigs":             sigs,
		"dataRoots":        dataRoots,
	}).Info("deposit-tx")

	err = task.waitTxOk(depositTx.Hash())
	if err != nil {
		return err
	}

	for _, pubkey := range validatorPubkeys {
		status, err := task.mustGetTrustNodePubkeyStatus(pubkey)
		if err != nil {
			return fmt.Errorf("mustGetTrustNodePubkeyStatus err: %s", err.Error())
		}
		if status == utils.ValidatorStatusUnInitial {
			return fmt.Errorf("validator %s not exist on chain", hex.EncodeToString(pubkey))
		}
	}

	return nil
}
