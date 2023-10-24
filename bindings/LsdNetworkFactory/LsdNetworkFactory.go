// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lsd_network_factory

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ILsdNetworkFactoryNetworkContracts is an auto generated low-level Go binding around an user-defined struct.
type ILsdNetworkFactoryNetworkContracts struct {
	FeePool         common.Address
	NetworkBalances common.Address
	NetworkProposal common.Address
	NodeDeposit     common.Address
	UserDeposit     common.Address
	NetworkWithdraw common.Address
	LsdToken        common.Address
	Block           *big.Int
}

// LsdNetworkFactoryMetaData contains all meta data concerning the LsdNetworkFactory contract.
var LsdNetworkFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDealedEpoch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDealedHeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyNotifyCycle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyVoted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountUnmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableDepositZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableRewardZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CycleNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositAmountLTMinAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LsdTokenAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyRemoved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeNotClaimable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotClaimable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotFactoryAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNetworkAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNetworkProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPubkeyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTrustNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProposalAlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProposalExecFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyStatusUnmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateChangeOverLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateValueUnmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReachCycleWithdrawLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReachPubkeyNumberLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReachUserWithdrawLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SecondsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SoloNodeDepositAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SoloNodeDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmitBalancesDisable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustNodeDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VoterNumberOverLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawIndexEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawIndexOver\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_feePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkBalances\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkProposal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeDeposit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_userDeposit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkWithdraw\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lsdToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structILsdNetworkFactory.NetworkContracts\",\"name\":\"contracts\",\"type\":\"tuple\"}],\"name\":\"LsdNetwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_lsdTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_lsdTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_networkAdmin\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"createLsdNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_lsdTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_lsdTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDelay\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"proposers\",\"type\":\"address[]\"}],\"name\":\"createLsdNetworkWithTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDepositAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"factoryClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePoolLogicAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factoryAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethDepositAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feePoolLogicAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkBalancesLogicAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkProposalLogicAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeDepositLogicAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_userDepositLogicAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkWithdrawLogicAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creater\",\"type\":\"address\"}],\"name\":\"lsdTokensOfCreater\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkBalancesLogicAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"networkContractsOfLsdToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_feePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkBalances\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkProposal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeDeposit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_userDeposit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkWithdraw\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lsdToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkProposalLogicAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkWithdrawLogicAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeDepositLogicAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_networkBalancesLogicAddress\",\"type\":\"address\"}],\"name\":\"setNetworkBalancesLogicAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_networkProposalLogicAddress\",\"type\":\"address\"}],\"name\":\"setNetworkProposalLogicAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_networkWithdrawLogicAddress\",\"type\":\"address\"}],\"name\":\"setNetworkWithdrawLogicAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeDepositLogicAddress\",\"type\":\"address\"}],\"name\":\"setNodeDepositLogicAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userDepositLogicAddress\",\"type\":\"address\"}],\"name\":\"setUserDepositLogicAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDepositLogicAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LsdNetworkFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use LsdNetworkFactoryMetaData.ABI instead.
var LsdNetworkFactoryABI = LsdNetworkFactoryMetaData.ABI

// LsdNetworkFactory is an auto generated Go binding around an Ethereum contract.
type LsdNetworkFactory struct {
	LsdNetworkFactoryCaller     // Read-only binding to the contract
	LsdNetworkFactoryTransactor // Write-only binding to the contract
	LsdNetworkFactoryFilterer   // Log filterer for contract events
}

// LsdNetworkFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LsdNetworkFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LsdNetworkFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LsdNetworkFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LsdNetworkFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LsdNetworkFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LsdNetworkFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LsdNetworkFactorySession struct {
	Contract     *LsdNetworkFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LsdNetworkFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LsdNetworkFactoryCallerSession struct {
	Contract *LsdNetworkFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LsdNetworkFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LsdNetworkFactoryTransactorSession struct {
	Contract     *LsdNetworkFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LsdNetworkFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LsdNetworkFactoryRaw struct {
	Contract *LsdNetworkFactory // Generic contract binding to access the raw methods on
}

// LsdNetworkFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LsdNetworkFactoryCallerRaw struct {
	Contract *LsdNetworkFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// LsdNetworkFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LsdNetworkFactoryTransactorRaw struct {
	Contract *LsdNetworkFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLsdNetworkFactory creates a new instance of LsdNetworkFactory, bound to a specific deployed contract.
func NewLsdNetworkFactory(address common.Address, backend bind.ContractBackend) (*LsdNetworkFactory, error) {
	contract, err := bindLsdNetworkFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LsdNetworkFactory{LsdNetworkFactoryCaller: LsdNetworkFactoryCaller{contract: contract}, LsdNetworkFactoryTransactor: LsdNetworkFactoryTransactor{contract: contract}, LsdNetworkFactoryFilterer: LsdNetworkFactoryFilterer{contract: contract}}, nil
}

// NewLsdNetworkFactoryCaller creates a new read-only instance of LsdNetworkFactory, bound to a specific deployed contract.
func NewLsdNetworkFactoryCaller(address common.Address, caller bind.ContractCaller) (*LsdNetworkFactoryCaller, error) {
	contract, err := bindLsdNetworkFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LsdNetworkFactoryCaller{contract: contract}, nil
}

// NewLsdNetworkFactoryTransactor creates a new write-only instance of LsdNetworkFactory, bound to a specific deployed contract.
func NewLsdNetworkFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LsdNetworkFactoryTransactor, error) {
	contract, err := bindLsdNetworkFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LsdNetworkFactoryTransactor{contract: contract}, nil
}

// NewLsdNetworkFactoryFilterer creates a new log filterer instance of LsdNetworkFactory, bound to a specific deployed contract.
func NewLsdNetworkFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LsdNetworkFactoryFilterer, error) {
	contract, err := bindLsdNetworkFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LsdNetworkFactoryFilterer{contract: contract}, nil
}

