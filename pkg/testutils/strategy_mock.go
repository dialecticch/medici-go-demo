// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testutils

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

// MockStrategyHarvest is an auto generated low-level Go binding around an user-defined struct.
type MockStrategyHarvest struct {
	Token  common.Address
	Amount *big.Int
}

// MockStrategyMetaData contains all meta data concerning the MockStrategy contract.
var MockStrategyMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrew\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_safe\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"compound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"safe\",\"type\":\"address\"}],\"name\":\"depositedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_safe\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"logDeposited\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"logHarvested\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"logWithdrew\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"simulateClaim\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structMockStrategy.Harvest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a3f4df7e": "NAME()",
		"0e5d528a": "compound(uint256,address,bytes)",
		"c89039c5": "depositToken()",
		"f4c8770e": "depositedAmount(uint256,address)",
		"f023a693": "harvest(uint256,address,bytes)",
		"d31cbad9": "logDeposited(uint256,address,uint256)",
		"edc4c14f": "logHarvested(uint256,address,address,uint256)",
		"ae280e7c": "logWithdrew(uint256,address,uint256)",
		"bdcceb0f": "simulateClaim(uint256,address,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610594806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063c89039c511610066578063c89039c514610119578063d31cbad91461013b578063edc4c14f1461014e578063f023a69314610098578063f4c8770e1461016157600080fd5b80630e5d528a14610098578063a3f4df7e146100ae578063ae280e7c146100e6578063bdcceb0f146100f9575b600080fd5b6100ac6100a6366004610376565b50505050565b005b604080518082018252600d81526c4d6f636b20537472617465677960981b602082015290516100dd91906103fd565b60405180910390f35b6100ac6100f436600461044b565b610185565b61010c610107366004610376565b6101cf565b6040516100dd9190610480565b60405173a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4881526020016100dd565b6100ac61014936600461044b565b6102bd565b6100ac61015c3660046104d8565b6102fe565b61017761016f36600461051c565b600a92915050565b6040519081526020016100dd565b60408051848152602081018390526001600160a01b038416917ff4c56bb5a568ebcb1087fa42400137fd62e67119180ddfe97906e108ecc0e63391015b60405180910390a2505050565b604080516002808252606082810190935260009190816020015b60408051808201909152600080825260208201528152602001906001900390816101e9579050509050604051806040016040528073a0b86991c6218b36c1d19d4a2e9eb0ce3606eb486001600160a01b03168152602001600a8152508160008151811061025857610258610548565b6020026020010181905250604051806040016040528073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26001600160a01b03168152602001600a815250816001815181106102a9576102a9610548565b602090810291909101015295945050505050565b60408051848152602081018390526001600160a01b038416917f1599c0fcf897af5babc2bfcf707f5dc050f841b044d97c3251ecec35b9abf80b91016101c2565b816001600160a01b0316836001600160a01b03167f0617cd812caff7604af7246f2c877b5e717fc45ab6b07769061a0ee05fe5b256868460405161034c929190918252602082015260400190565b60405180910390a350505050565b80356001600160a01b038116811461037157600080fd5b919050565b6000806000806060858703121561038c57600080fd5b8435935061039c6020860161035a565b9250604085013567ffffffffffffffff808211156103b957600080fd5b818701915087601f8301126103cd57600080fd5b8135818111156103dc57600080fd5b8860208285010111156103ee57600080fd5b95989497505060200194505050565b600060208083528351808285015260005b8181101561042a5785810183015185820160400152820161040e565b506000604082860101526040601f19601f8301168501019250505092915050565b60008060006060848603121561046057600080fd5b833592506104706020850161035a565b9150604084013590509250925092565b602080825282518282018190526000919060409081850190868401855b828110156104cb57815180516001600160a01b0316855286015186850152928401929085019060010161049d565b5091979650505050505050565b600080600080608085870312156104ee57600080fd5b843593506104fe6020860161035a565b925061050c6040860161035a565b9396929550929360600135925050565b6000806040838503121561052f57600080fd5b8235915061053f6020840161035a565b90509250929050565b634e487b7160e01b600052603260045260246000fdfea264697066735822122097be2c201de6b885af4893d8d366f5c4b9ee9ae6888e781d0e4f3a286573e76164736f6c63430008100033",
}

// MockStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use MockStrategyMetaData.ABI instead.
var MockStrategyABI = MockStrategyMetaData.ABI

// Deprecated: Use MockStrategyMetaData.Sigs instead.
// MockStrategyFuncSigs maps the 4-byte function signature to its string representation.
var MockStrategyFuncSigs = MockStrategyMetaData.Sigs

// MockStrategyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockStrategyMetaData.Bin instead.
var MockStrategyBin = MockStrategyMetaData.Bin

// DeployMockStrategy deploys a new Ethereum contract, binding an instance of MockStrategy to it.
func DeployMockStrategy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockStrategy, error) {
	parsed, err := MockStrategyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockStrategyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockStrategy{MockStrategyCaller: MockStrategyCaller{contract: contract}, MockStrategyTransactor: MockStrategyTransactor{contract: contract}, MockStrategyFilterer: MockStrategyFilterer{contract: contract}}, nil
}

// MockStrategy is an auto generated Go binding around an Ethereum contract.
type MockStrategy struct {
	MockStrategyCaller     // Read-only binding to the contract
	MockStrategyTransactor // Write-only binding to the contract
	MockStrategyFilterer   // Log filterer for contract events
}

// MockStrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockStrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockStrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockStrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockStrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockStrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockStrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockStrategySession struct {
	Contract     *MockStrategy     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockStrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockStrategyCallerSession struct {
	Contract *MockStrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MockStrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockStrategyTransactorSession struct {
	Contract     *MockStrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MockStrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockStrategyRaw struct {
	Contract *MockStrategy // Generic contract binding to access the raw methods on
}

// MockStrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockStrategyCallerRaw struct {
	Contract *MockStrategyCaller // Generic read-only contract binding to access the raw methods on
}

// MockStrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockStrategyTransactorRaw struct {
	Contract *MockStrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockStrategy creates a new instance of MockStrategy, bound to a specific deployed contract.
func NewMockStrategy(address common.Address, backend bind.ContractBackend) (*MockStrategy, error) {
	contract, err := bindMockStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockStrategy{MockStrategyCaller: MockStrategyCaller{contract: contract}, MockStrategyTransactor: MockStrategyTransactor{contract: contract}, MockStrategyFilterer: MockStrategyFilterer{contract: contract}}, nil
}

// NewMockStrategyCaller creates a new read-only instance of MockStrategy, bound to a specific deployed contract.
func NewMockStrategyCaller(address common.Address, caller bind.ContractCaller) (*MockStrategyCaller, error) {
	contract, err := bindMockStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockStrategyCaller{contract: contract}, nil
}

// NewMockStrategyTransactor creates a new write-only instance of MockStrategy, bound to a specific deployed contract.
func NewMockStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*MockStrategyTransactor, error) {
	contract, err := bindMockStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockStrategyTransactor{contract: contract}, nil
}

// NewMockStrategyFilterer creates a new log filterer instance of MockStrategy, bound to a specific deployed contract.
func NewMockStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*MockStrategyFilterer, error) {
	contract, err := bindMockStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockStrategyFilterer{contract: contract}, nil
}

// bindMockStrategy binds a generic wrapper to an already deployed contract.
func bindMockStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockStrategyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockStrategy *MockStrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockStrategy.Contract.MockStrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockStrategy *MockStrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockStrategy.Contract.MockStrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockStrategy *MockStrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockStrategy.Contract.MockStrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockStrategy *MockStrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockStrategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockStrategy *MockStrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockStrategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockStrategy *MockStrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockStrategy.Contract.contract.Transact(opts, method, params...)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_MockStrategy *MockStrategyCaller) DepositToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockStrategy.contract.Call(opts, &out, "depositToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_MockStrategy *MockStrategySession) DepositToken() (common.Address, error) {
	return _MockStrategy.Contract.DepositToken(&_MockStrategy.CallOpts)
}

