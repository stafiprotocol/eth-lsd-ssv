// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package network_withdraw

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

// NetworkWithdrawMetaData contains all meta data concerning the NetworkWithdraw contract.
var NetworkWithdrawMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumINetworkWithdraw.DistributeType\",\"name\":\"_distributeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_dealedHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_userAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nodeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_platformAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_maxClaimableWithdrawIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_mvAmount\",\"type\":\"uint256\"}],\"name\":\"DistributeRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"EtherDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_claimableReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_claimableDeposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumINetworkWithdraw.ClaimType\",\"name\":\"_claimType\",\"type\":\"uint8\"}],\"name\":\"NodeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_withdrawCycle\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_ejectedStartWithdrawCycle\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_ejectedValidators\",\"type\":\"uint256[]\"}],\"name\":\"NotifyValidatorExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_dealedEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_nodeRewardsFileCid\",\"type\":\"string\"}],\"name\":\"SetMerkleRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_userWithdrawLimitPerCycle\",\"type\":\"uint256\"}],\"name\":\"SetUserWithdrawLimitPerCycle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_seconds\",\"type\":\"uint256\"}],\"name\":\"SetWithdrawCycleSeconds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_withdrawLimitPerCycle\",\"type\":\"uint256\"}],\"name\":\"SetWithdrawLimitPerCycle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lsdTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_withdrawIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_instantly\",\"type\":\"bool\"}],\"name\":\"Unstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_withdrawIndexList\",\"type\":\"uint256[]\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentWithdrawCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEthAndUpdateTotalMissingAmount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumINetworkWithdraw.DistributeType\",\"name\":\"_distributeType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_dealedHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_platformAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxClaimableWithdrawIndex\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ejectedStartCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ejectedValidatorsAtCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryCommissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"getEjectedValidatorsAtCycle\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUnclaimedWithdrawalsOfUser\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lsdTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_userDepositAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkProposalAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feePoolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factoryAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestDistributePriorityFeeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestDistributeWithdrawalsHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestMerkleRootEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lsdTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxClaimableWithdrawIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkBalancesAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkProposalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextWithdrawIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalExitDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"enumINetworkWithdraw.ClaimType\",\"name\":\"_claimType\",\"type\":\"uint8\"}],\"name\":\"nodeClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeRewardsFileCid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ejectedStartCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_validatorIndexList\",\"type\":\"uint256[]\"}],\"name\":\"notifyValidatorExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"platformClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformCommissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dealedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_nodeRewardsFileCid\",\"type\":\"string\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_userWithdrawLimitPerCycle\",\"type\":\"uint256\"}],\"name\":\"setUserWithdrawLimitAmountPerCycle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawCycleSeconds\",\"type\":\"uint256\"}],\"name\":\"setWithdrawCycleSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawLimitPerCycle\",\"type\":\"uint256\"}],\"name\":\"setWithdrawLimitAmountPerCycle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalClaimedDepositOfNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalClaimedRewardOfNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMissingAmountForWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPlatformClaimedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPlatformCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalWithdrawAmountAtCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lsdTokenAmount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"updateMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDepositAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userWithdrawAmountAtCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userWithdrawLimitAmountPerCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_withdrawIndexList\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawCycleSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLimitAmountPerCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// NetworkWithdrawABI is the input ABI used to generate the binding from.
// Deprecated: Use NetworkWithdrawMetaData.ABI instead.
var NetworkWithdrawABI = NetworkWithdrawMetaData.ABI

// NetworkWithdraw is an auto generated Go binding around an Ethereum contract.
type NetworkWithdraw struct {
	NetworkWithdrawCaller     // Read-only binding to the contract
	NetworkWithdrawTransactor // Write-only binding to the contract
	NetworkWithdrawFilterer   // Log filterer for contract events
}

// NetworkWithdrawCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkWithdrawCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkWithdrawTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkWithdrawTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkWithdrawFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NetworkWithdrawFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkWithdrawSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkWithdrawSession struct {
	Contract     *NetworkWithdraw  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NetworkWithdrawCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkWithdrawCallerSession struct {
	Contract *NetworkWithdrawCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NetworkWithdrawTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkWithdrawTransactorSession struct {
	Contract     *NetworkWithdrawTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NetworkWithdrawRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkWithdrawRaw struct {
	Contract *NetworkWithdraw // Generic contract binding to access the raw methods on
}

// NetworkWithdrawCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkWithdrawCallerRaw struct {
	Contract *NetworkWithdrawCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkWithdrawTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkWithdrawTransactorRaw struct {
	Contract *NetworkWithdrawTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetworkWithdraw creates a new instance of NetworkWithdraw, bound to a specific deployed contract.
func NewNetworkWithdraw(address common.Address, backend bind.ContractBackend) (*NetworkWithdraw, error) {
	contract, err := bindNetworkWithdraw(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NetworkWithdraw{NetworkWithdrawCaller: NetworkWithdrawCaller{contract: contract}, NetworkWithdrawTransactor: NetworkWithdrawTransactor{contract: contract}, NetworkWithdrawFilterer: NetworkWithdrawFilterer{contract: contract}}, nil
}

// NewNetworkWithdrawCaller creates a new read-only instance of NetworkWithdraw, bound to a specific deployed contract.
func NewNetworkWithdrawCaller(address common.Address, caller bind.ContractCaller) (*NetworkWithdrawCaller, error) {
	contract, err := bindNetworkWithdraw(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkWithdrawCaller{contract: contract}, nil
}

// NewNetworkWithdrawTransactor creates a new write-only instance of NetworkWithdraw, bound to a specific deployed contract.
func NewNetworkWithdrawTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkWithdrawTransactor, error) {
	contract, err := bindNetworkWithdraw(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkWithdrawTransactor{contract: contract}, nil
}

// NewNetworkWithdrawFilterer creates a new log filterer instance of NetworkWithdraw, bound to a specific deployed contract.
func NewNetworkWithdrawFilterer(address common.Address, filterer bind.ContractFilterer) (*NetworkWithdrawFilterer, error) {
	contract, err := bindNetworkWithdraw(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NetworkWithdrawFilterer{contract: contract}, nil
}

// bindNetworkWithdraw binds a generic wrapper to an already deployed contract.
func bindNetworkWithdraw(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkWithdrawABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkWithdraw *NetworkWithdrawRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkWithdraw.Contract.NetworkWithdrawCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkWithdraw *NetworkWithdrawRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.NetworkWithdrawTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkWithdraw *NetworkWithdrawRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.NetworkWithdrawTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkWithdraw *NetworkWithdrawCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkWithdraw.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkWithdraw *NetworkWithdrawTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkWithdraw *NetworkWithdrawTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.contract.Transact(opts, method, params...)
}

// CurrentWithdrawCycle is a free data retrieval call binding the contract method 0xdb17815b.
//
// Solidity: function currentWithdrawCycle() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) CurrentWithdrawCycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "currentWithdrawCycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentWithdrawCycle is a free data retrieval call binding the contract method 0xdb17815b.
//
// Solidity: function currentWithdrawCycle() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) CurrentWithdrawCycle() (*big.Int, error) {
	return _NetworkWithdraw.Contract.CurrentWithdrawCycle(&_NetworkWithdraw.CallOpts)
}

// CurrentWithdrawCycle is a free data retrieval call binding the contract method 0xdb17815b.
//
// Solidity: function currentWithdrawCycle() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) CurrentWithdrawCycle() (*big.Int, error) {
	return _NetworkWithdraw.Contract.CurrentWithdrawCycle(&_NetworkWithdraw.CallOpts)
}

// EjectedStartCycle is a free data retrieval call binding the contract method 0x8a699828.
//
// Solidity: function ejectedStartCycle() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) EjectedStartCycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "ejectedStartCycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EjectedStartCycle is a free data retrieval call binding the contract method 0x8a699828.
//
// Solidity: function ejectedStartCycle() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) EjectedStartCycle() (*big.Int, error) {
	return _NetworkWithdraw.Contract.EjectedStartCycle(&_NetworkWithdraw.CallOpts)
}

// EjectedStartCycle is a free data retrieval call binding the contract method 0x8a699828.
//
// Solidity: function ejectedStartCycle() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) EjectedStartCycle() (*big.Int, error) {
	return _NetworkWithdraw.Contract.EjectedStartCycle(&_NetworkWithdraw.CallOpts)
}

// EjectedValidatorsAtCycle is a free data retrieval call binding the contract method 0x261a792d.
//
// Solidity: function ejectedValidatorsAtCycle(uint256 , uint256 ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) EjectedValidatorsAtCycle(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "ejectedValidatorsAtCycle", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EjectedValidatorsAtCycle is a free data retrieval call binding the contract method 0x261a792d.
//
// Solidity: function ejectedValidatorsAtCycle(uint256 , uint256 ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) EjectedValidatorsAtCycle(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _NetworkWithdraw.Contract.EjectedValidatorsAtCycle(&_NetworkWithdraw.CallOpts, arg0, arg1)
}

// EjectedValidatorsAtCycle is a free data retrieval call binding the contract method 0x261a792d.
//
// Solidity: function ejectedValidatorsAtCycle(uint256 , uint256 ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) EjectedValidatorsAtCycle(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _NetworkWithdraw.Contract.EjectedValidatorsAtCycle(&_NetworkWithdraw.CallOpts, arg0, arg1)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawCaller) FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "factoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawSession) FactoryAddress() (common.Address, error) {
	return _NetworkWithdraw.Contract.FactoryAddress(&_NetworkWithdraw.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) FactoryAddress() (common.Address, error) {
	return _NetworkWithdraw.Contract.FactoryAddress(&_NetworkWithdraw.CallOpts)
}

// FactoryCommissionRate is a free data retrieval call binding the contract method 0xc2156b4b.
//
// Solidity: function factoryCommissionRate() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) FactoryCommissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "factoryCommissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FactoryCommissionRate is a free data retrieval call binding the contract method 0xc2156b4b.
//
// Solidity: function factoryCommissionRate() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) FactoryCommissionRate() (*big.Int, error) {
	return _NetworkWithdraw.Contract.FactoryCommissionRate(&_NetworkWithdraw.CallOpts)
}

// FactoryCommissionRate is a free data retrieval call binding the contract method 0xc2156b4b.
//
// Solidity: function factoryCommissionRate() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) FactoryCommissionRate() (*big.Int, error) {
	return _NetworkWithdraw.Contract.FactoryCommissionRate(&_NetworkWithdraw.CallOpts)
}

