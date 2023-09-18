// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package node_deposit

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

// NodeDepositMetaData contains all meta data concerning the NodeDeposit contract.
var NodeDepositMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumINodeDeposit.NodeType\",\"name\":\"nodeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"OffBoarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"enumINodeDeposit.PubkeyStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"SetPubkeyStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustNodeAddress\",\"type\":\"address\"}],\"name\":\"addTrustNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_validatorPubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_validatorSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDepositAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"getPubkeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeyList\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPubkeysLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userDepositAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethDepositAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkProposalAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_withdrawCredentials\",\"type\":\"bytes\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lightNodeDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lightNodeDepositEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkProposalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeInfoOf\",\"outputs\":[{\"internalType\":\"enumINodeDeposit.NodeType\",\"name\":\"_nodeType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_removed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_pubkeyNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pubkeyInfoOf\",\"outputs\":[{\"internalType\":\"enumINodeDeposit.PubkeyStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nodeDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_depositBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_depositSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pubkeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustNodeAddress\",\"type\":\"address\"}],\"name\":\"removeTrustNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setLightNodeDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setLightNodeDepositEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"},{\"internalType\":\"enumINodeDeposit.PubkeyStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setNodePubkeyStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setTrustNodeDepositEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setTrustNodePubkeyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_withdrawCredentials\",\"type\":\"bytes\"}],\"name\":\"setWithdrawCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_validatorPubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_validatorSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustNodeDepositEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustNodePubkeyNumberLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDepositAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bool[]\",\"name\":\"_matchs\",\"type\":\"bool[]\"}],\"name\":\"voteWithdrawCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawCredentials\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NodeDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeDepositMetaData.ABI instead.
var NodeDepositABI = NodeDepositMetaData.ABI

// NodeDeposit is an auto generated Go binding around an Ethereum contract.
type NodeDeposit struct {
	NodeDepositCaller     // Read-only binding to the contract
	NodeDepositTransactor // Write-only binding to the contract
	NodeDepositFilterer   // Log filterer for contract events
}

// NodeDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeDepositSession struct {
	Contract     *NodeDeposit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeDepositCallerSession struct {
	Contract *NodeDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeDepositTransactorSession struct {
	Contract     *NodeDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeDepositRaw struct {
	Contract *NodeDeposit // Generic contract binding to access the raw methods on
}

// NodeDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeDepositCallerRaw struct {
	Contract *NodeDepositCaller // Generic read-only contract binding to access the raw methods on
}

// NodeDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeDepositTransactorRaw struct {
	Contract *NodeDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeDeposit creates a new instance of NodeDeposit, bound to a specific deployed contract.
func NewNodeDeposit(address common.Address, backend bind.ContractBackend) (*NodeDeposit, error) {
	contract, err := bindNodeDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeDeposit{NodeDepositCaller: NodeDepositCaller{contract: contract}, NodeDepositTransactor: NodeDepositTransactor{contract: contract}, NodeDepositFilterer: NodeDepositFilterer{contract: contract}}, nil
}

// NewNodeDepositCaller creates a new read-only instance of NodeDeposit, bound to a specific deployed contract.
func NewNodeDepositCaller(address common.Address, caller bind.ContractCaller) (*NodeDepositCaller, error) {
	contract, err := bindNodeDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeDepositCaller{contract: contract}, nil
}

// NewNodeDepositTransactor creates a new write-only instance of NodeDeposit, bound to a specific deployed contract.
func NewNodeDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeDepositTransactor, error) {
	contract, err := bindNodeDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeDepositTransactor{contract: contract}, nil
}

// NewNodeDepositFilterer creates a new log filterer instance of NodeDeposit, bound to a specific deployed contract.
func NewNodeDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeDepositFilterer, error) {
	contract, err := bindNodeDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeDepositFilterer{contract: contract}, nil
}

// bindNodeDeposit binds a generic wrapper to an already deployed contract.
func bindNodeDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeDepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeDeposit *NodeDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeDeposit.Contract.NodeDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeDeposit *NodeDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeDeposit.Contract.NodeDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeDeposit *NodeDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeDeposit.Contract.NodeDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeDeposit *NodeDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeDeposit *NodeDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeDeposit *NodeDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeDeposit.Contract.contract.Transact(opts, method, params...)
}

// EthDepositAddress is a free data retrieval call binding the contract method 0xb420feb2.
//
// Solidity: function ethDepositAddress() view returns(address)
func (_NodeDeposit *NodeDepositCaller) EthDepositAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "ethDepositAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthDepositAddress is a free data retrieval call binding the contract method 0xb420feb2.
//
// Solidity: function ethDepositAddress() view returns(address)
func (_NodeDeposit *NodeDepositSession) EthDepositAddress() (common.Address, error) {
	return _NodeDeposit.Contract.EthDepositAddress(&_NodeDeposit.CallOpts)
}

// EthDepositAddress is a free data retrieval call binding the contract method 0xb420feb2.
//
// Solidity: function ethDepositAddress() view returns(address)
func (_NodeDeposit *NodeDepositCallerSession) EthDepositAddress() (common.Address, error) {
	return _NodeDeposit.Contract.EthDepositAddress(&_NodeDeposit.CallOpts)
}

// GetPubkeys is a free data retrieval call binding the contract method 0xa57fd0f7.
//
// Solidity: function getPubkeys(uint256 _start, uint256 _end) view returns(bytes[] pubkeyList)
func (_NodeDeposit *NodeDepositCaller) GetPubkeys(opts *bind.CallOpts, _start *big.Int, _end *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "getPubkeys", _start, _end)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetPubkeys is a free data retrieval call binding the contract method 0xa57fd0f7.
//
// Solidity: function getPubkeys(uint256 _start, uint256 _end) view returns(bytes[] pubkeyList)
func (_NodeDeposit *NodeDepositSession) GetPubkeys(_start *big.Int, _end *big.Int) ([][]byte, error) {
	return _NodeDeposit.Contract.GetPubkeys(&_NodeDeposit.CallOpts, _start, _end)
}