// DepositToken is a free data retrieval call binding the contract method 0xc89039c5.
//
// Solidity: function depositToken() view returns(address)
func (_MockStrategy *MockStrategyCallerSession) DepositToken() (common.Address, error) {
	return _MockStrategy.Contract.DepositToken(&_MockStrategy.CallOpts)
}

// DepositedAmount is a free data retrieval call binding the contract method 0xf4c8770e.
//
// Solidity: function depositedAmount(uint256 pool, address safe) view returns(uint256)
func (_MockStrategy *MockStrategyCaller) DepositedAmount(opts *bind.CallOpts, pool *big.Int, safe common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockStrategy.contract.Call(opts, &out, "depositedAmount", pool, safe)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositedAmount is a free data retrieval call binding the contract method 0xf4c8770e.
//
// Solidity: function depositedAmount(uint256 pool, address safe) view returns(uint256)
func (_MockStrategy *MockStrategySession) DepositedAmount(pool *big.Int, safe common.Address) (*big.Int, error) {
	return _MockStrategy.Contract.DepositedAmount(&_MockStrategy.CallOpts, pool, safe)
}

// DepositedAmount is a free data retrieval call binding the contract method 0xf4c8770e.
//
// Solidity: function depositedAmount(uint256 pool, address safe) view returns(uint256)
func (_MockStrategy *MockStrategyCallerSession) DepositedAmount(pool *big.Int, safe common.Address) (*big.Int, error) {
	return _MockStrategy.Contract.DepositedAmount(&_MockStrategy.CallOpts, pool, safe)
}

// NAME is a paid mutator transaction binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() returns(string)
func (_MockStrategy *MockStrategyTransactor) NAME(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockStrategy.contract.Transact(opts, "NAME")
}

// NAME is a paid mutator transaction binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() returns(string)
func (_MockStrategy *MockStrategySession) NAME() (*types.Transaction, error) {
	return _MockStrategy.Contract.NAME(&_MockStrategy.TransactOpts)
}

// NAME is a paid mutator transaction binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() returns(string)
func (_MockStrategy *MockStrategyTransactorSession) NAME() (*types.Transaction, error) {
	return _MockStrategy.Contract.NAME(&_MockStrategy.TransactOpts)
}

// Compound is a paid mutator transaction binding the contract method 0x0e5d528a.
//
// Solidity: function compound(uint256 pool, address _safe, bytes data) returns()
func (_MockStrategy *MockStrategyTransactor) Compound(opts *bind.TransactOpts, pool *big.Int, _safe common.Address, data []byte) (*types.Transaction, error) {
	return _MockStrategy.contract.Transact(opts, "compound", pool, _safe, data)
}

// Compound is a paid mutator transaction binding the contract method 0x0e5d528a.
//
// Solidity: function compound(uint256 pool, address _safe, bytes data) returns()
func (_MockStrategy *MockStrategySession) Compound(pool *big.Int, _safe common.Address, data []byte) (*types.Transaction, error) {
	return _MockStrategy.Contract.Compound(&_MockStrategy.TransactOpts, pool, _safe, data)
}

// Compound is a paid mutator transaction binding the contract method 0x0e5d528a.
//
// Solidity: function compound(uint256 pool, address _safe, bytes data) returns()
func (_MockStrategy *MockStrategyTransactorSession) Compound(pool *big.Int, _safe common.Address, data []byte) (*types.Transaction, error) {
	return _MockStrategy.Contract.Compound(&_MockStrategy.TransactOpts, pool, _safe, data)
}

// Harvest is a paid mutator transaction binding the contract method 0xf023a693.
//
// Solidity: function harvest(uint256 pool, address _safe, bytes data) returns()
func (_MockStrategy *MockStrategyTransactor) Harvest(opts *bind.TransactOpts, pool *big.Int, _safe common.Address, data []byte) (*types.Transaction, error) {
	return _MockStrategy.contract.Transact(opts, "harvest", pool, _safe, data)
}

// Harvest is a paid mutator transaction binding the contract method 0xf023a693.
//
// Solidity: function harvest(uint256 pool, address _safe, bytes data) returns()
func (_MockStrategy *MockStrategySession) Harvest(pool *big.Int, _safe common.Address, data []byte) (*types.Transaction, error) {
	return _MockStrategy.Contract.Harvest(&_MockStrategy.TransactOpts, pool, _safe, data)
}

// Harvest is a paid mutator transaction binding the contract method 0xf023a693.
//
// Solidity: function harvest(uint256 pool, address _safe, bytes data) returns()
func (_MockStrategy *MockStrategyTransactorSession) Harvest(pool *big.Int, _safe common.Address, data []byte) (*types.Transaction, error) {
	return _MockStrategy.Contract.Harvest(&_MockStrategy.TransactOpts, pool, _safe, data)
}

// LogDeposited is a paid mutator transaction binding the contract method 0xd31cbad9.
//
// Solidity: function logDeposited(uint256 pool, address safe, uint256 amount) returns()
func (_MockStrategy *MockStrategyTransactor) LogDeposited(opts *bind.TransactOpts, pool *big.Int, safe common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockStrategy.contract.Transact(opts, "logDeposited", pool, safe, amount)
}

// LogDeposited is a paid mutator transaction binding the contract method 0xd31cbad9.
//
// Solidity: function logDeposited(uint256 pool, address safe, uint256 amount) returns()
func (_MockStrategy *MockStrategySession) LogDeposited(pool *big.Int, safe common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockStrategy.Contract.LogDeposited(&_MockStrategy.TransactOpts, pool, safe, amount)
}

// LogDeposited is a paid mutator transaction binding the contract method 0xd31cbad9.
//
// Solidity: function logDeposited(uint256 pool, address safe, uint256 amount) returns()
func (_MockStrategy *MockStrategyTransactorSession) LogDeposited(pool *big.Int, safe common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockStrategy.Contract.LogDeposited(&_MockStrategy.TransactOpts, pool, safe, amount)
}

// LogHarvested is a paid mutator transaction binding the contract method 0xedc4c14f.
//
// Solidity: function logHarvested(uint256 pool, address safe, address token, uint256 amount) returns()
func (_MockStrategy *MockStrategyTransactor) LogHarvested(opts *bind.TransactOpts, pool *big.Int, safe common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockStrategy.contract.Transact(opts, "logHarvested", pool, safe, token, amount)
}

// LogHarvested is a paid mutator transaction binding the contract method 0xedc4c14f.
//
// Solidity: function logHarvested(uint256 pool, address safe, address token, uint256 amount) returns()
func (_MockStrategy *MockStrategySession) LogHarvested(pool *big.Int, safe common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockStrategy.Contract.LogHarvested(&_MockStrategy.TransactOpts, pool, safe, token, amount)
}

// LogHarvested is a paid mutator transaction binding the contract method 0xedc4c14f.
//
// Solidity: function logHarvested(uint256 pool, address safe, address token, uint256 amount) returns()
func (_MockStrategy *MockStrategyTransactorSession) LogHarvested(pool *big.Int, safe common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockStrategy.Contract.LogHarvested(&_MockStrategy.TransactOpts, pool, safe, token, amount)
}

// LogWithdrew is a paid mutator transaction binding the contract method 0xae280e7c.
//
// Solidity: function logWithdrew(uint256 pool, address safe, uint256 amount) returns()
func (_MockStrategy *MockStrategyTransactor) LogWithdrew(opts *bind.TransactOpts, pool *big.Int, safe common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockStrategy.contract.Transact(opts, "logWithdrew", pool, safe, amount)
}

// LogWithdrew is a paid mutator transaction binding the contract method 0xae280e7c.
//
// Solidity: function logWithdrew(uint256 pool, address safe, uint256 amount) returns()
func (_MockStrategy *MockStrategySession) LogWithdrew(pool *big.Int, safe common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockStrategy.Contract.LogWithdrew(&_MockStrategy.TransactOpts, pool, safe, amount)
}

// LogWithdrew is a paid mutator transaction binding the contract method 0xae280e7c.
//
// Solidity: function logWithdrew(uint256 pool, address safe, uint256 amount) returns()
func (_MockStrategy *MockStrategyTransactorSession) LogWithdrew(pool *big.Int, safe common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockStrategy.Contract.LogWithdrew(&_MockStrategy.TransactOpts, pool, safe, amount)
}

// SimulateClaim is a paid mutator transaction binding the contract method 0xbdcceb0f.
//
// Solidity: function simulateClaim(uint256 pool, address safe, bytes data) returns((address,uint256)[])
func (_MockStrategy *MockStrategyTransactor) SimulateClaim(opts *bind.TransactOpts, pool *big.Int, safe common.Address, data []byte) (*types.Transaction, error) {
	return _MockStrategy.contract.Transact(opts, "simulateClaim", pool, safe, data)
}

// SimulateClaim is a paid mutator transaction binding the contract method 0xbdcceb0f.
//
// Solidity: function simulateClaim(uint256 pool, address safe, bytes data) returns((address,uint256)[])
func (_MockStrategy *MockStrategySession) SimulateClaim(pool *big.Int, safe common.Address, data []byte) (*types.Transaction, error) {
	return _MockStrategy.Contract.SimulateClaim(&_MockStrategy.TransactOpts, pool, safe, data)
}

// SimulateClaim is a paid mutator transaction binding the contract method 0xbdcceb0f.
//
// Solidity: function simulateClaim(uint256 pool, address safe, bytes data) returns((address,uint256)[])
func (_MockStrategy *MockStrategyTransactorSession) SimulateClaim(pool *big.Int, safe common.Address, data []byte) (*types.Transaction, error) {
	return _MockStrategy.Contract.SimulateClaim(&_MockStrategy.TransactOpts, pool, safe, data)
}

// MockStrategyDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the MockStrategy contract.
type MockStrategyDepositedIterator struct {
	Event *MockStrategyDeposited // Event containing the contract specifics and raw log

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
func (it *MockStrategyDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockStrategyDeposited)
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
		it.Event = new(MockStrategyDeposited)
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
func (it *MockStrategyDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockStrategyDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockStrategyDeposited represents a Deposited event raised by the MockStrategy contract.
type MockStrategyDeposited struct {
	Pool   *big.Int
	Safe   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1599c0fcf897af5babc2bfcf707f5dc050f841b044d97c3251ecec35b9abf80b.
//
// Solidity: event Deposited(uint256 pool, address indexed safe, uint256 amount)
func (_MockStrategy *MockStrategyFilterer) FilterDeposited(opts *bind.FilterOpts, safe []common.Address) (*MockStrategyDepositedIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _MockStrategy.contract.FilterLogs(opts, "Deposited", safeRule)
	if err != nil {
		return nil, err
	}
	return &MockStrategyDepositedIterator{contract: _MockStrategy.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1599c0fcf897af5babc2bfcf707f5dc050f841b044d97c3251ecec35b9abf80b.
//
// Solidity: event Deposited(uint256 pool, address indexed safe, uint256 amount)
func (_MockStrategy *MockStrategyFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *MockStrategyDeposited, safe []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _MockStrategy.contract.WatchLogs(opts, "Deposited", safeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockStrategyDeposited)
				if err := _MockStrategy.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_MockStrategy *MockStrategyFilterer) ParseDeposited(log types.Log) (*MockStrategyDeposited, error) {
	event := new(MockStrategyDeposited)
	if err := _MockStrategy.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockStrategyHarvestedIterator is returned from FilterHarvested and is used to iterate over the raw logs and unpacked data for Harvested events raised by the MockStrategy contract.
type MockStrategyHarvestedIterator struct {
	Event *MockStrategyHarvested // Event containing the contract specifics and raw log

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
func (it *MockStrategyHarvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockStrategyHarvested)
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
		it.Event = new(MockStrategyHarvested)
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
func (it *MockStrategyHarvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockStrategyHarvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockStrategyHarvested represents a Harvested event raised by the MockStrategy contract.
type MockStrategyHarvested struct {
	Pool   *big.Int
	Safe   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHarvested is a free log retrieval operation binding the contract event 0x0617cd812caff7604af7246f2c877b5e717fc45ab6b07769061a0ee05fe5b256.
//
// Solidity: event Harvested(uint256 pool, address indexed safe, address indexed token, uint256 amount)
func (_MockStrategy *MockStrategyFilterer) FilterHarvested(opts *bind.FilterOpts, safe []common.Address, token []common.Address) (*MockStrategyHarvestedIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MockStrategy.contract.FilterLogs(opts, "Harvested", safeRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MockStrategyHarvestedIterator{contract: _MockStrategy.contract, event: "Harvested", logs: logs, sub: sub}, nil
}

// WatchHarvested is a free log subscription operation binding the contract event 0x0617cd812caff7604af7246f2c877b5e717fc45ab6b07769061a0ee05fe5b256.
//
// Solidity: event Harvested(uint256 pool, address indexed safe, address indexed token, uint256 amount)
func (_MockStrategy *MockStrategyFilterer) WatchHarvested(opts *bind.WatchOpts, sink chan<- *MockStrategyHarvested, safe []common.Address, token []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MockStrategy.contract.WatchLogs(opts, "Harvested", safeRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockStrategyHarvested)
				if err := _MockStrategy.contract.UnpackLog(event, "Harvested", log); err != nil {
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
func (_MockStrategy *MockStrategyFilterer) ParseHarvested(log types.Log) (*MockStrategyHarvested, error) {
	event := new(MockStrategyHarvested)
	if err := _MockStrategy.contract.UnpackLog(event, "Harvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockStrategyWithdrewIterator is returned from FilterWithdrew and is used to iterate over the raw logs and unpacked data for Withdrew events raised by the MockStrategy contract.
type MockStrategyWithdrewIterator struct {
	Event *MockStrategyWithdrew // Event containing the contract specifics and raw log

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
func (it *MockStrategyWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockStrategyWithdrew)
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
		it.Event = new(MockStrategyWithdrew)
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
func (it *MockStrategyWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockStrategyWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockStrategyWithdrew represents a Withdrew event raised by the MockStrategy contract.
type MockStrategyWithdrew struct {
	Pool   *big.Int
	Safe   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrew is a free log retrieval operation binding the contract event 0xf4c56bb5a568ebcb1087fa42400137fd62e67119180ddfe97906e108ecc0e633.
//
// Solidity: event Withdrew(uint256 pool, address indexed safe, uint256 amount)
func (_MockStrategy *MockStrategyFilterer) FilterWithdrew(opts *bind.FilterOpts, safe []common.Address) (*MockStrategyWithdrewIterator, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _MockStrategy.contract.FilterLogs(opts, "Withdrew", safeRule)
	if err != nil {
		return nil, err
	}
	return &MockStrategyWithdrewIterator{contract: _MockStrategy.contract, event: "Withdrew", logs: logs, sub: sub}, nil
}

// WatchWithdrew is a free log subscription operation binding the contract event 0xf4c56bb5a568ebcb1087fa42400137fd62e67119180ddfe97906e108ecc0e633.
//
// Solidity: event Withdrew(uint256 pool, address indexed safe, uint256 amount)
func (_MockStrategy *MockStrategyFilterer) WatchWithdrew(opts *bind.WatchOpts, sink chan<- *MockStrategyWithdrew, safe []common.Address) (event.Subscription, error) {

	var safeRule []interface{}
	for _, safeItem := range safe {
		safeRule = append(safeRule, safeItem)
	}

	logs, sub, err := _MockStrategy.contract.WatchLogs(opts, "Withdrew", safeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockStrategyWithdrew)
				if err := _MockStrategy.contract.UnpackLog(event, "Withdrew", log); err != nil {
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
func (_MockStrategy *MockStrategyFilterer) ParseWithdrew(log types.Log) (*MockStrategyWithdrew, error) {
	event := new(MockStrategyWithdrew)
	if err := _MockStrategy.contract.UnpackLog(event, "Withdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