// bindLsdNetworkFactory binds a generic wrapper to an already deployed contract.
func bindLsdNetworkFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LsdNetworkFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LsdNetworkFactory *LsdNetworkFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LsdNetworkFactory.Contract.LsdNetworkFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LsdNetworkFactory *LsdNetworkFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.LsdNetworkFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LsdNetworkFactory *LsdNetworkFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.LsdNetworkFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LsdNetworkFactory *LsdNetworkFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LsdNetworkFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.contract.Transact(opts, method, params...)
}

// EthDepositAddress is a free data retrieval call binding the contract method 0xb420feb2.
//
// Solidity: function ethDepositAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCaller) EthDepositAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LsdNetworkFactory.contract.Call(opts, &out, "ethDepositAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthDepositAddress is a free data retrieval call binding the contract method 0xb420feb2.
//
// Solidity: function ethDepositAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactorySession) EthDepositAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.EthDepositAddress(&_LsdNetworkFactory.CallOpts)
}

// EthDepositAddress is a free data retrieval call binding the contract method 0xb420feb2.
//
// Solidity: function ethDepositAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCallerSession) EthDepositAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.EthDepositAddress(&_LsdNetworkFactory.CallOpts)
}

// FactoryAdmin is a free data retrieval call binding the contract method 0x17d8ec7f.
//
// Solidity: function factoryAdmin() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCaller) FactoryAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LsdNetworkFactory.contract.Call(opts, &out, "factoryAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryAdmin is a free data retrieval call binding the contract method 0x17d8ec7f.
//
// Solidity: function factoryAdmin() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactorySession) FactoryAdmin() (common.Address, error) {
	return _LsdNetworkFactory.Contract.FactoryAdmin(&_LsdNetworkFactory.CallOpts)
}

// FactoryAdmin is a free data retrieval call binding the contract method 0x17d8ec7f.
//
// Solidity: function factoryAdmin() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCallerSession) FactoryAdmin() (common.Address, error) {
	return _LsdNetworkFactory.Contract.FactoryAdmin(&_LsdNetworkFactory.CallOpts)
}

// FeePoolLogicAddress is a free data retrieval call binding the contract method 0x93ba4506.
//
// Solidity: function feePoolLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCaller) FeePoolLogicAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LsdNetworkFactory.contract.Call(opts, &out, "feePoolLogicAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeePoolLogicAddress is a free data retrieval call binding the contract method 0x93ba4506.
//
// Solidity: function feePoolLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactorySession) FeePoolLogicAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.FeePoolLogicAddress(&_LsdNetworkFactory.CallOpts)
}

// FeePoolLogicAddress is a free data retrieval call binding the contract method 0x93ba4506.
//
// Solidity: function feePoolLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCallerSession) FeePoolLogicAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.FeePoolLogicAddress(&_LsdNetworkFactory.CallOpts)
}

// LsdTokensOfCreater is a free data retrieval call binding the contract method 0x15ae831a.
//
// Solidity: function lsdTokensOfCreater(address _creater) view returns(address[])
func (_LsdNetworkFactory *LsdNetworkFactoryCaller) LsdTokensOfCreater(opts *bind.CallOpts, _creater common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _LsdNetworkFactory.contract.Call(opts, &out, "lsdTokensOfCreater", _creater)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// LsdTokensOfCreater is a free data retrieval call binding the contract method 0x15ae831a.
//
// Solidity: function lsdTokensOfCreater(address _creater) view returns(address[])
func (_LsdNetworkFactory *LsdNetworkFactorySession) LsdTokensOfCreater(_creater common.Address) ([]common.Address, error) {
	return _LsdNetworkFactory.Contract.LsdTokensOfCreater(&_LsdNetworkFactory.CallOpts, _creater)
}

// LsdTokensOfCreater is a free data retrieval call binding the contract method 0x15ae831a.
//
// Solidity: function lsdTokensOfCreater(address _creater) view returns(address[])
func (_LsdNetworkFactory *LsdNetworkFactoryCallerSession) LsdTokensOfCreater(_creater common.Address) ([]common.Address, error) {
	return _LsdNetworkFactory.Contract.LsdTokensOfCreater(&_LsdNetworkFactory.CallOpts, _creater)
}

// NetworkBalancesLogicAddress is a free data retrieval call binding the contract method 0xc0c66c3d.
//
// Solidity: function networkBalancesLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCaller) NetworkBalancesLogicAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LsdNetworkFactory.contract.Call(opts, &out, "networkBalancesLogicAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkBalancesLogicAddress is a free data retrieval call binding the contract method 0xc0c66c3d.
//
// Solidity: function networkBalancesLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactorySession) NetworkBalancesLogicAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.NetworkBalancesLogicAddress(&_LsdNetworkFactory.CallOpts)
}