// GetPubkeys is a free data retrieval call binding the contract method 0xa57fd0f7.
//
// Solidity: function getPubkeys(uint256 _start, uint256 _end) view returns(bytes[] pubkeyList)
func (_NodeDeposit *NodeDepositCallerSession) GetPubkeys(_start *big.Int, _end *big.Int) ([][]byte, error) {
	return _NodeDeposit.Contract.GetPubkeys(&_NodeDeposit.CallOpts, _start, _end)
}

// GetPubkeysLength is a free data retrieval call binding the contract method 0x6fe2d088.
//
// Solidity: function getPubkeysLength() view returns(uint256)
func (_NodeDeposit *NodeDepositCaller) GetPubkeysLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "getPubkeysLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPubkeysLength is a free data retrieval call binding the contract method 0x6fe2d088.
//
// Solidity: function getPubkeysLength() view returns(uint256)
func (_NodeDeposit *NodeDepositSession) GetPubkeysLength() (*big.Int, error) {
	return _NodeDeposit.Contract.GetPubkeysLength(&_NodeDeposit.CallOpts)
}

// GetPubkeysLength is a free data retrieval call binding the contract method 0x6fe2d088.
//
// Solidity: function getPubkeysLength() view returns(uint256)
func (_NodeDeposit *NodeDepositCallerSession) GetPubkeysLength() (*big.Int, error) {
	return _NodeDeposit.Contract.GetPubkeysLength(&_NodeDeposit.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_NodeDeposit *NodeDepositCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_NodeDeposit *NodeDepositSession) Initialized() (bool, error) {
	return _NodeDeposit.Contract.Initialized(&_NodeDeposit.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_NodeDeposit *NodeDepositCallerSession) Initialized() (bool, error) {
	return _NodeDeposit.Contract.Initialized(&_NodeDeposit.CallOpts)
}

// LightNodeDepositAmount is a free data retrieval call binding the contract method 0x0e96b92e.
//
// Solidity: function lightNodeDepositAmount() view returns(uint256)
func (_NodeDeposit *NodeDepositCaller) LightNodeDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "lightNodeDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LightNodeDepositAmount is a free data retrieval call binding the contract method 0x0e96b92e.
//
// Solidity: function lightNodeDepositAmount() view returns(uint256)
func (_NodeDeposit *NodeDepositSession) LightNodeDepositAmount() (*big.Int, error) {
	return _NodeDeposit.Contract.LightNodeDepositAmount(&_NodeDeposit.CallOpts)
}

// LightNodeDepositAmount is a free data retrieval call binding the contract method 0x0e96b92e.
//
// Solidity: function lightNodeDepositAmount() view returns(uint256)
func (_NodeDeposit *NodeDepositCallerSession) LightNodeDepositAmount() (*big.Int, error) {
	return _NodeDeposit.Contract.LightNodeDepositAmount(&_NodeDeposit.CallOpts)
}

// LightNodeDepositEnabled is a free data retrieval call binding the contract method 0xb25f1fd6.
//
// Solidity: function lightNodeDepositEnabled() view returns(bool)
func (_NodeDeposit *NodeDepositCaller) LightNodeDepositEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "lightNodeDepositEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LightNodeDepositEnabled is a free data retrieval call binding the contract method 0xb25f1fd6.
//
// Solidity: function lightNodeDepositEnabled() view returns(bool)
func (_NodeDeposit *NodeDepositSession) LightNodeDepositEnabled() (bool, error) {
	return _NodeDeposit.Contract.LightNodeDepositEnabled(&_NodeDeposit.CallOpts)
}

// LightNodeDepositEnabled is a free data retrieval call binding the contract method 0xb25f1fd6.
//
// Solidity: function lightNodeDepositEnabled() view returns(bool)
func (_NodeDeposit *NodeDepositCallerSession) LightNodeDepositEnabled() (bool, error) {
	return _NodeDeposit.Contract.LightNodeDepositEnabled(&_NodeDeposit.CallOpts)
}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_NodeDeposit *NodeDepositCaller) NetworkProposalAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "networkProposalAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_NodeDeposit *NodeDepositSession) NetworkProposalAddress() (common.Address, error) {
	return _NodeDeposit.Contract.NetworkProposalAddress(&_NodeDeposit.CallOpts)
}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_NodeDeposit *NodeDepositCallerSession) NetworkProposalAddress() (common.Address, error) {
	return _NodeDeposit.Contract.NetworkProposalAddress(&_NodeDeposit.CallOpts)
}

