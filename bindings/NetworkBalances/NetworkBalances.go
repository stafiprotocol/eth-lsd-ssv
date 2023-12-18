// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package network_balances

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

// NetworkBalancesMetaData contains all meta data concerning the NetworkBalances contract.
var NetworkBalancesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDealedEpoch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDealedHeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyNotifiedCycle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyVoted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountUnmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableDepositZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableRewardZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimableWithdrawIndexOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommissionRateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CycleNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositAmountGTMaxAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositAmountLTMinAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LsdTokenAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyRemoved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeNotClaimable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorizedLsdToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotClaimable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPubkeyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTrustNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProposalExecFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyNumberOverLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyStatusUnmatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateChangeOverLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SoloNodeDepositAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SoloNodeDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmitBalancesDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"TooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustNodeDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserDepositDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VotersNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawIndexEmpty\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lsdTokenSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BalancesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"balancesSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalLsdToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lsdTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"getLsdTokenValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_networkProposalAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkProposalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateChangeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setRateChangeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setUpdateBalancesEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalLsdToken\",\"type\":\"uint256\"}],\"name\":\"submitBalances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitBalancesEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateBalancesEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NetworkBalancesABI is the input ABI used to generate the binding from.
// Deprecated: Use NetworkBalancesMetaData.ABI instead.
var NetworkBalancesABI = NetworkBalancesMetaData.ABI

// NetworkBalances is an auto generated Go binding around an Ethereum contract.
type NetworkBalances struct {
	NetworkBalancesCaller     // Read-only binding to the contract
	NetworkBalancesTransactor // Write-only binding to the contract
	NetworkBalancesFilterer   // Log filterer for contract events
}

// NetworkBalancesCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkBalancesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkBalancesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkBalancesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkBalancesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NetworkBalancesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkBalancesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkBalancesSession struct {
	Contract     *NetworkBalances  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NetworkBalancesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkBalancesCallerSession struct {
	Contract *NetworkBalancesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NetworkBalancesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkBalancesTransactorSession struct {
	Contract     *NetworkBalancesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NetworkBalancesRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkBalancesRaw struct {
	Contract *NetworkBalances // Generic contract binding to access the raw methods on
}

// NetworkBalancesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkBalancesCallerRaw struct {
	Contract *NetworkBalancesCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkBalancesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkBalancesTransactorRaw struct {
	Contract *NetworkBalancesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetworkBalances creates a new instance of NetworkBalances, bound to a specific deployed contract.
func NewNetworkBalances(address common.Address, backend bind.ContractBackend) (*NetworkBalances, error) {
	contract, err := bindNetworkBalances(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NetworkBalances{NetworkBalancesCaller: NetworkBalancesCaller{contract: contract}, NetworkBalancesTransactor: NetworkBalancesTransactor{contract: contract}, NetworkBalancesFilterer: NetworkBalancesFilterer{contract: contract}}, nil
}

// NewNetworkBalancesCaller creates a new read-only instance of NetworkBalances, bound to a specific deployed contract.
func NewNetworkBalancesCaller(address common.Address, caller bind.ContractCaller) (*NetworkBalancesCaller, error) {
	contract, err := bindNetworkBalances(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkBalancesCaller{contract: contract}, nil
}

// NewNetworkBalancesTransactor creates a new write-only instance of NetworkBalances, bound to a specific deployed contract.
func NewNetworkBalancesTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkBalancesTransactor, error) {
	contract, err := bindNetworkBalances(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkBalancesTransactor{contract: contract}, nil
}

// NewNetworkBalancesFilterer creates a new log filterer instance of NetworkBalances, bound to a specific deployed contract.
func NewNetworkBalancesFilterer(address common.Address, filterer bind.ContractFilterer) (*NetworkBalancesFilterer, error) {
	contract, err := bindNetworkBalances(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NetworkBalancesFilterer{contract: contract}, nil
}

// bindNetworkBalances binds a generic wrapper to an already deployed contract.
func bindNetworkBalances(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NetworkBalancesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkBalances *NetworkBalancesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkBalances.Contract.NetworkBalancesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkBalances *NetworkBalancesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkBalances.Contract.NetworkBalancesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkBalances *NetworkBalancesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkBalances.Contract.NetworkBalancesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkBalances *NetworkBalancesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkBalances.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkBalances *NetworkBalancesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkBalances.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkBalances *NetworkBalancesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkBalances.Contract.contract.Transact(opts, method, params...)
}

// BalancesSnapshot is a free data retrieval call binding the contract method 0xcc386bf5.
//
// Solidity: function balancesSnapshot() view returns(uint256 _block, uint256 _totalEth, uint256 _totalLsdToken)
func (_NetworkBalances *NetworkBalancesCaller) BalancesSnapshot(opts *bind.CallOpts) (struct {
	Block         *big.Int
	TotalEth      *big.Int
	TotalLsdToken *big.Int
}, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "balancesSnapshot")

	outstruct := new(struct {
		Block         *big.Int
		TotalEth      *big.Int
		TotalLsdToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Block = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalEth = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalLsdToken = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BalancesSnapshot is a free data retrieval call binding the contract method 0xcc386bf5.
//
// Solidity: function balancesSnapshot() view returns(uint256 _block, uint256 _totalEth, uint256 _totalLsdToken)
func (_NetworkBalances *NetworkBalancesSession) BalancesSnapshot() (struct {
	Block         *big.Int
	TotalEth      *big.Int
	TotalLsdToken *big.Int
}, error) {
	return _NetworkBalances.Contract.BalancesSnapshot(&_NetworkBalances.CallOpts)
}

// BalancesSnapshot is a free data retrieval call binding the contract method 0xcc386bf5.
//
// Solidity: function balancesSnapshot() view returns(uint256 _block, uint256 _totalEth, uint256 _totalLsdToken)
func (_NetworkBalances *NetworkBalancesCallerSession) BalancesSnapshot() (struct {
	Block         *big.Int
	TotalEth      *big.Int
	TotalLsdToken *big.Int
}, error) {
	return _NetworkBalances.Contract.BalancesSnapshot(&_NetworkBalances.CallOpts)
}

// GetEthValue is a free data retrieval call binding the contract method 0x8b32fa23.
//
// Solidity: function getEthValue(uint256 _lsdTokenAmount) view returns(uint256)
func (_NetworkBalances *NetworkBalancesCaller) GetEthValue(opts *bind.CallOpts, _lsdTokenAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "getEthValue", _lsdTokenAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthValue is a free data retrieval call binding the contract method 0x8b32fa23.
//
// Solidity: function getEthValue(uint256 _lsdTokenAmount) view returns(uint256)
func (_NetworkBalances *NetworkBalancesSession) GetEthValue(_lsdTokenAmount *big.Int) (*big.Int, error) {
	return _NetworkBalances.Contract.GetEthValue(&_NetworkBalances.CallOpts, _lsdTokenAmount)
}

// GetEthValue is a free data retrieval call binding the contract method 0x8b32fa23.
//
// Solidity: function getEthValue(uint256 _lsdTokenAmount) view returns(uint256)
func (_NetworkBalances *NetworkBalancesCallerSession) GetEthValue(_lsdTokenAmount *big.Int) (*big.Int, error) {
	return _NetworkBalances.Contract.GetEthValue(&_NetworkBalances.CallOpts, _lsdTokenAmount)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCaller) GetExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "getExchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_NetworkBalances *NetworkBalancesSession) GetExchangeRate() (*big.Int, error) {
	return _NetworkBalances.Contract.GetExchangeRate(&_NetworkBalances.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCallerSession) GetExchangeRate() (*big.Int, error) {
	return _NetworkBalances.Contract.GetExchangeRate(&_NetworkBalances.CallOpts)
}

// GetLsdTokenValue is a free data retrieval call binding the contract method 0xf66ace79.
//
// Solidity: function getLsdTokenValue(uint256 _ethAmount) view returns(uint256)
func (_NetworkBalances *NetworkBalancesCaller) GetLsdTokenValue(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "getLsdTokenValue", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLsdTokenValue is a free data retrieval call binding the contract method 0xf66ace79.
//
// Solidity: function getLsdTokenValue(uint256 _ethAmount) view returns(uint256)
func (_NetworkBalances *NetworkBalancesSession) GetLsdTokenValue(_ethAmount *big.Int) (*big.Int, error) {
	return _NetworkBalances.Contract.GetLsdTokenValue(&_NetworkBalances.CallOpts, _ethAmount)
}

// GetLsdTokenValue is a free data retrieval call binding the contract method 0xf66ace79.
//
// Solidity: function getLsdTokenValue(uint256 _ethAmount) view returns(uint256)
func (_NetworkBalances *NetworkBalancesCallerSession) GetLsdTokenValue(_ethAmount *big.Int) (*big.Int, error) {
	return _NetworkBalances.Contract.GetLsdTokenValue(&_NetworkBalances.CallOpts, _ethAmount)
}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_NetworkBalances *NetworkBalancesCaller) NetworkProposalAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "networkProposalAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_NetworkBalances *NetworkBalancesSession) NetworkProposalAddress() (common.Address, error) {
	return _NetworkBalances.Contract.NetworkProposalAddress(&_NetworkBalances.CallOpts)
}

// NetworkProposalAddress is a free data retrieval call binding the contract method 0xb4701c09.
//
// Solidity: function networkProposalAddress() view returns(address)
func (_NetworkBalances *NetworkBalancesCallerSession) NetworkProposalAddress() (common.Address, error) {
	return _NetworkBalances.Contract.NetworkProposalAddress(&_NetworkBalances.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NetworkBalances *NetworkBalancesCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NetworkBalances *NetworkBalancesSession) ProxiableUUID() ([32]byte, error) {
	return _NetworkBalances.Contract.ProxiableUUID(&_NetworkBalances.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NetworkBalances *NetworkBalancesCallerSession) ProxiableUUID() ([32]byte, error) {
	return _NetworkBalances.Contract.ProxiableUUID(&_NetworkBalances.CallOpts)
}

// RateChangeLimit is a free data retrieval call binding the contract method 0xc0152b71.
//
// Solidity: function rateChangeLimit() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCaller) RateChangeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "rateChangeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RateChangeLimit is a free data retrieval call binding the contract method 0xc0152b71.
//
// Solidity: function rateChangeLimit() view returns(uint256)
func (_NetworkBalances *NetworkBalancesSession) RateChangeLimit() (*big.Int, error) {
	return _NetworkBalances.Contract.RateChangeLimit(&_NetworkBalances.CallOpts)
}

// RateChangeLimit is a free data retrieval call binding the contract method 0xc0152b71.
//
// Solidity: function rateChangeLimit() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCallerSession) RateChangeLimit() (*big.Int, error) {
	return _NetworkBalances.Contract.RateChangeLimit(&_NetworkBalances.CallOpts)
}

// SubmitBalancesEnabled is a free data retrieval call binding the contract method 0x5414817f.
//
// Solidity: function submitBalancesEnabled() view returns(bool)
func (_NetworkBalances *NetworkBalancesCaller) SubmitBalancesEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "submitBalancesEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SubmitBalancesEnabled is a free data retrieval call binding the contract method 0x5414817f.
//
// Solidity: function submitBalancesEnabled() view returns(bool)
func (_NetworkBalances *NetworkBalancesSession) SubmitBalancesEnabled() (bool, error) {
	return _NetworkBalances.Contract.SubmitBalancesEnabled(&_NetworkBalances.CallOpts)
}

// SubmitBalancesEnabled is a free data retrieval call binding the contract method 0x5414817f.
//
// Solidity: function submitBalancesEnabled() view returns(bool)
func (_NetworkBalances *NetworkBalancesCallerSession) SubmitBalancesEnabled() (bool, error) {
	return _NetworkBalances.Contract.SubmitBalancesEnabled(&_NetworkBalances.CallOpts)
}

// UpdateBalancesEpochs is a free data retrieval call binding the contract method 0x8fad5ce5.
//
// Solidity: function updateBalancesEpochs() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCaller) UpdateBalancesEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "updateBalancesEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdateBalancesEpochs is a free data retrieval call binding the contract method 0x8fad5ce5.
//
// Solidity: function updateBalancesEpochs() view returns(uint256)
func (_NetworkBalances *NetworkBalancesSession) UpdateBalancesEpochs() (*big.Int, error) {
	return _NetworkBalances.Contract.UpdateBalancesEpochs(&_NetworkBalances.CallOpts)
}

// UpdateBalancesEpochs is a free data retrieval call binding the contract method 0x8fad5ce5.
//
// Solidity: function updateBalancesEpochs() view returns(uint256)
func (_NetworkBalances *NetworkBalancesCallerSession) UpdateBalancesEpochs() (*big.Int, error) {
	return _NetworkBalances.Contract.UpdateBalancesEpochs(&_NetworkBalances.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkBalances *NetworkBalancesCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NetworkBalances.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkBalances *NetworkBalancesSession) Version() (uint8, error) {
	return _NetworkBalances.Contract.Version(&_NetworkBalances.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_NetworkBalances *NetworkBalancesCallerSession) Version() (uint8, error) {
	return _NetworkBalances.Contract.Version(&_NetworkBalances.CallOpts)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _networkProposalAddress) returns()
func (_NetworkBalances *NetworkBalancesTransactor) Init(opts *bind.TransactOpts, _networkProposalAddress common.Address) (*types.Transaction, error) {
	return _NetworkBalances.contract.Transact(opts, "init", _networkProposalAddress)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _networkProposalAddress) returns()
func (_NetworkBalances *NetworkBalancesSession) Init(_networkProposalAddress common.Address) (*types.Transaction, error) {
	return _NetworkBalances.Contract.Init(&_NetworkBalances.TransactOpts, _networkProposalAddress)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _networkProposalAddress) returns()
func (_NetworkBalances *NetworkBalancesTransactorSession) Init(_networkProposalAddress common.Address) (*types.Transaction, error) {
	return _NetworkBalances.Contract.Init(&_NetworkBalances.TransactOpts, _networkProposalAddress)
}

// Reinit is a paid mutator transaction binding the contract method 0xc482ceaf.
//
// Solidity: function reinit() returns()
func (_NetworkBalances *NetworkBalancesTransactor) Reinit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkBalances.contract.Transact(opts, "reinit")
}

// Reinit is a paid mutator transaction binding the contract method 0xc482ceaf.
//
// Solidity: function reinit() returns()
func (_NetworkBalances *NetworkBalancesSession) Reinit() (*types.Transaction, error) {
	return _NetworkBalances.Contract.Reinit(&_NetworkBalances.TransactOpts)
}

// Reinit is a paid mutator transaction binding the contract method 0xc482ceaf.
//
// Solidity: function reinit() returns()
func (_NetworkBalances *NetworkBalancesTransactorSession) Reinit() (*types.Transaction, error) {
	return _NetworkBalances.Contract.Reinit(&_NetworkBalances.TransactOpts)
}

// SetRateChangeLimit is a paid mutator transaction binding the contract method 0x19826e71.
//
// Solidity: function setRateChangeLimit(uint256 _value) returns()
func (_NetworkBalances *NetworkBalancesTransactor) SetRateChangeLimit(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _NetworkBalances.contract.Transact(opts, "setRateChangeLimit", _value)
}

// SetRateChangeLimit is a paid mutator transaction binding the contract method 0x19826e71.
//
// Solidity: function setRateChangeLimit(uint256 _value) returns()
func (_NetworkBalances *NetworkBalancesSession) SetRateChangeLimit(_value *big.Int) (*types.Transaction, error) {
	return _NetworkBalances.Contract.SetRateChangeLimit(&_NetworkBalances.TransactOpts, _value)
}

// SetRateChangeLimit is a paid mutator transaction binding the contract method 0x19826e71.
//
// Solidity: function setRateChangeLimit(uint256 _value) returns()
func (_NetworkBalances *NetworkBalancesTransactorSession) SetRateChangeLimit(_value *big.Int) (*types.Transaction, error) {
	return _NetworkBalances.Contract.SetRateChangeLimit(&_NetworkBalances.TransactOpts, _value)
}

// SetUpdateBalancesEpochs is a paid mutator transaction binding the contract method 0xd5480be0.
//
// Solidity: function setUpdateBalancesEpochs(uint256 _value) returns()
func (_NetworkBalances *NetworkBalancesTransactor) SetUpdateBalancesEpochs(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _NetworkBalances.contract.Transact(opts, "setUpdateBalancesEpochs", _value)
}

// SetUpdateBalancesEpochs is a paid mutator transaction binding the contract method 0xd5480be0.
//
// Solidity: function setUpdateBalancesEpochs(uint256 _value) returns()
func (_NetworkBalances *NetworkBalancesSession) SetUpdateBalancesEpochs(_value *big.Int) (*types.Transaction, error) {
	return _NetworkBalances.Contract.SetUpdateBalancesEpochs(&_NetworkBalances.TransactOpts, _value)
}

// SetUpdateBalancesEpochs is a paid mutator transaction binding the contract method 0xd5480be0.
//
// Solidity: function setUpdateBalancesEpochs(uint256 _value) returns()
func (_NetworkBalances *NetworkBalancesTransactorSession) SetUpdateBalancesEpochs(_value *big.Int) (*types.Transaction, error) {
	return _NetworkBalances.Contract.SetUpdateBalancesEpochs(&_NetworkBalances.TransactOpts, _value)
}

// SubmitBalances is a paid mutator transaction binding the contract method 0x4c5212a5.
//
// Solidity: function submitBalances(uint256 _block, uint256 _totalEth, uint256 _totalLsdToken) returns()
func (_NetworkBalances *NetworkBalancesTransactor) SubmitBalances(opts *bind.TransactOpts, _block *big.Int, _totalEth *big.Int, _totalLsdToken *big.Int) (*types.Transaction, error) {
	return _NetworkBalances.contract.Transact(opts, "submitBalances", _block, _totalEth, _totalLsdToken)
}

// SubmitBalances is a paid mutator transaction binding the contract method 0x4c5212a5.
//
// Solidity: function submitBalances(uint256 _block, uint256 _totalEth, uint256 _totalLsdToken) returns()
func (_NetworkBalances *NetworkBalancesSession) SubmitBalances(_block *big.Int, _totalEth *big.Int, _totalLsdToken *big.Int) (*types.Transaction, error) {
	return _NetworkBalances.Contract.SubmitBalances(&_NetworkBalances.TransactOpts, _block, _totalEth, _totalLsdToken)
}

// SubmitBalances is a paid mutator transaction binding the contract method 0x4c5212a5.
//
// Solidity: function submitBalances(uint256 _block, uint256 _totalEth, uint256 _totalLsdToken) returns()
func (_NetworkBalances *NetworkBalancesTransactorSession) SubmitBalances(_block *big.Int, _totalEth *big.Int, _totalLsdToken *big.Int) (*types.Transaction, error) {
	return _NetworkBalances.Contract.SubmitBalances(&_NetworkBalances.TransactOpts, _block, _totalEth, _totalLsdToken)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NetworkBalances *NetworkBalancesTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _NetworkBalances.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NetworkBalances *NetworkBalancesSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NetworkBalances.Contract.UpgradeTo(&_NetworkBalances.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NetworkBalances *NetworkBalancesTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NetworkBalances.Contract.UpgradeTo(&_NetworkBalances.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NetworkBalances *NetworkBalancesTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NetworkBalances.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NetworkBalances *NetworkBalancesSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NetworkBalances.Contract.UpgradeToAndCall(&_NetworkBalances.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NetworkBalances *NetworkBalancesTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NetworkBalances.Contract.UpgradeToAndCall(&_NetworkBalances.TransactOpts, newImplementation, data)
}

// NetworkBalancesAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the NetworkBalances contract.
type NetworkBalancesAdminChangedIterator struct {
	Event *NetworkBalancesAdminChanged // Event containing the contract specifics and raw log

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
func (it *NetworkBalancesAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkBalancesAdminChanged)
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
		it.Event = new(NetworkBalancesAdminChanged)
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
func (it *NetworkBalancesAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkBalancesAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkBalancesAdminChanged represents a AdminChanged event raised by the NetworkBalances contract.
type NetworkBalancesAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NetworkBalances *NetworkBalancesFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*NetworkBalancesAdminChangedIterator, error) {

	logs, sub, err := _NetworkBalances.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &NetworkBalancesAdminChangedIterator{contract: _NetworkBalances.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NetworkBalances *NetworkBalancesFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *NetworkBalancesAdminChanged) (event.Subscription, error) {

	logs, sub, err := _NetworkBalances.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkBalancesAdminChanged)
				if err := _NetworkBalances.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_NetworkBalances *NetworkBalancesFilterer) ParseAdminChanged(log types.Log) (*NetworkBalancesAdminChanged, error) {
	event := new(NetworkBalancesAdminChanged)
	if err := _NetworkBalances.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkBalancesBalancesUpdatedIterator is returned from FilterBalancesUpdated and is used to iterate over the raw logs and unpacked data for BalancesUpdated events raised by the NetworkBalances contract.
type NetworkBalancesBalancesUpdatedIterator struct {
	Event *NetworkBalancesBalancesUpdated // Event containing the contract specifics and raw log

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
func (it *NetworkBalancesBalancesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkBalancesBalancesUpdated)
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
		it.Event = new(NetworkBalancesBalancesUpdated)
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
func (it *NetworkBalancesBalancesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkBalancesBalancesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkBalancesBalancesUpdated represents a BalancesUpdated event raised by the NetworkBalances contract.
type NetworkBalancesBalancesUpdated struct {
	Block          *big.Int
	TotalEth       *big.Int
	LsdTokenSupply *big.Int
	Time           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBalancesUpdated is a free log retrieval operation binding the contract event 0x6d217af0d189a7286db93d91c87a48b3b03a6366652338f42f05064790ae0497.
//
// Solidity: event BalancesUpdated(uint256 block, uint256 totalEth, uint256 lsdTokenSupply, uint256 time)
func (_NetworkBalances *NetworkBalancesFilterer) FilterBalancesUpdated(opts *bind.FilterOpts) (*NetworkBalancesBalancesUpdatedIterator, error) {

	logs, sub, err := _NetworkBalances.contract.FilterLogs(opts, "BalancesUpdated")
	if err != nil {
		return nil, err
	}
	return &NetworkBalancesBalancesUpdatedIterator{contract: _NetworkBalances.contract, event: "BalancesUpdated", logs: logs, sub: sub}, nil
}

// WatchBalancesUpdated is a free log subscription operation binding the contract event 0x6d217af0d189a7286db93d91c87a48b3b03a6366652338f42f05064790ae0497.
//
// Solidity: event BalancesUpdated(uint256 block, uint256 totalEth, uint256 lsdTokenSupply, uint256 time)
func (_NetworkBalances *NetworkBalancesFilterer) WatchBalancesUpdated(opts *bind.WatchOpts, sink chan<- *NetworkBalancesBalancesUpdated) (event.Subscription, error) {

	logs, sub, err := _NetworkBalances.contract.WatchLogs(opts, "BalancesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkBalancesBalancesUpdated)
				if err := _NetworkBalances.contract.UnpackLog(event, "BalancesUpdated", log); err != nil {
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

// ParseBalancesUpdated is a log parse operation binding the contract event 0x6d217af0d189a7286db93d91c87a48b3b03a6366652338f42f05064790ae0497.
//
// Solidity: event BalancesUpdated(uint256 block, uint256 totalEth, uint256 lsdTokenSupply, uint256 time)
func (_NetworkBalances *NetworkBalancesFilterer) ParseBalancesUpdated(log types.Log) (*NetworkBalancesBalancesUpdated, error) {
	event := new(NetworkBalancesBalancesUpdated)
	if err := _NetworkBalances.contract.UnpackLog(event, "BalancesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkBalancesBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the NetworkBalances contract.
type NetworkBalancesBeaconUpgradedIterator struct {
	Event *NetworkBalancesBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *NetworkBalancesBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkBalancesBeaconUpgraded)
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
		it.Event = new(NetworkBalancesBeaconUpgraded)
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
func (it *NetworkBalancesBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkBalancesBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkBalancesBeaconUpgraded represents a BeaconUpgraded event raised by the NetworkBalances contract.
type NetworkBalancesBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NetworkBalances *NetworkBalancesFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*NetworkBalancesBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NetworkBalances.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &NetworkBalancesBeaconUpgradedIterator{contract: _NetworkBalances.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NetworkBalances *NetworkBalancesFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *NetworkBalancesBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NetworkBalances.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkBalancesBeaconUpgraded)
				if err := _NetworkBalances.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_NetworkBalances *NetworkBalancesFilterer) ParseBeaconUpgraded(log types.Log) (*NetworkBalancesBeaconUpgraded, error) {
	event := new(NetworkBalancesBeaconUpgraded)
	if err := _NetworkBalances.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkBalancesInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NetworkBalances contract.
type NetworkBalancesInitializedIterator struct {
	Event *NetworkBalancesInitialized // Event containing the contract specifics and raw log

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
func (it *NetworkBalancesInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkBalancesInitialized)
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
		it.Event = new(NetworkBalancesInitialized)
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
func (it *NetworkBalancesInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkBalancesInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkBalancesInitialized represents a Initialized event raised by the NetworkBalances contract.
type NetworkBalancesInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NetworkBalances *NetworkBalancesFilterer) FilterInitialized(opts *bind.FilterOpts) (*NetworkBalancesInitializedIterator, error) {

	logs, sub, err := _NetworkBalances.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NetworkBalancesInitializedIterator{contract: _NetworkBalances.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NetworkBalances *NetworkBalancesFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NetworkBalancesInitialized) (event.Subscription, error) {

	logs, sub, err := _NetworkBalances.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkBalancesInitialized)
				if err := _NetworkBalances.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NetworkBalances *NetworkBalancesFilterer) ParseInitialized(log types.Log) (*NetworkBalancesInitialized, error) {
	event := new(NetworkBalancesInitialized)
	if err := _NetworkBalances.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkBalancesUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NetworkBalances contract.
type NetworkBalancesUpgradedIterator struct {
	Event *NetworkBalancesUpgraded // Event containing the contract specifics and raw log

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
func (it *NetworkBalancesUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkBalancesUpgraded)
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
		it.Event = new(NetworkBalancesUpgraded)
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
func (it *NetworkBalancesUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkBalancesUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkBalancesUpgraded represents a Upgraded event raised by the NetworkBalances contract.
type NetworkBalancesUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NetworkBalances *NetworkBalancesFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NetworkBalancesUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NetworkBalances.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NetworkBalancesUpgradedIterator{contract: _NetworkBalances.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NetworkBalances *NetworkBalancesFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NetworkBalancesUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NetworkBalances.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkBalancesUpgraded)
				if err := _NetworkBalances.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_NetworkBalances *NetworkBalancesFilterer) ParseUpgraded(log types.Log) (*NetworkBalancesUpgraded, error) {
	event := new(NetworkBalancesUpgraded)
	if err := _NetworkBalances.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