// FeePoolAddress is a free data retrieval call binding the contract method 0x4319ebe4.
//
// Solidity: function feePoolAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawCaller) FeePoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "feePoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeePoolAddress is a free data retrieval call binding the contract method 0x4319ebe4.
//
// Solidity: function feePoolAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawSession) FeePoolAddress() (common.Address, error) {
	return _NetworkWithdraw.Contract.FeePoolAddress(&_NetworkWithdraw.CallOpts)
}

// FeePoolAddress is a free data retrieval call binding the contract method 0x4319ebe4.
//
// Solidity: function feePoolAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) FeePoolAddress() (common.Address, error) {
	return _NetworkWithdraw.Contract.FeePoolAddress(&_NetworkWithdraw.CallOpts)
}

// GetEjectedValidatorsAtCycle is a free data retrieval call binding the contract method 0x2c0f4166.
//
// Solidity: function getEjectedValidatorsAtCycle(uint256 cycle) view returns(uint256[])
func (_NetworkWithdraw *NetworkWithdrawCaller) GetEjectedValidatorsAtCycle(opts *bind.CallOpts, cycle *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "getEjectedValidatorsAtCycle", cycle)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetEjectedValidatorsAtCycle is a free data retrieval call binding the contract method 0x2c0f4166.
//
// Solidity: function getEjectedValidatorsAtCycle(uint256 cycle) view returns(uint256[])
func (_NetworkWithdraw *NetworkWithdrawSession) GetEjectedValidatorsAtCycle(cycle *big.Int) ([]*big.Int, error) {
	return _NetworkWithdraw.Contract.GetEjectedValidatorsAtCycle(&_NetworkWithdraw.CallOpts, cycle)
}

// GetEjectedValidatorsAtCycle is a free data retrieval call binding the contract method 0x2c0f4166.
//
// Solidity: function getEjectedValidatorsAtCycle(uint256 cycle) view returns(uint256[])
func (_NetworkWithdraw *NetworkWithdrawCallerSession) GetEjectedValidatorsAtCycle(cycle *big.Int) ([]*big.Int, error) {
	return _NetworkWithdraw.Contract.GetEjectedValidatorsAtCycle(&_NetworkWithdraw.CallOpts, cycle)
}

// GetUnclaimedWithdrawalsOfUser is a free data retrieval call binding the contract method 0xfd6b5a49.
//
// Solidity: function getUnclaimedWithdrawalsOfUser(address user) view returns(uint256[])
func (_NetworkWithdraw *NetworkWithdrawCaller) GetUnclaimedWithdrawalsOfUser(opts *bind.CallOpts, user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "getUnclaimedWithdrawalsOfUser", user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUnclaimedWithdrawalsOfUser is a free data retrieval call binding the contract method 0xfd6b5a49.
//
// Solidity: function getUnclaimedWithdrawalsOfUser(address user) view returns(uint256[])
func (_NetworkWithdraw *NetworkWithdrawSession) GetUnclaimedWithdrawalsOfUser(user common.Address) ([]*big.Int, error) {
	return _NetworkWithdraw.Contract.GetUnclaimedWithdrawalsOfUser(&_NetworkWithdraw.CallOpts, user)
}

// GetUnclaimedWithdrawalsOfUser is a free data retrieval call binding the contract method 0xfd6b5a49.
//
// Solidity: function getUnclaimedWithdrawalsOfUser(address user) view returns(uint256[])
func (_NetworkWithdraw *NetworkWithdrawCallerSession) GetUnclaimedWithdrawalsOfUser(user common.Address) ([]*big.Int, error) {
	return _NetworkWithdraw.Contract.GetUnclaimedWithdrawalsOfUser(&_NetworkWithdraw.CallOpts, user)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_NetworkWithdraw *NetworkWithdrawCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_NetworkWithdraw *NetworkWithdrawSession) Initialized() (bool, error) {
	return _NetworkWithdraw.Contract.Initialized(&_NetworkWithdraw.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) Initialized() (bool, error) {
	return _NetworkWithdraw.Contract.Initialized(&_NetworkWithdraw.CallOpts)
}

// LatestDistributePriorityFeeHeight is a free data retrieval call binding the contract method 0x4dff8430.
//
// Solidity: function latestDistributePriorityFeeHeight() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) LatestDistributePriorityFeeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "latestDistributePriorityFeeHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestDistributePriorityFeeHeight is a free data retrieval call binding the contract method 0x4dff8430.
//
// Solidity: function latestDistributePriorityFeeHeight() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) LatestDistributePriorityFeeHeight() (*big.Int, error) {
	return _NetworkWithdraw.Contract.LatestDistributePriorityFeeHeight(&_NetworkWithdraw.CallOpts)
}

// LatestDistributePriorityFeeHeight is a free data retrieval call binding the contract method 0x4dff8430.
//
// Solidity: function latestDistributePriorityFeeHeight() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) LatestDistributePriorityFeeHeight() (*big.Int, error) {
	return _NetworkWithdraw.Contract.LatestDistributePriorityFeeHeight(&_NetworkWithdraw.CallOpts)
}

// LatestDistributeWithdrawalsHeight is a free data retrieval call binding the contract method 0x9fa1f5ba.
//
// Solidity: function latestDistributeWithdrawalsHeight() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) LatestDistributeWithdrawalsHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "latestDistributeWithdrawalsHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestDistributeWithdrawalsHeight is a free data retrieval call binding the contract method 0x9fa1f5ba.
//
// Solidity: function latestDistributeWithdrawalsHeight() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) LatestDistributeWithdrawalsHeight() (*big.Int, error) {
	return _NetworkWithdraw.Contract.LatestDistributeWithdrawalsHeight(&_NetworkWithdraw.CallOpts)
}

// LatestDistributeWithdrawalsHeight is a free data retrieval call binding the contract method 0x9fa1f5ba.
//
// Solidity: function latestDistributeWithdrawalsHeight() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) LatestDistributeWithdrawalsHeight() (*big.Int, error) {
	return _NetworkWithdraw.Contract.LatestDistributeWithdrawalsHeight(&_NetworkWithdraw.CallOpts)
}

// LatestMerkleRootEpoch is a free data retrieval call binding the contract method 0xb5ca7410.
//
// Solidity: function latestMerkleRootEpoch() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) LatestMerkleRootEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "latestMerkleRootEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestMerkleRootEpoch is a free data retrieval call binding the contract method 0xb5ca7410.
//
// Solidity: function latestMerkleRootEpoch() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) LatestMerkleRootEpoch() (*big.Int, error) {
	return _NetworkWithdraw.Contract.LatestMerkleRootEpoch(&_NetworkWithdraw.CallOpts)
}

// LatestMerkleRootEpoch is a free data retrieval call binding the contract method 0xb5ca7410.
//
// Solidity: function latestMerkleRootEpoch() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) LatestMerkleRootEpoch() (*big.Int, error) {
	return _NetworkWithdraw.Contract.LatestMerkleRootEpoch(&_NetworkWithdraw.CallOpts)
}

// LsdTokenAddress is a free data retrieval call binding the contract method 0x87505b9d.
//
// Solidity: function lsdTokenAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawCaller) LsdTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "lsdTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LsdTokenAddress is a free data retrieval call binding the contract method 0x87505b9d.
//
// Solidity: function lsdTokenAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawSession) LsdTokenAddress() (common.Address, error) {
	return _NetworkWithdraw.Contract.LsdTokenAddress(&_NetworkWithdraw.CallOpts)
}

// LsdTokenAddress is a free data retrieval call binding the contract method 0x87505b9d.
//
// Solidity: function lsdTokenAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) LsdTokenAddress() (common.Address, error) {
	return _NetworkWithdraw.Contract.LsdTokenAddress(&_NetworkWithdraw.CallOpts)
}

// MaxClaimableWithdrawIndex is a free data retrieval call binding the contract method 0x0a64041b.
//
// Solidity: function maxClaimableWithdrawIndex() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) MaxClaimableWithdrawIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "maxClaimableWithdrawIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxClaimableWithdrawIndex is a free data retrieval call binding the contract method 0x0a64041b.
//
// Solidity: function maxClaimableWithdrawIndex() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) MaxClaimableWithdrawIndex() (*big.Int, error) {
	return _NetworkWithdraw.Contract.MaxClaimableWithdrawIndex(&_NetworkWithdraw.CallOpts)
}

// MaxClaimableWithdrawIndex is a free data retrieval call binding the contract method 0x0a64041b.
//
// Solidity: function maxClaimableWithdrawIndex() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) MaxClaimableWithdrawIndex() (*big.Int, error) {
	return _NetworkWithdraw.Contract.MaxClaimableWithdrawIndex(&_NetworkWithdraw.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_NetworkWithdraw *NetworkWithdrawCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_NetworkWithdraw *NetworkWithdrawSession) MerkleRoot() ([32]byte, error) {
	return _NetworkWithdraw.Contract.MerkleRoot(&_NetworkWithdraw.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) MerkleRoot() ([32]byte, error) {
	return _NetworkWithdraw.Contract.MerkleRoot(&_NetworkWithdraw.CallOpts)
}

// NetworkBalancesAddress is a free data retrieval call binding the contract method 0x38fcf092.
//
// Solidity: function networkBalancesAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawCaller) NetworkBalancesAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "networkBalancesAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkBalancesAddress is a free data retrieval call binding the contract method 0x38fcf092.
//
// Solidity: function networkBalancesAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawSession) NetworkBalancesAddress() (common.Address, error) {
	return _NetworkWithdraw.Contract.NetworkBalancesAddress(&_NetworkWithdraw.CallOpts)
}

// NetworkBalancesAddress is a free data retrieval call binding the contract method 0x38fcf092.
//
// Solidity: function networkBalancesAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) NetworkBalancesAddress() (common.Address, error) {
	return _NetworkWithdraw.Contract.NetworkBalancesAddress(&_NetworkWithdraw.CallOpts)
}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawCaller) NetworkProposalAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "networkProposalAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawSession) NetworkProposalAddress() (common.Address, error) {
	return _NetworkWithdraw.Contract.NetworkProposalAddress(&_NetworkWithdraw.CallOpts)
}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) NetworkProposalAddress() (common.Address, error) {
	return _NetworkWithdraw.Contract.NetworkProposalAddress(&_NetworkWithdraw.CallOpts)
}

// NextWithdrawIndex is a free data retrieval call binding the contract method 0x7e4dc15c.
//
// Solidity: function nextWithdrawIndex() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) NextWithdrawIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "nextWithdrawIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextWithdrawIndex is a free data retrieval call binding the contract method 0x7e4dc15c.
//
// Solidity: function nextWithdrawIndex() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) NextWithdrawIndex() (*big.Int, error) {
	return _NetworkWithdraw.Contract.NextWithdrawIndex(&_NetworkWithdraw.CallOpts)
}