// NodeInfoOf is a free data retrieval call binding the contract method 0x0fdefd36.
//
// Solidity: function nodeInfoOf(address ) view returns(uint8 _nodeType, bool _removed, uint256 _pubkeyNumber)
func (_NodeDeposit *NodeDepositCaller) NodeInfoOf(opts *bind.CallOpts, arg0 common.Address) (struct {
	NodeType     uint8
	Removed      bool
	PubkeyNumber *big.Int
}, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "nodeInfoOf", arg0)

	outstruct := new(struct {
		NodeType     uint8
		Removed      bool
		PubkeyNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NodeType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Removed = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.PubkeyNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NodeInfoOf is a free data retrieval call binding the contract method 0x0fdefd36.
//
// Solidity: function nodeInfoOf(address ) view returns(uint8 _nodeType, bool _removed, uint256 _pubkeyNumber)
func (_NodeDeposit *NodeDepositSession) NodeInfoOf(arg0 common.Address) (struct {
	NodeType     uint8
	Removed      bool
	PubkeyNumber *big.Int
}, error) {
	return _NodeDeposit.Contract.NodeInfoOf(&_NodeDeposit.CallOpts, arg0)
}

// NodeInfoOf is a free data retrieval call binding the contract method 0x0fdefd36.
//
// Solidity: function nodeInfoOf(address ) view returns(uint8 _nodeType, bool _removed, uint256 _pubkeyNumber)
func (_NodeDeposit *NodeDepositCallerSession) NodeInfoOf(arg0 common.Address) (struct {
	NodeType     uint8
	Removed      bool
	PubkeyNumber *big.Int
}, error) {
	return _NodeDeposit.Contract.NodeInfoOf(&_NodeDeposit.CallOpts, arg0)
}

// PubkeyInfoOf is a free data retrieval call binding the contract method 0xd6c9f4c0.
//
// Solidity: function pubkeyInfoOf(bytes ) view returns(uint8 _status, address _owner, uint256 _nodeDepositAmount, uint256 _depositBlock, bytes _depositSignature)
func (_NodeDeposit *NodeDepositCaller) PubkeyInfoOf(opts *bind.CallOpts, arg0 []byte) (struct {
	Status            uint8
	Owner             common.Address
	NodeDepositAmount *big.Int
	DepositBlock      *big.Int
	DepositSignature  []byte
}, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "pubkeyInfoOf", arg0)

	outstruct := new(struct {
		Status            uint8
		Owner             common.Address
		NodeDepositAmount *big.Int
		DepositBlock      *big.Int
		DepositSignature  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NodeDepositAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DepositBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DepositSignature = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// PubkeyInfoOf is a free data retrieval call binding the contract method 0xd6c9f4c0.
//
// Solidity: function pubkeyInfoOf(bytes ) view returns(uint8 _status, address _owner, uint256 _nodeDepositAmount, uint256 _depositBlock, bytes _depositSignature)
func (_NodeDeposit *NodeDepositSession) PubkeyInfoOf(arg0 []byte) (struct {
	Status            uint8
	Owner             common.Address
	NodeDepositAmount *big.Int
	DepositBlock      *big.Int
	DepositSignature  []byte
}, error) {
	return _NodeDeposit.Contract.PubkeyInfoOf(&_NodeDeposit.CallOpts, arg0)
}

// PubkeyInfoOf is a free data retrieval call binding the contract method 0xd6c9f4c0.
//
// Solidity: function pubkeyInfoOf(bytes ) view returns(uint8 _status, address _owner, uint256 _nodeDepositAmount, uint256 _depositBlock, bytes _depositSignature)
func (_NodeDeposit *NodeDepositCallerSession) PubkeyInfoOf(arg0 []byte) (struct {
	Status            uint8
	Owner             common.Address
	NodeDepositAmount *big.Int
	DepositBlock      *big.Int
	DepositSignature  []byte
}, error) {
	return _NodeDeposit.Contract.PubkeyInfoOf(&_NodeDeposit.CallOpts, arg0)
}

// Pubkeys is a free data retrieval call binding the contract method 0x7bd3e3f6.
//
// Solidity: function pubkeys(uint256 ) view returns(bytes)
func (_NodeDeposit *NodeDepositCaller) Pubkeys(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "pubkeys", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Pubkeys is a free data retrieval call binding the contract method 0x7bd3e3f6.
//
// Solidity: function pubkeys(uint256 ) view returns(bytes)
func (_NodeDeposit *NodeDepositSession) Pubkeys(arg0 *big.Int) ([]byte, error) {
	return _NodeDeposit.Contract.Pubkeys(&_NodeDeposit.CallOpts, arg0)
}

// Pubkeys is a free data retrieval call binding the contract method 0x7bd3e3f6.
//
// Solidity: function pubkeys(uint256 ) view returns(bytes)
func (_NodeDeposit *NodeDepositCallerSession) Pubkeys(arg0 *big.Int) ([]byte, error) {
	return _NodeDeposit.Contract.Pubkeys(&_NodeDeposit.CallOpts, arg0)
}

// TrustNodeDepositEnabled is a free data retrieval call binding the contract method 0xc2e6d64c.
//
// Solidity: function trustNodeDepositEnabled() view returns(bool)
func (_NodeDeposit *NodeDepositCaller) TrustNodeDepositEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "trustNodeDepositEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TrustNodeDepositEnabled is a free data retrieval call binding the contract method 0xc2e6d64c.
//
// Solidity: function trustNodeDepositEnabled() view returns(bool)
func (_NodeDeposit *NodeDepositSession) TrustNodeDepositEnabled() (bool, error) {
	return _NodeDeposit.Contract.TrustNodeDepositEnabled(&_NodeDeposit.CallOpts)
}

// TrustNodeDepositEnabled is a free data retrieval call binding the contract method 0xc2e6d64c.
//
// Solidity: function trustNodeDepositEnabled() view returns(bool)
func (_NodeDeposit *NodeDepositCallerSession) TrustNodeDepositEnabled() (bool, error) {
	return _NodeDeposit.Contract.TrustNodeDepositEnabled(&_NodeDeposit.CallOpts)
}

// TrustNodePubkeyNumberLimit is a free data retrieval call binding the contract method 0x7fe1852a.
//
// Solidity: function trustNodePubkeyNumberLimit() view returns(uint256)
func (_NodeDeposit *NodeDepositCaller) TrustNodePubkeyNumberLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "trustNodePubkeyNumberLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrustNodePubkeyNumberLimit is a free data retrieval call binding the contract method 0x7fe1852a.
//
// Solidity: function trustNodePubkeyNumberLimit() view returns(uint256)
func (_NodeDeposit *NodeDepositSession) TrustNodePubkeyNumberLimit() (*big.Int, error) {
	return _NodeDeposit.Contract.TrustNodePubkeyNumberLimit(&_NodeDeposit.CallOpts)
}

// TrustNodePubkeyNumberLimit is a free data retrieval call binding the contract method 0x7fe1852a.
//
// Solidity: function trustNodePubkeyNumberLimit() view returns(uint256)
func (_NodeDeposit *NodeDepositCallerSession) TrustNodePubkeyNumberLimit() (*big.Int, error) {
	return _NodeDeposit.Contract.TrustNodePubkeyNumberLimit(&_NodeDeposit.CallOpts)
}

// UserDepositAddress is a free data retrieval call binding the contract method 0x46773830.
//
// Solidity: function userDepositAddress() view returns(address)
func (_NodeDeposit *NodeDepositCaller) UserDepositAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "userDepositAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserDepositAddress is a free data retrieval call binding the contract method 0x46773830.
//
// Solidity: function userDepositAddress() view returns(address)
func (_NodeDeposit *NodeDepositSession) UserDepositAddress() (common.Address, error) {
	return _NodeDeposit.Contract.UserDepositAddress(&_NodeDeposit.CallOpts)
}

// UserDepositAddress is a free data retrieval call binding the contract method 0x46773830.
//
// Solidity: function userDepositAddress() view returns(address)
func (_NodeDeposit *NodeDepositCallerSession) UserDepositAddress() (common.Address, error) {
	return _NodeDeposit.Contract.UserDepositAddress(&_NodeDeposit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NodeDeposit *NodeDepositCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NodeDeposit *NodeDepositSession) Version() (uint8, error) {
	return _NodeDeposit.Contract.Version(&_NodeDeposit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NodeDeposit *NodeDepositCallerSession) Version() (uint8, error) {
	return _NodeDeposit.Contract.Version(&_NodeDeposit.CallOpts)
}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes)
func (_NodeDeposit *NodeDepositCaller) WithdrawCredentials(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "withdrawCredentials")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes)
func (_NodeDeposit *NodeDepositSession) WithdrawCredentials() ([]byte, error) {
	return _NodeDeposit.Contract.WithdrawCredentials(&_NodeDeposit.CallOpts)
}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes)
func (_NodeDeposit *NodeDepositCallerSession) WithdrawCredentials() ([]byte, error) {
	return _NodeDeposit.Contract.WithdrawCredentials(&_NodeDeposit.CallOpts)
}

// AddTrustNode is a paid mutator transaction binding the contract method 0xea5c80f1.
//
// Solidity: function addTrustNode(address _trustNodeAddress) returns()
func (_NodeDeposit *NodeDepositTransactor) AddTrustNode(opts *bind.TransactOpts, _trustNodeAddress common.Address) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "addTrustNode", _trustNodeAddress)
}

// AddTrustNode is a paid mutator transaction binding the contract method 0xea5c80f1.
//
// Solidity: function addTrustNode(address _trustNodeAddress) returns()
func (_NodeDeposit *NodeDepositSession) AddTrustNode(_trustNodeAddress common.Address) (*types.Transaction, error) {
	return _NodeDeposit.Contract.AddTrustNode(&_NodeDeposit.TransactOpts, _trustNodeAddress)
}

// AddTrustNode is a paid mutator transaction binding the contract method 0xea5c80f1.
//
// Solidity: function addTrustNode(address _trustNodeAddress) returns()
func (_NodeDeposit *NodeDepositTransactorSession) AddTrustNode(_trustNodeAddress common.Address) (*types.Transaction, error) {
	return _NodeDeposit.Contract.AddTrustNode(&_NodeDeposit.TransactOpts, _trustNodeAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) payable returns()
func (_NodeDeposit *NodeDepositTransactor) Deposit(opts *bind.TransactOpts, _validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "deposit", _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) payable returns()
func (_NodeDeposit *NodeDepositSession) Deposit(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.Deposit(&_NodeDeposit.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Deposit is a paid mutator transaction binding the contract method 0xd46221f9.
//
// Solidity: function deposit(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) payable returns()
func (_NodeDeposit *NodeDepositTransactorSession) Deposit(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.Deposit(&_NodeDeposit.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_NodeDeposit *NodeDepositTransactor) DepositEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "depositEth")
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_NodeDeposit *NodeDepositSession) DepositEth() (*types.Transaction, error) {
	return _NodeDeposit.Contract.DepositEth(&_NodeDeposit.TransactOpts)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns()
func (_NodeDeposit *NodeDepositTransactorSession) DepositEth() (*types.Transaction, error) {
	return _NodeDeposit.Contract.DepositEth(&_NodeDeposit.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xdfdfbd0b.
//
// Solidity: function init(address _userDepositAddress, address _ethDepositAddress, address _networkProposalAddress, bytes _withdrawCredentials) returns()
func (_NodeDeposit *NodeDepositTransactor) Init(opts *bind.TransactOpts, _userDepositAddress common.Address, _ethDepositAddress common.Address, _networkProposalAddress common.Address, _withdrawCredentials []byte) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "init", _userDepositAddress, _ethDepositAddress, _networkProposalAddress, _withdrawCredentials)
}

// Init is a paid mutator transaction binding the contract method 0xdfdfbd0b.
//
// Solidity: function init(address _userDepositAddress, address _ethDepositAddress, address _networkProposalAddress, bytes _withdrawCredentials) returns()
func (_NodeDeposit *NodeDepositSession) Init(_userDepositAddress common.Address, _ethDepositAddress common.Address, _networkProposalAddress common.Address, _withdrawCredentials []byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.Init(&_NodeDeposit.TransactOpts, _userDepositAddress, _ethDepositAddress, _networkProposalAddress, _withdrawCredentials)
}

// Init is a paid mutator transaction binding the contract method 0xdfdfbd0b.
//
// Solidity: function init(address _userDepositAddress, address _ethDepositAddress, address _networkProposalAddress, bytes _withdrawCredentials) returns()
func (_NodeDeposit *NodeDepositTransactorSession) Init(_userDepositAddress common.Address, _ethDepositAddress common.Address, _networkProposalAddress common.Address, _withdrawCredentials []byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.Init(&_NodeDeposit.TransactOpts, _userDepositAddress, _ethDepositAddress, _networkProposalAddress, _withdrawCredentials)
}

// RemoveTrustNode is a paid mutator transaction binding the contract method 0x5ffffad6.
//
// Solidity: function removeTrustNode(address _trustNodeAddress) returns()
func (_NodeDeposit *NodeDepositTransactor) RemoveTrustNode(opts *bind.TransactOpts, _trustNodeAddress common.Address) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "removeTrustNode", _trustNodeAddress)
}

// RemoveTrustNode is a paid mutator transaction binding the contract method 0x5ffffad6.
//
// Solidity: function removeTrustNode(address _trustNodeAddress) returns()
func (_NodeDeposit *NodeDepositSession) RemoveTrustNode(_trustNodeAddress common.Address) (*types.Transaction, error) {
	return _NodeDeposit.Contract.RemoveTrustNode(&_NodeDeposit.TransactOpts, _trustNodeAddress)
}

// RemoveTrustNode is a paid mutator transaction binding the contract method 0x5ffffad6.
//
// Solidity: function removeTrustNode(address _trustNodeAddress) returns()
func (_NodeDeposit *NodeDepositTransactorSession) RemoveTrustNode(_trustNodeAddress common.Address) (*types.Transaction, error) {
	return _NodeDeposit.Contract.RemoveTrustNode(&_NodeDeposit.TransactOpts, _trustNodeAddress)
}

// SetLightNodeDepositAmount is a paid mutator transaction binding the contract method 0x3fc50c92.
//
// Solidity: function setLightNodeDepositAmount(uint256 _amount) returns()
func (_NodeDeposit *NodeDepositTransactor) SetLightNodeDepositAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "setLightNodeDepositAmount", _amount)
}

// SetLightNodeDepositAmount is a paid mutator transaction binding the contract method 0x3fc50c92.
//
// Solidity: function setLightNodeDepositAmount(uint256 _amount) returns()
func (_NodeDeposit *NodeDepositSession) SetLightNodeDepositAmount(_amount *big.Int) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetLightNodeDepositAmount(&_NodeDeposit.TransactOpts, _amount)
}

