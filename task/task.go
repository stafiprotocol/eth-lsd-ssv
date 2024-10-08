package task

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"reflect"
	"runtime"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/forta-network/go-multicall"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v4/config/params"
	types "github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	"github.com/prysmaticlabs/prysm/v4/encoding/bytesutil"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/stafiprotocol/chainbridge/utils/crypto/secp256k1"
	erc20 "github.com/stafiprotocol/eth-lsd-ssv-client/bindings/Erc20"
	lsd_network_factory "github.com/stafiprotocol/eth-lsd-ssv-client/bindings/LsdNetworkFactory"
	network_proposal "github.com/stafiprotocol/eth-lsd-ssv-client/bindings/NetworkProposal"
	network_withdraw "github.com/stafiprotocol/eth-lsd-ssv-client/bindings/NetworkWithdraw"
	node_deposit "github.com/stafiprotocol/eth-lsd-ssv-client/bindings/NodeDeposit"
	ssv_network "github.com/stafiprotocol/eth-lsd-ssv-client/bindings/SsvNetwork"
	ssv_network_views "github.com/stafiprotocol/eth-lsd-ssv-client/bindings/SsvNetworkViews"
	user_deposit "github.com/stafiprotocol/eth-lsd-ssv-client/bindings/UserDeposit"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/config"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/connection"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/connection/beacon"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/constants"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/crypto/bls"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/keyshare"
	"github.com/stafiprotocol/eth-lsd-ssv-client/pkg/utils"
)

var (
	minAmountNeedStake   = decimal.NewFromBigInt(big.NewInt(31), 18)
	minAmountNeedDeposit = decimal.NewFromBigInt(big.NewInt(32), 18)

	trustNodeDepositAmount = decimal.NewFromBigInt(big.NewInt(1), 18)
	trustNodeStakeAmount   = decimal.NewFromBigInt(big.NewInt(31), 18)

	blocksOfOneYear = decimal.NewFromInt(2629800)
)

const minNumberOfTargetOperators = 4

var (
	domainVoluntaryExit  = bytesutil.Uint32ToBytes4(0x04000000)
	shardCommitteePeriod = types.Epoch(256) // ShardCommitteePeriod is the minimum amount of epochs a validator must participate before exiting.
)

const (
	valStatusUnInitiated = uint8(0)

	// on stafi
	valStatusDeposited = uint8(1)
	valStatusMatch     = uint8(2)
	valStatusUnmatch   = uint8(3)
	valStatusStaked    = uint8(4)

	// on ssv
	valStatusRegistedOnSsvValid   = uint8(1)
	valStatusRegistedOnSsvInvalid = uint8(2)
	valStatusRemovedOnSsv         = uint8(3)

	// on beacon
	valStatusActiveOnBeacon = uint8(1)
	valStatusExitedOnBeacon = uint8(2)
)

// only support stafi trust node account now !!!
// 0. find next key index and cache validator status on start
// 1. update validator status(on execution/ssv/beacon) periodically
// 2. check stakepool balance periodically, call stake/deposit if match
// 3. register validator on ssv, if status is staked on stafi contract
// 4. remove validator on ssv, if status is exited on beacon
type Task struct {
	stop            chan struct{}
	eth1StartHeight uint64
	eth1Endpoint    string
	eth2Endpoint    string

	trustNodeKeyPair *secp256k1.Keypair
	ssvKeyPair       *secp256k1.Keypair

	gasLimit            *big.Int
	maxGasPrice         *big.Int
	poolReservedBalance *big.Int
	seed                []byte
	isViewMode          bool
	targetOperatorIds   []uint64

	ssvNetworkContractAddress      common.Address
	ssvNetworkViewsContractAddress common.Address
	ssvTokenContractAddress        common.Address

	lsdNetworkFactoryAddress common.Address
	lsdTokenAddress          common.Address

	ssvApiNetwork string
	chain         constants.Chain

	connectionOfTrustNodeAccount *connection.Connection
	connectionOfSsvAccount       *connection.Connection

	eth1WithdrawalAddress      common.Address
	feeRecipientAddressOnStafi common.Address
	latestRegistrationNonce    uint64
	latestTxBlock              uint64

	eth1Client *ethclient.Client
	eth2Config beacon.Eth2Config

	ssvNetworkContract      *ssv_network.SsvNetwork
	ssvNetworkViewsContract *ssv_network_views.SsvNetworkViews
	ssvTokenContract        *erc20.Erc20

	lsdNetworkFactoryContract *lsd_network_factory.LsdNetworkFactory
	userDepositContract       *user_deposit.UserDeposit
	nodeDepositContract       *node_deposit.NodeDeposit
	networkWithdrawContract   *network_withdraw.NetworkWithdraw
	networkProposalContract   *network_proposal.NetworkProposal

	multicaler *multicall.Caller

	ssvNetworkAbi abi.ABI

	nextKeyIndex                int
	dealtEth1Block              uint64 // for offchain state
	validatorsPerOperatorLimit  uint64
	ValidatorsPerTrustNodeLimit uint64
	ValidatorsLimitByGas        uint64 // gas = 162917*n+268921

	validatorsByKeyIndex      map[int]*Validator    // key index => validator, cache all validators(pending/active/exist) by keyIndex
	validatorsByPubkey        map[string]*Validator // pubkey => validator, cache all validators(pending/active/exist) by pubkey
	validatorsByValIndex      map[uint64]*Validator // val index => validator
	validatorsByValIndexMutex sync.RWMutex

	targetOperators map[uint64]*keyshare.Operator

	// ssv offchain state
	clusters                 map[string]*Cluster // cluster key => cluster
	feeRecipientAddressOnSsv common.Address

	handlers     []func() error
	handlersName []string
}

