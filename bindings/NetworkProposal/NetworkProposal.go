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
	_ = abi.ConvertType
)

// NetworkProposalMetaData contains all meta data concerning the NetworkProposal contract.
var NetworkProposalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDealedEpoch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDealedHeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyNotifiedCycle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyVoted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountUnmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableDepositZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableRewardZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableWithdrawIndexOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommissionRateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CycleNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositAmountGTMaxAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositAmountLTMinAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LsdTokenAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyRemoved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeNotClaimable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorizedLsdToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotClaimable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPubkeyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTrustNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProposalExecFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyNumberOverLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyStatusUnmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateChangeOverLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SoloNodeDepositAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SoloNodeDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmitBalancesDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"TooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustNodeDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawIndexEmpty\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_proposalId\",\"type\":\"bytes32\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_proposalId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"VoteProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"VoterManagementTakenOver\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tos\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_callDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_proposalFactors\",\"type\":\"uint256[]\"}],\"name\":\"batchExecProposals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_proposalFactor\",\"type\":\"uint256\"}],\"name\":\"execProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposalId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_initialThreshold\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_adminAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voterManagerAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"enumINetworkProposal.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"_yesVotes\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"_yesVotesTotal\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"removeVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_newVoters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"replaceVoters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVoterManager\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_newVoters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"takeoverVoterManagement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVoterManager\",\"type\":\"address\"}],\"name\":\"transferVoterManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	parsed, err := NetworkProposalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NetworkProposal *NetworkProposalCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NetworkProposal.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NetworkProposal *NetworkProposalSession) ProxiableUUID() ([32]byte, error) {
	return _NetworkProposal.Contract.ProxiableUUID(&_NetworkProposal.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NetworkProposal *NetworkProposalCallerSession) ProxiableUUID() ([32]byte, error) {
	return _NetworkProposal.Contract.ProxiableUUID(&_NetworkProposal.CallOpts)
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

// VoterManager is a free data retrieval call binding the contract method 0x65910a4f.
//
// Solidity: function voterManager() view returns(address)
func (_NetworkProposal *NetworkProposalCaller) VoterManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NetworkProposal.contract.Call(opts, &out, "voterManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoterManager is a free data retrieval call binding the contract method 0x65910a4f.
//
// Solidity: function voterManager() view returns(address)
func (_NetworkProposal *NetworkProposalSession) VoterManager() (common.Address, error) {
	return _NetworkProposal.Contract.VoterManager(&_NetworkProposal.CallOpts)
}

// VoterManager is a free data retrieval call binding the contract method 0x65910a4f.
//
// Solidity: function voterManager() view returns(address)
func (_NetworkProposal *NetworkProposalCallerSession) VoterManager() (common.Address, error) {
	return _NetworkProposal.Contract.VoterManager(&_NetworkProposal.CallOpts)
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

// BatchExecProposals is a paid mutator transaction binding the contract method 0x6a32390d.
//
// Solidity: function batchExecProposals(address[] _tos, bytes[] _callDatas, uint256[] _proposalFactors) returns()
func (_NetworkProposal *NetworkProposalTransactor) BatchExecProposals(opts *bind.TransactOpts, _tos []common.Address, _callDatas [][]byte, _proposalFactors []*big.Int) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "batchExecProposals", _tos, _callDatas, _proposalFactors)
}

// BatchExecProposals is a paid mutator transaction binding the contract method 0x6a32390d.
//
// Solidity: function batchExecProposals(address[] _tos, bytes[] _callDatas, uint256[] _proposalFactors) returns()
func (_NetworkProposal *NetworkProposalSession) BatchExecProposals(_tos []common.Address, _callDatas [][]byte, _proposalFactors []*big.Int) (*types.Transaction, error) {
	return _NetworkProposal.Contract.BatchExecProposals(&_NetworkProposal.TransactOpts, _tos, _callDatas, _proposalFactors)
}

// BatchExecProposals is a paid mutator transaction binding the contract method 0x6a32390d.
//
// Solidity: function batchExecProposals(address[] _tos, bytes[] _callDatas, uint256[] _proposalFactors) returns()
func (_NetworkProposal *NetworkProposalTransactorSession) BatchExecProposals(_tos []common.Address, _callDatas [][]byte, _proposalFactors []*big.Int) (*types.Transaction, error) {
	return _NetworkProposal.Contract.BatchExecProposals(&_NetworkProposal.TransactOpts, _tos, _callDatas, _proposalFactors)
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

// ExecProposal is a paid mutator transaction binding the contract method 0x5eb80ef6.
//
// Solidity: function execProposal(address _to, bytes _callData, uint256 _proposalFactor) returns()
func (_NetworkProposal *NetworkProposalTransactor) ExecProposal(opts *bind.TransactOpts, _to common.Address, _callData []byte, _proposalFactor *big.Int) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "execProposal", _to, _callData, _proposalFactor)
}