// SetLightNodeDepositAmount is a paid mutator transaction binding the contract method 0x3fc50c92.
//
// Solidity: function setLightNodeDepositAmount(uint256 _amount) returns()
func (_NodeDeposit *NodeDepositTransactorSession) SetLightNodeDepositAmount(_amount *big.Int) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetLightNodeDepositAmount(&_NodeDeposit.TransactOpts, _amount)
}

// SetLightNodeDepositEnabled is a paid mutator transaction binding the contract method 0xadf1d8d6.
//
// Solidity: function setLightNodeDepositEnabled(bool _value) returns()
func (_NodeDeposit *NodeDepositTransactor) SetLightNodeDepositEnabled(opts *bind.TransactOpts, _value bool) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "setLightNodeDepositEnabled", _value)
}

// SetLightNodeDepositEnabled is a paid mutator transaction binding the contract method 0xadf1d8d6.
//
// Solidity: function setLightNodeDepositEnabled(bool _value) returns()
func (_NodeDeposit *NodeDepositSession) SetLightNodeDepositEnabled(_value bool) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetLightNodeDepositEnabled(&_NodeDeposit.TransactOpts, _value)
}

// SetLightNodeDepositEnabled is a paid mutator transaction binding the contract method 0xadf1d8d6.
//
// Solidity: function setLightNodeDepositEnabled(bool _value) returns()
func (_NodeDeposit *NodeDepositTransactorSession) SetLightNodeDepositEnabled(_value bool) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetLightNodeDepositEnabled(&_NodeDeposit.TransactOpts, _value)
}