type Cluster struct {
	operatorIds   []uint64
	latestCluster *ssv_network.ISSVNetworkCoreCluster

	balance decimal.Decimal

	managingValidators map[int]struct{} // key index

	latestUpdateClusterBlockNumber    uint64
	latestValidatorAddedBlockNumber   uint64
	latestValidatorRemovedBlockNumber uint64
}

type Validator struct {
	privateKey *bls.PrivateKey
	keyIndex   int

	statusOnStafi  uint8
	statusOnSsv    uint8
	statusOnBeacon uint8

	validatorIndex uint64
	exitEpoch      uint64

	clusterKey string

	removedFromSsvOnBlock uint64
}

func NewTask(cfg *config.Config, seed []byte, isViewMode bool, trustNodeKeyPair, ssvKeyPair *secp256k1.Keypair) (*Task, error) {
	if !common.IsHexAddress(cfg.Contracts.SsvNetworkAddress) {
		return nil, fmt.Errorf("ssvnetwork contract address fmt err")
	}
	if !common.IsHexAddress(cfg.Contracts.SsvNetworkViewsAddress) {
		return nil, fmt.Errorf("ssvnetworkviews contract address fmt err")
	}
	if !common.IsHexAddress(cfg.Contracts.SsvTokenAddress) {
		return nil, fmt.Errorf("SsvTokenAddress contract address fmt err")
	}
	if !common.IsHexAddress(cfg.Contracts.LsdFactoryAddress) {
		return nil, fmt.Errorf("LsdNetworkFactory contract address fmt err")
	}
	if !common.IsHexAddress(cfg.Contracts.LsdTokenAddress) {
		return nil, fmt.Errorf("LsdToken contract address fmt err")
	}

	if len(cfg.TargetOperators) < minNumberOfTargetOperators {
		return nil, fmt.Errorf("target operators number %d less than %d", len(cfg.TargetOperators), minNumberOfTargetOperators)
	}

	gasLimitDeci, err := decimal.NewFromString(cfg.GasLimit)
	if err != nil {
		return nil, fmt.Errorf("parse config gasLimit error: %w", err)
	}

	if gasLimitDeci.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("gas limit is zero")
	}

	// gas = 162917*n+268921
	n := gasLimitDeci.Sub(decimal.NewFromInt(268921)).Div(decimal.NewFromInt(162917))
	validatorsLimitByGas := n.BigInt().Uint64()
	if validatorsLimitByGas == 0 {
		return nil, fmt.Errorf("gasLimit %s too small", cfg.GasLimit)
	}
	logrus.Debug("validatorsLimitByGas ", validatorsLimitByGas)

	maxGasPriceDeci, err := decimal.NewFromString(cfg.MaxGasPrice)
	if err != nil {
		return nil, fmt.Errorf("parse config maxGasPrice error: %w", err)
	}
	if maxGasPriceDeci.LessThanOrEqual(decimal.Zero) {
		return nil, fmt.Errorf("max gas price is zero")
	}

	poolReservedBalance := big.NewInt(0)
	if len(cfg.PoolReservedBalance) > 0 {
		reservedBalance, err := decimal.NewFromString(cfg.PoolReservedBalance)
		if err != nil {
			return nil, err
		}
		if maxGasPriceDeci.IsNegative() {
			return nil, fmt.Errorf("PoolReservedBalance is negative")
		}
		poolReservedBalance = reservedBalance.BigInt()
	}

	eth1client, err := ethclient.Dial(cfg.Eth1Endpoint)
	if err != nil {
		return nil, err
	}

	s := &Task{
		stop:                 make(chan struct{}),
		eth1Endpoint:         cfg.Eth1Endpoint,
		eth2Endpoint:         cfg.Eth2Endpoint,
		eth1Client:           eth1client,
		trustNodeKeyPair:     trustNodeKeyPair,
		ssvKeyPair:           ssvKeyPair,
		seed:                 seed,
		isViewMode:           isViewMode,
		gasLimit:             gasLimitDeci.BigInt(),
		maxGasPrice:          maxGasPriceDeci.BigInt(),
		poolReservedBalance:  poolReservedBalance,
		eth1StartHeight:      utils.TheMergeBlockNumber,
		targetOperatorIds:    cfg.TargetOperators,
		ValidatorsLimitByGas: validatorsLimitByGas,

		ssvNetworkContractAddress:      common.HexToAddress(cfg.Contracts.SsvNetworkAddress),
		ssvNetworkViewsContractAddress: common.HexToAddress(cfg.Contracts.SsvNetworkViewsAddress),
		ssvTokenContractAddress:        common.HexToAddress(cfg.Contracts.SsvTokenAddress),
		lsdNetworkFactoryAddress:       common.HexToAddress(cfg.Contracts.LsdFactoryAddress),
		lsdTokenAddress:                common.HexToAddress(cfg.Contracts.LsdTokenAddress),

		validatorsByKeyIndex: make(map[int]*Validator),
		validatorsByPubkey:   make(map[string]*Validator),
		validatorsByValIndex: make(map[uint64]*Validator),

		targetOperators: make(map[uint64]*keyshare.Operator),

		clusters: make(map[string]*Cluster),
	}

	return s, nil
}