// ExecProposal is a paid mutator transaction binding the contract method 0x5eb80ef6.
//
// Solidity: function execProposal(address _to, bytes _callData, uint256 _proposalFactor) returns()
func (_NetworkProposal *NetworkProposalSession) ExecProposal(_to common.Address, _callData []byte, _proposalFactor *big.Int) (*types.Transaction, error) {
	return _NetworkProposal.Contract.ExecProposal(&_NetworkProposal.TransactOpts, _to, _callData, _proposalFactor)
}

// ExecProposal is a paid mutator transaction binding the contract method 0x5eb80ef6.
//
// Solidity: function execProposal(address _to, bytes _callData, uint256 _proposalFactor) returns()
func (_NetworkProposal *NetworkProposalTransactorSession) ExecProposal(_to common.Address, _callData []byte, _proposalFactor *big.Int) (*types.Transaction, error) {
	return _NetworkProposal.Contract.ExecProposal(&_NetworkProposal.TransactOpts, _to, _callData, _proposalFactor)
}

// Init is a paid mutator transaction binding the contract method 0x6d844b5c.
//
// Solidity: function init(address[] _voters, uint256 _initialThreshold, address _adminAddress, address _voterManagerAddress) returns()
func (_NetworkProposal *NetworkProposalTransactor) Init(opts *bind.TransactOpts, _voters []common.Address, _initialThreshold *big.Int, _adminAddress common.Address, _voterManagerAddress common.Address) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "init", _voters, _initialThreshold, _adminAddress, _voterManagerAddress)
}

// Init is a paid mutator transaction binding the contract method 0x6d844b5c.
//
// Solidity: function init(address[] _voters, uint256 _initialThreshold, address _adminAddress, address _voterManagerAddress) returns()
func (_NetworkProposal *NetworkProposalSession) Init(_voters []common.Address, _initialThreshold *big.Int, _adminAddress common.Address, _voterManagerAddress common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.Init(&_NetworkProposal.TransactOpts, _voters, _initialThreshold, _adminAddress, _voterManagerAddress)
}

// Init is a paid mutator transaction binding the contract method 0x6d844b5c.
//
// Solidity: function init(address[] _voters, uint256 _initialThreshold, address _adminAddress, address _voterManagerAddress) returns()
func (_NetworkProposal *NetworkProposalTransactorSession) Init(_voters []common.Address, _initialThreshold *big.Int, _adminAddress common.Address, _voterManagerAddress common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.Init(&_NetworkProposal.TransactOpts, _voters, _initialThreshold, _adminAddress, _voterManagerAddress)
}

// Reinit is a paid mutator transaction binding the contract method 0xc482ceaf.
//
// Solidity: function reinit() returns()
func (_NetworkProposal *NetworkProposalTransactor) Reinit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "reinit")
}

// Reinit is a paid mutator transaction binding the contract method 0xc482ceaf.
//
// Solidity: function reinit() returns()
func (_NetworkProposal *NetworkProposalSession) Reinit() (*types.Transaction, error) {
	return _NetworkProposal.Contract.Reinit(&_NetworkProposal.TransactOpts)
}

