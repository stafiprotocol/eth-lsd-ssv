// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package user_deposit

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

// UserDepositMetaData contains all meta data concerning the UserDeposit contract.
var UserDepositMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"DepositReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"DepositRecycled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ExcessWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lsdTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeDepositAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkWithdrawAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkProposalAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkBalancesAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lsdTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkBalancesAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkProposalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkWithdrawAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeDepositAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recycleNetworkWithdrawDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setDepositEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setMinimumDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawExcessBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UserDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use UserDepositMetaData.ABI instead.
var UserDepositABI = UserDepositMetaData.ABI

// UserDeposit is an auto generated Go binding around an Ethereum contract.
type UserDeposit struct {
	UserDepositCaller     // Read-only binding to the contract
	UserDepositTransactor // Write-only binding to the contract
	UserDepositFilterer   // Log filterer for contract events
}

// UserDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserDepositSession struct {
	Contract     *UserDeposit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserDepositCallerSession struct {
	Contract *UserDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UserDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserDepositTransactorSession struct {
	Contract     *UserDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UserDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserDepositRaw struct {
	Contract *UserDeposit // Generic contract binding to access the raw methods on
}

// UserDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserDepositCallerRaw struct {
	Contract *UserDepositCaller // Generic read-only contract binding to access the raw methods on
}

// UserDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserDepositTransactorRaw struct {
	Contract *UserDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserDeposit creates a new instance of UserDeposit, bound to a specific deployed contract.
func NewUserDeposit(address common.Address, backend bind.ContractBackend) (*UserDeposit, error) {
	contract, err := bindUserDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserDeposit{UserDepositCaller: UserDepositCaller{contract: contract}, UserDepositTransactor: UserDepositTransactor{contract: contract}, UserDepositFilterer: UserDepositFilterer{contract: contract}}, nil
}

// NewUserDepositCaller creates a new read-only instance of UserDeposit, bound to a specific deployed contract.
func NewUserDepositCaller(address common.Address, caller bind.ContractCaller) (*UserDepositCaller, error) {
	contract, err := bindUserDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserDepositCaller{contract: contract}, nil
}

// NewUserDepositTransactor creates a new write-only instance of UserDeposit, bound to a specific deployed contract.
func NewUserDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*UserDepositTransactor, error) {
	contract, err := bindUserDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserDepositTransactor{contract: contract}, nil
}

// NewUserDepositFilterer creates a new log filterer instance of UserDeposit, bound to a specific deployed contract.
func NewUserDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*UserDepositFilterer, error) {
	contract, err := bindUserDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserDepositFilterer{contract: contract}, nil
}

// bindUserDeposit binds a generic wrapper to an already deployed contract.
func bindUserDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserDepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserDeposit *UserDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserDeposit.Contract.UserDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserDeposit *UserDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDeposit.Contract.UserDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserDeposit *UserDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserDeposit.Contract.UserDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserDeposit *UserDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserDeposit *UserDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserDeposit *UserDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserDeposit.Contract.contract.Transact(opts, method, params...)
}