func (task *Task) Start() error {
	var err error
	task.connectionOfTrustNodeAccount, err = connection.NewConnection(task.eth1Endpoint, task.eth2Endpoint, task.trustNodeKeyPair,
		task.gasLimit, task.maxGasPrice)
	if err != nil {
		return err
	}
	task.connectionOfSsvAccount, err = connection.NewConnection(task.eth1Endpoint, task.eth2Endpoint, task.ssvKeyPair,
		task.gasLimit, task.maxGasPrice)
	if err != nil {
		return err
	}
	chainId, err := task.eth1Client.ChainID(context.Background())
	if err != nil {
		return err
	}

	task.eth2Config, err = task.connectionOfTrustNodeAccount.Eth2Client().GetEth2Config()
	if err != nil {
		return err
	}

	switch chainId.Uint64() {
	case 1: //mainnet
		task.chain = constants.GetChain(constants.ChainMAINNET)
		if !bytes.Equal(task.eth2Config.GenesisForkVersion, params.MainnetConfig().GenesisForkVersion) {
			return fmt.Errorf("endpoint network not match")
		}
		task.dealtEth1Block = 17705353
		task.ssvApiNetwork = "mainnet"

	case 11155111: // sepolia
		task.chain = constants.GetChain(constants.ChainSEPOLIA)
		if !bytes.Equal(task.eth2Config.GenesisForkVersion, params.SepoliaConfig().GenesisForkVersion) {
			return fmt.Errorf("endpoint network not match")
		}
		task.dealtEth1Block = 9354882
		task.ssvApiNetwork = "prater"
	case 17000: // holesky
		task.chain = constants.GetChain(constants.ChainHOLESKY)
		if !bytes.Equal(task.eth2Config.GenesisForkVersion, params.HoleskyConfig().GenesisForkVersion) {
			return fmt.Errorf("endpoint network not match")
		}
		task.dealtEth1Block = 549480
		task.ssvApiNetwork = "holesky"
	case 5: // goerli
		task.chain = constants.GetChain(constants.ChainGOERLI)
		if !bytes.Equal(task.eth2Config.GenesisForkVersion, params.PraterConfig().GenesisForkVersion) {
			return fmt.Errorf("endpoint network not match")
		}
		task.dealtEth1Block = 9403883
		task.ssvApiNetwork = "prater"

	default:
		return fmt.Errorf("unsupport chainId: %d", chainId.Int64())
	}
	if err != nil {
		return err
	}

	task.ssvNetworkAbi, err = abi.JSON(strings.NewReader(ssv_network.SsvNetworkABI))
	if err != nil {
		return err
	}

	err = task.initContract()
	if err != nil {
		return err
	}

	caller, err := multicall.Dial(context.Background(), task.eth1Endpoint)
	if err != nil {
		return err
	}
	task.multicaler = caller

	latestBlockNumber, err := task.connectionOfTrustNodeAccount.Eth1LatestBlock()
	if err != nil {
		return err
	}

	if err = task.checkTrustNode(); err != nil {
		return err
	}

	// check target operator id
	notActiveOperators := make([]uint64, 0)
	for _, opId := range task.targetOperatorIds {
		if _, exist := task.targetOperators[opId]; exist {
			return fmt.Errorf("duplicate operator id: %d", opId)
		}

		owner, fee, validatorCount, _, isPrivate, _, err := task.ssvNetworkViewsContract.GetOperatorById(nil, opId)
		if err != nil {
			return errors.Wrap(err, "ssvNetworkViewsContract.GetOperatorById failed")
		}
		if isPrivate {
			return fmt.Errorf("target operator %d is private", opId)
		}

		// fetch and check acitve status from api
		operatorFromApi, err := utils.MustGetOperatorDetail(task.ssvApiNetwork, opId)
		if err != nil {
			return err
		}
		if operatorFromApi.IsActive != 1 {
			notActiveOperators = append(notActiveOperators, opId)
			continue
		}

		// fetch pubkey
		operatorAddedEvent, err := task.loadOperatorAddedEvent(opId, owner, latestBlockNumber)
		// operatorAddedEvent, err := task.ssvNetworkContract.FilterOperatorAdded(&bind.FilterOpts{
		// 	Start:   181612,
		// 	End:     nil,
		// 	Context: context.Background(),
		// }, []uint64{opId}, []common.Address{owner})
		if err != nil {
			return err
		}
		if operatorAddedEvent == nil {
			return fmt.Errorf("filter operator pubkey failed, opId %d owner %s", opId, owner.String())
		}
		pubkey, err := unpackOperatorPublicKey(operatorAddedEvent.Event.PublicKey)
		if err != nil {
			return err
		}
		task.targetOperators[opId] = &keyshare.Operator{
			Id:             opId,
			PublicKey:      string(pubkey),
			Fee:            decimal.NewFromBigInt(fee, 0),
			Active:         true,
			ValidatorCount: uint64(validatorCount),
		}
	}
	if len(notActiveOperators) > 0 {
		return fmt.Errorf("target operators %v is not active", notActiveOperators)
	}

	err = task.initValNextKeyIndex()
	if err != nil {
		return err
	}
	logrus.Infof("nextKeyIndex: %d", task.nextKeyIndex)
	logrus.Infof("init state...")

	// retry := 0
	// for {
	// 	if retry > utils.RetryLimit {
	// 		return fmt.Errorf("init state reach retry limit, err: %s", err.Error())
	// 	}
	// 	select {
	// 	case <-task.stop:
	// 		logrus.Info("task has stopped")
	// 		return nil
	// 	default:
	// 	}

	// 	err = task.updateSsvOffchainState()
	// 	if err != nil {
	// 		retry++
	// 		logrus.Warnf("updateSsvOffchainState err: %s", err.Error())
	// 		continue
	// 	}
	// 	err = task.updateValStatus()
	// 	if err != nil {
	// 		logrus.Warnf("updateValStatus err: %s", err.Error())
	// 		retry++
	// 		continue
	// 	}
	// 	err = task.updateOperatorStatus()
	// 	if err != nil {
	// 		logrus.Warnf("updateValStatus err: %s", err.Error())
	// 		retry++
	// 		continue
	// 	}

	// 	break
	// }

	task.appendHandlers(
		task.checkAndRepairValNexKeyIndex,
		task.updateSsvOffchainState,
		task.updateValStatus,
		task.updateOperatorStatus,
	)
	if !task.isViewMode {
		task.appendHandlers(
			task.checkAndStake, //stafi
			task.checkAndDeposit,
			task.checkAndSetFeeRecipient, // ssv
			task.checkAndReactiveOnSSV,
			task.checkAndOnboardOnSSV,
			task.checkAndOffboardOnSSV,
			task.checkAndWithdrawOnSSV,
		)

		utils.SafeGo(task.ejectorService)
	}

	utils.SafeGo(task.ssvService)

	return nil
}

