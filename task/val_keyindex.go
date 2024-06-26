package task

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/constants"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/credential"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/utils"
)

func (task *Task) initValNextKeyIndex() error {
	task.nextKeyIndex = 0
	return task.checkAndRepairValNexKeyIndex()
}

func (task *Task) checkAndRepairValNexKeyIndex() error {
	retry := 0
	for {
		if retry > utils.RetryLimit {
			return fmt.Errorf("findNextKeyIndex reach retry limit")
		}
		credential, err := credential.NewCredential(task.copySeed(), task.nextKeyIndex, nil, constants.Chain{}, task.eth1WithdrawalAddress)
		if err != nil {
			return err
		}
		pubkey := credential.SigningPK().Marshal()

		pubkeyStatus, err := task.mustGetTrustNodePubkeyStatus(pubkey)
		if err != nil {
			logrus.Warnf("GetTrustNodePubkeyStatus err: %s", err.Error())
			time.Sleep(utils.RetryInterval)
			retry++
			continue
		}

		if pubkeyStatus == utils.ValidatorStatusUnInitial {
			break
		}

		var valStatus uint8
		switch pubkeyStatus {
		case utils.ValidatorStatusUnInitial:
			return fmt.Errorf("should not happen here")
		case utils.ValidatorStatusDeposited:
			valStatus = valStatusDeposited
		case utils.ValidatorStatusWithdrawMatch:
			valStatus = valStatusMatch
		case utils.ValidatorStatusWithdrawUnmatch:
			valStatus = valStatusUnmatch
		case utils.ValidatorStatusStaked:
			valStatus = valStatusStaked
		default:
			return fmt.Errorf("validator %s at index %d unknown status %d", hex.EncodeToString(pubkey), task.nextKeyIndex, pubkeyStatus)
		}

		val := &Validator{
			privateKey:    credential.SigningSk,
			statusOnStafi: valStatus,
			keyIndex:      task.nextKeyIndex,
		}
		task.validatorsByKeyIndex[task.nextKeyIndex] = val
		task.validatorsByPubkey[hex.EncodeToString(pubkey)] = val

		logrus.WithFields(logrus.Fields{
			"keyIndex":      task.nextKeyIndex,
			"pubkey":        hex.EncodeToString(pubkey),
			"statusOnStafi": pubkeyStatus,
		}).Debug("validator key info")

		task.nextKeyIndex++
	}

	return nil
}