// DepositEnabled is a free data retrieval call binding the contract method 0x2eebe78e.
//
// Solidity: function depositEnabled() view returns(bool)
func (_UserDeposit *UserDepositCaller) DepositEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UserDeposit.contract.Call(opts, &out, "depositEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositEnabled is a free data retrieval call binding the contract method 0x2eebe78e.
//
// Solidity: function depositEnabled() view returns(bool)
func (_UserDeposit *UserDepositSession) DepositEnabled() (bool, error) {
	return _UserDeposit.Contract.DepositEnabled(&_UserDeposit.CallOpts)
}

// DepositEnabled is a free data retrieval call binding the contract method 0x2eebe78e.
//
// Solidity: function depositEnabled() view returns(bool)
func (_UserDeposit *UserDepositCallerSession) DepositEnabled() (bool, error) {
	return _UserDeposit.Contract.DepositEnabled(&_UserDeposit.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_UserDeposit *UserDepositCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserDeposit.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_UserDeposit *UserDepositSession) GetBalance() (*big.Int, error) {
	return _UserDeposit.Contract.GetBalance(&_UserDeposit.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_UserDeposit *UserDepositCallerSession) GetBalance() (*big.Int, error) {
	return _UserDeposit.Contract.GetBalance(&_UserDeposit.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_UserDeposit *UserDepositCaller) GetRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserDeposit.contract.Call(opts, &out, "getRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_UserDeposit *UserDepositSession) GetRate() (*big.Int, error) {
	return _UserDeposit.Contract.GetRate(&_UserDeposit.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_UserDeposit *UserDepositCallerSession) GetRate() (*big.Int, error) {
	return _UserDeposit.Contract.GetRate(&_UserDeposit.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_UserDeposit *UserDepositCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UserDeposit.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_UserDeposit *UserDepositSession) Initialized() (bool, error) {
	return _UserDeposit.Contract.Initialized(&_UserDeposit.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_UserDeposit *UserDepositCallerSession) Initialized() (bool, error) {
	return _UserDeposit.Contract.Initialized(&_UserDeposit.CallOpts)
}

// LsdTokenAddress is a free data retrieval call binding the contract method 0x87505b9d.
//
// Solidity: function lsdTokenAddress() view returns(address)
func (_UserDeposit *UserDepositCaller) LsdTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserDeposit.contract.Call(opts, &out, "lsdTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LsdTokenAddress is a free data retrieval call binding the contract method 0x87505b9d.
//
// Solidity: function lsdTokenAddress() view returns(address)
func (_UserDeposit *UserDepositSession) LsdTokenAddress() (common.Address, error) {
	return _UserDeposit.Contract.LsdTokenAddress(&_UserDeposit.CallOpts)
}

// LsdTokenAddress is a free data retrieval call binding the contract method 0x87505b9d.
//
// Solidity: function lsdTokenAddress() view returns(address)
func (_UserDeposit *UserDepositCallerSession) LsdTokenAddress() (common.Address, error) {
	return _UserDeposit.Contract.LsdTokenAddress(&_UserDeposit.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_UserDeposit *UserDepositCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserDeposit.contract.Call(opts, &out, "minDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_UserDeposit *UserDepositSession) MinDeposit() (*big.Int, error) {
	return _UserDeposit.Contract.MinDeposit(&_UserDeposit.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_UserDeposit *UserDepositCallerSession) MinDeposit() (*big.Int, error) {
	return _UserDeposit.Contract.MinDeposit(&_UserDeposit.CallOpts)
}

// NetworkBalancesAddress is a free data retrieval call binding the contract method 0x38fcf092.
//
// Solidity: function networkBalancesAddress() view returns(address)
func (_UserDeposit *UserDepositCaller) NetworkBalancesAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserDeposit.contract.Call(opts, &out, "networkBalancesAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkBalancesAddress is a free data retrieval call binding the contract method 0x38fcf092.
//
// Solidity: function networkBalancesAddress() view returns(address)
func (_UserDeposit *UserDepositSession) NetworkBalancesAddress() (common.Address, error) {
	return _UserDeposit.Contract.NetworkBalancesAddress(&_UserDeposit.CallOpts)
}

// NetworkBalancesAddress is a free data retrieval call binding the contract method 0x38fcf092.
//
// Solidity: function networkBalancesAddress() view returns(address)
func (_UserDeposit *UserDepositCallerSession) NetworkBalancesAddress() (common.Address, error) {
	return _UserDeposit.Contract.NetworkBalancesAddress(&_UserDeposit.CallOpts)
}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_UserDeposit *UserDepositCaller) NetworkProposalAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserDeposit.contract.Call(opts, &out, "networkProposalAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_UserDeposit *UserDepositSession) NetworkProposalAddress() (common.Address, error) {
	return _UserDeposit.Contract.NetworkProposalAddress(&_UserDeposit.CallOpts)
}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_UserDeposit *UserDepositCallerSession) NetworkProposalAddress() (common.Address, error) {
	return _UserDeposit.Contract.NetworkProposalAddress(&_UserDeposit.CallOpts)
}

// NetworkWithdrawAddress is a free data retrieval call binding the contract method 0x1587ef1b.
//
// Solidity: function networkWithdrawAddress() view returns(address)
func (_UserDeposit *UserDepositCaller) NetworkWithdrawAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserDeposit.contract.Call(opts, &out, "networkWithdrawAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkWithdrawAddress is a free data retrieval call binding the contract method 0x1587ef1b.
//
// Solidity: function networkWithdrawAddress() view returns(address)
func (_UserDeposit *UserDepositSession) NetworkWithdrawAddress() (common.Address, error) {
	return _UserDeposit.Contract.NetworkWithdrawAddress(&_UserDeposit.CallOpts)
}

// NetworkWithdrawAddress is a free data retrieval call binding the contract method 0x1587ef1b.
//
// Solidity: function networkWithdrawAddress() view returns(address)
func (_UserDeposit *UserDepositCallerSession) NetworkWithdrawAddress() (common.Address, error) {
	return _UserDeposit.Contract.NetworkWithdrawAddress(&_UserDeposit.CallOpts)
}

// NodeDepositAddress is a free data retrieval call binding the contract method 0xc39d2723.
//
// Solidity: function nodeDepositAddress() view returns(address)
func (_UserDeposit *UserDepositCaller) NodeDepositAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserDeposit.contract.Call(opts, &out, "nodeDepositAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeDepositAddress is a free data retrieval call binding the contract method 0xc39d2723.
//
// Solidity: function nodeDepositAddress() view returns(address)
func (_UserDeposit *UserDepositSession) NodeDepositAddress() (common.Address, error) {
	return _UserDeposit.Contract.NodeDepositAddress(&_UserDeposit.CallOpts)
}

// NodeDepositAddress is a free data retrieval call binding the contract method 0xc39d2723.
//
// Solidity: function nodeDepositAddress() view returns(address)
func (_UserDeposit *UserDepositCallerSession) NodeDepositAddress() (common.Address, error) {
	return _UserDeposit.Contract.NodeDepositAddress(&_UserDeposit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_UserDeposit *UserDepositCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _UserDeposit.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_UserDeposit *UserDepositSession) Version() (uint8, error) {
	return _UserDeposit.Contract.Version(&_UserDeposit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_UserDeposit *UserDepositCallerSession) Version() (uint8, error) {
	return _UserDeposit.Contract.Version(&_UserDeposit.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_UserDeposit *UserDepositTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDeposit.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_UserDeposit *UserDepositSession) Deposit() (*types.Transaction, error) {
	return _UserDeposit.Contract.Deposit(&_UserDeposit.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_UserDeposit *UserDepositTransactorSession) Deposit() (*types.Transaction, error) {
	return _UserDeposit.Contract.Deposit(&_UserDeposit.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _lsdTokenAddress, address _nodeDepositAddress, address _networkWithdrawAddress, address _networkProposalAddress, address _networkBalancesAddress) returns()
func (_UserDeposit *UserDepositTransactor) Init(opts *bind.TransactOpts, _lsdTokenAddress common.Address, _nodeDepositAddress common.Address, _networkWithdrawAddress common.Address, _networkProposalAddress common.Address, _networkBalancesAddress common.Address) (*types.Transaction, error) {
	return _UserDeposit.contract.Transact(opts, "init", _lsdTokenAddress, _nodeDepositAddress, _networkWithdrawAddress, _networkProposalAddress, _networkBalancesAddress)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _lsdTokenAddress, address _nodeDepositAddress, address _networkWithdrawAddress, address _networkProposalAddress, address _networkBalancesAddress) returns()
func (_UserDeposit *UserDepositSession) Init(_lsdTokenAddress common.Address, _nodeDepositAddress common.Address, _networkWithdrawAddress common.Address, _networkProposalAddress common.Address, _networkBalancesAddress common.Address) (*types.Transaction, error) {
	return _UserDeposit.Contract.Init(&_UserDeposit.TransactOpts, _lsdTokenAddress, _nodeDepositAddress, _networkWithdrawAddress, _networkProposalAddress, _networkBalancesAddress)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _lsdTokenAddress, address _nodeDepositAddress, address _networkWithdrawAddress, address _networkProposalAddress, address _networkBalancesAddress) returns()
func (_UserDeposit *UserDepositTransactorSession) Init(_lsdTokenAddress common.Address, _nodeDepositAddress common.Address, _networkWithdrawAddress common.Address, _networkProposalAddress common.Address, _networkBalancesAddress common.Address) (*types.Transaction, error) {
	return _UserDeposit.Contract.Init(&_UserDeposit.TransactOpts, _lsdTokenAddress, _nodeDepositAddress, _networkWithdrawAddress, _networkProposalAddress, _networkBalancesAddress)
}

// RecycleNetworkWithdrawDeposit is a paid mutator transaction binding the contract method 0x3d0062f8.
//
// Solidity: function recycleNetworkWithdrawDeposit() payable returns()
func (_UserDeposit *UserDepositTransactor) RecycleNetworkWithdrawDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDeposit.contract.Transact(opts, "recycleNetworkWithdrawDeposit")
}

// RecycleNetworkWithdrawDeposit is a paid mutator transaction binding the contract method 0x3d0062f8.
//
// Solidity: function recycleNetworkWithdrawDeposit() payable returns()
func (_UserDeposit *UserDepositSession) RecycleNetworkWithdrawDeposit() (*types.Transaction, error) {
	return _UserDeposit.Contract.RecycleNetworkWithdrawDeposit(&_UserDeposit.TransactOpts)
}

// RecycleNetworkWithdrawDeposit is a paid mutator transaction binding the contract method 0x3d0062f8.
//
// Solidity: function recycleNetworkWithdrawDeposit() payable returns()
func (_UserDeposit *UserDepositTransactorSession) RecycleNetworkWithdrawDeposit() (*types.Transaction, error) {
	return _UserDeposit.Contract.RecycleNetworkWithdrawDeposit(&_UserDeposit.TransactOpts)
}

// SetDepositEnabled is a paid mutator transaction binding the contract method 0x5b17d04b.
//
// Solidity: function setDepositEnabled(bool _value) returns()
func (_UserDeposit *UserDepositTransactor) SetDepositEnabled(opts *bind.TransactOpts, _value bool) (*types.Transaction, error) {
	return _UserDeposit.contract.Transact(opts, "setDepositEnabled", _value)
}

// SetDepositEnabled is a paid mutator transaction binding the contract method 0x5b17d04b.
//
// Solidity: function setDepositEnabled(bool _value) returns()
func (_UserDeposit *UserDepositSession) SetDepositEnabled(_value bool) (*types.Transaction, error) {
	return _UserDeposit.Contract.SetDepositEnabled(&_UserDeposit.TransactOpts, _value)
}

// SetDepositEnabled is a paid mutator transaction binding the contract method 0x5b17d04b.
//
// Solidity: function setDepositEnabled(bool _value) returns()
func (_UserDeposit *UserDepositTransactorSession) SetDepositEnabled(_value bool) (*types.Transaction, error) {
	return _UserDeposit.Contract.SetDepositEnabled(&_UserDeposit.TransactOpts, _value)
}

// SetMinimumDeposit is a paid mutator transaction binding the contract method 0xe78ec42e.
//
// Solidity: function setMinimumDeposit(uint256 _value) returns()
func (_UserDeposit *UserDepositTransactor) SetMinimumDeposit(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _UserDeposit.contract.Transact(opts, "setMinimumDeposit", _value)
}

// SetMinimumDeposit is a paid mutator transaction binding the contract method 0xe78ec42e.
//
// Solidity: function setMinimumDeposit(uint256 _value) returns()
func (_UserDeposit *UserDepositSession) SetMinimumDeposit(_value *big.Int) (*types.Transaction, error) {
	return _UserDeposit.Contract.SetMinimumDeposit(&_UserDeposit.TransactOpts, _value)
}

// SetMinimumDeposit is a paid mutator transaction binding the contract method 0xe78ec42e.
//
// Solidity: function setMinimumDeposit(uint256 _value) returns()
func (_UserDeposit *UserDepositTransactorSession) SetMinimumDeposit(_value *big.Int) (*types.Transaction, error) {
	return _UserDeposit.Contract.SetMinimumDeposit(&_UserDeposit.TransactOpts, _value)
}

// WithdrawExcessBalance is a paid mutator transaction binding the contract method 0x63a5db9e.
//
// Solidity: function withdrawExcessBalance(uint256 _amount) returns()
func (_UserDeposit *UserDepositTransactor) WithdrawExcessBalance(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _UserDeposit.contract.Transact(opts, "withdrawExcessBalance", _amount)
}

// WithdrawExcessBalance is a paid mutator transaction binding the contract method 0x63a5db9e.
//
// Solidity: function withdrawExcessBalance(uint256 _amount) returns()
func (_UserDeposit *UserDepositSession) WithdrawExcessBalance(_amount *big.Int) (*types.Transaction, error) {
	return _UserDeposit.Contract.WithdrawExcessBalance(&_UserDeposit.TransactOpts, _amount)
}

// WithdrawExcessBalance is a paid mutator transaction binding the contract method 0x63a5db9e.
//
// Solidity: function withdrawExcessBalance(uint256 _amount) returns()
func (_UserDeposit *UserDepositTransactorSession) WithdrawExcessBalance(_amount *big.Int) (*types.Transaction, error) {
	return _UserDeposit.Contract.WithdrawExcessBalance(&_UserDeposit.TransactOpts, _amount)
}

// UserDepositDepositReceivedIterator is returned from FilterDepositReceived and is used to iterate over the raw logs and unpacked data for DepositReceived events raised by the UserDeposit contract.
type UserDepositDepositReceivedIterator struct {
	Event *UserDepositDepositReceived // Event containing the contract specifics and raw log

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
func (it *UserDepositDepositReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserDepositDepositReceived)
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
		it.Event = new(UserDepositDepositReceived)
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
func (it *UserDepositDepositReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserDepositDepositReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserDepositDepositReceived represents a DepositReceived event raised by the UserDeposit contract.
type UserDepositDepositReceived struct {
	From   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositReceived is a free log retrieval operation binding the contract event 0x7aa1a8eb998c779420645fc14513bf058edb347d95c2fc2e6845bdc22f888631.
//
// Solidity: event DepositReceived(address indexed from, uint256 amount, uint256 time)
func (_UserDeposit *UserDepositFilterer) FilterDepositReceived(opts *bind.FilterOpts, from []common.Address) (*UserDepositDepositReceivedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _UserDeposit.contract.FilterLogs(opts, "DepositReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return &UserDepositDepositReceivedIterator{contract: _UserDeposit.contract, event: "DepositReceived", logs: logs, sub: sub}, nil
}

// WatchDepositReceived is a free log subscription operation binding the contract event 0x7aa1a8eb998c779420645fc14513bf058edb347d95c2fc2e6845bdc22f888631.
//
// Solidity: event DepositReceived(address indexed from, uint256 amount, uint256 time)
func (_UserDeposit *UserDepositFilterer) WatchDepositReceived(opts *bind.WatchOpts, sink chan<- *UserDepositDepositReceived, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _UserDeposit.contract.WatchLogs(opts, "DepositReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserDepositDepositReceived)
				if err := _UserDeposit.contract.UnpackLog(event, "DepositReceived", log); err != nil {
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

// ParseDepositReceived is a log parse operation binding the contract event 0x7aa1a8eb998c779420645fc14513bf058edb347d95c2fc2e6845bdc22f888631.
//
// Solidity: event DepositReceived(address indexed from, uint256 amount, uint256 time)
func (_UserDeposit *UserDepositFilterer) ParseDepositReceived(log types.Log) (*UserDepositDepositReceived, error) {
	event := new(UserDepositDepositReceived)
	if err := _UserDeposit.contract.UnpackLog(event, "DepositReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserDepositDepositRecycledIterator is returned from FilterDepositRecycled and is used to iterate over the raw logs and unpacked data for DepositRecycled events raised by the UserDeposit contract.
type UserDepositDepositRecycledIterator struct {
	Event *UserDepositDepositRecycled // Event containing the contract specifics and raw log

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
func (it *UserDepositDepositRecycledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserDepositDepositRecycled)
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
		it.Event = new(UserDepositDepositRecycled)
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
func (it *UserDepositDepositRecycledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserDepositDepositRecycledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserDepositDepositRecycled represents a DepositRecycled event raised by the UserDeposit contract.
type UserDepositDepositRecycled struct {
	From   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositRecycled is a free log retrieval operation binding the contract event 0x3a6614e80d02b57255cbb1f8305fbeca53d7e05a4b779d406279196608512925.
//
// Solidity: event DepositRecycled(address indexed from, uint256 amount, uint256 time)
func (_UserDeposit *UserDepositFilterer) FilterDepositRecycled(opts *bind.FilterOpts, from []common.Address) (*UserDepositDepositRecycledIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _UserDeposit.contract.FilterLogs(opts, "DepositRecycled", fromRule)
	if err != nil {
		return nil, err
	}
	return &UserDepositDepositRecycledIterator{contract: _UserDeposit.contract, event: "DepositRecycled", logs: logs, sub: sub}, nil
}

// WatchDepositRecycled is a free log subscription operation binding the contract event 0x3a6614e80d02b57255cbb1f8305fbeca53d7e05a4b779d406279196608512925.
//
// Solidity: event DepositRecycled(address indexed from, uint256 amount, uint256 time)
func (_UserDeposit *UserDepositFilterer) WatchDepositRecycled(opts *bind.WatchOpts, sink chan<- *UserDepositDepositRecycled, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _UserDeposit.contract.WatchLogs(opts, "DepositRecycled", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserDepositDepositRecycled)
				if err := _UserDeposit.contract.UnpackLog(event, "DepositRecycled", log); err != nil {
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

// ParseDepositRecycled is a log parse operation binding the contract event 0x3a6614e80d02b57255cbb1f8305fbeca53d7e05a4b779d406279196608512925.
//
// Solidity: event DepositRecycled(address indexed from, uint256 amount, uint256 time)
func (_UserDeposit *UserDepositFilterer) ParseDepositRecycled(log types.Log) (*UserDepositDepositRecycled, error) {
	event := new(UserDepositDepositRecycled)
	if err := _UserDeposit.contract.UnpackLog(event, "DepositRecycled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserDepositExcessWithdrawnIterator is returned from FilterExcessWithdrawn and is used to iterate over the raw logs and unpacked data for ExcessWithdrawn events raised by the UserDeposit contract.
type UserDepositExcessWithdrawnIterator struct {
	Event *UserDepositExcessWithdrawn // Event containing the contract specifics and raw log

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
func (it *UserDepositExcessWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserDepositExcessWithdrawn)
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
		it.Event = new(UserDepositExcessWithdrawn)
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
func (it *UserDepositExcessWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserDepositExcessWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserDepositExcessWithdrawn represents a ExcessWithdrawn event raised by the UserDeposit contract.
type UserDepositExcessWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExcessWithdrawn is a free log retrieval operation binding the contract event 0x992f462cfb62e164bd03bf07baf2cffce83fbd9370cae10635842b2020012120.
//
// Solidity: event ExcessWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_UserDeposit *UserDepositFilterer) FilterExcessWithdrawn(opts *bind.FilterOpts, to []common.Address) (*UserDepositExcessWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UserDeposit.contract.FilterLogs(opts, "ExcessWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &UserDepositExcessWithdrawnIterator{contract: _UserDeposit.contract, event: "ExcessWithdrawn", logs: logs, sub: sub}, nil
}

// WatchExcessWithdrawn is a free log subscription operation binding the contract event 0x992f462cfb62e164bd03bf07baf2cffce83fbd9370cae10635842b2020012120.
//
// Solidity: event ExcessWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_UserDeposit *UserDepositFilterer) WatchExcessWithdrawn(opts *bind.WatchOpts, sink chan<- *UserDepositExcessWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UserDeposit.contract.WatchLogs(opts, "ExcessWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserDepositExcessWithdrawn)
				if err := _UserDeposit.contract.UnpackLog(event, "ExcessWithdrawn", log); err != nil {
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

// ParseExcessWithdrawn is a log parse operation binding the contract event 0x992f462cfb62e164bd03bf07baf2cffce83fbd9370cae10635842b2020012120.
//
// Solidity: event ExcessWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_UserDeposit *UserDepositFilterer) ParseExcessWithdrawn(log types.Log) (*UserDepositExcessWithdrawn, error) {
	event := new(UserDepositExcessWithdrawn)
	if err := _UserDeposit.contract.UnpackLog(event, "ExcessWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