func (task *Task) Stop() {
	close(task.stop)
}

func (task *Task) loadOperatorAddedEvent(opId uint64, owner common.Address, endBlock uint64) (*ssv_network.SsvNetworkOperatorAddedIterator, error) {
	var start uint64 = 180000
	for start <= endBlock {
		end := start + fetchEventBlockLimit
		if end > endBlock {
			end = endBlock
		}
		if start > endBlock {
			break
		}

		operatorAddedEvent, err := task.ssvNetworkContract.FilterOperatorAdded(&bind.FilterOpts{
			Start:   start,
			End:     &end,
			Context: context.Background(),
		}, []uint64{opId}, []common.Address{owner})
		if err != nil {
			return nil, errors.Wrapf(err, "ssvNetworkContract.FilterOperatorAdded failed, opId %d owner %s", opId, owner.String())
		}
		if operatorAddedEvent.Next() {
			return operatorAddedEvent, nil
		}
		start = end + 1
	}
	return nil, fmt.Errorf("filter operator pubkey failed, opId %d owner %s", opId, owner.String())
}

func (task *Task) initContract() error {
	var err error

	task.lsdNetworkFactoryContract, err = lsd_network_factory.NewLsdNetworkFactory(task.lsdNetworkFactoryAddress, task.eth1Client)
	if err != nil {
		return err
	}

	networkContracts, err := task.lsdNetworkFactoryContract.NetworkContractsOfLsdToken(nil, task.lsdTokenAddress)
	if err != nil {
		return err
	}
	logrus.Infof("networkContracts: %+v", networkContracts)
	task.feeRecipientAddressOnStafi = networkContracts.FeePool
	task.eth1WithdrawalAddress = networkContracts.NetworkWithdraw

	task.eth1StartHeight = networkContracts.Block.Uint64()

	task.nodeDepositContract, err = node_deposit.NewNodeDeposit(networkContracts.NodeDeposit, task.eth1Client)
	if err != nil {
		return err
	}

	task.userDepositContract, err = user_deposit.NewUserDeposit(networkContracts.UserDeposit, task.eth1Client)
	if err != nil {
		return err
	}
	task.networkWithdrawContract, err = network_withdraw.NewNetworkWithdraw(networkContracts.NetworkWithdraw, task.eth1Client)
	if err != nil {
		return err
	}

	task.networkProposalContract, err = network_proposal.NewNetworkProposal(networkContracts.NetworkProposal, task.eth1Client)
	if err != nil {
		return err
	}

	task.ssvNetworkContract, err = ssv_network.NewSsvNetwork(task.ssvNetworkContractAddress, task.eth1Client)
	if err != nil {
		return err
	}
	task.ssvNetworkViewsContract, err = ssv_network_views.NewSsvNetworkViews(task.ssvNetworkViewsContractAddress, task.eth1Client)
	if err != nil {
		return err
	}

	task.ssvTokenContract, err = erc20.NewErc20(task.ssvTokenContractAddress, task.eth1Client)
	if err != nil {
		return err
	}
	return nil
}