// NetworkBalancesLogicAddress is a free data retrieval call binding the contract method 0xc0c66c3d.
//
// Solidity: function networkBalancesLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCallerSession) NetworkBalancesLogicAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.NetworkBalancesLogicAddress(&_LsdNetworkFactory.CallOpts)
}

// NetworkContractsOfLsdToken is a free data retrieval call binding the contract method 0x89e56b80.
//
// Solidity: function networkContractsOfLsdToken(address ) view returns(address _feePool, address _networkBalances, address _networkProposal, address _nodeDeposit, address _userDeposit, address _networkWithdraw, address _lsdToken, uint256 _block)
func (_LsdNetworkFactory *LsdNetworkFactoryCaller) NetworkContractsOfLsdToken(opts *bind.CallOpts, arg0 common.Address) (struct {
	FeePool         common.Address
	NetworkBalances common.Address
	NetworkProposal common.Address
	NodeDeposit     common.Address
	UserDeposit     common.Address
	NetworkWithdraw common.Address
	LsdToken        common.Address
	Block           *big.Int
}, error) {
	var out []interface{}
	err := _LsdNetworkFactory.contract.Call(opts, &out, "networkContractsOfLsdToken", arg0)

	outstruct := new(struct {
		FeePool         common.Address
		NetworkBalances common.Address
		NetworkProposal common.Address
		NodeDeposit     common.Address
		UserDeposit     common.Address
		NetworkWithdraw common.Address
		LsdToken        common.Address
		Block           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeePool = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NetworkBalances = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NetworkProposal = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.NodeDeposit = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.UserDeposit = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.NetworkWithdraw = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.LsdToken = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Block = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NetworkContractsOfLsdToken is a free data retrieval call binding the contract method 0x89e56b80.
//
// Solidity: function networkContractsOfLsdToken(address ) view returns(address _feePool, address _networkBalances, address _networkProposal, address _nodeDeposit, address _userDeposit, address _networkWithdraw, address _lsdToken, uint256 _block)
func (_LsdNetworkFactory *LsdNetworkFactorySession) NetworkContractsOfLsdToken(arg0 common.Address) (struct {
	FeePool         common.Address
	NetworkBalances common.Address
	NetworkProposal common.Address
	NodeDeposit     common.Address
	UserDeposit     common.Address
	NetworkWithdraw common.Address
	LsdToken        common.Address
	Block           *big.Int
}, error) {
	return _LsdNetworkFactory.Contract.NetworkContractsOfLsdToken(&_LsdNetworkFactory.CallOpts, arg0)
}

// NetworkContractsOfLsdToken is a free data retrieval call binding the contract method 0x89e56b80.
//
// Solidity: function networkContractsOfLsdToken(address ) view returns(address _feePool, address _networkBalances, address _networkProposal, address _nodeDeposit, address _userDeposit, address _networkWithdraw, address _lsdToken, uint256 _block)
func (_LsdNetworkFactory *LsdNetworkFactoryCallerSession) NetworkContractsOfLsdToken(arg0 common.Address) (struct {
	FeePool         common.Address
	NetworkBalances common.Address
	NetworkProposal common.Address
	NodeDeposit     common.Address
	UserDeposit     common.Address
	NetworkWithdraw common.Address
	LsdToken        common.Address
	Block           *big.Int
}, error) {
	return _LsdNetworkFactory.Contract.NetworkContractsOfLsdToken(&_LsdNetworkFactory.CallOpts, arg0)
}

// NetworkProposalLogicAddress is a free data retrieval call binding the contract method 0x7a3ddd32.
//
// Solidity: function networkProposalLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCaller) NetworkProposalLogicAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LsdNetworkFactory.contract.Call(opts, &out, "networkProposalLogicAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkProposalLogicAddress is a free data retrieval call binding the contract method 0x7a3ddd32.
//
// Solidity: function networkProposalLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactorySession) NetworkProposalLogicAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.NetworkProposalLogicAddress(&_LsdNetworkFactory.CallOpts)
}

// NetworkProposalLogicAddress is a free data retrieval call binding the contract method 0x7a3ddd32.
//
// Solidity: function networkProposalLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCallerSession) NetworkProposalLogicAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.NetworkProposalLogicAddress(&_LsdNetworkFactory.CallOpts)
}

// NetworkWithdrawLogicAddress is a free data retrieval call binding the contract method 0x8e96ce62.
//
// Solidity: function networkWithdrawLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCaller) NetworkWithdrawLogicAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LsdNetworkFactory.contract.Call(opts, &out, "networkWithdrawLogicAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkWithdrawLogicAddress is a free data retrieval call binding the contract method 0x8e96ce62.
//
// Solidity: function networkWithdrawLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactorySession) NetworkWithdrawLogicAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.NetworkWithdrawLogicAddress(&_LsdNetworkFactory.CallOpts)
}