// SetNodePubkeyStatus is a paid mutator transaction binding the contract method 0xd1dd8614.
//
// Solidity: function setNodePubkeyStatus(bytes _validatorPubkey, uint8 _status) returns()
func (_NodeDeposit *NodeDepositTransactor) SetNodePubkeyStatus(opts *bind.TransactOpts, _validatorPubkey []byte, _status uint8) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "setNodePubkeyStatus", _validatorPubkey, _status)
}

// SetNodePubkeyStatus is a paid mutator transaction binding the contract method 0xd1dd8614.
//
// Solidity: function setNodePubkeyStatus(bytes _validatorPubkey, uint8 _status) returns()
func (_NodeDeposit *NodeDepositSession) SetNodePubkeyStatus(_validatorPubkey []byte, _status uint8) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetNodePubkeyStatus(&_NodeDeposit.TransactOpts, _validatorPubkey, _status)
}

// SetNodePubkeyStatus is a paid mutator transaction binding the contract method 0xd1dd8614.
//
// Solidity: function setNodePubkeyStatus(bytes _validatorPubkey, uint8 _status) returns()
func (_NodeDeposit *NodeDepositTransactorSession) SetNodePubkeyStatus(_validatorPubkey []byte, _status uint8) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetNodePubkeyStatus(&_NodeDeposit.TransactOpts, _validatorPubkey, _status)
}

