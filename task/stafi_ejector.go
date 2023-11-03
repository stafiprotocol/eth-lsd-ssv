package task

import (
	"math"
	"math/big"

	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v3/beacon-chain/core/signing"
	primTypes "github.com/prysmaticlabs/prysm/v3/consensus-types/primitives"
	"github.com/prysmaticlabs/prysm/v3/crypto/bls"
	ethpb "github.com/prysmaticlabs/prysm/v3/proto/prysm/v1alpha1"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/connection/beacon"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/connection/types"
)

func (task *Task) checkCycle(cycle int64) error {
	logrus.Debugf("checkCycle %d", cycle)
	ejectedValidators, err := task.networkWithdrawContract.GetEjectedValidatorsAtCycle(nil, big.NewInt(cycle))
	if err != nil {
		return err
	}

	for _, ejectedValidator := range ejectedValidators {
		var validator *Validator
		var exist bool

		task.validatorsByValIndexMutex.RLock()
		if validator, exist = task.validatorsByValIndex[ejectedValidator.Uint64()]; !exist {
			task.validatorsByValIndexMutex.RUnlock()
			continue
		}
		task.validatorsByValIndexMutex.RUnlock()

		if validator.exitEpoch != 0 {
			continue
		}

		if validator.statusOnBeacon == valStatusExitedOnBeacon {
			continue
		}

		logrus.Infof("validator %d ejected at cycle %d", validator.validatorIndex, cycle)
		// check beacon sync status
		syncStatus, err := task.connectionOfTrustNodeAccount.Eth2Client().GetSyncStatus()
		if err != nil {
			return err
		}
		if syncStatus.Syncing {
			return errors.New("could not perform exit: beacon node is syncing.")
		}
		beaconHead, err := task.connectionOfTrustNodeAccount.Eth2Client().GetBeaconHead()
		if err != nil {
			return err
		}
		// check exited before
		pubkey := types.BytesToValidatorPubkey(validator.privateKey.PublicKey().Marshal())
		status, err := task.connectionOfTrustNodeAccount.GetValidatorStatus(pubkey, &beacon.ValidatorStatusOptions{Epoch: &beaconHead.Epoch})
		if err != nil {
			return err
		}
		// will skip if already sign exit
		if status.ExitEpoch != math.MaxUint64 {
			logrus.Infof("validator %d will exit at epoch %d", validator.validatorIndex, status.ExitEpoch)
			validator.exitEpoch = status.ExitEpoch
			continue
		}

		currentEpoch := primTypes.Epoch(beaconHead.Epoch)

		// not active
		if uint64(currentEpoch) < status.ActivationEpoch {
			logrus.Warnf("validator %d is not active and can't exit, will skip, active epoch: %d current epoch: %d", validator.validatorIndex, status.ActivationEpoch, currentEpoch)
			continue
		}
		if currentEpoch < primTypes.Epoch(status.ActivationEpoch)+shardCommitteePeriod {
			logrus.Warnf("validator %d is not active long enough and can't exit, will skip, active epoch: %d current epoch: %d", validator.validatorIndex, status.ActivationEpoch, currentEpoch)
			continue
		}

		// will sign and broadcast exit msg
		exit := &ethpb.VoluntaryExit{Epoch: currentEpoch, ValidatorIndex: primTypes.ValidatorIndex(validator.validatorIndex)}

		domain, err := task.connectionOfTrustNodeAccount.Eth2Client().GetDomainData(domainVoluntaryExit[:], uint64(exit.Epoch))
		if err != nil {
			return errors.Wrap(err, "Get domainData failed")
		}
		exitRoot, err := signing.ComputeSigningRoot(exit, domain)
		if err != nil {
			return errors.Wrap(err, "ComputeSigningRoot failed")
		}

		secretKey, err := bls.SecretKeyFromBytes(validator.privateKey.Marshal())
		if err != nil {
			return errors.Wrap(err, "failed to initialize keys caches from account keystore")
		}
		sig := secretKey.Sign(exitRoot[:])

		err = task.connectionOfTrustNodeAccount.Eth2Client().ExitValidator(validator.validatorIndex, uint64(currentEpoch), types.BytesToValidatorSignature(sig.Marshal()))
		if err != nil {
			return err
		}

		logrus.Infof("validator %d broadcast voluntary exit ok", validator.validatorIndex)

	}
	return nil
}