// NetworkWithdrawLogicAddress is a free data retrieval call binding the contract method 0x8e96ce62.
//
// Solidity: function networkWithdrawLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCallerSession) NetworkWithdrawLogicAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.NetworkWithdrawLogicAddress(&_LsdNetworkFactory.CallOpts)
}

// NodeDepositLogicAddress is a free data retrieval call binding the contract method 0xdb7e2d18.
//
// Solidity: function nodeDepositLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCaller) NodeDepositLogicAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LsdNetworkFactory.contract.Call(opts, &out, "nodeDepositLogicAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeDepositLogicAddress is a free data retrieval call binding the contract method 0xdb7e2d18.
//
// Solidity: function nodeDepositLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactorySession) NodeDepositLogicAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.NodeDepositLogicAddress(&_LsdNetworkFactory.CallOpts)
}

// NodeDepositLogicAddress is a free data retrieval call binding the contract method 0xdb7e2d18.
//
// Solidity: function nodeDepositLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCallerSession) NodeDepositLogicAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.NodeDepositLogicAddress(&_LsdNetworkFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LsdNetworkFactory *LsdNetworkFactoryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LsdNetworkFactory.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LsdNetworkFactory *LsdNetworkFactorySession) ProxiableUUID() ([32]byte, error) {
	return _LsdNetworkFactory.Contract.ProxiableUUID(&_LsdNetworkFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LsdNetworkFactory *LsdNetworkFactoryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LsdNetworkFactory.Contract.ProxiableUUID(&_LsdNetworkFactory.CallOpts)
}

// UserDepositLogicAddress is a free data retrieval call binding the contract method 0x30d7338a.
//
// Solidity: function userDepositLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCaller) UserDepositLogicAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LsdNetworkFactory.contract.Call(opts, &out, "userDepositLogicAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserDepositLogicAddress is a free data retrieval call binding the contract method 0x30d7338a.
//
// Solidity: function userDepositLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactorySession) UserDepositLogicAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.UserDepositLogicAddress(&_LsdNetworkFactory.CallOpts)
}

// UserDepositLogicAddress is a free data retrieval call binding the contract method 0x30d7338a.
//
// Solidity: function userDepositLogicAddress() view returns(address)
func (_LsdNetworkFactory *LsdNetworkFactoryCallerSession) UserDepositLogicAddress() (common.Address, error) {
	return _LsdNetworkFactory.Contract.UserDepositLogicAddress(&_LsdNetworkFactory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_LsdNetworkFactory *LsdNetworkFactoryCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LsdNetworkFactory.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_LsdNetworkFactory *LsdNetworkFactorySession) Version() (uint8, error) {
	return _LsdNetworkFactory.Contract.Version(&_LsdNetworkFactory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_LsdNetworkFactory *LsdNetworkFactoryCallerSession) Version() (uint8, error) {
	return _LsdNetworkFactory.Contract.Version(&_LsdNetworkFactory.CallOpts)
}

// CreateLsdNetwork is a paid mutator transaction binding the contract method 0xb0bc87e7.
//
// Solidity: function createLsdNetwork(string _lsdTokenName, string _lsdTokenSymbol, address _networkAdmin, address[] _voters, uint256 _threshold) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) CreateLsdNetwork(opts *bind.TransactOpts, _lsdTokenName string, _lsdTokenSymbol string, _networkAdmin common.Address, _voters []common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.Transact(opts, "createLsdNetwork", _lsdTokenName, _lsdTokenSymbol, _networkAdmin, _voters, _threshold)
}

// CreateLsdNetwork is a paid mutator transaction binding the contract method 0xb0bc87e7.
//
// Solidity: function createLsdNetwork(string _lsdTokenName, string _lsdTokenSymbol, address _networkAdmin, address[] _voters, uint256 _threshold) returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) CreateLsdNetwork(_lsdTokenName string, _lsdTokenSymbol string, _networkAdmin common.Address, _voters []common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.CreateLsdNetwork(&_LsdNetworkFactory.TransactOpts, _lsdTokenName, _lsdTokenSymbol, _networkAdmin, _voters, _threshold)
}

// CreateLsdNetwork is a paid mutator transaction binding the contract method 0xb0bc87e7.
//
// Solidity: function createLsdNetwork(string _lsdTokenName, string _lsdTokenSymbol, address _networkAdmin, address[] _voters, uint256 _threshold) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) CreateLsdNetwork(_lsdTokenName string, _lsdTokenSymbol string, _networkAdmin common.Address, _voters []common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.CreateLsdNetwork(&_LsdNetworkFactory.TransactOpts, _lsdTokenName, _lsdTokenSymbol, _networkAdmin, _voters, _threshold)
}