// SetTrustNodeDepositEnabled is a paid mutator transaction binding the contract method 0xcce0835e.
//
// Solidity: function setTrustNodeDepositEnabled(bool _value) returns()
func (_NodeDeposit *NodeDepositTransactor) SetTrustNodeDepositEnabled(opts *bind.TransactOpts, _value bool) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "setTrustNodeDepositEnabled", _value)
}

// SetTrustNodeDepositEnabled is a paid mutator transaction binding the contract method 0xcce0835e.
//
// Solidity: function setTrustNodeDepositEnabled(bool _value) returns()
func (_NodeDeposit *NodeDepositSession) SetTrustNodeDepositEnabled(_value bool) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetTrustNodeDepositEnabled(&_NodeDeposit.TransactOpts, _value)
}

// SetTrustNodeDepositEnabled is a paid mutator transaction binding the contract method 0xcce0835e.
//
// Solidity: function setTrustNodeDepositEnabled(bool _value) returns()
func (_NodeDeposit *NodeDepositTransactorSession) SetTrustNodeDepositEnabled(_value bool) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetTrustNodeDepositEnabled(&_NodeDeposit.TransactOpts, _value)
}

// SetTrustNodePubkeyLimit is a paid mutator transaction binding the contract method 0x7a0c1aa4.
//
// Solidity: function setTrustNodePubkeyLimit(uint256 _value) returns()
func (_NodeDeposit *NodeDepositTransactor) SetTrustNodePubkeyLimit(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "setTrustNodePubkeyLimit", _value)
}

// SetTrustNodePubkeyLimit is a paid mutator transaction binding the contract method 0x7a0c1aa4.
//
// Solidity: function setTrustNodePubkeyLimit(uint256 _value) returns()
func (_NodeDeposit *NodeDepositSession) SetTrustNodePubkeyLimit(_value *big.Int) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetTrustNodePubkeyLimit(&_NodeDeposit.TransactOpts, _value)
}

// SetTrustNodePubkeyLimit is a paid mutator transaction binding the contract method 0x7a0c1aa4.
//
// Solidity: function setTrustNodePubkeyLimit(uint256 _value) returns()
func (_NodeDeposit *NodeDepositTransactorSession) SetTrustNodePubkeyLimit(_value *big.Int) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetTrustNodePubkeyLimit(&_NodeDeposit.TransactOpts, _value)
}

// SetWithdrawCredentials is a paid mutator transaction binding the contract method 0xe4304698.
//
// Solidity: function setWithdrawCredentials(bytes _withdrawCredentials) returns()
func (_NodeDeposit *NodeDepositTransactor) SetWithdrawCredentials(opts *bind.TransactOpts, _withdrawCredentials []byte) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "setWithdrawCredentials", _withdrawCredentials)
}

// SetWithdrawCredentials is a paid mutator transaction binding the contract method 0xe4304698.
//
// Solidity: function setWithdrawCredentials(bytes _withdrawCredentials) returns()
func (_NodeDeposit *NodeDepositSession) SetWithdrawCredentials(_withdrawCredentials []byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetWithdrawCredentials(&_NodeDeposit.TransactOpts, _withdrawCredentials)
}

// SetWithdrawCredentials is a paid mutator transaction binding the contract method 0xe4304698.
//
// Solidity: function setWithdrawCredentials(bytes _withdrawCredentials) returns()
func (_NodeDeposit *NodeDepositTransactorSession) SetWithdrawCredentials(_withdrawCredentials []byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetWithdrawCredentials(&_NodeDeposit.TransactOpts, _withdrawCredentials)
}

// Stake is a paid mutator transaction binding the contract method 0xca2b87af.
//
// Solidity: function stake(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_NodeDeposit *NodeDepositTransactor) Stake(opts *bind.TransactOpts, _validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "stake", _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Stake is a paid mutator transaction binding the contract method 0xca2b87af.
//
// Solidity: function stake(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_NodeDeposit *NodeDepositSession) Stake(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.Stake(&_NodeDeposit.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// Stake is a paid mutator transaction binding the contract method 0xca2b87af.
//
// Solidity: function stake(bytes[] _validatorPubkeys, bytes[] _validatorSignatures, bytes32[] _depositDataRoots) returns()
func (_NodeDeposit *NodeDepositTransactorSession) Stake(_validatorPubkeys [][]byte, _validatorSignatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.Stake(&_NodeDeposit.TransactOpts, _validatorPubkeys, _validatorSignatures, _depositDataRoots)
}

// VoteWithdrawCredentials is a paid mutator transaction binding the contract method 0xa4b4334b.
//
// Solidity: function voteWithdrawCredentials(bytes[] _pubkeys, bool[] _matchs) returns()
func (_NodeDeposit *NodeDepositTransactor) VoteWithdrawCredentials(opts *bind.TransactOpts, _pubkeys [][]byte, _matchs []bool) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "voteWithdrawCredentials", _pubkeys, _matchs)
}

// VoteWithdrawCredentials is a paid mutator transaction binding the contract method 0xa4b4334b.
//
// Solidity: function voteWithdrawCredentials(bytes[] _pubkeys, bool[] _matchs) returns()
func (_NodeDeposit *NodeDepositSession) VoteWithdrawCredentials(_pubkeys [][]byte, _matchs []bool) (*types.Transaction, error) {
	return _NodeDeposit.Contract.VoteWithdrawCredentials(&_NodeDeposit.TransactOpts, _pubkeys, _matchs)
}