// Reinit is a paid mutator transaction binding the contract method 0xc482ceaf.
//
// Solidity: function reinit() returns()
func (_NetworkProposal *NetworkProposalTransactorSession) Reinit() (*types.Transaction, error) {
	return _NetworkProposal.Contract.Reinit(&_NetworkProposal.TransactOpts)
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

// ReplaceVoters is a paid mutator transaction binding the contract method 0x02126d4a.
//
// Solidity: function replaceVoters(address[] _newVoters, uint256 _threshold) returns()
func (_NetworkProposal *NetworkProposalTransactor) ReplaceVoters(opts *bind.TransactOpts, _newVoters []common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "replaceVoters", _newVoters, _threshold)
}

// ReplaceVoters is a paid mutator transaction binding the contract method 0x02126d4a.
//
// Solidity: function replaceVoters(address[] _newVoters, uint256 _threshold) returns()
func (_NetworkProposal *NetworkProposalSession) ReplaceVoters(_newVoters []common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _NetworkProposal.Contract.ReplaceVoters(&_NetworkProposal.TransactOpts, _newVoters, _threshold)
}

// ReplaceVoters is a paid mutator transaction binding the contract method 0x02126d4a.
//
// Solidity: function replaceVoters(address[] _newVoters, uint256 _threshold) returns()
func (_NetworkProposal *NetworkProposalTransactorSession) ReplaceVoters(_newVoters []common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _NetworkProposal.Contract.ReplaceVoters(&_NetworkProposal.TransactOpts, _newVoters, _threshold)
}

// TakeoverVoterManagement is a paid mutator transaction binding the contract method 0x07ed5464.
//
// Solidity: function takeoverVoterManagement(address _newVoterManager, address[] _newVoters, uint256 _threshold) returns()
func (_NetworkProposal *NetworkProposalTransactor) TakeoverVoterManagement(opts *bind.TransactOpts, _newVoterManager common.Address, _newVoters []common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "takeoverVoterManagement", _newVoterManager, _newVoters, _threshold)
}

// TakeoverVoterManagement is a paid mutator transaction binding the contract method 0x07ed5464.
//
// Solidity: function takeoverVoterManagement(address _newVoterManager, address[] _newVoters, uint256 _threshold) returns()
func (_NetworkProposal *NetworkProposalSession) TakeoverVoterManagement(_newVoterManager common.Address, _newVoters []common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _NetworkProposal.Contract.TakeoverVoterManagement(&_NetworkProposal.TransactOpts, _newVoterManager, _newVoters, _threshold)
}

// TakeoverVoterManagement is a paid mutator transaction binding the contract method 0x07ed5464.
//
// Solidity: function takeoverVoterManagement(address _newVoterManager, address[] _newVoters, uint256 _threshold) returns()
func (_NetworkProposal *NetworkProposalTransactorSession) TakeoverVoterManagement(_newVoterManager common.Address, _newVoters []common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _NetworkProposal.Contract.TakeoverVoterManagement(&_NetworkProposal.TransactOpts, _newVoterManager, _newVoters, _threshold)
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

// TransferVoterManager is a paid mutator transaction binding the contract method 0xb88c35be.
//
// Solidity: function transferVoterManager(address _newVoterManager) returns()
func (_NetworkProposal *NetworkProposalTransactor) TransferVoterManager(opts *bind.TransactOpts, _newVoterManager common.Address) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "transferVoterManager", _newVoterManager)
}

// TransferVoterManager is a paid mutator transaction binding the contract method 0xb88c35be.
//
// Solidity: function transferVoterManager(address _newVoterManager) returns()
func (_NetworkProposal *NetworkProposalSession) TransferVoterManager(_newVoterManager common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.TransferVoterManager(&_NetworkProposal.TransactOpts, _newVoterManager)
}

