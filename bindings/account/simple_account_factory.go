// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package account

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

// SimpleAccountFactoryMetaData contains all meta data concerning the SimpleAccountFactory contract.
var SimpleAccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"accountImplementation\",\"outputs\":[{\"internalType\":\"contractSimpleAccount\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"contractSimpleAccount\",\"name\":\"ret\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SimpleAccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleAccountFactoryMetaData.ABI instead.
var SimpleAccountFactoryABI = SimpleAccountFactoryMetaData.ABI

// SimpleAccountFactory is an auto generated Go binding around an Ethereum contract.
type SimpleAccountFactory struct {
	SimpleAccountFactoryCaller     // Read-only binding to the contract
	SimpleAccountFactoryTransactor // Write-only binding to the contract
	SimpleAccountFactoryFilterer   // Log filterer for contract events
}

// SimpleAccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleAccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleAccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleAccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleAccountFactorySession struct {
	Contract     *SimpleAccountFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SimpleAccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleAccountFactoryCallerSession struct {
	Contract *SimpleAccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SimpleAccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleAccountFactoryTransactorSession struct {
	Contract     *SimpleAccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SimpleAccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleAccountFactoryRaw struct {
	Contract *SimpleAccountFactory // Generic contract binding to access the raw methods on
}

// SimpleAccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleAccountFactoryCallerRaw struct {
	Contract *SimpleAccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleAccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleAccountFactoryTransactorRaw struct {
	Contract *SimpleAccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleAccountFactory creates a new instance of SimpleAccountFactory, bound to a specific deployed contract.
func NewSimpleAccountFactory(address common.Address, backend bind.ContractBackend) (*SimpleAccountFactory, error) {
	contract, err := bindSimpleAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleAccountFactory{SimpleAccountFactoryCaller: SimpleAccountFactoryCaller{contract: contract}, SimpleAccountFactoryTransactor: SimpleAccountFactoryTransactor{contract: contract}, SimpleAccountFactoryFilterer: SimpleAccountFactoryFilterer{contract: contract}}, nil
}

// NewSimpleAccountFactoryCaller creates a new read-only instance of SimpleAccountFactory, bound to a specific deployed contract.
func NewSimpleAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*SimpleAccountFactoryCaller, error) {
	contract, err := bindSimpleAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAccountFactoryCaller{contract: contract}, nil
}

// NewSimpleAccountFactoryTransactor creates a new write-only instance of SimpleAccountFactory, bound to a specific deployed contract.
func NewSimpleAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleAccountFactoryTransactor, error) {
	contract, err := bindSimpleAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAccountFactoryTransactor{contract: contract}, nil
}

// NewSimpleAccountFactoryFilterer creates a new log filterer instance of SimpleAccountFactory, bound to a specific deployed contract.
func NewSimpleAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleAccountFactoryFilterer, error) {
	contract, err := bindSimpleAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleAccountFactoryFilterer{contract: contract}, nil
}

// bindSimpleAccountFactory binds a generic wrapper to an already deployed contract.
func bindSimpleAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleAccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAccountFactory *SimpleAccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleAccountFactory.Contract.SimpleAccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAccountFactory *SimpleAccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAccountFactory.Contract.SimpleAccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAccountFactory *SimpleAccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleAccountFactory.Contract.SimpleAccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAccountFactory *SimpleAccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleAccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAccountFactory *SimpleAccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAccountFactory *SimpleAccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleAccountFactory.Contract.contract.Transact(opts, method, params...)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_SimpleAccountFactory *SimpleAccountFactoryCaller) AccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleAccountFactory.contract.Call(opts, &out, "accountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_SimpleAccountFactory *SimpleAccountFactorySession) AccountImplementation() (common.Address, error) {
	return _SimpleAccountFactory.Contract.AccountImplementation(&_SimpleAccountFactory.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_SimpleAccountFactory *SimpleAccountFactoryCallerSession) AccountImplementation() (common.Address, error) {
	return _SimpleAccountFactory.Contract.AccountImplementation(&_SimpleAccountFactory.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_SimpleAccountFactory *SimpleAccountFactoryCaller) GetAddress(opts *bind.CallOpts, owner common.Address, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SimpleAccountFactory.contract.Call(opts, &out, "getAddress", owner, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_SimpleAccountFactory *SimpleAccountFactorySession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _SimpleAccountFactory.Contract.GetAddress(&_SimpleAccountFactory.CallOpts, owner, salt)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_SimpleAccountFactory *SimpleAccountFactoryCallerSession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _SimpleAccountFactory.Contract.GetAddress(&_SimpleAccountFactory.CallOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_SimpleAccountFactory *SimpleAccountFactoryTransactor) CreateAccount(opts *bind.TransactOpts, owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _SimpleAccountFactory.contract.Transact(opts, "createAccount", owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_SimpleAccountFactory *SimpleAccountFactorySession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _SimpleAccountFactory.Contract.CreateAccount(&_SimpleAccountFactory.TransactOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_SimpleAccountFactory *SimpleAccountFactoryTransactorSession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _SimpleAccountFactory.Contract.CreateAccount(&_SimpleAccountFactory.TransactOpts, owner, salt)
}