// CreateLsdNetworkWithTimelock is a paid mutator transaction binding the contract method 0x6ec78cc5.
//
// Solidity: function createLsdNetworkWithTimelock(string _lsdTokenName, string _lsdTokenSymbol, address[] _voters, uint256 _threshold, uint256 minDelay, address[] proposers) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) CreateLsdNetworkWithTimelock(opts *bind.TransactOpts, _lsdTokenName string, _lsdTokenSymbol string, _voters []common.Address, _threshold *big.Int, minDelay *big.Int, proposers []common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.Transact(opts, "createLsdNetworkWithTimelock", _lsdTokenName, _lsdTokenSymbol, _voters, _threshold, minDelay, proposers)
}

// CreateLsdNetworkWithTimelock is a paid mutator transaction binding the contract method 0x6ec78cc5.
//
// Solidity: function createLsdNetworkWithTimelock(string _lsdTokenName, string _lsdTokenSymbol, address[] _voters, uint256 _threshold, uint256 minDelay, address[] proposers) returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) CreateLsdNetworkWithTimelock(_lsdTokenName string, _lsdTokenSymbol string, _voters []common.Address, _threshold *big.Int, minDelay *big.Int, proposers []common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.CreateLsdNetworkWithTimelock(&_LsdNetworkFactory.TransactOpts, _lsdTokenName, _lsdTokenSymbol, _voters, _threshold, minDelay, proposers)
}

// CreateLsdNetworkWithTimelock is a paid mutator transaction binding the contract method 0x6ec78cc5.
//
// Solidity: function createLsdNetworkWithTimelock(string _lsdTokenName, string _lsdTokenSymbol, address[] _voters, uint256 _threshold, uint256 minDelay, address[] proposers) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) CreateLsdNetworkWithTimelock(_lsdTokenName string, _lsdTokenSymbol string, _voters []common.Address, _threshold *big.Int, minDelay *big.Int, proposers []common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.CreateLsdNetworkWithTimelock(&_LsdNetworkFactory.TransactOpts, _lsdTokenName, _lsdTokenSymbol, _voters, _threshold, minDelay, proposers)
}

// FactoryClaim is a paid mutator transaction binding the contract method 0x531998ea.
//
// Solidity: function factoryClaim(address _recipient) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) FactoryClaim(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.Transact(opts, "factoryClaim", _recipient)
}

// FactoryClaim is a paid mutator transaction binding the contract method 0x531998ea.
//
// Solidity: function factoryClaim(address _recipient) returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) FactoryClaim(_recipient common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.FactoryClaim(&_LsdNetworkFactory.TransactOpts, _recipient)
}

// FactoryClaim is a paid mutator transaction binding the contract method 0x531998ea.
//
// Solidity: function factoryClaim(address _recipient) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) FactoryClaim(_recipient common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.FactoryClaim(&_LsdNetworkFactory.TransactOpts, _recipient)
}

// Init is a paid mutator transaction binding the contract method 0x525240c0.
//
// Solidity: function init(address _factoryAdmin, address _ethDepositAddress, address _feePoolLogicAddress, address _networkBalancesLogicAddress, address _networkProposalLogicAddress, address _nodeDepositLogicAddress, address _userDepositLogicAddress, address _networkWithdrawLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) Init(opts *bind.TransactOpts, _factoryAdmin common.Address, _ethDepositAddress common.Address, _feePoolLogicAddress common.Address, _networkBalancesLogicAddress common.Address, _networkProposalLogicAddress common.Address, _nodeDepositLogicAddress common.Address, _userDepositLogicAddress common.Address, _networkWithdrawLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.Transact(opts, "init", _factoryAdmin, _ethDepositAddress, _feePoolLogicAddress, _networkBalancesLogicAddress, _networkProposalLogicAddress, _nodeDepositLogicAddress, _userDepositLogicAddress, _networkWithdrawLogicAddress)
}

// Init is a paid mutator transaction binding the contract method 0x525240c0.
//
// Solidity: function init(address _factoryAdmin, address _ethDepositAddress, address _feePoolLogicAddress, address _networkBalancesLogicAddress, address _networkProposalLogicAddress, address _nodeDepositLogicAddress, address _userDepositLogicAddress, address _networkWithdrawLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) Init(_factoryAdmin common.Address, _ethDepositAddress common.Address, _feePoolLogicAddress common.Address, _networkBalancesLogicAddress common.Address, _networkProposalLogicAddress common.Address, _nodeDepositLogicAddress common.Address, _userDepositLogicAddress common.Address, _networkWithdrawLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.Init(&_LsdNetworkFactory.TransactOpts, _factoryAdmin, _ethDepositAddress, _feePoolLogicAddress, _networkBalancesLogicAddress, _networkProposalLogicAddress, _nodeDepositLogicAddress, _userDepositLogicAddress, _networkWithdrawLogicAddress)
}

