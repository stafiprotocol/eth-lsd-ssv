// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package network_proposal

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

// NetworkProposalMetaData contains all meta data concerning the NetworkProposal contract.
var NetworkProposalMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalId\",\"type\":\"bytes32\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VoteProposal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposalId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_initialThreshold\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_adminAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"enumINetworkProposal.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"_yesVotes\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"_yesVotesTotal\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"removeVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposalId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"shouldExecute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NetworkProposalABI is the input ABI used to generate the binding from.
// Deprecated: Use NetworkProposalMetaData.ABI instead.
var NetworkProposalABI = NetworkProposalMetaData.ABI

// NetworkProposal is an auto generated Go binding around an Ethereum contract.
type NetworkProposal struct {
	NetworkProposalCaller     // Read-only binding to the contract
	NetworkProposalTransactor // Write-only binding to the contract
	NetworkProposalFilterer   // Log filterer for contract events
}

// NetworkProposalCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkProposalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkProposalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkProposalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkProposalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NetworkProposalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkProposalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkProposalSession struct {
	Contract     *NetworkProposal  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NetworkProposalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkProposalCallerSession struct {
	Contract *NetworkProposalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NetworkProposalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkProposalTransactorSession struct {
	Contract     *NetworkProposalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NetworkProposalRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkProposalRaw struct {
	Contract *NetworkProposal // Generic contract binding to access the raw methods on
}

// NetworkProposalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkProposalCallerRaw struct {
	Contract *NetworkProposalCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkProposalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkProposalTransactorRaw struct {
	Contract *NetworkProposalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetworkProposal creates a new instance of NetworkProposal, bound to a specific deployed contract.
func NewNetworkProposal(address common.Address, backend bind.ContractBackend) (*NetworkProposal, error) {
	contract, err := bindNetworkProposal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NetworkProposal{NetworkProposalCaller: NetworkProposalCaller{contract: contract}, NetworkProposalTransactor: NetworkProposalTransactor{contract: contract}, NetworkProposalFilterer: NetworkProposalFilterer{contract: contract}}, nil
}

// NewNetworkProposalCaller creates a new read-only instance of NetworkProposal, bound to a specific deployed contract.
func NewNetworkProposalCaller(address common.Address, caller bind.ContractCaller) (*NetworkProposalCaller, error) {
	contract, err := bindNetworkProposal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkProposalCaller{contract: contract}, nil
}

// NewNetworkProposalTransactor creates a new write-only instance of NetworkProposal, bound to a specific deployed contract.
func NewNetworkProposalTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkProposalTransactor, error) {
	contract, err := bindNetworkProposal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkProposalTransactor{contract: contract}, nil
}

// NewNetworkProposalFilterer creates a new log filterer instance of NetworkProposal, bound to a specific deployed contract.
func NewNetworkProposalFilterer(address common.Address, filterer bind.ContractFilterer) (*NetworkProposalFilterer, error) {
	contract, err := bindNetworkProposal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NetworkProposalFilterer{contract: contract}, nil
}

// bindNetworkProposal binds a generic wrapper to an already deployed contract.
func bindNetworkProposal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetworkProposalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkProposal *NetworkProposalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkProposal.Contract.NetworkProposalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkProposal *NetworkProposalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkProposal.Contract.NetworkProposalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkProposal *NetworkProposalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkProposal.Contract.NetworkProposalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkProposal *NetworkProposalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkProposal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkProposal *NetworkProposalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkProposal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkProposal *NetworkProposalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkProposal.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_NetworkProposal *NetworkProposalCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NetworkProposal.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_NetworkProposal *NetworkProposalSession) Admin() (common.Address, error) {
	return _NetworkProposal.Contract.Admin(&_NetworkProposal.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_NetworkProposal *NetworkProposalCallerSession) Admin() (common.Address, error) {
	return _NetworkProposal.Contract.Admin(&_NetworkProposal.CallOpts)
}

// GetVoterIndex is a free data retrieval call binding the contract method 0x7941743a.
//
// Solidity: function getVoterIndex(address _voter) view returns(uint256)
func (_NetworkProposal *NetworkProposalCaller) GetVoterIndex(opts *bind.CallOpts, _voter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NetworkProposal.contract.Call(opts, &out, "getVoterIndex", _voter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoterIndex is a free data retrieval call binding the contract method 0x7941743a.
//
// Solidity: function getVoterIndex(address _voter) view returns(uint256)
func (_NetworkProposal *NetworkProposalSession) GetVoterIndex(_voter common.Address) (*big.Int, error) {
	return _NetworkProposal.Contract.GetVoterIndex(&_NetworkProposal.CallOpts, _voter)
}

// GetVoterIndex is a free data retrieval call binding the contract method 0x7941743a.
//
// Solidity: function getVoterIndex(address _voter) view returns(uint256)
func (_NetworkProposal *NetworkProposalCallerSession) GetVoterIndex(_voter common.Address) (*big.Int, error) {
	return _NetworkProposal.Contract.GetVoterIndex(&_NetworkProposal.CallOpts, _voter)
}

// HasVoted is a free data retrieval call binding the contract method 0xaadc3b72.
//
// Solidity: function hasVoted(bytes32 _proposalId, address _voter) view returns(bool)
func (_NetworkProposal *NetworkProposalCaller) HasVoted(opts *bind.CallOpts, _proposalId [32]byte, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _NetworkProposal.contract.Call(opts, &out, "hasVoted", _proposalId, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0xaadc3b72.
//
// Solidity: function hasVoted(bytes32 _proposalId, address _voter) view returns(bool)
func (_NetworkProposal *NetworkProposalSession) HasVoted(_proposalId [32]byte, _voter common.Address) (bool, error) {
	return _NetworkProposal.Contract.HasVoted(&_NetworkProposal.CallOpts, _proposalId, _voter)
}

// HasVoted is a free data retrieval call binding the contract method 0xaadc3b72.
//
// Solidity: function hasVoted(bytes32 _proposalId, address _voter) view returns(bool)
func (_NetworkProposal *NetworkProposalCallerSession) HasVoted(_proposalId [32]byte, _voter common.Address) (bool, error) {
	return _NetworkProposal.Contract.HasVoted(&_NetworkProposal.CallOpts, _proposalId, _voter)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_NetworkProposal *NetworkProposalCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NetworkProposal.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_NetworkProposal *NetworkProposalSession) Initialized() (bool, error) {
	return _NetworkProposal.Contract.Initialized(&_NetworkProposal.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_NetworkProposal *NetworkProposalCallerSession) Initialized() (bool, error) {
	return _NetworkProposal.Contract.Initialized(&_NetworkProposal.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _sender) view returns(bool)
func (_NetworkProposal *NetworkProposalCaller) IsAdmin(opts *bind.CallOpts, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _NetworkProposal.contract.Call(opts, &out, "isAdmin", _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _sender) view returns(bool)
func (_NetworkProposal *NetworkProposalSession) IsAdmin(_sender common.Address) (bool, error) {
	return _NetworkProposal.Contract.IsAdmin(&_NetworkProposal.CallOpts, _sender)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _sender) view returns(bool)
func (_NetworkProposal *NetworkProposalCallerSession) IsAdmin(_sender common.Address) (bool, error) {
	return _NetworkProposal.Contract.IsAdmin(&_NetworkProposal.CallOpts, _sender)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address _sender) view returns(bool)
func (_NetworkProposal *NetworkProposalCaller) IsVoter(opts *bind.CallOpts, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _NetworkProposal.contract.Call(opts, &out, "isVoter", _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address _sender) view returns(bool)
func (_NetworkProposal *NetworkProposalSession) IsVoter(_sender common.Address) (bool, error) {
	return _NetworkProposal.Contract.IsVoter(&_NetworkProposal.CallOpts, _sender)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address _sender) view returns(bool)
func (_NetworkProposal *NetworkProposalCallerSession) IsVoter(_sender common.Address) (bool, error) {
	return _NetworkProposal.Contract.IsVoter(&_NetworkProposal.CallOpts, _sender)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(uint8 _status, uint16 _yesVotes, uint8 _yesVotesTotal)
func (_NetworkProposal *NetworkProposalCaller) Proposals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status        uint8
	YesVotes      uint16
	YesVotesTotal uint8
}, error) {
	var out []interface{}
	err := _NetworkProposal.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Status        uint8
		YesVotes      uint16
		YesVotesTotal uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.YesVotes = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.YesVotesTotal = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(uint8 _status, uint16 _yesVotes, uint8 _yesVotesTotal)
func (_NetworkProposal *NetworkProposalSession) Proposals(arg0 [32]byte) (struct {
	Status        uint8
	YesVotes      uint16
	YesVotesTotal uint8
}, error) {
	return _NetworkProposal.Contract.Proposals(&_NetworkProposal.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(uint8 _status, uint16 _yesVotes, uint8 _yesVotesTotal)
func (_NetworkProposal *NetworkProposalCallerSession) Proposals(arg0 [32]byte) (struct {
	Status        uint8
	YesVotes      uint16
	YesVotesTotal uint8
}, error) {
	return _NetworkProposal.Contract.Proposals(&_NetworkProposal.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_NetworkProposal *NetworkProposalCaller) Threshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NetworkProposal.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_NetworkProposal *NetworkProposalSession) Threshold() (uint8, error) {
	return _NetworkProposal.Contract.Threshold(&_NetworkProposal.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint8)
func (_NetworkProposal *NetworkProposalCallerSession) Threshold() (uint8, error) {
	return _NetworkProposal.Contract.Threshold(&_NetworkProposal.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkProposal *NetworkProposalCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NetworkProposal.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkProposal *NetworkProposalSession) Version() (uint8, error) {
	return _NetworkProposal.Contract.Version(&_NetworkProposal.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkProposal *NetworkProposalCallerSession) Version() (uint8, error) {
	return _NetworkProposal.Contract.Version(&_NetworkProposal.CallOpts)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(address _voter) returns()
func (_NetworkProposal *NetworkProposalTransactor) AddVoter(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "addVoter", _voter)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(address _voter) returns()
func (_NetworkProposal *NetworkProposalSession) AddVoter(_voter common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.AddVoter(&_NetworkProposal.TransactOpts, _voter)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(address _voter) returns()
func (_NetworkProposal *NetworkProposalTransactorSession) AddVoter(_voter common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.AddVoter(&_NetworkProposal.TransactOpts, _voter)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _newThreshold) returns()
func (_NetworkProposal *NetworkProposalTransactor) ChangeThreshold(opts *bind.TransactOpts, _newThreshold *big.Int) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "changeThreshold", _newThreshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _newThreshold) returns()
func (_NetworkProposal *NetworkProposalSession) ChangeThreshold(_newThreshold *big.Int) (*types.Transaction, error) {
	return _NetworkProposal.Contract.ChangeThreshold(&_NetworkProposal.TransactOpts, _newThreshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _newThreshold) returns()
func (_NetworkProposal *NetworkProposalTransactorSession) ChangeThreshold(_newThreshold *big.Int) (*types.Transaction, error) {
	return _NetworkProposal.Contract.ChangeThreshold(&_NetworkProposal.TransactOpts, _newThreshold)
}

// Init is a paid mutator transaction binding the contract method 0xba50db3b.
//
// Solidity: function init(address[] _voters, uint256 _initialThreshold, address _adminAddress) returns()
func (_NetworkProposal *NetworkProposalTransactor) Init(opts *bind.TransactOpts, _voters []common.Address, _initialThreshold *big.Int, _adminAddress common.Address) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "init", _voters, _initialThreshold, _adminAddress)
}

// Init is a paid mutator transaction binding the contract method 0xba50db3b.
//
// Solidity: function init(address[] _voters, uint256 _initialThreshold, address _adminAddress) returns()
func (_NetworkProposal *NetworkProposalSession) Init(_voters []common.Address, _initialThreshold *big.Int, _adminAddress common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.Init(&_NetworkProposal.TransactOpts, _voters, _initialThreshold, _adminAddress)
}

// Init is a paid mutator transaction binding the contract method 0xba50db3b.
//
// Solidity: function init(address[] _voters, uint256 _initialThreshold, address _adminAddress) returns()
func (_NetworkProposal *NetworkProposalTransactorSession) Init(_voters []common.Address, _initialThreshold *big.Int, _adminAddress common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.Init(&_NetworkProposal.TransactOpts, _voters, _initialThreshold, _adminAddress)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(address _voter) returns()
func (_NetworkProposal *NetworkProposalTransactor) RemoveVoter(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "removeVoter", _voter)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(address _voter) returns()
func (_NetworkProposal *NetworkProposalSession) RemoveVoter(_voter common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.RemoveVoter(&_NetworkProposal.TransactOpts, _voter)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(address _voter) returns()
func (_NetworkProposal *NetworkProposalTransactorSession) RemoveVoter(_voter common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.RemoveVoter(&_NetworkProposal.TransactOpts, _voter)
}

// ShouldExecute is a paid mutator transaction binding the contract method 0xcd23ef14.
//
// Solidity: function shouldExecute(bytes32 _proposalId, address _voter) returns(bool)
func (_NetworkProposal *NetworkProposalTransactor) ShouldExecute(opts *bind.TransactOpts, _proposalId [32]byte, _voter common.Address) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "shouldExecute", _proposalId, _voter)
}

// ShouldExecute is a paid mutator transaction binding the contract method 0xcd23ef14.
//
// Solidity: function shouldExecute(bytes32 _proposalId, address _voter) returns(bool)
func (_NetworkProposal *NetworkProposalSession) ShouldExecute(_proposalId [32]byte, _voter common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.ShouldExecute(&_NetworkProposal.TransactOpts, _proposalId, _voter)
}

// ShouldExecute is a paid mutator transaction binding the contract method 0xcd23ef14.
//
// Solidity: function shouldExecute(bytes32 _proposalId, address _voter) returns(bool)
func (_NetworkProposal *NetworkProposalTransactorSession) ShouldExecute(_proposalId [32]byte, _voter common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.ShouldExecute(&_NetworkProposal.TransactOpts, _proposalId, _voter)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address _newAdmin) returns()
func (_NetworkProposal *NetworkProposalTransactor) TransferAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "transferAdmin", _newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address _newAdmin) returns()
func (_NetworkProposal *NetworkProposalSession) TransferAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.TransferAdmin(&_NetworkProposal.TransactOpts, _newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address _newAdmin) returns()
func (_NetworkProposal *NetworkProposalTransactorSession) TransferAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.TransferAdmin(&_NetworkProposal.TransactOpts, _newAdmin)
}

// NetworkProposalProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the NetworkProposal contract.
type NetworkProposalProposalExecutedIterator struct {
	Event *NetworkProposalProposalExecuted // Event containing the contract specifics and raw log

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
func (it *NetworkProposalProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkProposalProposalExecuted)
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
		it.Event = new(NetworkProposalProposalExecuted)
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
func (it *NetworkProposalProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkProposalProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkProposalProposalExecuted represents a ProposalExecuted event raised by the NetworkProposal contract.
type NetworkProposalProposalExecuted struct {
	ProposalId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalId)
func (_NetworkProposal *NetworkProposalFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalId [][32]byte) (*NetworkProposalProposalExecutedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _NetworkProposal.contract.FilterLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &NetworkProposalProposalExecutedIterator{contract: _NetworkProposal.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalId)
func (_NetworkProposal *NetworkProposalFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *NetworkProposalProposalExecuted, proposalId [][32]byte) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _NetworkProposal.contract.WatchLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkProposalProposalExecuted)
				if err := _NetworkProposal.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalId)
func (_NetworkProposal *NetworkProposalFilterer) ParseProposalExecuted(log types.Log) (*NetworkProposalProposalExecuted, error) {
	event := new(NetworkProposalProposalExecuted)
	if err := _NetworkProposal.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkProposalVoteProposalIterator is returned from FilterVoteProposal and is used to iterate over the raw logs and unpacked data for VoteProposal events raised by the NetworkProposal contract.
type NetworkProposalVoteProposalIterator struct {
	Event *NetworkProposalVoteProposal // Event containing the contract specifics and raw log

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
func (it *NetworkProposalVoteProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkProposalVoteProposal)
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
		it.Event = new(NetworkProposalVoteProposal)
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
func (it *NetworkProposalVoteProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkProposalVoteProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkProposalVoteProposal represents a VoteProposal event raised by the NetworkProposal contract.
type NetworkProposalVoteProposal struct {
	ProposalId [32]byte
	Voter      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteProposal is a free log retrieval operation binding the contract event 0x3b58f01618556cdc5e9f7b0f1f6dccbac40024bc1043f589bd4a324e3414cfad.
//
// Solidity: event VoteProposal(bytes32 indexed proposalId, address voter)
func (_NetworkProposal *NetworkProposalFilterer) FilterVoteProposal(opts *bind.FilterOpts, proposalId [][32]byte) (*NetworkProposalVoteProposalIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _NetworkProposal.contract.FilterLogs(opts, "VoteProposal", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &NetworkProposalVoteProposalIterator{contract: _NetworkProposal.contract, event: "VoteProposal", logs: logs, sub: sub}, nil
}

// WatchVoteProposal is a free log subscription operation binding the contract event 0x3b58f01618556cdc5e9f7b0f1f6dccbac40024bc1043f589bd4a324e3414cfad.
//
// Solidity: event VoteProposal(bytes32 indexed proposalId, address voter)
func (_NetworkProposal *NetworkProposalFilterer) WatchVoteProposal(opts *bind.WatchOpts, sink chan<- *NetworkProposalVoteProposal, proposalId [][32]byte) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _NetworkProposal.contract.WatchLogs(opts, "VoteProposal", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkProposalVoteProposal)
				if err := _NetworkProposal.contract.UnpackLog(event, "VoteProposal", log); err != nil {
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

// ParseVoteProposal is a log parse operation binding the contract event 0x3b58f01618556cdc5e9f7b0f1f6dccbac40024bc1043f589bd4a324e3414cfad.
//
// Solidity: event VoteProposal(bytes32 indexed proposalId, address voter)
func (_NetworkProposal *NetworkProposalFilterer) ParseVoteProposal(log types.Log) (*NetworkProposalVoteProposal, error) {
	event := new(NetworkProposalVoteProposal)
	if err := _NetworkProposal.contract.UnpackLog(event, "VoteProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