// TransferVoterManager is a paid mutator transaction binding the contract method 0xb88c35be.
//
// Solidity: function transferVoterManager(address _newVoterManager) returns()
func (_NetworkProposal *NetworkProposalTransactorSession) TransferVoterManager(_newVoterManager common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.TransferVoterManager(&_NetworkProposal.TransactOpts, _newVoterManager)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NetworkProposal *NetworkProposalTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NetworkProposal *NetworkProposalSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.UpgradeTo(&_NetworkProposal.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NetworkProposal *NetworkProposalTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NetworkProposal.Contract.UpgradeTo(&_NetworkProposal.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NetworkProposal *NetworkProposalTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NetworkProposal.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NetworkProposal *NetworkProposalSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NetworkProposal.Contract.UpgradeToAndCall(&_NetworkProposal.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NetworkProposal *NetworkProposalTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NetworkProposal.Contract.UpgradeToAndCall(&_NetworkProposal.TransactOpts, newImplementation, data)
}

// NetworkProposalAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the NetworkProposal contract.
type NetworkProposalAdminChangedIterator struct {
	Event *NetworkProposalAdminChanged // Event containing the contract specifics and raw log

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
func (it *NetworkProposalAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkProposalAdminChanged)
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
		it.Event = new(NetworkProposalAdminChanged)
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
func (it *NetworkProposalAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkProposalAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkProposalAdminChanged represents a AdminChanged event raised by the NetworkProposal contract.
type NetworkProposalAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NetworkProposal *NetworkProposalFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*NetworkProposalAdminChangedIterator, error) {

	logs, sub, err := _NetworkProposal.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &NetworkProposalAdminChangedIterator{contract: _NetworkProposal.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NetworkProposal *NetworkProposalFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *NetworkProposalAdminChanged) (event.Subscription, error) {

	logs, sub, err := _NetworkProposal.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkProposalAdminChanged)
				if err := _NetworkProposal.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_NetworkProposal *NetworkProposalFilterer) ParseAdminChanged(log types.Log) (*NetworkProposalAdminChanged, error) {
	event := new(NetworkProposalAdminChanged)
	if err := _NetworkProposal.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkProposalBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the NetworkProposal contract.
type NetworkProposalBeaconUpgradedIterator struct {
	Event *NetworkProposalBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *NetworkProposalBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkProposalBeaconUpgraded)
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
		it.Event = new(NetworkProposalBeaconUpgraded)
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
func (it *NetworkProposalBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkProposalBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkProposalBeaconUpgraded represents a BeaconUpgraded event raised by the NetworkProposal contract.
type NetworkProposalBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NetworkProposal *NetworkProposalFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*NetworkProposalBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NetworkProposal.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &NetworkProposalBeaconUpgradedIterator{contract: _NetworkProposal.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NetworkProposal *NetworkProposalFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *NetworkProposalBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NetworkProposal.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkProposalBeaconUpgraded)
				if err := _NetworkProposal.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_NetworkProposal *NetworkProposalFilterer) ParseBeaconUpgraded(log types.Log) (*NetworkProposalBeaconUpgraded, error) {
	event := new(NetworkProposalBeaconUpgraded)
	if err := _NetworkProposal.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkProposalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NetworkProposal contract.
type NetworkProposalInitializedIterator struct {
	Event *NetworkProposalInitialized // Event containing the contract specifics and raw log

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
func (it *NetworkProposalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkProposalInitialized)
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
		it.Event = new(NetworkProposalInitialized)
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
func (it *NetworkProposalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkProposalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkProposalInitialized represents a Initialized event raised by the NetworkProposal contract.
type NetworkProposalInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NetworkProposal *NetworkProposalFilterer) FilterInitialized(opts *bind.FilterOpts) (*NetworkProposalInitializedIterator, error) {

	logs, sub, err := _NetworkProposal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NetworkProposalInitializedIterator{contract: _NetworkProposal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NetworkProposal *NetworkProposalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NetworkProposalInitialized) (event.Subscription, error) {

	logs, sub, err := _NetworkProposal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkProposalInitialized)
				if err := _NetworkProposal.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NetworkProposal *NetworkProposalFilterer) ParseInitialized(log types.Log) (*NetworkProposalInitialized, error) {
	event := new(NetworkProposalInitialized)
	if err := _NetworkProposal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
// Solidity: event ProposalExecuted(bytes32 indexed _proposalId)
func (_NetworkProposal *NetworkProposalFilterer) FilterProposalExecuted(opts *bind.FilterOpts, _proposalId [][32]byte) (*NetworkProposalProposalExecutedIterator, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _NetworkProposal.contract.FilterLogs(opts, "ProposalExecuted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &NetworkProposalProposalExecutedIterator{contract: _NetworkProposal.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 indexed _proposalId)
func (_NetworkProposal *NetworkProposalFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *NetworkProposalProposalExecuted, _proposalId [][32]byte) (event.Subscription, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _NetworkProposal.contract.WatchLogs(opts, "ProposalExecuted", _proposalIdRule)
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
// Solidity: event ProposalExecuted(bytes32 indexed _proposalId)
func (_NetworkProposal *NetworkProposalFilterer) ParseProposalExecuted(log types.Log) (*NetworkProposalProposalExecuted, error) {
	event := new(NetworkProposalProposalExecuted)
	if err := _NetworkProposal.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkProposalUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NetworkProposal contract.
type NetworkProposalUpgradedIterator struct {
	Event *NetworkProposalUpgraded // Event containing the contract specifics and raw log

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
func (it *NetworkProposalUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkProposalUpgraded)
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
		it.Event = new(NetworkProposalUpgraded)
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
func (it *NetworkProposalUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkProposalUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkProposalUpgraded represents a Upgraded event raised by the NetworkProposal contract.
type NetworkProposalUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NetworkProposal *NetworkProposalFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NetworkProposalUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NetworkProposal.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NetworkProposalUpgradedIterator{contract: _NetworkProposal.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NetworkProposal *NetworkProposalFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NetworkProposalUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NetworkProposal.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkProposalUpgraded)
				if err := _NetworkProposal.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_NetworkProposal *NetworkProposalFilterer) ParseUpgraded(log types.Log) (*NetworkProposalUpgraded, error) {
	event := new(NetworkProposalUpgraded)
	if err := _NetworkProposal.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
// Solidity: event VoteProposal(bytes32 indexed _proposalId, address _voter)
func (_NetworkProposal *NetworkProposalFilterer) FilterVoteProposal(opts *bind.FilterOpts, _proposalId [][32]byte) (*NetworkProposalVoteProposalIterator, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _NetworkProposal.contract.FilterLogs(opts, "VoteProposal", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &NetworkProposalVoteProposalIterator{contract: _NetworkProposal.contract, event: "VoteProposal", logs: logs, sub: sub}, nil
}

// WatchVoteProposal is a free log subscription operation binding the contract event 0x3b58f01618556cdc5e9f7b0f1f6dccbac40024bc1043f589bd4a324e3414cfad.
//
// Solidity: event VoteProposal(bytes32 indexed _proposalId, address _voter)
func (_NetworkProposal *NetworkProposalFilterer) WatchVoteProposal(opts *bind.WatchOpts, sink chan<- *NetworkProposalVoteProposal, _proposalId [][32]byte) (event.Subscription, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _NetworkProposal.contract.WatchLogs(opts, "VoteProposal", _proposalIdRule)
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
// Solidity: event VoteProposal(bytes32 indexed _proposalId, address _voter)
func (_NetworkProposal *NetworkProposalFilterer) ParseVoteProposal(log types.Log) (*NetworkProposalVoteProposal, error) {
	event := new(NetworkProposalVoteProposal)
	if err := _NetworkProposal.contract.UnpackLog(event, "VoteProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkProposalVoterManagementTakenOverIterator is returned from FilterVoterManagementTakenOver and is used to iterate over the raw logs and unpacked data for VoterManagementTakenOver events raised by the NetworkProposal contract.
type NetworkProposalVoterManagementTakenOverIterator struct {
	Event *NetworkProposalVoterManagementTakenOver // Event containing the contract specifics and raw log

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
func (it *NetworkProposalVoterManagementTakenOverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkProposalVoterManagementTakenOver)
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
		it.Event = new(NetworkProposalVoterManagementTakenOver)
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
func (it *NetworkProposalVoterManagementTakenOverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkProposalVoterManagementTakenOverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkProposalVoterManagementTakenOver represents a VoterManagementTakenOver event raised by the NetworkProposal contract.
type NetworkProposalVoterManagementTakenOver struct {
	OldManager common.Address
	NewManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoterManagementTakenOver is a free log retrieval operation binding the contract event 0x77cc016861b4e095bccd1c6d24dce5f502f9906d955ddb2b96fb9ea44be9f8a7.
//
// Solidity: event VoterManagementTakenOver(address indexed _oldManager, address indexed _newManager)
func (_NetworkProposal *NetworkProposalFilterer) FilterVoterManagementTakenOver(opts *bind.FilterOpts, _oldManager []common.Address, _newManager []common.Address) (*NetworkProposalVoterManagementTakenOverIterator, error) {

	var _oldManagerRule []interface{}
	for _, _oldManagerItem := range _oldManager {
		_oldManagerRule = append(_oldManagerRule, _oldManagerItem)
	}
	var _newManagerRule []interface{}
	for _, _newManagerItem := range _newManager {
		_newManagerRule = append(_newManagerRule, _newManagerItem)
	}

	logs, sub, err := _NetworkProposal.contract.FilterLogs(opts, "VoterManagementTakenOver", _oldManagerRule, _newManagerRule)
	if err != nil {
		return nil, err
	}
	return &NetworkProposalVoterManagementTakenOverIterator{contract: _NetworkProposal.contract, event: "VoterManagementTakenOver", logs: logs, sub: sub}, nil
}

// WatchVoterManagementTakenOver is a free log subscription operation binding the contract event 0x77cc016861b4e095bccd1c6d24dce5f502f9906d955ddb2b96fb9ea44be9f8a7.
//
// Solidity: event VoterManagementTakenOver(address indexed _oldManager, address indexed _newManager)
func (_NetworkProposal *NetworkProposalFilterer) WatchVoterManagementTakenOver(opts *bind.WatchOpts, sink chan<- *NetworkProposalVoterManagementTakenOver, _oldManager []common.Address, _newManager []common.Address) (event.Subscription, error) {

	var _oldManagerRule []interface{}
	for _, _oldManagerItem := range _oldManager {
		_oldManagerRule = append(_oldManagerRule, _oldManagerItem)
	}
	var _newManagerRule []interface{}
	for _, _newManagerItem := range _newManager {
		_newManagerRule = append(_newManagerRule, _newManagerItem)
	}

	logs, sub, err := _NetworkProposal.contract.WatchLogs(opts, "VoterManagementTakenOver", _oldManagerRule, _newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkProposalVoterManagementTakenOver)
				if err := _NetworkProposal.contract.UnpackLog(event, "VoterManagementTakenOver", log); err != nil {
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

// ParseVoterManagementTakenOver is a log parse operation binding the contract event 0x77cc016861b4e095bccd1c6d24dce5f502f9906d955ddb2b96fb9ea44be9f8a7.
//
// Solidity: event VoterManagementTakenOver(address indexed _oldManager, address indexed _newManager)
func (_NetworkProposal *NetworkProposalFilterer) ParseVoterManagementTakenOver(log types.Log) (*NetworkProposalVoterManagementTakenOver, error) {
	event := new(NetworkProposalVoterManagementTakenOver)
	if err := _NetworkProposal.contract.UnpackLog(event, "VoterManagementTakenOver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