// Init is a paid mutator transaction binding the contract method 0x525240c0.
//
// Solidity: function init(address _factoryAdmin, address _ethDepositAddress, address _feePoolLogicAddress, address _networkBalancesLogicAddress, address _networkProposalLogicAddress, address _nodeDepositLogicAddress, address _userDepositLogicAddress, address _networkWithdrawLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) Init(_factoryAdmin common.Address, _ethDepositAddress common.Address, _feePoolLogicAddress common.Address, _networkBalancesLogicAddress common.Address, _networkProposalLogicAddress common.Address, _nodeDepositLogicAddress common.Address, _userDepositLogicAddress common.Address, _networkWithdrawLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.Init(&_LsdNetworkFactory.TransactOpts, _factoryAdmin, _ethDepositAddress, _feePoolLogicAddress, _networkBalancesLogicAddress, _networkProposalLogicAddress, _nodeDepositLogicAddress, _userDepositLogicAddress, _networkWithdrawLogicAddress)
}

// Reinit is a paid mutator transaction binding the contract method 0xc482ceaf.
//
// Solidity: function reinit() returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) Reinit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.Transact(opts, "reinit")
}

// Reinit is a paid mutator transaction binding the contract method 0xc482ceaf.
//
// Solidity: function reinit() returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) Reinit() (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.Reinit(&_LsdNetworkFactory.TransactOpts)
}

// Reinit is a paid mutator transaction binding the contract method 0xc482ceaf.
//
// Solidity: function reinit() returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) Reinit() (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.Reinit(&_LsdNetworkFactory.TransactOpts)
}

// SetNetworkBalancesLogicAddress is a paid mutator transaction binding the contract method 0xc083f1c2.
//
// Solidity: function setNetworkBalancesLogicAddress(address _networkBalancesLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) SetNetworkBalancesLogicAddress(opts *bind.TransactOpts, _networkBalancesLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.Transact(opts, "setNetworkBalancesLogicAddress", _networkBalancesLogicAddress)
}

// SetNetworkBalancesLogicAddress is a paid mutator transaction binding the contract method 0xc083f1c2.
//
// Solidity: function setNetworkBalancesLogicAddress(address _networkBalancesLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) SetNetworkBalancesLogicAddress(_networkBalancesLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.SetNetworkBalancesLogicAddress(&_LsdNetworkFactory.TransactOpts, _networkBalancesLogicAddress)
}

// SetNetworkBalancesLogicAddress is a paid mutator transaction binding the contract method 0xc083f1c2.
//
// Solidity: function setNetworkBalancesLogicAddress(address _networkBalancesLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) SetNetworkBalancesLogicAddress(_networkBalancesLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.SetNetworkBalancesLogicAddress(&_LsdNetworkFactory.TransactOpts, _networkBalancesLogicAddress)
}

// SetNetworkProposalLogicAddress is a paid mutator transaction binding the contract method 0x38b952fa.
//
// Solidity: function setNetworkProposalLogicAddress(address _networkProposalLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) SetNetworkProposalLogicAddress(opts *bind.TransactOpts, _networkProposalLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.Transact(opts, "setNetworkProposalLogicAddress", _networkProposalLogicAddress)
}

// SetNetworkProposalLogicAddress is a paid mutator transaction binding the contract method 0x38b952fa.
//
// Solidity: function setNetworkProposalLogicAddress(address _networkProposalLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) SetNetworkProposalLogicAddress(_networkProposalLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.SetNetworkProposalLogicAddress(&_LsdNetworkFactory.TransactOpts, _networkProposalLogicAddress)
}

// SetNetworkProposalLogicAddress is a paid mutator transaction binding the contract method 0x38b952fa.
//
// Solidity: function setNetworkProposalLogicAddress(address _networkProposalLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) SetNetworkProposalLogicAddress(_networkProposalLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.SetNetworkProposalLogicAddress(&_LsdNetworkFactory.TransactOpts, _networkProposalLogicAddress)
}

// SetNetworkWithdrawLogicAddress is a paid mutator transaction binding the contract method 0x6dd611a3.
//
// Solidity: function setNetworkWithdrawLogicAddress(address _networkWithdrawLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) SetNetworkWithdrawLogicAddress(opts *bind.TransactOpts, _networkWithdrawLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.Transact(opts, "setNetworkWithdrawLogicAddress", _networkWithdrawLogicAddress)
}

// SetNetworkWithdrawLogicAddress is a paid mutator transaction binding the contract method 0x6dd611a3.
//
// Solidity: function setNetworkWithdrawLogicAddress(address _networkWithdrawLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) SetNetworkWithdrawLogicAddress(_networkWithdrawLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.SetNetworkWithdrawLogicAddress(&_LsdNetworkFactory.TransactOpts, _networkWithdrawLogicAddress)
}

// SetNetworkWithdrawLogicAddress is a paid mutator transaction binding the contract method 0x6dd611a3.
//
// Solidity: function setNetworkWithdrawLogicAddress(address _networkWithdrawLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) SetNetworkWithdrawLogicAddress(_networkWithdrawLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.SetNetworkWithdrawLogicAddress(&_LsdNetworkFactory.TransactOpts, _networkWithdrawLogicAddress)
}

