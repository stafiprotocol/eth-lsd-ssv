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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDealedEpoch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDealedHeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyNotifiedCycle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyVoted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountUnmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableDepositZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableRewardZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableWithdrawIndexOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommissionRateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CycleNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositAmountLTMinAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LsdTokenAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyRemoved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeNotClaimable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorizedLsdToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotClaimable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPubkeyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTrustNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProposalExecFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyNumberOverLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyStatusUnmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateChangeOverLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SecondsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SoloNodeDepositAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SoloNodeDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmitBalancesDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustNodeDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawIndexEmpty\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumINodeDeposit.NodeType\",\"name\":\"nodeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"enumINodeDeposit.PubkeyStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"SetPubkeyStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustNodeAddress\",\"type\":\"address\"}],\"name\":\"addTrustNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_validatorPubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_validatorSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDepositAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"getNodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"nodeList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"getPubkeysOfNode\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userDepositAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethDepositAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_networkProposalAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_withdrawCredentials\",\"type\":\"bytes\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkProposalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeInfoOf\",\"outputs\":[{\"internalType\":\"enumINodeDeposit.NodeType\",\"name\":\"_nodeType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_removed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pubkeyInfoOf\",\"outputs\":[{\"internalType\":\"enumINodeDeposit.PubkeyStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nodeDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_depositBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pubkeysOfNode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustNodeAddress\",\"type\":\"address\"}],\"name\":\"removeTrustNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"},{\"internalType\":\"enumINodeDeposit.PubkeyStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setNodePubkeyStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setSoloNodeDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setSoloNodeDepositEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"setTrustNodeDepositEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setTrustNodePubkeyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_withdrawCredentials\",\"type\":\"bytes\"}],\"name\":\"setWithdrawCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"soloNodeDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"soloNodeDepositEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_validatorPubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_validatorSignatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustNodeDepositEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustNodePubkeyNumberLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDepositAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"_match\",\"type\":\"bool\"}],\"name\":\"voteWithdrawCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawCredentials\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetNodes is a free data retrieval call binding the contract method 0x038d67e8.
//
// Solidity: function getNodes(uint256 _start, uint256 _end) view returns(address[] nodeList)
func (_NodeDeposit *NodeDepositCaller) GetNodes(opts *bind.CallOpts, _start *big.Int, _end *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "getNodes", _start, _end)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetNodes is a free data retrieval call binding the contract method 0x038d67e8.
//
// Solidity: function getNodes(uint256 _start, uint256 _end) view returns(address[] nodeList)
func (_NodeDeposit *NodeDepositSession) GetNodes(_start *big.Int, _end *big.Int) ([]common.Address, error) {
	return _NodeDeposit.Contract.GetNodes(&_NodeDeposit.CallOpts, _start, _end)
}

// GetNodes is a free data retrieval call binding the contract method 0x038d67e8.
//
// Solidity: function getNodes(uint256 _start, uint256 _end) view returns(address[] nodeList)
func (_NodeDeposit *NodeDepositCallerSession) GetNodes(_start *big.Int, _end *big.Int) ([]common.Address, error) {
	return _NodeDeposit.Contract.GetNodes(&_NodeDeposit.CallOpts, _start, _end)
}

// GetNodesLength is a free data retrieval call binding the contract method 0x1821aa62.
//
// Solidity: function getNodesLength() view returns(uint256)
func (_NodeDeposit *NodeDepositCaller) GetNodesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "getNodesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodesLength is a free data retrieval call binding the contract method 0x1821aa62.
//
// Solidity: function getNodesLength() view returns(uint256)
func (_NodeDeposit *NodeDepositSession) GetNodesLength() (*big.Int, error) {
	return _NodeDeposit.Contract.GetNodesLength(&_NodeDeposit.CallOpts)
}

// GetNodesLength is a free data retrieval call binding the contract method 0x1821aa62.
//
// Solidity: function getNodesLength() view returns(uint256)
func (_NodeDeposit *NodeDepositCallerSession) GetNodesLength() (*big.Int, error) {
	return _NodeDeposit.Contract.GetNodesLength(&_NodeDeposit.CallOpts)
}

// GetPubkeysOfNode is a free data retrieval call binding the contract method 0xac1464a3.
//
// Solidity: function getPubkeysOfNode(address _node) view returns(bytes[])
func (_NodeDeposit *NodeDepositCaller) GetPubkeysOfNode(opts *bind.CallOpts, _node common.Address) ([][]byte, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "getPubkeysOfNode", _node)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetPubkeysOfNode is a free data retrieval call binding the contract method 0xac1464a3.
//
// Solidity: function getPubkeysOfNode(address _node) view returns(bytes[])
func (_NodeDeposit *NodeDepositSession) GetPubkeysOfNode(_node common.Address) ([][]byte, error) {
	return _NodeDeposit.Contract.GetPubkeysOfNode(&_NodeDeposit.CallOpts, _node)
}

// GetPubkeysOfNode is a free data retrieval call binding the contract method 0xac1464a3.
//
// Solidity: function getPubkeysOfNode(address _node) view returns(bytes[])
func (_NodeDeposit *NodeDepositCallerSession) GetPubkeysOfNode(_node common.Address) ([][]byte, error) {
	return _NodeDeposit.Contract.GetPubkeysOfNode(&_NodeDeposit.CallOpts, _node)
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
// Solidity: function nodeInfoOf(address ) view returns(uint8 _nodeType, bool _removed)
func (_NodeDeposit *NodeDepositCaller) NodeInfoOf(opts *bind.CallOpts, arg0 common.Address) (struct {
	NodeType uint8
	Removed  bool
}, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "nodeInfoOf", arg0)

	outstruct := new(struct {
		NodeType uint8
		Removed  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NodeType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Removed = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// NodeInfoOf is a free data retrieval call binding the contract method 0x0fdefd36.
//
// Solidity: function nodeInfoOf(address ) view returns(uint8 _nodeType, bool _removed)
func (_NodeDeposit *NodeDepositSession) NodeInfoOf(arg0 common.Address) (struct {
	NodeType uint8
	Removed  bool
}, error) {
	return _NodeDeposit.Contract.NodeInfoOf(&_NodeDeposit.CallOpts, arg0)
}

// NodeInfoOf is a free data retrieval call binding the contract method 0x0fdefd36.
//
// Solidity: function nodeInfoOf(address ) view returns(uint8 _nodeType, bool _removed)
func (_NodeDeposit *NodeDepositCallerSession) NodeInfoOf(arg0 common.Address) (struct {
	NodeType uint8
	Removed  bool
}, error) {
	return _NodeDeposit.Contract.NodeInfoOf(&_NodeDeposit.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(address)
func (_NodeDeposit *NodeDepositCaller) Nodes(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "nodes", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(address)
func (_NodeDeposit *NodeDepositSession) Nodes(arg0 *big.Int) (common.Address, error) {
	return _NodeDeposit.Contract.Nodes(&_NodeDeposit.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(address)
func (_NodeDeposit *NodeDepositCallerSession) Nodes(arg0 *big.Int) (common.Address, error) {
	return _NodeDeposit.Contract.Nodes(&_NodeDeposit.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NodeDeposit *NodeDepositCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NodeDeposit *NodeDepositSession) ProxiableUUID() ([32]byte, error) {
	return _NodeDeposit.Contract.ProxiableUUID(&_NodeDeposit.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NodeDeposit *NodeDepositCallerSession) ProxiableUUID() ([32]byte, error) {
	return _NodeDeposit.Contract.ProxiableUUID(&_NodeDeposit.CallOpts)
}

// PubkeyInfoOf is a free data retrieval call binding the contract method 0xd6c9f4c0.
//
// Solidity: function pubkeyInfoOf(bytes ) view returns(uint8 _status, address _owner, uint256 _nodeDepositAmount, uint256 _depositBlock)
func (_NodeDeposit *NodeDepositCaller) PubkeyInfoOf(opts *bind.CallOpts, arg0 []byte) (struct {
	Status            uint8
	Owner             common.Address
	NodeDepositAmount *big.Int
	DepositBlock      *big.Int
}, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "pubkeyInfoOf", arg0)

	outstruct := new(struct {
		Status            uint8
		Owner             common.Address
		NodeDepositAmount *big.Int
		DepositBlock      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NodeDepositAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DepositBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PubkeyInfoOf is a free data retrieval call binding the contract method 0xd6c9f4c0.
//
// Solidity: function pubkeyInfoOf(bytes ) view returns(uint8 _status, address _owner, uint256 _nodeDepositAmount, uint256 _depositBlock)
func (_NodeDeposit *NodeDepositSession) PubkeyInfoOf(arg0 []byte) (struct {
	Status            uint8
	Owner             common.Address
	NodeDepositAmount *big.Int
	DepositBlock      *big.Int
}, error) {
	return _NodeDeposit.Contract.PubkeyInfoOf(&_NodeDeposit.CallOpts, arg0)
}

// PubkeyInfoOf is a free data retrieval call binding the contract method 0xd6c9f4c0.
//
// Solidity: function pubkeyInfoOf(bytes ) view returns(uint8 _status, address _owner, uint256 _nodeDepositAmount, uint256 _depositBlock)
func (_NodeDeposit *NodeDepositCallerSession) PubkeyInfoOf(arg0 []byte) (struct {
	Status            uint8
	Owner             common.Address
	NodeDepositAmount *big.Int
	DepositBlock      *big.Int
}, error) {
	return _NodeDeposit.Contract.PubkeyInfoOf(&_NodeDeposit.CallOpts, arg0)
}

// PubkeysOfNode is a free data retrieval call binding the contract method 0xa6b9ba17.
//
// Solidity: function pubkeysOfNode(address , uint256 ) view returns(bytes)
func (_NodeDeposit *NodeDepositCaller) PubkeysOfNode(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "pubkeysOfNode", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PubkeysOfNode is a free data retrieval call binding the contract method 0xa6b9ba17.
//
// Solidity: function pubkeysOfNode(address , uint256 ) view returns(bytes)
func (_NodeDeposit *NodeDepositSession) PubkeysOfNode(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _NodeDeposit.Contract.PubkeysOfNode(&_NodeDeposit.CallOpts, arg0, arg1)
}

// PubkeysOfNode is a free data retrieval call binding the contract method 0xa6b9ba17.
//
// Solidity: function pubkeysOfNode(address , uint256 ) view returns(bytes)
func (_NodeDeposit *NodeDepositCallerSession) PubkeysOfNode(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _NodeDeposit.Contract.PubkeysOfNode(&_NodeDeposit.CallOpts, arg0, arg1)
}

// SoloNodeDepositAmount is a free data retrieval call binding the contract method 0x0fc6d510.
//
// Solidity: function soloNodeDepositAmount() view returns(uint256)
func (_NodeDeposit *NodeDepositCaller) SoloNodeDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "soloNodeDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SoloNodeDepositAmount is a free data retrieval call binding the contract method 0x0fc6d510.
//
// Solidity: function soloNodeDepositAmount() view returns(uint256)
func (_NodeDeposit *NodeDepositSession) SoloNodeDepositAmount() (*big.Int, error) {
	return _NodeDeposit.Contract.SoloNodeDepositAmount(&_NodeDeposit.CallOpts)
}

// SoloNodeDepositAmount is a free data retrieval call binding the contract method 0x0fc6d510.
//
// Solidity: function soloNodeDepositAmount() view returns(uint256)
func (_NodeDeposit *NodeDepositCallerSession) SoloNodeDepositAmount() (*big.Int, error) {
	return _NodeDeposit.Contract.SoloNodeDepositAmount(&_NodeDeposit.CallOpts)
}

// SoloNodeDepositEnabled is a free data retrieval call binding the contract method 0x70a70b8f.
//
// Solidity: function soloNodeDepositEnabled() view returns(bool)
func (_NodeDeposit *NodeDepositCaller) SoloNodeDepositEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NodeDeposit.contract.Call(opts, &out, "soloNodeDepositEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SoloNodeDepositEnabled is a free data retrieval call binding the contract method 0x70a70b8f.
//
// Solidity: function soloNodeDepositEnabled() view returns(bool)
func (_NodeDeposit *NodeDepositSession) SoloNodeDepositEnabled() (bool, error) {
	return _NodeDeposit.Contract.SoloNodeDepositEnabled(&_NodeDeposit.CallOpts)
}

// SoloNodeDepositEnabled is a free data retrieval call binding the contract method 0x70a70b8f.
//
// Solidity: function soloNodeDepositEnabled() view returns(bool)
func (_NodeDeposit *NodeDepositCallerSession) SoloNodeDepositEnabled() (bool, error) {
	return _NodeDeposit.Contract.SoloNodeDepositEnabled(&_NodeDeposit.CallOpts)
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

// Reinit is a paid mutator transaction binding the contract method 0xc482ceaf.
//
// Solidity: function reinit() returns()
func (_NodeDeposit *NodeDepositTransactor) Reinit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "reinit")
}

// Reinit is a paid mutator transaction binding the contract method 0xc482ceaf.
//
// Solidity: function reinit() returns()
func (_NodeDeposit *NodeDepositSession) Reinit() (*types.Transaction, error) {
	return _NodeDeposit.Contract.Reinit(&_NodeDeposit.TransactOpts)
}

// Reinit is a paid mutator transaction binding the contract method 0xc482ceaf.
//
// Solidity: function reinit() returns()
func (_NodeDeposit *NodeDepositTransactorSession) Reinit() (*types.Transaction, error) {
	return _NodeDeposit.Contract.Reinit(&_NodeDeposit.TransactOpts)
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

// SetSoloNodeDepositAmount is a paid mutator transaction binding the contract method 0x7ac75644.
//
// Solidity: function setSoloNodeDepositAmount(uint256 _amount) returns()
func (_NodeDeposit *NodeDepositTransactor) SetSoloNodeDepositAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "setSoloNodeDepositAmount", _amount)
}

// SetSoloNodeDepositAmount is a paid mutator transaction binding the contract method 0x7ac75644.
//
// Solidity: function setSoloNodeDepositAmount(uint256 _amount) returns()
func (_NodeDeposit *NodeDepositSession) SetSoloNodeDepositAmount(_amount *big.Int) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetSoloNodeDepositAmount(&_NodeDeposit.TransactOpts, _amount)
}

// SetSoloNodeDepositAmount is a paid mutator transaction binding the contract method 0x7ac75644.
//
// Solidity: function setSoloNodeDepositAmount(uint256 _amount) returns()
func (_NodeDeposit *NodeDepositTransactorSession) SetSoloNodeDepositAmount(_amount *big.Int) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetSoloNodeDepositAmount(&_NodeDeposit.TransactOpts, _amount)
}

// SetSoloNodeDepositEnabled is a paid mutator transaction binding the contract method 0xb89358a7.
//
// Solidity: function setSoloNodeDepositEnabled(bool _value) returns()
func (_NodeDeposit *NodeDepositTransactor) SetSoloNodeDepositEnabled(opts *bind.TransactOpts, _value bool) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "setSoloNodeDepositEnabled", _value)
}

// SetSoloNodeDepositEnabled is a paid mutator transaction binding the contract method 0xb89358a7.
//
// Solidity: function setSoloNodeDepositEnabled(bool _value) returns()
func (_NodeDeposit *NodeDepositSession) SetSoloNodeDepositEnabled(_value bool) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetSoloNodeDepositEnabled(&_NodeDeposit.TransactOpts, _value)
}

// SetSoloNodeDepositEnabled is a paid mutator transaction binding the contract method 0xb89358a7.
//
// Solidity: function setSoloNodeDepositEnabled(bool _value) returns()
func (_NodeDeposit *NodeDepositTransactorSession) SetSoloNodeDepositEnabled(_value bool) (*types.Transaction, error) {
	return _NodeDeposit.Contract.SetSoloNodeDepositEnabled(&_NodeDeposit.TransactOpts, _value)
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

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NodeDeposit *NodeDepositTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NodeDeposit *NodeDepositSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NodeDeposit.Contract.UpgradeTo(&_NodeDeposit.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NodeDeposit *NodeDepositTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NodeDeposit.Contract.UpgradeTo(&_NodeDeposit.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NodeDeposit *NodeDepositTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NodeDeposit *NodeDepositSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.UpgradeToAndCall(&_NodeDeposit.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NodeDeposit *NodeDepositTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NodeDeposit.Contract.UpgradeToAndCall(&_NodeDeposit.TransactOpts, newImplementation, data)
}

// VoteWithdrawCredentials is a paid mutator transaction binding the contract method 0x5952629b.
//
// Solidity: function voteWithdrawCredentials(bytes _pubkey, bool _match) returns()
func (_NodeDeposit *NodeDepositTransactor) VoteWithdrawCredentials(opts *bind.TransactOpts, _pubkey []byte, _match bool) (*types.Transaction, error) {
	return _NodeDeposit.contract.Transact(opts, "voteWithdrawCredentials", _pubkey, _match)
}

// VoteWithdrawCredentials is a paid mutator transaction binding the contract method 0x5952629b.
//
// Solidity: function voteWithdrawCredentials(bytes _pubkey, bool _match) returns()
func (_NodeDeposit *NodeDepositSession) VoteWithdrawCredentials(_pubkey []byte, _match bool) (*types.Transaction, error) {
	return _NodeDeposit.Contract.VoteWithdrawCredentials(&_NodeDeposit.TransactOpts, _pubkey, _match)
}

// VoteWithdrawCredentials is a paid mutator transaction binding the contract method 0x5952629b.
//
// Solidity: function voteWithdrawCredentials(bytes _pubkey, bool _match) returns()
func (_NodeDeposit *NodeDepositTransactorSession) VoteWithdrawCredentials(_pubkey []byte, _match bool) (*types.Transaction, error) {
	return _NodeDeposit.Contract.VoteWithdrawCredentials(&_NodeDeposit.TransactOpts, _pubkey, _match)
}

// NodeDepositAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the NodeDeposit contract.
type NodeDepositAdminChangedIterator struct {
	Event *NodeDepositAdminChanged // Event containing the contract specifics and raw log

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
func (it *NodeDepositAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDepositAdminChanged)
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
		it.Event = new(NodeDepositAdminChanged)
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
func (it *NodeDepositAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDepositAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDepositAdminChanged represents a AdminChanged event raised by the NodeDeposit contract.
type NodeDepositAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NodeDeposit *NodeDepositFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*NodeDepositAdminChangedIterator, error) {

	logs, sub, err := _NodeDeposit.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &NodeDepositAdminChangedIterator{contract: _NodeDeposit.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NodeDeposit *NodeDepositFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *NodeDepositAdminChanged) (event.Subscription, error) {

	logs, sub, err := _NodeDeposit.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDepositAdminChanged)
				if err := _NodeDeposit.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_NodeDeposit *NodeDepositFilterer) ParseAdminChanged(log types.Log) (*NodeDepositAdminChanged, error) {
	event := new(NodeDepositAdminChanged)
	if err := _NodeDeposit.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeDepositBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the NodeDeposit contract.
type NodeDepositBeaconUpgradedIterator struct {
	Event *NodeDepositBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *NodeDepositBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDepositBeaconUpgraded)
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
		it.Event = new(NodeDepositBeaconUpgraded)
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
func (it *NodeDepositBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDepositBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDepositBeaconUpgraded represents a BeaconUpgraded event raised by the NodeDeposit contract.
type NodeDepositBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NodeDeposit *NodeDepositFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*NodeDepositBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NodeDeposit.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &NodeDepositBeaconUpgradedIterator{contract: _NodeDeposit.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NodeDeposit *NodeDepositFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *NodeDepositBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NodeDeposit.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDepositBeaconUpgraded)
				if err := _NodeDeposit.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_NodeDeposit *NodeDepositFilterer) ParseBeaconUpgraded(log types.Log) (*NodeDepositBeaconUpgraded, error) {
	event := new(NodeDepositBeaconUpgraded)
	if err := _NodeDeposit.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// NodeDepositInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NodeDeposit contract.
type NodeDepositInitializedIterator struct {
	Event *NodeDepositInitialized // Event containing the contract specifics and raw log

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
func (it *NodeDepositInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDepositInitialized)
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
		it.Event = new(NodeDepositInitialized)
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
func (it *NodeDepositInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDepositInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDepositInitialized represents a Initialized event raised by the NodeDeposit contract.
type NodeDepositInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodeDeposit *NodeDepositFilterer) FilterInitialized(opts *bind.FilterOpts) (*NodeDepositInitializedIterator, error) {

	logs, sub, err := _NodeDeposit.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodeDepositInitializedIterator{contract: _NodeDeposit.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodeDeposit *NodeDepositFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodeDepositInitialized) (event.Subscription, error) {

	logs, sub, err := _NodeDeposit.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDepositInitialized)
				if err := _NodeDeposit.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NodeDeposit *NodeDepositFilterer) ParseInitialized(log types.Log) (*NodeDepositInitialized, error) {
	event := new(NodeDepositInitialized)
	if err := _NodeDeposit.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// NodeDepositUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NodeDeposit contract.
type NodeDepositUpgradedIterator struct {
	Event *NodeDepositUpgraded // Event containing the contract specifics and raw log

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
func (it *NodeDepositUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeDepositUpgraded)
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
		it.Event = new(NodeDepositUpgraded)
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
func (it *NodeDepositUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeDepositUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeDepositUpgraded represents a Upgraded event raised by the NodeDeposit contract.
type NodeDepositUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NodeDeposit *NodeDepositFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NodeDepositUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NodeDeposit.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NodeDepositUpgradedIterator{contract: _NodeDeposit.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NodeDeposit *NodeDepositFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NodeDepositUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NodeDeposit.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeDepositUpgraded)
				if err := _NodeDeposit.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_NodeDeposit *NodeDepositFilterer) ParseUpgraded(log types.Log) (*NodeDepositUpgraded, error) {
	event := new(NodeDepositUpgraded)
	if err := _NodeDeposit.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
