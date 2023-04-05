// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package strategy

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

// AbstractStrategyHarvest is an auto generated low-level Go binding around an user-defined struct.
type AbstractStrategyHarvest struct {
	Token  common.Address
	Amount *big.Int
}

// StrategyMetaData contains all meta data concerning the Strategy contract.
var StrategyMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractGnosisSafe\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractGnosisSafe\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractGnosisSafe\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountFrom\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountTo\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractGnosisSafe\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrew\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authRegistry\",\"outputs\":[{\"internalType\":\"contractAuthRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"contractGnosisSafe\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"contractGnosisSafe\",\"name\":\"safe\",\"type\":\"address\"}],\"name\":\"depositedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extRegistry\",\"outputs\":[{\"internalType\":\"contractExtRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"contractGnosisSafe\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"contractGnosisSafe\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"compound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"}],\"name\":\"poolName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"contractGnosisSafe\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"simulateClaim\",\"outputs\":[{\"components\":[{\"internalType\":\"contractERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAbstractStrategy.Harvest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"contractGnosisSafe\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyMetaData.ABI instead.
var StrategyABI = StrategyMetaData.ABI

// Strategy is an auto generated Go binding around an Ethereum contract.
type Strategy struct {
	StrategyCaller     // Read-only binding to the contract
	StrategyTransactor // Write-only binding to the contract
	StrategyFilterer   // Log filterer for contract events
}

// StrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategySession struct {
	Contract     *Strategy         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyCallerSession struct {
	Contract *StrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyTransactorSession struct {
	Contract     *StrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyRaw struct {
	Contract *Strategy // Generic contract binding to access the raw methods on
}

// StrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyCallerRaw struct {
	Contract *StrategyCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyTransactorRaw struct {
	Contract *StrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategy creates a new instance of Strategy, bound to a specific deployed contract.
func NewStrategy(address common.Address, backend bind.ContractBackend) (*Strategy, error) {
	contract, err := bindStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strategy{StrategyCaller: StrategyCaller{contract: contract}, StrategyTransactor: StrategyTransactor{contract: contract}, StrategyFilterer: StrategyFilterer{contract: contract}}, nil
}

// NewStrategyCaller creates a new read-only instance of Strategy, bound to a specific deployed contract.
func NewStrategyCaller(address common.Address, caller bind.ContractCaller) (*StrategyCaller, error) {
	contract, err := bindStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyCaller{contract: contract}, nil
}

// NewStrategyTransactor creates a new write-only instance of Strategy, bound to a specific deployed contract.
func NewStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyTransactor, error) {
	contract, err := bindStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyTransactor{contract: contract}, nil
}

// NewStrategyFilterer creates a new log filterer instance of Strategy, bound to a specific deployed contract.
func NewStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyFilterer, error) {
	contract, err := bindStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyFilterer{contract: contract}, nil
}

// bindStrategy binds a generic wrapper to an already deployed contract.
func bindStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StrategyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strategy *StrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strategy.Contract.StrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strategy *StrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.Contract.StrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strategy *StrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strategy.Contract.StrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strategy *StrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strategy *StrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strategy *StrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strategy.Contract.contract.Transact(opts, method, params...)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() pure returns(string)
func (_Strategy *StrategyCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() pure returns(string)
func (_Strategy *StrategySession) NAME() (string, error) {
	return _Strategy.Contract.NAME(&_Strategy.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() pure returns(string)
func (_Strategy *StrategyCallerSession) NAME() (string, error) {
	return _Strategy.Contract.NAME(&_Strategy.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() pure returns(string)
func (_Strategy *StrategyCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() pure returns(string)
func (_Strategy *StrategySession) VERSION() (string, error) {
	return _Strategy.Contract.VERSION(&_Strategy.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() pure returns(string)
func (_Strategy *StrategyCallerSession) VERSION() (string, error) {
	return _Strategy.Contract.VERSION(&_Strategy.CallOpts)
}

// AuthRegistry is a free data retrieval call binding the contract method 0x629e025a.
//
// Solidity: function authRegistry() view returns(address)
func (_Strategy *StrategyCaller) AuthRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "authRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthRegistry is a free data retrieval call binding the contract method 0x629e025a.
//
// Solidity: function authRegistry() view returns(address)
func (_Strategy *StrategySession) AuthRegistry() (common.Address, error) {
	return _Strategy.Contract.AuthRegistry(&_Strategy.CallOpts)
}

// AuthRegistry is a free data retrieval call binding the contract method 0x629e025a.
//
// Solidity: function authRegistry() view returns(address)
func (_Strategy *StrategyCallerSession) AuthRegistry() (common.Address, error) {
	return _Strategy.Contract.AuthRegistry(&_Strategy.CallOpts)
}

// DepositToken is a free data retrieval call binding the contract method 0x6215be77.
//
// Solidity: function depositToken(uint256 pool) view returns(address)
func (_Strategy *StrategyCaller) DepositToken(opts *bind.CallOpts, pool *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "depositToken", pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositToken is a free data retrieval call binding the contract method 0x6215be77.
//
// Solidity: function depositToken(uint256 pool) view returns(address)
func (_Strategy *StrategySession) DepositToken(pool *big.Int) (common.Address, error) {
	return _Strategy.Contract.DepositToken(&_Strategy.CallOpts, pool)
}

// DepositToken is a free data retrieval call binding the contract method 0x6215be77.
//
// Solidity: function depositToken(uint256 pool) view returns(address)
func (_Strategy *StrategyCallerSession) DepositToken(pool *big.Int) (common.Address, error) {
	return _Strategy.Contract.DepositToken(&_Strategy.CallOpts, pool)
}

// DepositedAmount is a free data retrieval call binding the contract method 0xf4c8770e.
//
// Solidity: function depositedAmount(uint256 pool, address safe) view returns(uint256)
func (_Strategy *StrategyCaller) DepositedAmount(opts *bind.CallOpts, pool *big.Int, safe common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "depositedAmount", pool, safe)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositedAmount is a free data retrieval call binding the contract method 0xf4c8770e.
//
// Solidity: function depositedAmount(uint256 pool, address safe) view returns(uint256)
func (_Strategy *StrategySession) DepositedAmount(pool *big.Int, safe common.Address) (*big.Int, error) {
	return _Strategy.Contract.DepositedAmount(&_Strategy.CallOpts, pool, safe)
}

// DepositedAmount is a free data retrieval call binding the contract method 0xf4c8770e.
//
// Solidity: function depositedAmount(uint256 pool, address safe) view returns(uint256)
func (_Strategy *StrategyCallerSession) DepositedAmount(pool *big.Int, safe common.Address) (*big.Int, error) {
	return _Strategy.Contract.DepositedAmount(&_Strategy.CallOpts, pool, safe)
}

// ExtRegistry is a free data retrieval call binding the contract method 0x8a854771.
//
// Solidity: function extRegistry() view returns(address)
func (_Strategy *StrategyCaller) ExtRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "extRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtRegistry is a free data retrieval call binding the contract method 0x8a854771.
//
// Solidity: function extRegistry() view returns(address)
func (_Strategy *StrategySession) ExtRegistry() (common.Address, error) {
	return _Strategy.Contract.ExtRegistry(&_Strategy.CallOpts)
}

// ExtRegistry is a free data retrieval call binding the contract method 0x8a854771.
//
// Solidity: function extRegistry() view returns(address)
func (_Strategy *StrategyCallerSession) ExtRegistry() (common.Address, error) {
	return _Strategy.Contract.ExtRegistry(&_Strategy.CallOpts)
}

// PoolName is a free data retrieval call binding the contract method 0xcccf3a02.
//
// Solidity: function poolName(uint256 pool) view returns(string)
func (_Strategy *StrategyCaller) PoolName(opts *bind.CallOpts, pool *big.Int) (string, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "poolName", pool)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PoolName is a free data retrieval call binding the contract method 0xcccf3a02.
//
// Solidity: function poolName(uint256 pool) view returns(string)
func (_Strategy *StrategySession) PoolName(pool *big.Int) (string, error) {
	return _Strategy.Contract.PoolName(&_Strategy.CallOpts, pool)
}

// PoolName is a free data retrieval call binding the contract method 0xcccf3a02.
//
// Solidity: function poolName(uint256 pool) view returns(string)
func (_Strategy *StrategyCallerSession) PoolName(pool *big.Int) (string, error) {
	return _Strategy.Contract.PoolName(&_Strategy.CallOpts, pool)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Strategy *StrategyCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Strategy *StrategySession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Strategy.Contract.SupportsInterface(&_Strategy.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Strategy *StrategyCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Strategy.Contract.SupportsInterface(&_Strategy.CallOpts, interfaceID)
}

// Compound is a paid mutator transaction binding the contract method 0x0e5d528a.
//
// Solidity: function compound(uint256 pool, address safe, bytes data) returns()
func (_Strategy *StrategyTransactor) Compound(opts *bind.TransactOpts, pool *big.Int, safe common.Address, data []byte) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "compound", pool, safe, data)
}

// Compound is a paid mutator transaction binding the contract method 0x0e5d528a.
//
// Solidity: function compound(uint256 pool, address safe, bytes data) returns()
func (_Strategy *StrategySession) Compound(pool *big.Int, safe common.Address, data []byte) (*types.Transaction, error) {
	return _Strategy.Contract.Compound(&_Strategy.TransactOpts, pool, safe, data)
}

// Compound is a paid mutator transaction binding the contract method 0x0e5d528a.
//
// Solidity: function compound(uint256 pool, address safe, bytes data) returns()
func (_Strategy *StrategyTransactorSession) Compound(pool *big.Int, safe common.Address, data []byte) (*types.Transaction, error) {
	return _Strategy.Contract.Compound(&_Strategy.TransactOpts, pool, safe, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x1423feba.
//
// Solidity: function deposit(uint256 pool, address safe, uint256 amount, bytes data) returns()
func (_Strategy *StrategyTransactor) Deposit(opts *bind.TransactOpts, pool *big.Int, safe common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "deposit", pool, safe, amount, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x1423feba.
//
// Solidity: function deposit(uint256 pool, address safe, uint256 amount, bytes data) returns()
func (_Strategy *StrategySession) Deposit(pool *big.Int, safe common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Strategy.Contract.Deposit(&_Strategy.TransactOpts, pool, safe, amount, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x1423feba.
//
// Solidity: function deposit(uint256 pool, address safe, uint256 amount, bytes data) returns()
func (_Strategy *StrategyTransactorSession) Deposit(pool *big.Int, safe common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Strategy.Contract.Deposit(&_Strategy.TransactOpts, pool, safe, amount, data)
}

// Harvest is a paid mutator transaction binding the contract method 0xf023a693.
//
// Solidity: function harvest(uint256 pool, address safe, bytes data) returns()
func (_Strategy *StrategyTransactor) Harvest(opts *bind.TransactOpts, pool *big.Int, safe common.Address, data []byte) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "harvest", pool, safe, data)
}

// Harvest is a paid mutator transaction binding the contract method 0xf023a693.
//
// Solidity: function harvest(uint256 pool, address safe, bytes data) returns()
func (_Strategy *StrategySession) Harvest(pool *big.Int, safe common.Address, data []byte) (*types.Transaction, error) {
	return _Strategy.Contract.Harvest(&_Strategy.TransactOpts, pool, safe, data)
}

// Harvest is a paid mutator transaction binding the contract method 0xf023a693.
//
// Solidity: function harvest(uint256 pool, address safe, bytes data) returns()
func (_Strategy *StrategyTransactorSession) Harvest(pool *big.Int, safe common.Address, data []byte) (*types.Transaction, error) {
	return _Strategy.Contract.Harvest(&_Strategy.TransactOpts, pool, safe, data)
}

// SimulateClaim is a paid mutator transaction binding the contract method 0xbdcceb0f.
//
// Solidity: function simulateClaim(uint256 pool, address safe, bytes data) returns((address,uint256)[])
func (_Strategy *StrategyTransactor) SimulateClaim(opts *bind.TransactOpts, pool *big.Int, safe common.Address, data []byte) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "simulateClaim", pool, safe, data)
}

// SimulateClaim is a paid mutator transaction binding the contract method 0xbdcceb0f.
//
// Solidity: function simulateClaim(uint256 pool, address safe, bytes data) returns((address,uint256)[])
func (_Strategy *StrategySession) SimulateClaim(pool *big.Int, safe common.Address, data []byte) (*types.Transaction, error) {
	return _Strategy.Contract.SimulateClaim(&_Strategy.TransactOpts, pool, safe, data)
}

// SimulateClaim is a paid mutator transaction binding the contract method 0xbdcceb0f.
//
// Solidity: function simulateClaim(uint256 pool, address safe, bytes data) returns((address,uint256)[])
func (_Strategy *StrategyTransactorSession) SimulateClaim(pool *big.Int, safe common.Address, data []byte) (*types.Transaction, error) {
	return _Strategy.Contract.SimulateClaim(&_Strategy.TransactOpts, pool, safe, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xdcbd7a53.
//
// Solidity: function withdraw(uint256 pool, address safe, uint256 amount, bytes data) returns()
func (_Strategy *StrategyTransactor) Withdraw(opts *bind.TransactOpts, pool *big.Int, safe common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "withdraw", pool, safe, amount, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xdcbd7a53.
//
// Solidity: function withdraw(uint256 pool, address safe, uint256 amount, bytes data) returns()
func (_Strategy *StrategySession) Withdraw(pool *big.Int, safe common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Strategy.Contract.Withdraw(&_Strategy.TransactOpts, pool, safe, amount, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xdcbd7a53.
//
// Solidity: function withdraw(uint256 pool, address safe, uint256 amount, bytes data) returns()
func (_Strategy *StrategyTransactorSession) Withdraw(pool *big.Int, safe common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Strategy.Contract.Withdraw(&_Strategy.TransactOpts, pool, safe, amount, data)
}

// StrategyDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Strategy contract.
type StrategyDepositedIterator struct {
	Event *StrategyDeposited // Event containing the contract specifics and raw log

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
func (it *StrategyDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyDeposited)
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
		it.Event = new(StrategyDeposited)
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
func (it *StrategyDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyDeposited represents a Deposited event raised by the Strategy contract.
type StrategyDeposited struct {
	Pool   *big.Int
	Safe   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1599c0fcf897af5babc2bfcf707f5dc050f841b044d97c3251ecec35b9abf80b.
//
// Solidity: event Deposited(uint256 pool, address indexed safe, uint256 amount)
func (_Strategy *StrategyFilterer) FilterDeposited(opts *bind.FilterOpts, safe []common.Address) (*StrategyDepositedIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "Deposited", safeRule)
	if err != nil {
		return nil, err
	}
	return &StrategyDepositedIterator{contract: _Strategy.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1599c0fcf897af5babc2bfcf707f5dc050f841b044d97c3251ecec35b9abf80b.
//
// Solidity: event Deposited(uint256 pool, address indexed safe, uint256 amount)
func (_Strategy *StrategyFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *StrategyDeposited, safe []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "Deposited", safeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyDeposited)
				if err := _Strategy.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x1599c0fcf897af5babc2bfcf707f5dc050f841b044d97c3251ecec35b9abf80b.
//
// Solidity: event Deposited(uint256 pool, address indexed safe, uint256 amount)
func (_Strategy *StrategyFilterer) ParseDeposited(log types.Log) (*StrategyDeposited, error) {
	event := new(StrategyDeposited)
	if err := _Strategy.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyHarvestedIterator is returned from FilterHarvested and is used to iterate over the raw logs and unpacked data for Harvested events raised by the Strategy contract.
type StrategyHarvestedIterator struct {
	Event *StrategyHarvested // Event containing the contract specifics and raw log

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
func (it *StrategyHarvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyHarvested)
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
		it.Event = new(StrategyHarvested)
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
func (it *StrategyHarvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyHarvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyHarvested represents a Harvested event raised by the Strategy contract.
type StrategyHarvested struct {
	Pool   *big.Int
	Safe   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHarvested is a free log retrieval operation binding the contract event 0x0617cd812caff7604af7246f2c877b5e717fc45ab6b07769061a0ee05fe5b256.
//
// Solidity: event Harvested(uint256 pool, address indexed safe, address indexed token, uint256 amount)
func (_Strategy *StrategyFilterer) FilterHarvested(opts *bind.FilterOpts, safe []common.Address, token []common.Address) (*StrategyHarvestedIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "Harvested", safeRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &StrategyHarvestedIterator{contract: _Strategy.contract, event: "Harvested", logs: logs, sub: sub}, nil
}

// WatchHarvested is a free log subscription operation binding the contract event 0x0617cd812caff7604af7246f2c877b5e717fc45ab6b07769061a0ee05fe5b256.
//
// Solidity: event Harvested(uint256 pool, address indexed safe, address indexed token, uint256 amount)
func (_Strategy *StrategyFilterer) WatchHarvested(opts *bind.WatchOpts, sink chan<- *StrategyHarvested, safe []common.Address, token []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "Harvested", safeRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyHarvested)
				if err := _Strategy.contract.UnpackLog(event, "Harvested", log); err != nil {
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

// ParseHarvested is a log parse operation binding the contract event 0x0617cd812caff7604af7246f2c877b5e717fc45ab6b07769061a0ee05fe5b256.
//
// Solidity: event Harvested(uint256 pool, address indexed safe, address indexed token, uint256 amount)
func (_Strategy *StrategyFilterer) ParseHarvested(log types.Log) (*StrategyHarvested, error) {
	event := new(StrategyHarvested)
	if err := _Strategy.contract.UnpackLog(event, "Harvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategySwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the Strategy contract.
type StrategySwappedIterator struct {
	Event *StrategySwapped // Event containing the contract specifics and raw log

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
func (it *StrategySwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategySwapped)
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
		it.Event = new(StrategySwapped)
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
func (it *StrategySwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategySwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategySwapped represents a Swapped event raised by the Strategy contract.
type StrategySwapped struct {
	Safe       common.Address
	From       common.Address
	To         common.Address
	AmountFrom *big.Int
	AmountTo   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0x6782190c91d4a7e8ad2a867deed6ec0a970cab8ff137ae2bd4abd92b3810f4d3.
//
// Solidity: event Swapped(address indexed safe, address from, address to, uint256 amountFrom, uint256 amountTo)
func (_Strategy *StrategyFilterer) FilterSwapped(opts *bind.FilterOpts, safe []common.Address) (*StrategySwappedIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "Swapped", safeRule)
	if err != nil {
		return nil, err
	}
	return &StrategySwappedIterator{contract: _Strategy.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0x6782190c91d4a7e8ad2a867deed6ec0a970cab8ff137ae2bd4abd92b3810f4d3.
//
// Solidity: event Swapped(address indexed safe, address from, address to, uint256 amountFrom, uint256 amountTo)
func (_Strategy *StrategyFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *StrategySwapped, safe []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "Swapped", safeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategySwapped)
				if err := _Strategy.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0x6782190c91d4a7e8ad2a867deed6ec0a970cab8ff137ae2bd4abd92b3810f4d3.
//
// Solidity: event Swapped(address indexed safe, address from, address to, uint256 amountFrom, uint256 amountTo)
func (_Strategy *StrategyFilterer) ParseSwapped(log types.Log) (*StrategySwapped, error) {
	event := new(StrategySwapped)
	if err := _Strategy.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyWithdrewIterator is returned from FilterWithdrew and is used to iterate over the raw logs and unpacked data for Withdrew events raised by the Strategy contract.
type StrategyWithdrewIterator struct {
	Event *StrategyWithdrew // Event containing the contract specifics and raw log

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
func (it *StrategyWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyWithdrew)
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
		it.Event = new(StrategyWithdrew)
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
func (it *StrategyWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyWithdrew represents a Withdrew event raised by the Strategy contract.
type StrategyWithdrew struct {
	Pool   *big.Int
	Safe   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrew is a free log retrieval operation binding the contract event 0xf4c56bb5a568ebcb1087fa42400137fd62e67119180ddfe97906e108ecc0e633.
//
// Solidity: event Withdrew(uint256 pool, address indexed safe, uint256 amount)
func (_Strategy *StrategyFilterer) FilterWithdrew(opts *bind.FilterOpts, safe []common.Address) (*StrategyWithdrewIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "Withdrew", safeRule)
	if err != nil {
		return nil, err
	}
	return &StrategyWithdrewIterator{contract: _Strategy.contract, event: "Withdrew", logs: logs, sub: sub}, nil
}

// WatchWithdrew is a free log subscription operation binding the contract event 0xf4c56bb5a568ebcb1087fa42400137fd62e67119180ddfe97906e108ecc0e633.
//
// Solidity: event Withdrew(uint256 pool, address indexed safe, uint256 amount)
func (_Strategy *StrategyFilterer) WatchWithdrew(opts *bind.WatchOpts, sink chan<- *StrategyWithdrew, safe []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "Withdrew", safeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyWithdrew)
				if err := _Strategy.contract.UnpackLog(event, "Withdrew", log); err != nil {
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

// ParseWithdrew is a log parse operation binding the contract event 0xf4c56bb5a568ebcb1087fa42400137fd62e67119180ddfe97906e108ecc0e633.
//
// Solidity: event Withdrew(uint256 pool, address indexed safe, uint256 amount)
func (_Strategy *StrategyFilterer) ParseWithdrew(log types.Log) (*StrategyWithdrew, error) {
	event := new(StrategyWithdrew)
	if err := _Strategy.contract.UnpackLog(event, "Withdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