// NextWithdrawIndex is a free data retrieval call binding the contract method 0x7e4dc15c.
//
// Solidity: function nextWithdrawIndex() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) NextWithdrawIndex() (*big.Int, error) {
	return _NetworkWithdraw.Contract.NextWithdrawIndex(&_NetworkWithdraw.CallOpts)
}

// NodeRewardsFileCid is a free data retrieval call binding the contract method 0xd57dc824.
//
// Solidity: function nodeRewardsFileCid() view returns(string)
func (_NetworkWithdraw *NetworkWithdrawCaller) NodeRewardsFileCid(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "nodeRewardsFileCid")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NodeRewardsFileCid is a free data retrieval call binding the contract method 0xd57dc824.
//
// Solidity: function nodeRewardsFileCid() view returns(string)
func (_NetworkWithdraw *NetworkWithdrawSession) NodeRewardsFileCid() (string, error) {
	return _NetworkWithdraw.Contract.NodeRewardsFileCid(&_NetworkWithdraw.CallOpts)
}

// NodeRewardsFileCid is a free data retrieval call binding the contract method 0xd57dc824.
//
// Solidity: function nodeRewardsFileCid() view returns(string)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) NodeRewardsFileCid() (string, error) {
	return _NetworkWithdraw.Contract.NodeRewardsFileCid(&_NetworkWithdraw.CallOpts)
}

// PlatformCommissionRate is a free data retrieval call binding the contract method 0x1da4dd0d.
//
// Solidity: function platformCommissionRate() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) PlatformCommissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "platformCommissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformCommissionRate is a free data retrieval call binding the contract method 0x1da4dd0d.
//
// Solidity: function platformCommissionRate() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) PlatformCommissionRate() (*big.Int, error) {
	return _NetworkWithdraw.Contract.PlatformCommissionRate(&_NetworkWithdraw.CallOpts)
}

// PlatformCommissionRate is a free data retrieval call binding the contract method 0x1da4dd0d.
//
// Solidity: function platformCommissionRate() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) PlatformCommissionRate() (*big.Int, error) {
	return _NetworkWithdraw.Contract.PlatformCommissionRate(&_NetworkWithdraw.CallOpts)
}

// TotalClaimedDepositOfNode is a free data retrieval call binding the contract method 0x6c570dc1.
//
// Solidity: function totalClaimedDepositOfNode(address ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) TotalClaimedDepositOfNode(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "totalClaimedDepositOfNode", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimedDepositOfNode is a free data retrieval call binding the contract method 0x6c570dc1.
//
// Solidity: function totalClaimedDepositOfNode(address ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) TotalClaimedDepositOfNode(arg0 common.Address) (*big.Int, error) {
	return _NetworkWithdraw.Contract.TotalClaimedDepositOfNode(&_NetworkWithdraw.CallOpts, arg0)
}

// TotalClaimedDepositOfNode is a free data retrieval call binding the contract method 0x6c570dc1.
//
// Solidity: function totalClaimedDepositOfNode(address ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) TotalClaimedDepositOfNode(arg0 common.Address) (*big.Int, error) {
	return _NetworkWithdraw.Contract.TotalClaimedDepositOfNode(&_NetworkWithdraw.CallOpts, arg0)
}

// TotalClaimedRewardOfNode is a free data retrieval call binding the contract method 0xbb2d840c.
//
// Solidity: function totalClaimedRewardOfNode(address ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) TotalClaimedRewardOfNode(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "totalClaimedRewardOfNode", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimedRewardOfNode is a free data retrieval call binding the contract method 0xbb2d840c.
//
// Solidity: function totalClaimedRewardOfNode(address ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) TotalClaimedRewardOfNode(arg0 common.Address) (*big.Int, error) {
	return _NetworkWithdraw.Contract.TotalClaimedRewardOfNode(&_NetworkWithdraw.CallOpts, arg0)
}

// TotalClaimedRewardOfNode is a free data retrieval call binding the contract method 0xbb2d840c.
//
// Solidity: function totalClaimedRewardOfNode(address ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) TotalClaimedRewardOfNode(arg0 common.Address) (*big.Int, error) {
	return _NetworkWithdraw.Contract.TotalClaimedRewardOfNode(&_NetworkWithdraw.CallOpts, arg0)
}

// TotalMissingAmountForWithdraw is a free data retrieval call binding the contract method 0x3c677dbe.
//
// Solidity: function totalMissingAmountForWithdraw() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) TotalMissingAmountForWithdraw(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "totalMissingAmountForWithdraw")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMissingAmountForWithdraw is a free data retrieval call binding the contract method 0x3c677dbe.
//
// Solidity: function totalMissingAmountForWithdraw() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) TotalMissingAmountForWithdraw() (*big.Int, error) {
	return _NetworkWithdraw.Contract.TotalMissingAmountForWithdraw(&_NetworkWithdraw.CallOpts)
}

// TotalMissingAmountForWithdraw is a free data retrieval call binding the contract method 0x3c677dbe.
//
// Solidity: function totalMissingAmountForWithdraw() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) TotalMissingAmountForWithdraw() (*big.Int, error) {
	return _NetworkWithdraw.Contract.TotalMissingAmountForWithdraw(&_NetworkWithdraw.CallOpts)
}

// TotalPlatformClaimedAmount is a free data retrieval call binding the contract method 0xb3594059.
//
// Solidity: function totalPlatformClaimedAmount() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) TotalPlatformClaimedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "totalPlatformClaimedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPlatformClaimedAmount is a free data retrieval call binding the contract method 0xb3594059.
//
// Solidity: function totalPlatformClaimedAmount() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) TotalPlatformClaimedAmount() (*big.Int, error) {
	return _NetworkWithdraw.Contract.TotalPlatformClaimedAmount(&_NetworkWithdraw.CallOpts)
}

// TotalPlatformClaimedAmount is a free data retrieval call binding the contract method 0xb3594059.
//
// Solidity: function totalPlatformClaimedAmount() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) TotalPlatformClaimedAmount() (*big.Int, error) {
	return _NetworkWithdraw.Contract.TotalPlatformClaimedAmount(&_NetworkWithdraw.CallOpts)
}

// TotalPlatformCommission is a free data retrieval call binding the contract method 0xfef25c0d.
//
// Solidity: function totalPlatformCommission() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) TotalPlatformCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "totalPlatformCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPlatformCommission is a free data retrieval call binding the contract method 0xfef25c0d.
//
// Solidity: function totalPlatformCommission() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) TotalPlatformCommission() (*big.Int, error) {
	return _NetworkWithdraw.Contract.TotalPlatformCommission(&_NetworkWithdraw.CallOpts)
}

// TotalPlatformCommission is a free data retrieval call binding the contract method 0xfef25c0d.
//
// Solidity: function totalPlatformCommission() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) TotalPlatformCommission() (*big.Int, error) {
	return _NetworkWithdraw.Contract.TotalPlatformCommission(&_NetworkWithdraw.CallOpts)
}

// TotalWithdrawAmountAtCycle is a free data retrieval call binding the contract method 0x8a726d78.
//
// Solidity: function totalWithdrawAmountAtCycle(uint256 ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) TotalWithdrawAmountAtCycle(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "totalWithdrawAmountAtCycle", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWithdrawAmountAtCycle is a free data retrieval call binding the contract method 0x8a726d78.
//
// Solidity: function totalWithdrawAmountAtCycle(uint256 ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) TotalWithdrawAmountAtCycle(arg0 *big.Int) (*big.Int, error) {
	return _NetworkWithdraw.Contract.TotalWithdrawAmountAtCycle(&_NetworkWithdraw.CallOpts, arg0)
}

// TotalWithdrawAmountAtCycle is a free data retrieval call binding the contract method 0x8a726d78.
//
// Solidity: function totalWithdrawAmountAtCycle(uint256 ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) TotalWithdrawAmountAtCycle(arg0 *big.Int) (*big.Int, error) {
	return _NetworkWithdraw.Contract.TotalWithdrawAmountAtCycle(&_NetworkWithdraw.CallOpts, arg0)
}

// UserDepositAddress is a free data retrieval call binding the contract method 0x46773830.
//
// Solidity: function userDepositAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawCaller) UserDepositAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "userDepositAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserDepositAddress is a free data retrieval call binding the contract method 0x46773830.
//
// Solidity: function userDepositAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawSession) UserDepositAddress() (common.Address, error) {
	return _NetworkWithdraw.Contract.UserDepositAddress(&_NetworkWithdraw.CallOpts)
}

// UserDepositAddress is a free data retrieval call binding the contract method 0x46773830.
//
// Solidity: function userDepositAddress() view returns(address)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) UserDepositAddress() (common.Address, error) {
	return _NetworkWithdraw.Contract.UserDepositAddress(&_NetworkWithdraw.CallOpts)
}

// UserWithdrawAmountAtCycle is a free data retrieval call binding the contract method 0xf5ff612d.
//
// Solidity: function userWithdrawAmountAtCycle(address , uint256 ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) UserWithdrawAmountAtCycle(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "userWithdrawAmountAtCycle", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserWithdrawAmountAtCycle is a free data retrieval call binding the contract method 0xf5ff612d.
//
// Solidity: function userWithdrawAmountAtCycle(address , uint256 ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) UserWithdrawAmountAtCycle(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NetworkWithdraw.Contract.UserWithdrawAmountAtCycle(&_NetworkWithdraw.CallOpts, arg0, arg1)
}

// UserWithdrawAmountAtCycle is a free data retrieval call binding the contract method 0xf5ff612d.
//
// Solidity: function userWithdrawAmountAtCycle(address , uint256 ) view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) UserWithdrawAmountAtCycle(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NetworkWithdraw.Contract.UserWithdrawAmountAtCycle(&_NetworkWithdraw.CallOpts, arg0, arg1)
}

// UserWithdrawLimitAmountPerCycle is a free data retrieval call binding the contract method 0xc2942579.
//
// Solidity: function userWithdrawLimitAmountPerCycle() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) UserWithdrawLimitAmountPerCycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "userWithdrawLimitAmountPerCycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserWithdrawLimitAmountPerCycle is a free data retrieval call binding the contract method 0xc2942579.
//
// Solidity: function userWithdrawLimitAmountPerCycle() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) UserWithdrawLimitAmountPerCycle() (*big.Int, error) {
	return _NetworkWithdraw.Contract.UserWithdrawLimitAmountPerCycle(&_NetworkWithdraw.CallOpts)
}