// SetNodeDepositLogicAddress is a paid mutator transaction binding the contract method 0x920f8f8f.
//
// Solidity: function setNodeDepositLogicAddress(address _nodeDepositLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) SetNodeDepositLogicAddress(opts *bind.TransactOpts, _nodeDepositLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.Transact(opts, "setNodeDepositLogicAddress", _nodeDepositLogicAddress)
}

// SetNodeDepositLogicAddress is a paid mutator transaction binding the contract method 0x920f8f8f.
//
// Solidity: function setNodeDepositLogicAddress(address _nodeDepositLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) SetNodeDepositLogicAddress(_nodeDepositLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.SetNodeDepositLogicAddress(&_LsdNetworkFactory.TransactOpts, _nodeDepositLogicAddress)
}

// SetNodeDepositLogicAddress is a paid mutator transaction binding the contract method 0x920f8f8f.
//
// Solidity: function setNodeDepositLogicAddress(address _nodeDepositLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) SetNodeDepositLogicAddress(_nodeDepositLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.SetNodeDepositLogicAddress(&_LsdNetworkFactory.TransactOpts, _nodeDepositLogicAddress)
}

// SetUserDepositLogicAddress is a paid mutator transaction binding the contract method 0x100f740d.
//
// Solidity: function setUserDepositLogicAddress(address _userDepositLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) SetUserDepositLogicAddress(opts *bind.TransactOpts, _userDepositLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.Transact(opts, "setUserDepositLogicAddress", _userDepositLogicAddress)
}

// SetUserDepositLogicAddress is a paid mutator transaction binding the contract method 0x100f740d.
//
// Solidity: function setUserDepositLogicAddress(address _userDepositLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) SetUserDepositLogicAddress(_userDepositLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.SetUserDepositLogicAddress(&_LsdNetworkFactory.TransactOpts, _userDepositLogicAddress)
}

// SetUserDepositLogicAddress is a paid mutator transaction binding the contract method 0x100f740d.
//
// Solidity: function setUserDepositLogicAddress(address _userDepositLogicAddress) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) SetUserDepositLogicAddress(_userDepositLogicAddress common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.SetUserDepositLogicAddress(&_LsdNetworkFactory.TransactOpts, _userDepositLogicAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newAdmin) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.Transact(opts, "transferOwnership", _newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newAdmin) returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) TransferOwnership(_newAdmin common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.TransferOwnership(&_LsdNetworkFactory.TransactOpts, _newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newAdmin) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) TransferOwnership(_newAdmin common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.TransferOwnership(&_LsdNetworkFactory.TransactOpts, _newAdmin)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.UpgradeTo(&_LsdNetworkFactory.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.UpgradeTo(&_LsdNetworkFactory.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.UpgradeToAndCall(&_LsdNetworkFactory.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.UpgradeToAndCall(&_LsdNetworkFactory.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LsdNetworkFactory.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LsdNetworkFactory *LsdNetworkFactorySession) Receive() (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.Receive(&_LsdNetworkFactory.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LsdNetworkFactory *LsdNetworkFactoryTransactorSession) Receive() (*types.Transaction, error) {
	return _LsdNetworkFactory.Contract.Receive(&_LsdNetworkFactory.TransactOpts)
}

// LsdNetworkFactoryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the LsdNetworkFactory contract.
type LsdNetworkFactoryAdminChangedIterator struct {
	Event *LsdNetworkFactoryAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LsdNetworkFactoryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LsdNetworkFactoryAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LsdNetworkFactoryAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LsdNetworkFactoryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LsdNetworkFactoryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LsdNetworkFactoryAdminChanged represents a AdminChanged event raised by the LsdNetworkFactory contract.
type LsdNetworkFactoryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*LsdNetworkFactoryAdminChangedIterator, error) {

	logs, sub, err := _LsdNetworkFactory.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &LsdNetworkFactoryAdminChangedIterator{contract: _LsdNetworkFactory.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LsdNetworkFactoryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _LsdNetworkFactory.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LsdNetworkFactoryAdminChanged)
				if err := _LsdNetworkFactory.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) ParseAdminChanged(log types.Log) (*LsdNetworkFactoryAdminChanged, error) {
	event := new(LsdNetworkFactoryAdminChanged)
	if err := _LsdNetworkFactory.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LsdNetworkFactoryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the LsdNetworkFactory contract.
type LsdNetworkFactoryBeaconUpgradedIterator struct {
	Event *LsdNetworkFactoryBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LsdNetworkFactoryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LsdNetworkFactoryBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LsdNetworkFactoryBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LsdNetworkFactoryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LsdNetworkFactoryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LsdNetworkFactoryBeaconUpgraded represents a BeaconUpgraded event raised by the LsdNetworkFactory contract.
type LsdNetworkFactoryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*LsdNetworkFactoryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LsdNetworkFactory.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &LsdNetworkFactoryBeaconUpgradedIterator{contract: _LsdNetworkFactory.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *LsdNetworkFactoryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LsdNetworkFactory.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LsdNetworkFactoryBeaconUpgraded)
				if err := _LsdNetworkFactory.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) ParseBeaconUpgraded(log types.Log) (*LsdNetworkFactoryBeaconUpgraded, error) {
	event := new(LsdNetworkFactoryBeaconUpgraded)
	if err := _LsdNetworkFactory.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LsdNetworkFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LsdNetworkFactory contract.
type LsdNetworkFactoryInitializedIterator struct {
	Event *LsdNetworkFactoryInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LsdNetworkFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LsdNetworkFactoryInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LsdNetworkFactoryInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LsdNetworkFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LsdNetworkFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LsdNetworkFactoryInitialized represents a Initialized event raised by the LsdNetworkFactory contract.
type LsdNetworkFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*LsdNetworkFactoryInitializedIterator, error) {

	logs, sub, err := _LsdNetworkFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LsdNetworkFactoryInitializedIterator{contract: _LsdNetworkFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LsdNetworkFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _LsdNetworkFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LsdNetworkFactoryInitialized)
				if err := _LsdNetworkFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) ParseInitialized(log types.Log) (*LsdNetworkFactoryInitialized, error) {
	event := new(LsdNetworkFactoryInitialized)
	if err := _LsdNetworkFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LsdNetworkFactoryLsdNetworkIterator is returned from FilterLsdNetwork and is used to iterate over the raw logs and unpacked data for LsdNetwork events raised by the LsdNetworkFactory contract.
type LsdNetworkFactoryLsdNetworkIterator struct {
	Event *LsdNetworkFactoryLsdNetwork // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LsdNetworkFactoryLsdNetworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LsdNetworkFactoryLsdNetwork)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LsdNetworkFactoryLsdNetwork)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LsdNetworkFactoryLsdNetworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LsdNetworkFactoryLsdNetworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LsdNetworkFactoryLsdNetwork represents a LsdNetwork event raised by the LsdNetworkFactory contract.
type LsdNetworkFactoryLsdNetwork struct {
	Contracts ILsdNetworkFactoryNetworkContracts
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLsdNetwork is a free log retrieval operation binding the contract event 0xdaea250f2c303367bec9ee390644e69eb6d279d72f3b9fc79c0814649227cbd7.
//
// Solidity: event LsdNetwork((address,address,address,address,address,address,address,uint256) contracts)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) FilterLsdNetwork(opts *bind.FilterOpts) (*LsdNetworkFactoryLsdNetworkIterator, error) {

	logs, sub, err := _LsdNetworkFactory.contract.FilterLogs(opts, "LsdNetwork")
	if err != nil {
		return nil, err
	}
	return &LsdNetworkFactoryLsdNetworkIterator{contract: _LsdNetworkFactory.contract, event: "LsdNetwork", logs: logs, sub: sub}, nil
}

// WatchLsdNetwork is a free log subscription operation binding the contract event 0xdaea250f2c303367bec9ee390644e69eb6d279d72f3b9fc79c0814649227cbd7.
//
// Solidity: event LsdNetwork((address,address,address,address,address,address,address,uint256) contracts)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) WatchLsdNetwork(opts *bind.WatchOpts, sink chan<- *LsdNetworkFactoryLsdNetwork) (event.Subscription, error) {

	logs, sub, err := _LsdNetworkFactory.contract.WatchLogs(opts, "LsdNetwork")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LsdNetworkFactoryLsdNetwork)
				if err := _LsdNetworkFactory.contract.UnpackLog(event, "LsdNetwork", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLsdNetwork is a log parse operation binding the contract event 0xdaea250f2c303367bec9ee390644e69eb6d279d72f3b9fc79c0814649227cbd7.
//
// Solidity: event LsdNetwork((address,address,address,address,address,address,address,uint256) contracts)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) ParseLsdNetwork(log types.Log) (*LsdNetworkFactoryLsdNetwork, error) {
	event := new(LsdNetworkFactoryLsdNetwork)
	if err := _LsdNetworkFactory.contract.UnpackLog(event, "LsdNetwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LsdNetworkFactoryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LsdNetworkFactory contract.
type LsdNetworkFactoryUpgradedIterator struct {
	Event *LsdNetworkFactoryUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LsdNetworkFactoryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LsdNetworkFactoryUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LsdNetworkFactoryUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LsdNetworkFactoryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LsdNetworkFactoryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LsdNetworkFactoryUpgraded represents a Upgraded event raised by the LsdNetworkFactory contract.
type LsdNetworkFactoryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LsdNetworkFactoryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LsdNetworkFactory.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LsdNetworkFactoryUpgradedIterator{contract: _LsdNetworkFactory.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LsdNetworkFactoryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LsdNetworkFactory.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LsdNetworkFactoryUpgraded)
				if err := _LsdNetworkFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LsdNetworkFactory *LsdNetworkFactoryFilterer) ParseUpgraded(log types.Log) (*LsdNetworkFactoryUpgraded, error) {
	event := new(LsdNetworkFactoryUpgraded)
	if err := _LsdNetworkFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