// check whether node is a member of trust nodes or not
func (task *Task) checkTrustNode() error {
	addr := task.trustNodeKeyPair.CommonAddress()
	info, err := task.nodeDepositContract.NodeInfoOf(nil, addr)
	if err != nil {
		return fmt.Errorf("fail to get info of node: %s err: %w", addr, err)
	}
	if info.NodeType != 2 {
		return fmt.Errorf("address: %s has not been registered as a trust node", addr)
	}
	if info.Removed {
		return fmt.Errorf("address: %s has been removed from trust nodes", addr)
	}
	return nil
}

func (task *Task) appendHandlers(handlers ...func() error) {
	for _, handler := range handlers {

		funcNameRaw := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()

		splits := strings.Split(funcNameRaw, "/")
		funcName := splits[len(splits)-1]

		task.handlersName = append(task.handlersName, funcName)
		task.handlers = append(task.handlers, handler)
	}
}

func (task *Task) waitTxOk(txHash common.Hash) error {
	blockNumber, err := utils.WaitTxOkCommon(task.eth1Client, txHash)
	if err != nil {
		return err
	}
	task.latestTxBlock = blockNumber
	return nil
}

func (task *Task) offchainStateIsLatest() bool {
	return task.dealtEth1Block > task.latestTxBlock
}