// UserWithdrawLimitAmountPerCycle is a free data retrieval call binding the contract method 0xc2942579.
//
// Solidity: function userWithdrawLimitAmountPerCycle() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) UserWithdrawLimitAmountPerCycle() (*big.Int, error) {
	return _NetworkWithdraw.Contract.UserWithdrawLimitAmountPerCycle(&_NetworkWithdraw.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkWithdraw *NetworkWithdrawCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkWithdraw *NetworkWithdrawSession) Version() (uint8, error) {
	return _NetworkWithdraw.Contract.Version(&_NetworkWithdraw.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) Version() (uint8, error) {
	return _NetworkWithdraw.Contract.Version(&_NetworkWithdraw.CallOpts)
}

// WithdrawCycleSeconds is a free data retrieval call binding the contract method 0x4a4b061b.
//
// Solidity: function withdrawCycleSeconds() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) WithdrawCycleSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "withdrawCycleSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawCycleSeconds is a free data retrieval call binding the contract method 0x4a4b061b.
//
// Solidity: function withdrawCycleSeconds() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) WithdrawCycleSeconds() (*big.Int, error) {
	return _NetworkWithdraw.Contract.WithdrawCycleSeconds(&_NetworkWithdraw.CallOpts)
}

// WithdrawCycleSeconds is a free data retrieval call binding the contract method 0x4a4b061b.
//
// Solidity: function withdrawCycleSeconds() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) WithdrawCycleSeconds() (*big.Int, error) {
	return _NetworkWithdraw.Contract.WithdrawCycleSeconds(&_NetworkWithdraw.CallOpts)
}

// WithdrawLimitAmountPerCycle is a free data retrieval call binding the contract method 0xdfbb4b72.
//
// Solidity: function withdrawLimitAmountPerCycle() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCaller) WithdrawLimitAmountPerCycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "withdrawLimitAmountPerCycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawLimitAmountPerCycle is a free data retrieval call binding the contract method 0xdfbb4b72.
//
// Solidity: function withdrawLimitAmountPerCycle() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawSession) WithdrawLimitAmountPerCycle() (*big.Int, error) {
	return _NetworkWithdraw.Contract.WithdrawLimitAmountPerCycle(&_NetworkWithdraw.CallOpts)
}

// WithdrawLimitAmountPerCycle is a free data retrieval call binding the contract method 0xdfbb4b72.
//
// Solidity: function withdrawLimitAmountPerCycle() view returns(uint256)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) WithdrawLimitAmountPerCycle() (*big.Int, error) {
	return _NetworkWithdraw.Contract.WithdrawLimitAmountPerCycle(&_NetworkWithdraw.CallOpts)
}

// WithdrawalAtIndex is a free data retrieval call binding the contract method 0xa8e1b8ef.
//
// Solidity: function withdrawalAtIndex(uint256 ) view returns(address _address, uint256 _amount)
func (_NetworkWithdraw *NetworkWithdrawCaller) WithdrawalAtIndex(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Address common.Address
	Amount  *big.Int
}, error) {
	var out []interface{}
	err := _NetworkWithdraw.contract.Call(opts, &out, "withdrawalAtIndex", arg0)

	outstruct := new(struct {
		Address common.Address
		Amount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Address = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// WithdrawalAtIndex is a free data retrieval call binding the contract method 0xa8e1b8ef.
//
// Solidity: function withdrawalAtIndex(uint256 ) view returns(address _address, uint256 _amount)
func (_NetworkWithdraw *NetworkWithdrawSession) WithdrawalAtIndex(arg0 *big.Int) (struct {
	Address common.Address
	Amount  *big.Int
}, error) {
	return _NetworkWithdraw.Contract.WithdrawalAtIndex(&_NetworkWithdraw.CallOpts, arg0)
}

// WithdrawalAtIndex is a free data retrieval call binding the contract method 0xa8e1b8ef.
//
// Solidity: function withdrawalAtIndex(uint256 ) view returns(address _address, uint256 _amount)
func (_NetworkWithdraw *NetworkWithdrawCallerSession) WithdrawalAtIndex(arg0 *big.Int) (struct {
	Address common.Address
	Amount  *big.Int
}, error) {
	return _NetworkWithdraw.Contract.WithdrawalAtIndex(&_NetworkWithdraw.CallOpts, arg0)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) DepositEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "depositEth")
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_NetworkWithdraw *NetworkWithdrawSession) DepositEth() (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.DepositEth(&_NetworkWithdraw.TransactOpts)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) DepositEth() (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.DepositEth(&_NetworkWithdraw.TransactOpts)
}

// DepositEthAndUpdateTotalMissingAmount is a paid mutator transaction binding the contract method 0x9d2f846e.
//
// Solidity: function depositEthAndUpdateTotalMissingAmount() payable returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) DepositEthAndUpdateTotalMissingAmount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "depositEthAndUpdateTotalMissingAmount")
}

// DepositEthAndUpdateTotalMissingAmount is a paid mutator transaction binding the contract method 0x9d2f846e.
//
// Solidity: function depositEthAndUpdateTotalMissingAmount() payable returns()
func (_NetworkWithdraw *NetworkWithdrawSession) DepositEthAndUpdateTotalMissingAmount() (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.DepositEthAndUpdateTotalMissingAmount(&_NetworkWithdraw.TransactOpts)
}

// DepositEthAndUpdateTotalMissingAmount is a paid mutator transaction binding the contract method 0x9d2f846e.
//
// Solidity: function depositEthAndUpdateTotalMissingAmount() payable returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) DepositEthAndUpdateTotalMissingAmount() (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.DepositEthAndUpdateTotalMissingAmount(&_NetworkWithdraw.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xc980ba89.
//
// Solidity: function distribute(uint8 _distributeType, uint256 _dealedHeight, uint256 _userAmount, uint256 _nodeAmount, uint256 _platformAmount, uint256 _maxClaimableWithdrawIndex) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) Distribute(opts *bind.TransactOpts, _distributeType uint8, _dealedHeight *big.Int, _userAmount *big.Int, _nodeAmount *big.Int, _platformAmount *big.Int, _maxClaimableWithdrawIndex *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "distribute", _distributeType, _dealedHeight, _userAmount, _nodeAmount, _platformAmount, _maxClaimableWithdrawIndex)
}

// Distribute is a paid mutator transaction binding the contract method 0xc980ba89.
//
// Solidity: function distribute(uint8 _distributeType, uint256 _dealedHeight, uint256 _userAmount, uint256 _nodeAmount, uint256 _platformAmount, uint256 _maxClaimableWithdrawIndex) returns()
func (_NetworkWithdraw *NetworkWithdrawSession) Distribute(_distributeType uint8, _dealedHeight *big.Int, _userAmount *big.Int, _nodeAmount *big.Int, _platformAmount *big.Int, _maxClaimableWithdrawIndex *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.Distribute(&_NetworkWithdraw.TransactOpts, _distributeType, _dealedHeight, _userAmount, _nodeAmount, _platformAmount, _maxClaimableWithdrawIndex)
}

// Distribute is a paid mutator transaction binding the contract method 0xc980ba89.
//
// Solidity: function distribute(uint8 _distributeType, uint256 _dealedHeight, uint256 _userAmount, uint256 _nodeAmount, uint256 _platformAmount, uint256 _maxClaimableWithdrawIndex) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) Distribute(_distributeType uint8, _dealedHeight *big.Int, _userAmount *big.Int, _nodeAmount *big.Int, _platformAmount *big.Int, _maxClaimableWithdrawIndex *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.Distribute(&_NetworkWithdraw.TransactOpts, _distributeType, _dealedHeight, _userAmount, _nodeAmount, _platformAmount, _maxClaimableWithdrawIndex)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _lsdTokenAddress, address _userDepositAddress, address _networkProposalAddress, address _feePoolAddress, address _factoryAddress) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) Init(opts *bind.TransactOpts, _lsdTokenAddress common.Address, _userDepositAddress common.Address, _networkProposalAddress common.Address, _feePoolAddress common.Address, _factoryAddress common.Address) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "init", _lsdTokenAddress, _userDepositAddress, _networkProposalAddress, _feePoolAddress, _factoryAddress)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _lsdTokenAddress, address _userDepositAddress, address _networkProposalAddress, address _feePoolAddress, address _factoryAddress) returns()
func (_NetworkWithdraw *NetworkWithdrawSession) Init(_lsdTokenAddress common.Address, _userDepositAddress common.Address, _networkProposalAddress common.Address, _feePoolAddress common.Address, _factoryAddress common.Address) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.Init(&_NetworkWithdraw.TransactOpts, _lsdTokenAddress, _userDepositAddress, _networkProposalAddress, _feePoolAddress, _factoryAddress)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _lsdTokenAddress, address _userDepositAddress, address _networkProposalAddress, address _feePoolAddress, address _factoryAddress) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) Init(_lsdTokenAddress common.Address, _userDepositAddress common.Address, _networkProposalAddress common.Address, _feePoolAddress common.Address, _factoryAddress common.Address) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.Init(&_NetworkWithdraw.TransactOpts, _lsdTokenAddress, _userDepositAddress, _networkProposalAddress, _feePoolAddress, _factoryAddress)
}

// NodeClaim is a paid mutator transaction binding the contract method 0xfdf435e9.
//
// Solidity: function nodeClaim(uint256 _index, address _account, uint256 _totalRewardAmount, uint256 _totalExitDepositAmount, bytes32[] _merkleProof, uint8 _claimType) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) NodeClaim(opts *bind.TransactOpts, _index *big.Int, _account common.Address, _totalRewardAmount *big.Int, _totalExitDepositAmount *big.Int, _merkleProof [][32]byte, _claimType uint8) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "nodeClaim", _index, _account, _totalRewardAmount, _totalExitDepositAmount, _merkleProof, _claimType)
}

// NodeClaim is a paid mutator transaction binding the contract method 0xfdf435e9.
//
// Solidity: function nodeClaim(uint256 _index, address _account, uint256 _totalRewardAmount, uint256 _totalExitDepositAmount, bytes32[] _merkleProof, uint8 _claimType) returns()
func (_NetworkWithdraw *NetworkWithdrawSession) NodeClaim(_index *big.Int, _account common.Address, _totalRewardAmount *big.Int, _totalExitDepositAmount *big.Int, _merkleProof [][32]byte, _claimType uint8) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.NodeClaim(&_NetworkWithdraw.TransactOpts, _index, _account, _totalRewardAmount, _totalExitDepositAmount, _merkleProof, _claimType)
}