// VoteWithdrawCredentials is a paid mutator transaction binding the contract method 0xa4b4334b.
//
// Solidity: function voteWithdrawCredentials(bytes[] _pubkeys, bool[] _matchs) returns()
func (_NodeDeposit *NodeDepositTransactorSession) VoteWithdrawCredentials(_pubkeys [][]byte, _matchs []bool) (*types.Transaction, error) {
	return _NodeDeposit.Contract.VoteWithdrawCredentials(&_NodeDeposit.TransactOpts, _pubkeys, _matchs)
}

// NodeDepositDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the NodeDeposit contract.
type NodeDepositDepositedIterator struct {
	Event *NodeDepositDeposited // Event containing the contract specifics and raw log

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
func (it *NodeDepositDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDepositDeposited)
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
		it.Event = new(NodeDepositDeposited)
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
func (it *NodeDepositDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDepositDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDepositDeposited represents a Deposited event raised by the NodeDeposit contract.
type NodeDepositDeposited struct {
	Node               common.Address
	NodeType           uint8
	Pubkey             []byte
	ValidatorSignature []byte
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x765ed32e3058b9b6f85a409a5e6c225161f8b02555a254720491a1e47d5dc131.
//
// Solidity: event Deposited(address node, uint8 nodeType, bytes pubkey, bytes validatorSignature, uint256 amount)
func (_NodeDeposit *NodeDepositFilterer) FilterDeposited(opts *bind.FilterOpts) (*NodeDepositDepositedIterator, error) {

	logs, sub, err := _NodeDeposit.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &NodeDepositDepositedIterator{contract: _NodeDeposit.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x765ed32e3058b9b6f85a409a5e6c225161f8b02555a254720491a1e47d5dc131.
//
// Solidity: event Deposited(address node, uint8 nodeType, bytes pubkey, bytes validatorSignature, uint256 amount)
func (_NodeDeposit *NodeDepositFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *NodeDepositDeposited) (event.Subscription, error) {

	logs, sub, err := _NodeDeposit.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDepositDeposited)
				if err := _NodeDeposit.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x765ed32e3058b9b6f85a409a5e6c225161f8b02555a254720491a1e47d5dc131.
//
// Solidity: event Deposited(address node, uint8 nodeType, bytes pubkey, bytes validatorSignature, uint256 amount)
func (_NodeDeposit *NodeDepositFilterer) ParseDeposited(log types.Log) (*NodeDepositDeposited, error) {
	event := new(NodeDepositDeposited)
	if err := _NodeDeposit.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeDepositEtherDepositedIterator is returned from FilterEtherDeposited and is used to iterate over the raw logs and unpacked data for EtherDeposited events raised by the NodeDeposit contract.
type NodeDepositEtherDepositedIterator struct {
	Event *NodeDepositEtherDeposited // Event containing the contract specifics and raw log

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
func (it *NodeDepositEtherDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDepositEtherDeposited)
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
		it.Event = new(NodeDepositEtherDeposited)
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
func (it *NodeDepositEtherDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDepositEtherDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDepositEtherDeposited represents a EtherDeposited event raised by the NodeDeposit contract.
type NodeDepositEtherDeposited struct {
	From   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherDeposited is a free log retrieval operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_NodeDeposit *NodeDepositFilterer) FilterEtherDeposited(opts *bind.FilterOpts, from []common.Address) (*NodeDepositEtherDepositedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _NodeDeposit.contract.FilterLogs(opts, "EtherDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return &NodeDepositEtherDepositedIterator{contract: _NodeDeposit.contract, event: "EtherDeposited", logs: logs, sub: sub}, nil
}

// WatchEtherDeposited is a free log subscription operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_NodeDeposit *NodeDepositFilterer) WatchEtherDeposited(opts *bind.WatchOpts, sink chan<- *NodeDepositEtherDeposited, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _NodeDeposit.contract.WatchLogs(opts, "EtherDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDepositEtherDeposited)
				if err := _NodeDeposit.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
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
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_NodeDeposit *NodeDepositFilterer) ParseEtherDeposited(log types.Log) (*NodeDepositEtherDeposited, error) {
	event := new(NodeDepositEtherDeposited)
	if err := _NodeDeposit.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeDepositOffBoardedIterator is returned from FilterOffBoarded and is used to iterate over the raw logs and unpacked data for OffBoarded events raised by the NodeDeposit contract.
type NodeDepositOffBoardedIterator struct {
	Event *NodeDepositOffBoarded // Event containing the contract specifics and raw log

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
func (it *NodeDepositOffBoardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDepositOffBoarded)
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
		it.Event = new(NodeDepositOffBoarded)
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
func (it *NodeDepositOffBoardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDepositOffBoardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDepositOffBoarded represents a OffBoarded event raised by the NodeDeposit contract.
type NodeDepositOffBoarded struct {
	Node   common.Address
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOffBoarded is a free log retrieval operation binding the contract event 0x9d5023d85497e8c264e3b053f8da9415f4db76eb5af3ecef3fe953fe9f981470.
//
// Solidity: event OffBoarded(address node, bytes pubkey)
func (_NodeDeposit *NodeDepositFilterer) FilterOffBoarded(opts *bind.FilterOpts) (*NodeDepositOffBoardedIterator, error) {

	logs, sub, err := _NodeDeposit.contract.FilterLogs(opts, "OffBoarded")
	if err != nil {
		return nil, err
	}
	return &NodeDepositOffBoardedIterator{contract: _NodeDeposit.contract, event: "OffBoarded", logs: logs, sub: sub}, nil
}

// WatchOffBoarded is a free log subscription operation binding the contract event 0x9d5023d85497e8c264e3b053f8da9415f4db76eb5af3ecef3fe953fe9f981470.
//
// Solidity: event OffBoarded(address node, bytes pubkey)
func (_NodeDeposit *NodeDepositFilterer) WatchOffBoarded(opts *bind.WatchOpts, sink chan<- *NodeDepositOffBoarded) (event.Subscription, error) {

	logs, sub, err := _NodeDeposit.contract.WatchLogs(opts, "OffBoarded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDepositOffBoarded)
				if err := _NodeDeposit.contract.UnpackLog(event, "OffBoarded", log); err != nil {
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

// ParseOffBoarded is a log parse operation binding the contract event 0x9d5023d85497e8c264e3b053f8da9415f4db76eb5af3ecef3fe953fe9f981470.
//
// Solidity: event OffBoarded(address node, bytes pubkey)
func (_NodeDeposit *NodeDepositFilterer) ParseOffBoarded(log types.Log) (*NodeDepositOffBoarded, error) {
	event := new(NodeDepositOffBoarded)
	if err := _NodeDeposit.contract.UnpackLog(event, "OffBoarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeDepositSetPubkeyStatusIterator is returned from FilterSetPubkeyStatus and is used to iterate over the raw logs and unpacked data for SetPubkeyStatus events raised by the NodeDeposit contract.
type NodeDepositSetPubkeyStatusIterator struct {
	Event *NodeDepositSetPubkeyStatus // Event containing the contract specifics and raw log

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
func (it *NodeDepositSetPubkeyStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDepositSetPubkeyStatus)
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
		it.Event = new(NodeDepositSetPubkeyStatus)
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
func (it *NodeDepositSetPubkeyStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDepositSetPubkeyStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDepositSetPubkeyStatus represents a SetPubkeyStatus event raised by the NodeDeposit contract.
type NodeDepositSetPubkeyStatus struct {
	Pubkey []byte
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPubkeyStatus is a free log retrieval operation binding the contract event 0x1ca9918016060bae952954327fda778333672e209a1713f8c77b27d0ea23d19e.
//
// Solidity: event SetPubkeyStatus(bytes pubkey, uint8 status)
func (_NodeDeposit *NodeDepositFilterer) FilterSetPubkeyStatus(opts *bind.FilterOpts) (*NodeDepositSetPubkeyStatusIterator, error) {

	logs, sub, err := _NodeDeposit.contract.FilterLogs(opts, "SetPubkeyStatus")
	if err != nil {
		return nil, err
	}
	return &NodeDepositSetPubkeyStatusIterator{contract: _NodeDeposit.contract, event: "SetPubkeyStatus", logs: logs, sub: sub}, nil
}

// WatchSetPubkeyStatus is a free log subscription operation binding the contract event 0x1ca9918016060bae952954327fda778333672e209a1713f8c77b27d0ea23d19e.
//
// Solidity: event SetPubkeyStatus(bytes pubkey, uint8 status)
func (_NodeDeposit *NodeDepositFilterer) WatchSetPubkeyStatus(opts *bind.WatchOpts, sink chan<- *NodeDepositSetPubkeyStatus) (event.Subscription, error) {

	logs, sub, err := _NodeDeposit.contract.WatchLogs(opts, "SetPubkeyStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDepositSetPubkeyStatus)
				if err := _NodeDeposit.contract.UnpackLog(event, "SetPubkeyStatus", log); err != nil {
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

// ParseSetPubkeyStatus is a log parse operation binding the contract event 0x1ca9918016060bae952954327fda778333672e209a1713f8c77b27d0ea23d19e.
//
// Solidity: event SetPubkeyStatus(bytes pubkey, uint8 status)
func (_NodeDeposit *NodeDepositFilterer) ParseSetPubkeyStatus(log types.Log) (*NodeDepositSetPubkeyStatus, error) {
	event := new(NodeDepositSetPubkeyStatus)
	if err := _NodeDeposit.contract.UnpackLog(event, "SetPubkeyStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeDepositStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the NodeDeposit contract.
type NodeDepositStakedIterator struct {
	Event *NodeDepositStaked // Event containing the contract specifics and raw log

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
func (it *NodeDepositStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDepositStaked)
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
		it.Event = new(NodeDepositStaked)
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
func (it *NodeDepositStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDepositStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDepositStaked represents a Staked event raised by the NodeDeposit contract.
type NodeDepositStaked struct {
	Node   common.Address
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xb384b282308c431ac3ea7756646550752d3f0dbfb418ef60bfeaf4edc9494815.
//
// Solidity: event Staked(address node, bytes pubkey)
func (_NodeDeposit *NodeDepositFilterer) FilterStaked(opts *bind.FilterOpts) (*NodeDepositStakedIterator, error) {

	logs, sub, err := _NodeDeposit.contract.FilterLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return &NodeDepositStakedIterator{contract: _NodeDeposit.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xb384b282308c431ac3ea7756646550752d3f0dbfb418ef60bfeaf4edc9494815.
//
// Solidity: event Staked(address node, bytes pubkey)
func (_NodeDeposit *NodeDepositFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *NodeDepositStaked) (event.Subscription, error) {

	logs, sub, err := _NodeDeposit.contract.WatchLogs(opts, "Staked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDepositStaked)
				if err := _NodeDeposit.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0xb384b282308c431ac3ea7756646550752d3f0dbfb418ef60bfeaf4edc9494815.
//
// Solidity: event Staked(address node, bytes pubkey)
func (_NodeDeposit *NodeDepositFilterer) ParseStaked(log types.Log) (*NodeDepositStaked, error) {
	event := new(NodeDepositStaked)
	if err := _NodeDeposit.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