// NodeClaim is a paid mutator transaction binding the contract method 0xfdf435e9.
//
// Solidity: function nodeClaim(uint256 _index, address _account, uint256 _totalRewardAmount, uint256 _totalExitDepositAmount, bytes32[] _merkleProof, uint8 _claimType) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) NodeClaim(_index *big.Int, _account common.Address, _totalRewardAmount *big.Int, _totalExitDepositAmount *big.Int, _merkleProof [][32]byte, _claimType uint8) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.NodeClaim(&_NetworkWithdraw.TransactOpts, _index, _account, _totalRewardAmount, _totalExitDepositAmount, _merkleProof, _claimType)
}

// NotifyValidatorExit is a paid mutator transaction binding the contract method 0x1e0f4aae.
//
// Solidity: function notifyValidatorExit(uint256 _withdrawCycle, uint256 _ejectedStartCycle, uint256[] _validatorIndexList) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) NotifyValidatorExit(opts *bind.TransactOpts, _withdrawCycle *big.Int, _ejectedStartCycle *big.Int, _validatorIndexList []*big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "notifyValidatorExit", _withdrawCycle, _ejectedStartCycle, _validatorIndexList)
}

// NotifyValidatorExit is a paid mutator transaction binding the contract method 0x1e0f4aae.
//
// Solidity: function notifyValidatorExit(uint256 _withdrawCycle, uint256 _ejectedStartCycle, uint256[] _validatorIndexList) returns()
func (_NetworkWithdraw *NetworkWithdrawSession) NotifyValidatorExit(_withdrawCycle *big.Int, _ejectedStartCycle *big.Int, _validatorIndexList []*big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.NotifyValidatorExit(&_NetworkWithdraw.TransactOpts, _withdrawCycle, _ejectedStartCycle, _validatorIndexList)
}

// NotifyValidatorExit is a paid mutator transaction binding the contract method 0x1e0f4aae.
//
// Solidity: function notifyValidatorExit(uint256 _withdrawCycle, uint256 _ejectedStartCycle, uint256[] _validatorIndexList) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) NotifyValidatorExit(_withdrawCycle *big.Int, _ejectedStartCycle *big.Int, _validatorIndexList []*big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.NotifyValidatorExit(&_NetworkWithdraw.TransactOpts, _withdrawCycle, _ejectedStartCycle, _validatorIndexList)
}

// PlatformClaim is a paid mutator transaction binding the contract method 0xaaf82770.
//
// Solidity: function platformClaim(address _recipient) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) PlatformClaim(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "platformClaim", _recipient)
}

// PlatformClaim is a paid mutator transaction binding the contract method 0xaaf82770.
//
// Solidity: function platformClaim(address _recipient) returns()
func (_NetworkWithdraw *NetworkWithdrawSession) PlatformClaim(_recipient common.Address) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.PlatformClaim(&_NetworkWithdraw.TransactOpts, _recipient)
}

// PlatformClaim is a paid mutator transaction binding the contract method 0xaaf82770.
//
// Solidity: function platformClaim(address _recipient) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) PlatformClaim(_recipient common.Address) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.PlatformClaim(&_NetworkWithdraw.TransactOpts, _recipient)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x12b81931.
//
// Solidity: function setMerkleRoot(uint256 _dealedEpoch, bytes32 _merkleRoot, string _nodeRewardsFileCid) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) SetMerkleRoot(opts *bind.TransactOpts, _dealedEpoch *big.Int, _merkleRoot [32]byte, _nodeRewardsFileCid string) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "setMerkleRoot", _dealedEpoch, _merkleRoot, _nodeRewardsFileCid)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x12b81931.
//
// Solidity: function setMerkleRoot(uint256 _dealedEpoch, bytes32 _merkleRoot, string _nodeRewardsFileCid) returns()
func (_NetworkWithdraw *NetworkWithdrawSession) SetMerkleRoot(_dealedEpoch *big.Int, _merkleRoot [32]byte, _nodeRewardsFileCid string) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.SetMerkleRoot(&_NetworkWithdraw.TransactOpts, _dealedEpoch, _merkleRoot, _nodeRewardsFileCid)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x12b81931.
//
// Solidity: function setMerkleRoot(uint256 _dealedEpoch, bytes32 _merkleRoot, string _nodeRewardsFileCid) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) SetMerkleRoot(_dealedEpoch *big.Int, _merkleRoot [32]byte, _nodeRewardsFileCid string) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.SetMerkleRoot(&_NetworkWithdraw.TransactOpts, _dealedEpoch, _merkleRoot, _nodeRewardsFileCid)
}

// SetUserWithdrawLimitAmountPerCycle is a paid mutator transaction binding the contract method 0xdd0cb13f.
//
// Solidity: function setUserWithdrawLimitAmountPerCycle(uint256 _userWithdrawLimitPerCycle) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) SetUserWithdrawLimitAmountPerCycle(opts *bind.TransactOpts, _userWithdrawLimitPerCycle *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "setUserWithdrawLimitAmountPerCycle", _userWithdrawLimitPerCycle)
}

// SetUserWithdrawLimitAmountPerCycle is a paid mutator transaction binding the contract method 0xdd0cb13f.
//
// Solidity: function setUserWithdrawLimitAmountPerCycle(uint256 _userWithdrawLimitPerCycle) returns()
func (_NetworkWithdraw *NetworkWithdrawSession) SetUserWithdrawLimitAmountPerCycle(_userWithdrawLimitPerCycle *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.SetUserWithdrawLimitAmountPerCycle(&_NetworkWithdraw.TransactOpts, _userWithdrawLimitPerCycle)
}

// SetUserWithdrawLimitAmountPerCycle is a paid mutator transaction binding the contract method 0xdd0cb13f.
//
// Solidity: function setUserWithdrawLimitAmountPerCycle(uint256 _userWithdrawLimitPerCycle) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) SetUserWithdrawLimitAmountPerCycle(_userWithdrawLimitPerCycle *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.SetUserWithdrawLimitAmountPerCycle(&_NetworkWithdraw.TransactOpts, _userWithdrawLimitPerCycle)
}

// SetWithdrawCycleSeconds is a paid mutator transaction binding the contract method 0x939d1ee4.
//
// Solidity: function setWithdrawCycleSeconds(uint256 _withdrawCycleSeconds) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) SetWithdrawCycleSeconds(opts *bind.TransactOpts, _withdrawCycleSeconds *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "setWithdrawCycleSeconds", _withdrawCycleSeconds)
}

// SetWithdrawCycleSeconds is a paid mutator transaction binding the contract method 0x939d1ee4.
//
// Solidity: function setWithdrawCycleSeconds(uint256 _withdrawCycleSeconds) returns()
func (_NetworkWithdraw *NetworkWithdrawSession) SetWithdrawCycleSeconds(_withdrawCycleSeconds *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.SetWithdrawCycleSeconds(&_NetworkWithdraw.TransactOpts, _withdrawCycleSeconds)
}

// SetWithdrawCycleSeconds is a paid mutator transaction binding the contract method 0x939d1ee4.
//
// Solidity: function setWithdrawCycleSeconds(uint256 _withdrawCycleSeconds) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) SetWithdrawCycleSeconds(_withdrawCycleSeconds *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.SetWithdrawCycleSeconds(&_NetworkWithdraw.TransactOpts, _withdrawCycleSeconds)
}

// SetWithdrawLimitAmountPerCycle is a paid mutator transaction binding the contract method 0xfe02cfa4.
//
// Solidity: function setWithdrawLimitAmountPerCycle(uint256 _withdrawLimitPerCycle) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) SetWithdrawLimitAmountPerCycle(opts *bind.TransactOpts, _withdrawLimitPerCycle *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "setWithdrawLimitAmountPerCycle", _withdrawLimitPerCycle)
}

// SetWithdrawLimitAmountPerCycle is a paid mutator transaction binding the contract method 0xfe02cfa4.
//
// Solidity: function setWithdrawLimitAmountPerCycle(uint256 _withdrawLimitPerCycle) returns()
func (_NetworkWithdraw *NetworkWithdrawSession) SetWithdrawLimitAmountPerCycle(_withdrawLimitPerCycle *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.SetWithdrawLimitAmountPerCycle(&_NetworkWithdraw.TransactOpts, _withdrawLimitPerCycle)
}

// SetWithdrawLimitAmountPerCycle is a paid mutator transaction binding the contract method 0xfe02cfa4.
//
// Solidity: function setWithdrawLimitAmountPerCycle(uint256 _withdrawLimitPerCycle) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) SetWithdrawLimitAmountPerCycle(_withdrawLimitPerCycle *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.SetWithdrawLimitAmountPerCycle(&_NetworkWithdraw.TransactOpts, _withdrawLimitPerCycle)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _lsdTokenAmount) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) Unstake(opts *bind.TransactOpts, _lsdTokenAmount *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "unstake", _lsdTokenAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _lsdTokenAmount) returns()
func (_NetworkWithdraw *NetworkWithdrawSession) Unstake(_lsdTokenAmount *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.Unstake(&_NetworkWithdraw.TransactOpts, _lsdTokenAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _lsdTokenAmount) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) Unstake(_lsdTokenAmount *big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.Unstake(&_NetworkWithdraw.TransactOpts, _lsdTokenAmount)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 _merkleRoot) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) UpdateMerkleRoot(opts *bind.TransactOpts, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "updateMerkleRoot", _merkleRoot)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 _merkleRoot) returns()
func (_NetworkWithdraw *NetworkWithdrawSession) UpdateMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.UpdateMerkleRoot(&_NetworkWithdraw.TransactOpts, _merkleRoot)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 _merkleRoot) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) UpdateMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.UpdateMerkleRoot(&_NetworkWithdraw.TransactOpts, _merkleRoot)
}

// Withdraw is a paid mutator transaction binding the contract method 0x983d95ce.
//
// Solidity: function withdraw(uint256[] _withdrawIndexList) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) Withdraw(opts *bind.TransactOpts, _withdrawIndexList []*big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.Transact(opts, "withdraw", _withdrawIndexList)
}

// Withdraw is a paid mutator transaction binding the contract method 0x983d95ce.
//
// Solidity: function withdraw(uint256[] _withdrawIndexList) returns()
func (_NetworkWithdraw *NetworkWithdrawSession) Withdraw(_withdrawIndexList []*big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.Withdraw(&_NetworkWithdraw.TransactOpts, _withdrawIndexList)
}

// Withdraw is a paid mutator transaction binding the contract method 0x983d95ce.
//
// Solidity: function withdraw(uint256[] _withdrawIndexList) returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) Withdraw(_withdrawIndexList []*big.Int) (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.Withdraw(&_NetworkWithdraw.TransactOpts, _withdrawIndexList)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NetworkWithdraw *NetworkWithdrawTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkWithdraw.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NetworkWithdraw *NetworkWithdrawSession) Receive() (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.Receive(&_NetworkWithdraw.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NetworkWithdraw *NetworkWithdrawTransactorSession) Receive() (*types.Transaction, error) {
	return _NetworkWithdraw.Contract.Receive(&_NetworkWithdraw.TransactOpts)
}

// NetworkWithdrawDistributeRewardsIterator is returned from FilterDistributeRewards and is used to iterate over the raw logs and unpacked data for DistributeRewards events raised by the NetworkWithdraw contract.
type NetworkWithdrawDistributeRewardsIterator struct {
	Event *NetworkWithdrawDistributeRewards // Event containing the contract specifics and raw log

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
func (it *NetworkWithdrawDistributeRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkWithdrawDistributeRewards)
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
		it.Event = new(NetworkWithdrawDistributeRewards)
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
func (it *NetworkWithdrawDistributeRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkWithdrawDistributeRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkWithdrawDistributeRewards represents a DistributeRewards event raised by the NetworkWithdraw contract.
type NetworkWithdrawDistributeRewards struct {
	DistributeType            uint8
	DealedHeight              *big.Int
	UserAmount                *big.Int
	NodeAmount                *big.Int
	PlatformAmount            *big.Int
	MaxClaimableWithdrawIndex *big.Int
	MvAmount                  *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterDistributeRewards is a free log retrieval operation binding the contract event 0xf10021cf129ec9c5003084ae330dba6d0bd143c47a2677c4d68939a19423206b.
//
// Solidity: event DistributeRewards(uint8 _distributeType, uint256 _dealedHeight, uint256 _userAmount, uint256 _nodeAmount, uint256 _platformAmount, uint256 _maxClaimableWithdrawIndex, uint256 _mvAmount)
func (_NetworkWithdraw *NetworkWithdrawFilterer) FilterDistributeRewards(opts *bind.FilterOpts) (*NetworkWithdrawDistributeRewardsIterator, error) {

	logs, sub, err := _NetworkWithdraw.contract.FilterLogs(opts, "DistributeRewards")
	if err != nil {
		return nil, err
	}
	return &NetworkWithdrawDistributeRewardsIterator{contract: _NetworkWithdraw.contract, event: "DistributeRewards", logs: logs, sub: sub}, nil
}

// WatchDistributeRewards is a free log subscription operation binding the contract event 0xf10021cf129ec9c5003084ae330dba6d0bd143c47a2677c4d68939a19423206b.
//
// Solidity: event DistributeRewards(uint8 _distributeType, uint256 _dealedHeight, uint256 _userAmount, uint256 _nodeAmount, uint256 _platformAmount, uint256 _maxClaimableWithdrawIndex, uint256 _mvAmount)
func (_NetworkWithdraw *NetworkWithdrawFilterer) WatchDistributeRewards(opts *bind.WatchOpts, sink chan<- *NetworkWithdrawDistributeRewards) (event.Subscription, error) {

	logs, sub, err := _NetworkWithdraw.contract.WatchLogs(opts, "DistributeRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkWithdrawDistributeRewards)
				if err := _NetworkWithdraw.contract.UnpackLog(event, "DistributeRewards", log); err != nil {
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

// ParseDistributeRewards is a log parse operation binding the contract event 0xf10021cf129ec9c5003084ae330dba6d0bd143c47a2677c4d68939a19423206b.
//
// Solidity: event DistributeRewards(uint8 _distributeType, uint256 _dealedHeight, uint256 _userAmount, uint256 _nodeAmount, uint256 _platformAmount, uint256 _maxClaimableWithdrawIndex, uint256 _mvAmount)
func (_NetworkWithdraw *NetworkWithdrawFilterer) ParseDistributeRewards(log types.Log) (*NetworkWithdrawDistributeRewards, error) {
	event := new(NetworkWithdrawDistributeRewards)
	if err := _NetworkWithdraw.contract.UnpackLog(event, "DistributeRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkWithdrawEtherDepositedIterator is returned from FilterEtherDeposited and is used to iterate over the raw logs and unpacked data for EtherDeposited events raised by the NetworkWithdraw contract.
type NetworkWithdrawEtherDepositedIterator struct {
	Event *NetworkWithdrawEtherDeposited // Event containing the contract specifics and raw log

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
func (it *NetworkWithdrawEtherDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkWithdrawEtherDeposited)
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
		it.Event = new(NetworkWithdrawEtherDeposited)
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
func (it *NetworkWithdrawEtherDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkWithdrawEtherDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkWithdrawEtherDeposited represents a EtherDeposited event raised by the NetworkWithdraw contract.
type NetworkWithdrawEtherDeposited struct {
	From   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherDeposited is a free log retrieval operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed _from, uint256 _amount, uint256 _time)
func (_NetworkWithdraw *NetworkWithdrawFilterer) FilterEtherDeposited(opts *bind.FilterOpts, _from []common.Address) (*NetworkWithdrawEtherDepositedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NetworkWithdraw.contract.FilterLogs(opts, "EtherDeposited", _fromRule)
	if err != nil {
		return nil, err
	}
	return &NetworkWithdrawEtherDepositedIterator{contract: _NetworkWithdraw.contract, event: "EtherDeposited", logs: logs, sub: sub}, nil
}

// WatchEtherDeposited is a free log subscription operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed _from, uint256 _amount, uint256 _time)
func (_NetworkWithdraw *NetworkWithdrawFilterer) WatchEtherDeposited(opts *bind.WatchOpts, sink chan<- *NetworkWithdrawEtherDeposited, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NetworkWithdraw.contract.WatchLogs(opts, "EtherDeposited", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkWithdrawEtherDeposited)
				if err := _NetworkWithdraw.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
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

// ParseEtherDeposited is a log parse operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed _from, uint256 _amount, uint256 _time)
func (_NetworkWithdraw *NetworkWithdrawFilterer) ParseEtherDeposited(log types.Log) (*NetworkWithdrawEtherDeposited, error) {
	event := new(NetworkWithdrawEtherDeposited)
	if err := _NetworkWithdraw.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkWithdrawNodeClaimedIterator is returned from FilterNodeClaimed and is used to iterate over the raw logs and unpacked data for NodeClaimed events raised by the NetworkWithdraw contract.
type NetworkWithdrawNodeClaimedIterator struct {
	Event *NetworkWithdrawNodeClaimed // Event containing the contract specifics and raw log

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
func (it *NetworkWithdrawNodeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkWithdrawNodeClaimed)
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
		it.Event = new(NetworkWithdrawNodeClaimed)
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
func (it *NetworkWithdrawNodeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkWithdrawNodeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkWithdrawNodeClaimed represents a NodeClaimed event raised by the NetworkWithdraw contract.
type NetworkWithdrawNodeClaimed struct {
	Index            *big.Int
	Account          common.Address
	ClaimableReward  *big.Int
	ClaimableDeposit *big.Int
	ClaimType        uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNodeClaimed is a free log retrieval operation binding the contract event 0x659e842f0209726671f562e8d7ee3a25d2cef92c2b85dcd268af93ef5d13582c.
//
// Solidity: event NodeClaimed(uint256 _index, address _account, uint256 _claimableReward, uint256 _claimableDeposit, uint8 _claimType)
func (_NetworkWithdraw *NetworkWithdrawFilterer) FilterNodeClaimed(opts *bind.FilterOpts) (*NetworkWithdrawNodeClaimedIterator, error) {

	logs, sub, err := _NetworkWithdraw.contract.FilterLogs(opts, "NodeClaimed")
	if err != nil {
		return nil, err
	}
	return &NetworkWithdrawNodeClaimedIterator{contract: _NetworkWithdraw.contract, event: "NodeClaimed", logs: logs, sub: sub}, nil
}

// WatchNodeClaimed is a free log subscription operation binding the contract event 0x659e842f0209726671f562e8d7ee3a25d2cef92c2b85dcd268af93ef5d13582c.
//
// Solidity: event NodeClaimed(uint256 _index, address _account, uint256 _claimableReward, uint256 _claimableDeposit, uint8 _claimType)
func (_NetworkWithdraw *NetworkWithdrawFilterer) WatchNodeClaimed(opts *bind.WatchOpts, sink chan<- *NetworkWithdrawNodeClaimed) (event.Subscription, error) {

	logs, sub, err := _NetworkWithdraw.contract.WatchLogs(opts, "NodeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkWithdrawNodeClaimed)
				if err := _NetworkWithdraw.contract.UnpackLog(event, "NodeClaimed", log); err != nil {
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

// ParseNodeClaimed is a log parse operation binding the contract event 0x659e842f0209726671f562e8d7ee3a25d2cef92c2b85dcd268af93ef5d13582c.
//
// Solidity: event NodeClaimed(uint256 _index, address _account, uint256 _claimableReward, uint256 _claimableDeposit, uint8 _claimType)
func (_NetworkWithdraw *NetworkWithdrawFilterer) ParseNodeClaimed(log types.Log) (*NetworkWithdrawNodeClaimed, error) {
	event := new(NetworkWithdrawNodeClaimed)
	if err := _NetworkWithdraw.contract.UnpackLog(event, "NodeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkWithdrawNotifyValidatorExitIterator is returned from FilterNotifyValidatorExit and is used to iterate over the raw logs and unpacked data for NotifyValidatorExit events raised by the NetworkWithdraw contract.
type NetworkWithdrawNotifyValidatorExitIterator struct {
	Event *NetworkWithdrawNotifyValidatorExit // Event containing the contract specifics and raw log

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
func (it *NetworkWithdrawNotifyValidatorExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkWithdrawNotifyValidatorExit)
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
		it.Event = new(NetworkWithdrawNotifyValidatorExit)
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
func (it *NetworkWithdrawNotifyValidatorExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkWithdrawNotifyValidatorExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkWithdrawNotifyValidatorExit represents a NotifyValidatorExit event raised by the NetworkWithdraw contract.
type NetworkWithdrawNotifyValidatorExit struct {
	WithdrawCycle             *big.Int
	EjectedStartWithdrawCycle *big.Int
	EjectedValidators         []*big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterNotifyValidatorExit is a free log retrieval operation binding the contract event 0xb83477449e27b4bab4f28c938d033b953557d6a1b9b4469a43d229f78ed6e55c.
//
// Solidity: event NotifyValidatorExit(uint256 _withdrawCycle, uint256 _ejectedStartWithdrawCycle, uint256[] _ejectedValidators)
func (_NetworkWithdraw *NetworkWithdrawFilterer) FilterNotifyValidatorExit(opts *bind.FilterOpts) (*NetworkWithdrawNotifyValidatorExitIterator, error) {

	logs, sub, err := _NetworkWithdraw.contract.FilterLogs(opts, "NotifyValidatorExit")
	if err != nil {
		return nil, err
	}
	return &NetworkWithdrawNotifyValidatorExitIterator{contract: _NetworkWithdraw.contract, event: "NotifyValidatorExit", logs: logs, sub: sub}, nil
}

// WatchNotifyValidatorExit is a free log subscription operation binding the contract event 0xb83477449e27b4bab4f28c938d033b953557d6a1b9b4469a43d229f78ed6e55c.
//
// Solidity: event NotifyValidatorExit(uint256 _withdrawCycle, uint256 _ejectedStartWithdrawCycle, uint256[] _ejectedValidators)
func (_NetworkWithdraw *NetworkWithdrawFilterer) WatchNotifyValidatorExit(opts *bind.WatchOpts, sink chan<- *NetworkWithdrawNotifyValidatorExit) (event.Subscription, error) {

	logs, sub, err := _NetworkWithdraw.contract.WatchLogs(opts, "NotifyValidatorExit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkWithdrawNotifyValidatorExit)
				if err := _NetworkWithdraw.contract.UnpackLog(event, "NotifyValidatorExit", log); err != nil {
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

// ParseNotifyValidatorExit is a log parse operation binding the contract event 0xb83477449e27b4bab4f28c938d033b953557d6a1b9b4469a43d229f78ed6e55c.
//
// Solidity: event NotifyValidatorExit(uint256 _withdrawCycle, uint256 _ejectedStartWithdrawCycle, uint256[] _ejectedValidators)
func (_NetworkWithdraw *NetworkWithdrawFilterer) ParseNotifyValidatorExit(log types.Log) (*NetworkWithdrawNotifyValidatorExit, error) {
	event := new(NetworkWithdrawNotifyValidatorExit)
	if err := _NetworkWithdraw.contract.UnpackLog(event, "NotifyValidatorExit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkWithdrawSetMerkleRootIterator is returned from FilterSetMerkleRoot and is used to iterate over the raw logs and unpacked data for SetMerkleRoot events raised by the NetworkWithdraw contract.
type NetworkWithdrawSetMerkleRootIterator struct {
	Event *NetworkWithdrawSetMerkleRoot // Event containing the contract specifics and raw log

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
func (it *NetworkWithdrawSetMerkleRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkWithdrawSetMerkleRoot)
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
		it.Event = new(NetworkWithdrawSetMerkleRoot)
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
func (it *NetworkWithdrawSetMerkleRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkWithdrawSetMerkleRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkWithdrawSetMerkleRoot represents a SetMerkleRoot event raised by the NetworkWithdraw contract.
type NetworkWithdrawSetMerkleRoot struct {
	DealedEpoch        *big.Int
	MerkleRoot         [32]byte
	NodeRewardsFileCid string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetMerkleRoot is a free log retrieval operation binding the contract event 0xec43b2424d0504da473794ad49016df3e06fb0d772bb403d724c9e5d53d8afb8.
//
// Solidity: event SetMerkleRoot(uint256 indexed _dealedEpoch, bytes32 _merkleRoot, string _nodeRewardsFileCid)
func (_NetworkWithdraw *NetworkWithdrawFilterer) FilterSetMerkleRoot(opts *bind.FilterOpts, _dealedEpoch []*big.Int) (*NetworkWithdrawSetMerkleRootIterator, error) {

	var _dealedEpochRule []interface{}
	for _, _dealedEpochItem := range _dealedEpoch {
		_dealedEpochRule = append(_dealedEpochRule, _dealedEpochItem)
	}

	logs, sub, err := _NetworkWithdraw.contract.FilterLogs(opts, "SetMerkleRoot", _dealedEpochRule)
	if err != nil {
		return nil, err
	}
	return &NetworkWithdrawSetMerkleRootIterator{contract: _NetworkWithdraw.contract, event: "SetMerkleRoot", logs: logs, sub: sub}, nil
}

// WatchSetMerkleRoot is a free log subscription operation binding the contract event 0xec43b2424d0504da473794ad49016df3e06fb0d772bb403d724c9e5d53d8afb8.
//
// Solidity: event SetMerkleRoot(uint256 indexed _dealedEpoch, bytes32 _merkleRoot, string _nodeRewardsFileCid)
func (_NetworkWithdraw *NetworkWithdrawFilterer) WatchSetMerkleRoot(opts *bind.WatchOpts, sink chan<- *NetworkWithdrawSetMerkleRoot, _dealedEpoch []*big.Int) (event.Subscription, error) {

	var _dealedEpochRule []interface{}
	for _, _dealedEpochItem := range _dealedEpoch {
		_dealedEpochRule = append(_dealedEpochRule, _dealedEpochItem)
	}

	logs, sub, err := _NetworkWithdraw.contract.WatchLogs(opts, "SetMerkleRoot", _dealedEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkWithdrawSetMerkleRoot)
				if err := _NetworkWithdraw.contract.UnpackLog(event, "SetMerkleRoot", log); err != nil {
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

// ParseSetMerkleRoot is a log parse operation binding the contract event 0xec43b2424d0504da473794ad49016df3e06fb0d772bb403d724c9e5d53d8afb8.
//
// Solidity: event SetMerkleRoot(uint256 indexed _dealedEpoch, bytes32 _merkleRoot, string _nodeRewardsFileCid)
func (_NetworkWithdraw *NetworkWithdrawFilterer) ParseSetMerkleRoot(log types.Log) (*NetworkWithdrawSetMerkleRoot, error) {
	event := new(NetworkWithdrawSetMerkleRoot)
	if err := _NetworkWithdraw.contract.UnpackLog(event, "SetMerkleRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkWithdrawSetUserWithdrawLimitPerCycleIterator is returned from FilterSetUserWithdrawLimitPerCycle and is used to iterate over the raw logs and unpacked data for SetUserWithdrawLimitPerCycle events raised by the NetworkWithdraw contract.
type NetworkWithdrawSetUserWithdrawLimitPerCycleIterator struct {
	Event *NetworkWithdrawSetUserWithdrawLimitPerCycle // Event containing the contract specifics and raw log

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
func (it *NetworkWithdrawSetUserWithdrawLimitPerCycleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkWithdrawSetUserWithdrawLimitPerCycle)
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
		it.Event = new(NetworkWithdrawSetUserWithdrawLimitPerCycle)
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
func (it *NetworkWithdrawSetUserWithdrawLimitPerCycleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkWithdrawSetUserWithdrawLimitPerCycleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkWithdrawSetUserWithdrawLimitPerCycle represents a SetUserWithdrawLimitPerCycle event raised by the NetworkWithdraw contract.
type NetworkWithdrawSetUserWithdrawLimitPerCycle struct {
	UserWithdrawLimitPerCycle *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetUserWithdrawLimitPerCycle is a free log retrieval operation binding the contract event 0x208ed057768c8997dde633000b59d3a1c2f498337c95bf9ecccc810f698d9194.
//
// Solidity: event SetUserWithdrawLimitPerCycle(uint256 _userWithdrawLimitPerCycle)
func (_NetworkWithdraw *NetworkWithdrawFilterer) FilterSetUserWithdrawLimitPerCycle(opts *bind.FilterOpts) (*NetworkWithdrawSetUserWithdrawLimitPerCycleIterator, error) {

	logs, sub, err := _NetworkWithdraw.contract.FilterLogs(opts, "SetUserWithdrawLimitPerCycle")
	if err != nil {
		return nil, err
	}
	return &NetworkWithdrawSetUserWithdrawLimitPerCycleIterator{contract: _NetworkWithdraw.contract, event: "SetUserWithdrawLimitPerCycle", logs: logs, sub: sub}, nil
}

// WatchSetUserWithdrawLimitPerCycle is a free log subscription operation binding the contract event 0x208ed057768c8997dde633000b59d3a1c2f498337c95bf9ecccc810f698d9194.
//
// Solidity: event SetUserWithdrawLimitPerCycle(uint256 _userWithdrawLimitPerCycle)
func (_NetworkWithdraw *NetworkWithdrawFilterer) WatchSetUserWithdrawLimitPerCycle(opts *bind.WatchOpts, sink chan<- *NetworkWithdrawSetUserWithdrawLimitPerCycle) (event.Subscription, error) {

	logs, sub, err := _NetworkWithdraw.contract.WatchLogs(opts, "SetUserWithdrawLimitPerCycle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkWithdrawSetUserWithdrawLimitPerCycle)
				if err := _NetworkWithdraw.contract.UnpackLog(event, "SetUserWithdrawLimitPerCycle", log); err != nil {
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

// ParseSetUserWithdrawLimitPerCycle is a log parse operation binding the contract event 0x208ed057768c8997dde633000b59d3a1c2f498337c95bf9ecccc810f698d9194.
//
// Solidity: event SetUserWithdrawLimitPerCycle(uint256 _userWithdrawLimitPerCycle)
func (_NetworkWithdraw *NetworkWithdrawFilterer) ParseSetUserWithdrawLimitPerCycle(log types.Log) (*NetworkWithdrawSetUserWithdrawLimitPerCycle, error) {
	event := new(NetworkWithdrawSetUserWithdrawLimitPerCycle)
	if err := _NetworkWithdraw.contract.UnpackLog(event, "SetUserWithdrawLimitPerCycle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkWithdrawSetWithdrawCycleSecondsIterator is returned from FilterSetWithdrawCycleSeconds and is used to iterate over the raw logs and unpacked data for SetWithdrawCycleSeconds events raised by the NetworkWithdraw contract.
type NetworkWithdrawSetWithdrawCycleSecondsIterator struct {
	Event *NetworkWithdrawSetWithdrawCycleSeconds // Event containing the contract specifics and raw log

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
func (it *NetworkWithdrawSetWithdrawCycleSecondsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkWithdrawSetWithdrawCycleSeconds)
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
		it.Event = new(NetworkWithdrawSetWithdrawCycleSeconds)
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
func (it *NetworkWithdrawSetWithdrawCycleSecondsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkWithdrawSetWithdrawCycleSecondsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkWithdrawSetWithdrawCycleSeconds represents a SetWithdrawCycleSeconds event raised by the NetworkWithdraw contract.
type NetworkWithdrawSetWithdrawCycleSeconds struct {
	Seconds *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetWithdrawCycleSeconds is a free log retrieval operation binding the contract event 0xa8bf31f5ff469d988ee50031edcb6b8a44b1cd010a1561d9a7a06d71c2193e6c.
//
// Solidity: event SetWithdrawCycleSeconds(uint256 _seconds)
func (_NetworkWithdraw *NetworkWithdrawFilterer) FilterSetWithdrawCycleSeconds(opts *bind.FilterOpts) (*NetworkWithdrawSetWithdrawCycleSecondsIterator, error) {

	logs, sub, err := _NetworkWithdraw.contract.FilterLogs(opts, "SetWithdrawCycleSeconds")
	if err != nil {
		return nil, err
	}
	return &NetworkWithdrawSetWithdrawCycleSecondsIterator{contract: _NetworkWithdraw.contract, event: "SetWithdrawCycleSeconds", logs: logs, sub: sub}, nil
}

// WatchSetWithdrawCycleSeconds is a free log subscription operation binding the contract event 0xa8bf31f5ff469d988ee50031edcb6b8a44b1cd010a1561d9a7a06d71c2193e6c.
//
// Solidity: event SetWithdrawCycleSeconds(uint256 _seconds)
func (_NetworkWithdraw *NetworkWithdrawFilterer) WatchSetWithdrawCycleSeconds(opts *bind.WatchOpts, sink chan<- *NetworkWithdrawSetWithdrawCycleSeconds) (event.Subscription, error) {

	logs, sub, err := _NetworkWithdraw.contract.WatchLogs(opts, "SetWithdrawCycleSeconds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkWithdrawSetWithdrawCycleSeconds)
				if err := _NetworkWithdraw.contract.UnpackLog(event, "SetWithdrawCycleSeconds", log); err != nil {
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

// ParseSetWithdrawCycleSeconds is a log parse operation binding the contract event 0xa8bf31f5ff469d988ee50031edcb6b8a44b1cd010a1561d9a7a06d71c2193e6c.
//
// Solidity: event SetWithdrawCycleSeconds(uint256 _seconds)
func (_NetworkWithdraw *NetworkWithdrawFilterer) ParseSetWithdrawCycleSeconds(log types.Log) (*NetworkWithdrawSetWithdrawCycleSeconds, error) {
	event := new(NetworkWithdrawSetWithdrawCycleSeconds)
	if err := _NetworkWithdraw.contract.UnpackLog(event, "SetWithdrawCycleSeconds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkWithdrawSetWithdrawLimitPerCycleIterator is returned from FilterSetWithdrawLimitPerCycle and is used to iterate over the raw logs and unpacked data for SetWithdrawLimitPerCycle events raised by the NetworkWithdraw contract.
type NetworkWithdrawSetWithdrawLimitPerCycleIterator struct {
	Event *NetworkWithdrawSetWithdrawLimitPerCycle // Event containing the contract specifics and raw log

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
func (it *NetworkWithdrawSetWithdrawLimitPerCycleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkWithdrawSetWithdrawLimitPerCycle)
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
		it.Event = new(NetworkWithdrawSetWithdrawLimitPerCycle)
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
func (it *NetworkWithdrawSetWithdrawLimitPerCycleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkWithdrawSetWithdrawLimitPerCycleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkWithdrawSetWithdrawLimitPerCycle represents a SetWithdrawLimitPerCycle event raised by the NetworkWithdraw contract.
type NetworkWithdrawSetWithdrawLimitPerCycle struct {
	WithdrawLimitPerCycle *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetWithdrawLimitPerCycle is a free log retrieval operation binding the contract event 0x8425a2e73ee4ea13649867c99971ddd01b64049295025867a4737f69c671358c.
//
// Solidity: event SetWithdrawLimitPerCycle(uint256 _withdrawLimitPerCycle)
func (_NetworkWithdraw *NetworkWithdrawFilterer) FilterSetWithdrawLimitPerCycle(opts *bind.FilterOpts) (*NetworkWithdrawSetWithdrawLimitPerCycleIterator, error) {

	logs, sub, err := _NetworkWithdraw.contract.FilterLogs(opts, "SetWithdrawLimitPerCycle")
	if err != nil {
		return nil, err
	}
	return &NetworkWithdrawSetWithdrawLimitPerCycleIterator{contract: _NetworkWithdraw.contract, event: "SetWithdrawLimitPerCycle", logs: logs, sub: sub}, nil
}

// WatchSetWithdrawLimitPerCycle is a free log subscription operation binding the contract event 0x8425a2e73ee4ea13649867c99971ddd01b64049295025867a4737f69c671358c.
//
// Solidity: event SetWithdrawLimitPerCycle(uint256 _withdrawLimitPerCycle)
func (_NetworkWithdraw *NetworkWithdrawFilterer) WatchSetWithdrawLimitPerCycle(opts *bind.WatchOpts, sink chan<- *NetworkWithdrawSetWithdrawLimitPerCycle) (event.Subscription, error) {

	logs, sub, err := _NetworkWithdraw.contract.WatchLogs(opts, "SetWithdrawLimitPerCycle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkWithdrawSetWithdrawLimitPerCycle)
				if err := _NetworkWithdraw.contract.UnpackLog(event, "SetWithdrawLimitPerCycle", log); err != nil {
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

// ParseSetWithdrawLimitPerCycle is a log parse operation binding the contract event 0x8425a2e73ee4ea13649867c99971ddd01b64049295025867a4737f69c671358c.
//
// Solidity: event SetWithdrawLimitPerCycle(uint256 _withdrawLimitPerCycle)
func (_NetworkWithdraw *NetworkWithdrawFilterer) ParseSetWithdrawLimitPerCycle(log types.Log) (*NetworkWithdrawSetWithdrawLimitPerCycle, error) {
	event := new(NetworkWithdrawSetWithdrawLimitPerCycle)
	if err := _NetworkWithdraw.contract.UnpackLog(event, "SetWithdrawLimitPerCycle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkWithdrawUnstakeIterator is returned from FilterUnstake and is used to iterate over the raw logs and unpacked data for Unstake events raised by the NetworkWithdraw contract.
type NetworkWithdrawUnstakeIterator struct {
	Event *NetworkWithdrawUnstake // Event containing the contract specifics and raw log

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
func (it *NetworkWithdrawUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkWithdrawUnstake)
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
		it.Event = new(NetworkWithdrawUnstake)
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
func (it *NetworkWithdrawUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkWithdrawUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkWithdrawUnstake represents a Unstake event raised by the NetworkWithdraw contract.
type NetworkWithdrawUnstake struct {
	From           common.Address
	LsdTokenAmount *big.Int
	EthAmount      *big.Int
	WithdrawIndex  *big.Int
	Instantly      bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUnstake is a free log retrieval operation binding the contract event 0xc7ccdcb2d25f572c6814e377dbb34ea4318a4b7d3cd890f5cfad699d75327c7c.
//
// Solidity: event Unstake(address indexed _from, uint256 _lsdTokenAmount, uint256 _ethAmount, uint256 _withdrawIndex, bool _instantly)
func (_NetworkWithdraw *NetworkWithdrawFilterer) FilterUnstake(opts *bind.FilterOpts, _from []common.Address) (*NetworkWithdrawUnstakeIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NetworkWithdraw.contract.FilterLogs(opts, "Unstake", _fromRule)
	if err != nil {
		return nil, err
	}
	return &NetworkWithdrawUnstakeIterator{contract: _NetworkWithdraw.contract, event: "Unstake", logs: logs, sub: sub}, nil
}

// WatchUnstake is a free log subscription operation binding the contract event 0xc7ccdcb2d25f572c6814e377dbb34ea4318a4b7d3cd890f5cfad699d75327c7c.
//
// Solidity: event Unstake(address indexed _from, uint256 _lsdTokenAmount, uint256 _ethAmount, uint256 _withdrawIndex, bool _instantly)
func (_NetworkWithdraw *NetworkWithdrawFilterer) WatchUnstake(opts *bind.WatchOpts, sink chan<- *NetworkWithdrawUnstake, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NetworkWithdraw.contract.WatchLogs(opts, "Unstake", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkWithdrawUnstake)
				if err := _NetworkWithdraw.contract.UnpackLog(event, "Unstake", log); err != nil {
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

// ParseUnstake is a log parse operation binding the contract event 0xc7ccdcb2d25f572c6814e377dbb34ea4318a4b7d3cd890f5cfad699d75327c7c.
//
// Solidity: event Unstake(address indexed _from, uint256 _lsdTokenAmount, uint256 _ethAmount, uint256 _withdrawIndex, bool _instantly)
func (_NetworkWithdraw *NetworkWithdrawFilterer) ParseUnstake(log types.Log) (*NetworkWithdrawUnstake, error) {
	event := new(NetworkWithdrawUnstake)
	if err := _NetworkWithdraw.contract.UnpackLog(event, "Unstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkWithdrawWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the NetworkWithdraw contract.
type NetworkWithdrawWithdrawIterator struct {
	Event *NetworkWithdrawWithdraw // Event containing the contract specifics and raw log

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
func (it *NetworkWithdrawWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkWithdrawWithdraw)
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
		it.Event = new(NetworkWithdrawWithdraw)
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
func (it *NetworkWithdrawWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkWithdrawWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkWithdrawWithdraw represents a Withdraw event raised by the NetworkWithdraw contract.
type NetworkWithdrawWithdraw struct {
	From              common.Address
	WithdrawIndexList []*big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x67e9df8b3c7743c9f1b625ba4f2b4e601206dbd46ed5c33c85a1242e4d23a2d1.
//
// Solidity: event Withdraw(address indexed _from, uint256[] _withdrawIndexList)
func (_NetworkWithdraw *NetworkWithdrawFilterer) FilterWithdraw(opts *bind.FilterOpts, _from []common.Address) (*NetworkWithdrawWithdrawIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NetworkWithdraw.contract.FilterLogs(opts, "Withdraw", _fromRule)
	if err != nil {
		return nil, err
	}
	return &NetworkWithdrawWithdrawIterator{contract: _NetworkWithdraw.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x67e9df8b3c7743c9f1b625ba4f2b4e601206dbd46ed5c33c85a1242e4d23a2d1.
//
// Solidity: event Withdraw(address indexed _from, uint256[] _withdrawIndexList)
func (_NetworkWithdraw *NetworkWithdrawFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *NetworkWithdrawWithdraw, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NetworkWithdraw.contract.WatchLogs(opts, "Withdraw", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkWithdrawWithdraw)
				if err := _NetworkWithdraw.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x67e9df8b3c7743c9f1b625ba4f2b4e601206dbd46ed5c33c85a1242e4d23a2d1.
//
// Solidity: event Withdraw(address indexed _from, uint256[] _withdrawIndexList)
func (_NetworkWithdraw *NetworkWithdrawFilterer) ParseWithdraw(log types.Log) (*NetworkWithdrawWithdraw, error) {
	event := new(NetworkWithdrawWithdraw)
	if err := _NetworkWithdraw.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
