// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package services

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

// XbtcMetaData contains all meta data concerning the Xbtc contract.
var XbtcMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"Blacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBlacklister\",\"type\":\"address\"}],\"name\":\"BlacklisterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMasterMinter\",\"type\":\"address\"}],\"name\":\"MasterMinterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minterAllowedAmount\",\"type\":\"uint256\"}],\"name\":\"MinterConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMinter\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"UnBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"blacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blacklister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minterAllowedAmount\",\"type\":\"uint256\"}],\"name\":\"configureMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currency\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"destroyBlackFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenCurrency\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"newMasterMinter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newBlacklister\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"minterAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"removeMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"unBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newBlacklister\",\"type\":\"address\"}],\"name\":\"updateBlacklister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMasterMinter\",\"type\":\"address\"}],\"name\":\"updateMasterMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"}]",
}

// XbtcABI is the input ABI used to generate the binding from.
// Deprecated: Use XbtcMetaData.ABI instead.
var XbtcABI = XbtcMetaData.ABI

// Xbtc is an auto generated Go binding around an Ethereum contract.
type Xbtc struct {
	XbtcCaller     // Read-only binding to the contract
	XbtcTransactor // Write-only binding to the contract
	XbtcFilterer   // Log filterer for contract events
}

// XbtcCaller is an auto generated read-only Go binding around an Ethereum contract.
type XbtcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XbtcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XbtcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XbtcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XbtcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XbtcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XbtcSession struct {
	Contract     *Xbtc             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XbtcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XbtcCallerSession struct {
	Contract *XbtcCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// XbtcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XbtcTransactorSession struct {
	Contract     *XbtcTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XbtcRaw is an auto generated low-level Go binding around an Ethereum contract.
type XbtcRaw struct {
	Contract *Xbtc // Generic contract binding to access the raw methods on
}

// XbtcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XbtcCallerRaw struct {
	Contract *XbtcCaller // Generic read-only contract binding to access the raw methods on
}

// XbtcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XbtcTransactorRaw struct {
	Contract *XbtcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXbtc creates a new instance of Xbtc, bound to a specific deployed contract.
func NewXbtc(address common.Address, backend bind.ContractBackend) (*Xbtc, error) {
	contract, err := bindXbtc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Xbtc{XbtcCaller: XbtcCaller{contract: contract}, XbtcTransactor: XbtcTransactor{contract: contract}, XbtcFilterer: XbtcFilterer{contract: contract}}, nil
}

// NewXbtcCaller creates a new read-only instance of Xbtc, bound to a specific deployed contract.
func NewXbtcCaller(address common.Address, caller bind.ContractCaller) (*XbtcCaller, error) {
	contract, err := bindXbtc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XbtcCaller{contract: contract}, nil
}

// NewXbtcTransactor creates a new write-only instance of Xbtc, bound to a specific deployed contract.
func NewXbtcTransactor(address common.Address, transactor bind.ContractTransactor) (*XbtcTransactor, error) {
	contract, err := bindXbtc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XbtcTransactor{contract: contract}, nil
}

// NewXbtcFilterer creates a new log filterer instance of Xbtc, bound to a specific deployed contract.
func NewXbtcFilterer(address common.Address, filterer bind.ContractFilterer) (*XbtcFilterer, error) {
	contract, err := bindXbtc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XbtcFilterer{contract: contract}, nil
}

// bindXbtc binds a generic wrapper to an already deployed contract.
func bindXbtc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := XbtcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Xbtc *XbtcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Xbtc.Contract.XbtcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Xbtc *XbtcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xbtc.Contract.XbtcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Xbtc *XbtcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Xbtc.Contract.XbtcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Xbtc *XbtcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Xbtc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Xbtc *XbtcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xbtc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Xbtc *XbtcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Xbtc.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Xbtc *XbtcCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Xbtc *XbtcSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Xbtc.Contract.Allowance(&_Xbtc.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Xbtc *XbtcCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Xbtc.Contract.Allowance(&_Xbtc.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Xbtc *XbtcCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Xbtc *XbtcSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Xbtc.Contract.BalanceOf(&_Xbtc.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Xbtc *XbtcCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Xbtc.Contract.BalanceOf(&_Xbtc.CallOpts, account)
}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Xbtc *XbtcCaller) Blacklister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "blacklister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Xbtc *XbtcSession) Blacklister() (common.Address, error) {
	return _Xbtc.Contract.Blacklister(&_Xbtc.CallOpts)
}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Xbtc *XbtcCallerSession) Blacklister() (common.Address, error) {
	return _Xbtc.Contract.Blacklister(&_Xbtc.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(string)
func (_Xbtc *XbtcCaller) Currency(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "currency")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(string)
func (_Xbtc *XbtcSession) Currency() (string, error) {
	return _Xbtc.Contract.Currency(&_Xbtc.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(string)
func (_Xbtc *XbtcCallerSession) Currency() (string, error) {
	return _Xbtc.Contract.Currency(&_Xbtc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Xbtc *XbtcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Xbtc *XbtcSession) Decimals() (uint8, error) {
	return _Xbtc.Contract.Decimals(&_Xbtc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Xbtc *XbtcCallerSession) Decimals() (uint8, error) {
	return _Xbtc.Contract.Decimals(&_Xbtc.CallOpts)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_Xbtc *XbtcCaller) IsBlacklisted(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "isBlacklisted", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_Xbtc *XbtcSession) IsBlacklisted(_account common.Address) (bool, error) {
	return _Xbtc.Contract.IsBlacklisted(&_Xbtc.CallOpts, _account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_Xbtc *XbtcCallerSession) IsBlacklisted(_account common.Address) (bool, error) {
	return _Xbtc.Contract.IsBlacklisted(&_Xbtc.CallOpts, _account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Xbtc *XbtcCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "isMinter", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Xbtc *XbtcSession) IsMinter(account common.Address) (bool, error) {
	return _Xbtc.Contract.IsMinter(&_Xbtc.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_Xbtc *XbtcCallerSession) IsMinter(account common.Address) (bool, error) {
	return _Xbtc.Contract.IsMinter(&_Xbtc.CallOpts, account)
}

// MasterMinter is a free data retrieval call binding the contract method 0x35d99f35.
//
// Solidity: function masterMinter() view returns(address)
func (_Xbtc *XbtcCaller) MasterMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "masterMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterMinter is a free data retrieval call binding the contract method 0x35d99f35.
//
// Solidity: function masterMinter() view returns(address)
func (_Xbtc *XbtcSession) MasterMinter() (common.Address, error) {
	return _Xbtc.Contract.MasterMinter(&_Xbtc.CallOpts)
}

// MasterMinter is a free data retrieval call binding the contract method 0x35d99f35.
//
// Solidity: function masterMinter() view returns(address)
func (_Xbtc *XbtcCallerSession) MasterMinter() (common.Address, error) {
	return _Xbtc.Contract.MasterMinter(&_Xbtc.CallOpts)
}

// MinterAllowance is a free data retrieval call binding the contract method 0x8a6db9c3.
//
// Solidity: function minterAllowance(address minter) view returns(uint256)
func (_Xbtc *XbtcCaller) MinterAllowance(opts *bind.CallOpts, minter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "minterAllowance", minter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinterAllowance is a free data retrieval call binding the contract method 0x8a6db9c3.
//
// Solidity: function minterAllowance(address minter) view returns(uint256)
func (_Xbtc *XbtcSession) MinterAllowance(minter common.Address) (*big.Int, error) {
	return _Xbtc.Contract.MinterAllowance(&_Xbtc.CallOpts, minter)
}

// MinterAllowance is a free data retrieval call binding the contract method 0x8a6db9c3.
//
// Solidity: function minterAllowance(address minter) view returns(uint256)
func (_Xbtc *XbtcCallerSession) MinterAllowance(minter common.Address) (*big.Int, error) {
	return _Xbtc.Contract.MinterAllowance(&_Xbtc.CallOpts, minter)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Xbtc *XbtcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Xbtc *XbtcSession) Name() (string, error) {
	return _Xbtc.Contract.Name(&_Xbtc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Xbtc *XbtcCallerSession) Name() (string, error) {
	return _Xbtc.Contract.Name(&_Xbtc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Xbtc *XbtcCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Xbtc *XbtcSession) Owner() (common.Address, error) {
	return _Xbtc.Contract.Owner(&_Xbtc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Xbtc *XbtcCallerSession) Owner() (common.Address, error) {
	return _Xbtc.Contract.Owner(&_Xbtc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Xbtc *XbtcCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Xbtc *XbtcSession) Paused() (bool, error) {
	return _Xbtc.Contract.Paused(&_Xbtc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Xbtc *XbtcCallerSession) Paused() (bool, error) {
	return _Xbtc.Contract.Paused(&_Xbtc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Xbtc *XbtcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Xbtc *XbtcSession) Symbol() (string, error) {
	return _Xbtc.Contract.Symbol(&_Xbtc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Xbtc *XbtcCallerSession) Symbol() (string, error) {
	return _Xbtc.Contract.Symbol(&_Xbtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Xbtc *XbtcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xbtc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Xbtc *XbtcSession) TotalSupply() (*big.Int, error) {
	return _Xbtc.Contract.TotalSupply(&_Xbtc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Xbtc *XbtcCallerSession) TotalSupply() (*big.Int, error) {
	return _Xbtc.Contract.TotalSupply(&_Xbtc.CallOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Xbtc *XbtcTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Xbtc *XbtcSession) Admin() (*types.Transaction, error) {
	return _Xbtc.Contract.Admin(&_Xbtc.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Xbtc *XbtcTransactorSession) Admin() (*types.Transaction, error) {
	return _Xbtc.Contract.Admin(&_Xbtc.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Xbtc *XbtcTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Xbtc *XbtcSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Xbtc.Contract.Approve(&_Xbtc.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Xbtc *XbtcTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Xbtc.Contract.Approve(&_Xbtc.TransactOpts, spender, value)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_Xbtc *XbtcTransactor) Blacklist(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "blacklist", _account)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_Xbtc *XbtcSession) Blacklist(_account common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.Blacklist(&_Xbtc.TransactOpts, _account)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_Xbtc *XbtcTransactorSession) Blacklist(_account common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.Blacklist(&_Xbtc.TransactOpts, _account)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 _amount) returns()
func (_Xbtc *XbtcTransactor) Burn(opts *bind.TransactOpts, from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "burn", from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 _amount) returns()
func (_Xbtc *XbtcSession) Burn(from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Xbtc.Contract.Burn(&_Xbtc.TransactOpts, from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 _amount) returns()
func (_Xbtc *XbtcTransactorSession) Burn(from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Xbtc.Contract.Burn(&_Xbtc.TransactOpts, from, _amount)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Xbtc *XbtcTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Xbtc *XbtcSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.ChangeAdmin(&_Xbtc.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Xbtc *XbtcTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.ChangeAdmin(&_Xbtc.TransactOpts, newAdmin)
}

// ConfigureMinter is a paid mutator transaction binding the contract method 0x4e44d956.
//
// Solidity: function configureMinter(address minter, uint256 minterAllowedAmount) returns(bool)
func (_Xbtc *XbtcTransactor) ConfigureMinter(opts *bind.TransactOpts, minter common.Address, minterAllowedAmount *big.Int) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "configureMinter", minter, minterAllowedAmount)
}

// ConfigureMinter is a paid mutator transaction binding the contract method 0x4e44d956.
//
// Solidity: function configureMinter(address minter, uint256 minterAllowedAmount) returns(bool)
func (_Xbtc *XbtcSession) ConfigureMinter(minter common.Address, minterAllowedAmount *big.Int) (*types.Transaction, error) {
	return _Xbtc.Contract.ConfigureMinter(&_Xbtc.TransactOpts, minter, minterAllowedAmount)
}

// ConfigureMinter is a paid mutator transaction binding the contract method 0x4e44d956.
//
// Solidity: function configureMinter(address minter, uint256 minterAllowedAmount) returns(bool)
func (_Xbtc *XbtcTransactorSession) ConfigureMinter(minter common.Address, minterAllowedAmount *big.Int) (*types.Transaction, error) {
	return _Xbtc.Contract.ConfigureMinter(&_Xbtc.TransactOpts, minter, minterAllowedAmount)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _from) returns()
func (_Xbtc *XbtcTransactor) DestroyBlackFunds(opts *bind.TransactOpts, _from common.Address) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "destroyBlackFunds", _from)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _from) returns()
func (_Xbtc *XbtcSession) DestroyBlackFunds(_from common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.DestroyBlackFunds(&_Xbtc.TransactOpts, _from)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _from) returns()
func (_Xbtc *XbtcTransactorSession) DestroyBlackFunds(_from common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.DestroyBlackFunds(&_Xbtc.TransactOpts, _from)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Xbtc *XbtcTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Xbtc *XbtcSession) Implementation() (*types.Transaction, error) {
	return _Xbtc.Contract.Implementation(&_Xbtc.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Xbtc *XbtcTransactorSession) Implementation() (*types.Transaction, error) {
	return _Xbtc.Contract.Implementation(&_Xbtc.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xafa6e0db.
//
// Solidity: function initialize(string tokenName, string tokenSymbol, string tokenCurrency, uint8 tokenDecimals, address newMasterMinter, address newBlacklister, address newOwner) returns()
func (_Xbtc *XbtcTransactor) Initialize(opts *bind.TransactOpts, tokenName string, tokenSymbol string, tokenCurrency string, tokenDecimals uint8, newMasterMinter common.Address, newBlacklister common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "initialize", tokenName, tokenSymbol, tokenCurrency, tokenDecimals, newMasterMinter, newBlacklister, newOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xafa6e0db.
//
// Solidity: function initialize(string tokenName, string tokenSymbol, string tokenCurrency, uint8 tokenDecimals, address newMasterMinter, address newBlacklister, address newOwner) returns()
func (_Xbtc *XbtcSession) Initialize(tokenName string, tokenSymbol string, tokenCurrency string, tokenDecimals uint8, newMasterMinter common.Address, newBlacklister common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.Initialize(&_Xbtc.TransactOpts, tokenName, tokenSymbol, tokenCurrency, tokenDecimals, newMasterMinter, newBlacklister, newOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xafa6e0db.
//
// Solidity: function initialize(string tokenName, string tokenSymbol, string tokenCurrency, uint8 tokenDecimals, address newMasterMinter, address newBlacklister, address newOwner) returns()
func (_Xbtc *XbtcTransactorSession) Initialize(tokenName string, tokenSymbol string, tokenCurrency string, tokenDecimals uint8, newMasterMinter common.Address, newBlacklister common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.Initialize(&_Xbtc.TransactOpts, tokenName, tokenSymbol, tokenCurrency, tokenDecimals, newMasterMinter, newBlacklister, newOwner)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_Xbtc *XbtcTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_Xbtc *XbtcSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Xbtc.Contract.Mint(&_Xbtc.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_Xbtc *XbtcTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Xbtc.Contract.Mint(&_Xbtc.TransactOpts, _to, _amount)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns(bool)
func (_Xbtc *XbtcTransactor) RemoveMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "removeMinter", minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns(bool)
func (_Xbtc *XbtcSession) RemoveMinter(minter common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.RemoveMinter(&_Xbtc.TransactOpts, minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns(bool)
func (_Xbtc *XbtcTransactorSession) RemoveMinter(minter common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.RemoveMinter(&_Xbtc.TransactOpts, minter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Xbtc *XbtcTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Xbtc *XbtcSession) RenounceOwnership() (*types.Transaction, error) {
	return _Xbtc.Contract.RenounceOwnership(&_Xbtc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Xbtc *XbtcTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Xbtc.Contract.RenounceOwnership(&_Xbtc.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Xbtc *XbtcTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Xbtc *XbtcSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Xbtc.Contract.Transfer(&_Xbtc.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Xbtc *XbtcTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Xbtc.Contract.Transfer(&_Xbtc.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Xbtc *XbtcTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Xbtc *XbtcSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Xbtc.Contract.TransferFrom(&_Xbtc.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Xbtc *XbtcTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Xbtc.Contract.TransferFrom(&_Xbtc.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Xbtc *XbtcTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Xbtc *XbtcSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.TransferOwnership(&_Xbtc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Xbtc *XbtcTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.TransferOwnership(&_Xbtc.TransactOpts, newOwner)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_Xbtc *XbtcTransactor) UnBlacklist(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "unBlacklist", _account)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_Xbtc *XbtcSession) UnBlacklist(_account common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.UnBlacklist(&_Xbtc.TransactOpts, _account)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_Xbtc *XbtcTransactorSession) UnBlacklist(_account common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.UnBlacklist(&_Xbtc.TransactOpts, _account)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_Xbtc *XbtcTransactor) UpdateBlacklister(opts *bind.TransactOpts, _newBlacklister common.Address) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "updateBlacklister", _newBlacklister)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_Xbtc *XbtcSession) UpdateBlacklister(_newBlacklister common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.UpdateBlacklister(&_Xbtc.TransactOpts, _newBlacklister)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_Xbtc *XbtcTransactorSession) UpdateBlacklister(_newBlacklister common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.UpdateBlacklister(&_Xbtc.TransactOpts, _newBlacklister)
}

// UpdateMasterMinter is a paid mutator transaction binding the contract method 0xaa20e1e4.
//
// Solidity: function updateMasterMinter(address _newMasterMinter) returns()
func (_Xbtc *XbtcTransactor) UpdateMasterMinter(opts *bind.TransactOpts, _newMasterMinter common.Address) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "updateMasterMinter", _newMasterMinter)
}

// UpdateMasterMinter is a paid mutator transaction binding the contract method 0xaa20e1e4.
//
// Solidity: function updateMasterMinter(address _newMasterMinter) returns()
func (_Xbtc *XbtcSession) UpdateMasterMinter(_newMasterMinter common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.UpdateMasterMinter(&_Xbtc.TransactOpts, _newMasterMinter)
}

// UpdateMasterMinter is a paid mutator transaction binding the contract method 0xaa20e1e4.
//
// Solidity: function updateMasterMinter(address _newMasterMinter) returns()
func (_Xbtc *XbtcTransactorSession) UpdateMasterMinter(_newMasterMinter common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.UpdateMasterMinter(&_Xbtc.TransactOpts, _newMasterMinter)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Xbtc *XbtcTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Xbtc *XbtcSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.UpgradeTo(&_Xbtc.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Xbtc *XbtcTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Xbtc.Contract.UpgradeTo(&_Xbtc.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Xbtc *XbtcTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Xbtc.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Xbtc *XbtcSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Xbtc.Contract.UpgradeToAndCall(&_Xbtc.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Xbtc *XbtcTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Xbtc.Contract.UpgradeToAndCall(&_Xbtc.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Xbtc *XbtcTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Xbtc.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Xbtc *XbtcSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Xbtc.Contract.Fallback(&_Xbtc.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Xbtc *XbtcTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Xbtc.Contract.Fallback(&_Xbtc.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Xbtc *XbtcTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xbtc.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Xbtc *XbtcSession) Receive() (*types.Transaction, error) {
	return _Xbtc.Contract.Receive(&_Xbtc.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Xbtc *XbtcTransactorSession) Receive() (*types.Transaction, error) {
	return _Xbtc.Contract.Receive(&_Xbtc.TransactOpts)
}

// XbtcAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Xbtc contract.
type XbtcAdminChangedIterator struct {
	Event *XbtcAdminChanged // Event containing the contract specifics and raw log

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
func (it *XbtcAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcAdminChanged)
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
		it.Event = new(XbtcAdminChanged)
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
func (it *XbtcAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcAdminChanged represents a AdminChanged event raised by the Xbtc contract.
type XbtcAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Xbtc *XbtcFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*XbtcAdminChangedIterator, error) {

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &XbtcAdminChangedIterator{contract: _Xbtc.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Xbtc *XbtcFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *XbtcAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcAdminChanged)
				if err := _Xbtc.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Xbtc *XbtcFilterer) ParseAdminChanged(log types.Log) (*XbtcAdminChanged, error) {
	event := new(XbtcAdminChanged)
	if err := _Xbtc.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Xbtc contract.
type XbtcApprovalIterator struct {
	Event *XbtcApproval // Event containing the contract specifics and raw log

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
func (it *XbtcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcApproval)
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
		it.Event = new(XbtcApproval)
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
func (it *XbtcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcApproval represents a Approval event raised by the Xbtc contract.
type XbtcApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Xbtc *XbtcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*XbtcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &XbtcApprovalIterator{contract: _Xbtc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Xbtc *XbtcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *XbtcApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcApproval)
				if err := _Xbtc.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Xbtc *XbtcFilterer) ParseApproval(log types.Log) (*XbtcApproval, error) {
	event := new(XbtcApproval)
	if err := _Xbtc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Xbtc contract.
type XbtcBeaconUpgradedIterator struct {
	Event *XbtcBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *XbtcBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcBeaconUpgraded)
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
		it.Event = new(XbtcBeaconUpgraded)
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
func (it *XbtcBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcBeaconUpgraded represents a BeaconUpgraded event raised by the Xbtc contract.
type XbtcBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Xbtc *XbtcFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*XbtcBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &XbtcBeaconUpgradedIterator{contract: _Xbtc.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Xbtc *XbtcFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *XbtcBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcBeaconUpgraded)
				if err := _Xbtc.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Xbtc *XbtcFilterer) ParseBeaconUpgraded(log types.Log) (*XbtcBeaconUpgraded, error) {
	event := new(XbtcBeaconUpgraded)
	if err := _Xbtc.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcBlacklistedIterator is returned from FilterBlacklisted and is used to iterate over the raw logs and unpacked data for Blacklisted events raised by the Xbtc contract.
type XbtcBlacklistedIterator struct {
	Event *XbtcBlacklisted // Event containing the contract specifics and raw log

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
func (it *XbtcBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcBlacklisted)
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
		it.Event = new(XbtcBlacklisted)
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
func (it *XbtcBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcBlacklisted represents a Blacklisted event raised by the Xbtc contract.
type XbtcBlacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlacklisted is a free log retrieval operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_Xbtc *XbtcFilterer) FilterBlacklisted(opts *bind.FilterOpts, _account []common.Address) (*XbtcBlacklistedIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "Blacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return &XbtcBlacklistedIterator{contract: _Xbtc.contract, event: "Blacklisted", logs: logs, sub: sub}, nil
}

// WatchBlacklisted is a free log subscription operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_Xbtc *XbtcFilterer) WatchBlacklisted(opts *bind.WatchOpts, sink chan<- *XbtcBlacklisted, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "Blacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcBlacklisted)
				if err := _Xbtc.contract.UnpackLog(event, "Blacklisted", log); err != nil {
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

// ParseBlacklisted is a log parse operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_Xbtc *XbtcFilterer) ParseBlacklisted(log types.Log) (*XbtcBlacklisted, error) {
	event := new(XbtcBlacklisted)
	if err := _Xbtc.contract.UnpackLog(event, "Blacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcBlacklisterChangedIterator is returned from FilterBlacklisterChanged and is used to iterate over the raw logs and unpacked data for BlacklisterChanged events raised by the Xbtc contract.
type XbtcBlacklisterChangedIterator struct {
	Event *XbtcBlacklisterChanged // Event containing the contract specifics and raw log

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
func (it *XbtcBlacklisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcBlacklisterChanged)
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
		it.Event = new(XbtcBlacklisterChanged)
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
func (it *XbtcBlacklisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcBlacklisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcBlacklisterChanged represents a BlacklisterChanged event raised by the Xbtc contract.
type XbtcBlacklisterChanged struct {
	NewBlacklister common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlacklisterChanged is a free log retrieval operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_Xbtc *XbtcFilterer) FilterBlacklisterChanged(opts *bind.FilterOpts, newBlacklister []common.Address) (*XbtcBlacklisterChangedIterator, error) {

	var newBlacklisterRule []interface{}
	for _, newBlacklisterItem := range newBlacklister {
		newBlacklisterRule = append(newBlacklisterRule, newBlacklisterItem)
	}

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "BlacklisterChanged", newBlacklisterRule)
	if err != nil {
		return nil, err
	}
	return &XbtcBlacklisterChangedIterator{contract: _Xbtc.contract, event: "BlacklisterChanged", logs: logs, sub: sub}, nil
}

// WatchBlacklisterChanged is a free log subscription operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_Xbtc *XbtcFilterer) WatchBlacklisterChanged(opts *bind.WatchOpts, sink chan<- *XbtcBlacklisterChanged, newBlacklister []common.Address) (event.Subscription, error) {

	var newBlacklisterRule []interface{}
	for _, newBlacklisterItem := range newBlacklister {
		newBlacklisterRule = append(newBlacklisterRule, newBlacklisterItem)
	}

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "BlacklisterChanged", newBlacklisterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcBlacklisterChanged)
				if err := _Xbtc.contract.UnpackLog(event, "BlacklisterChanged", log); err != nil {
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

// ParseBlacklisterChanged is a log parse operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_Xbtc *XbtcFilterer) ParseBlacklisterChanged(log types.Log) (*XbtcBlacklisterChanged, error) {
	event := new(XbtcBlacklisterChanged)
	if err := _Xbtc.contract.UnpackLog(event, "BlacklisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Xbtc contract.
type XbtcBurnIterator struct {
	Event *XbtcBurn // Event containing the contract specifics and raw log

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
func (it *XbtcBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcBurn)
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
		it.Event = new(XbtcBurn)
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
func (it *XbtcBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcBurn represents a Burn event raised by the Xbtc contract.
type XbtcBurn struct {
	Burner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_Xbtc *XbtcFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*XbtcBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &XbtcBurnIterator{contract: _Xbtc.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_Xbtc *XbtcFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *XbtcBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcBurn)
				if err := _Xbtc.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_Xbtc *XbtcFilterer) ParseBurn(log types.Log) (*XbtcBurn, error) {
	event := new(XbtcBurn)
	if err := _Xbtc.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Xbtc contract.
type XbtcInitializedIterator struct {
	Event *XbtcInitialized // Event containing the contract specifics and raw log

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
func (it *XbtcInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcInitialized)
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
		it.Event = new(XbtcInitialized)
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
func (it *XbtcInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcInitialized represents a Initialized event raised by the Xbtc contract.
type XbtcInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Xbtc *XbtcFilterer) FilterInitialized(opts *bind.FilterOpts) (*XbtcInitializedIterator, error) {

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &XbtcInitializedIterator{contract: _Xbtc.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Xbtc *XbtcFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *XbtcInitialized) (event.Subscription, error) {

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcInitialized)
				if err := _Xbtc.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Xbtc *XbtcFilterer) ParseInitialized(log types.Log) (*XbtcInitialized, error) {
	event := new(XbtcInitialized)
	if err := _Xbtc.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcMasterMinterChangedIterator is returned from FilterMasterMinterChanged and is used to iterate over the raw logs and unpacked data for MasterMinterChanged events raised by the Xbtc contract.
type XbtcMasterMinterChangedIterator struct {
	Event *XbtcMasterMinterChanged // Event containing the contract specifics and raw log

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
func (it *XbtcMasterMinterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcMasterMinterChanged)
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
		it.Event = new(XbtcMasterMinterChanged)
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
func (it *XbtcMasterMinterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcMasterMinterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcMasterMinterChanged represents a MasterMinterChanged event raised by the Xbtc contract.
type XbtcMasterMinterChanged struct {
	NewMasterMinter common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMasterMinterChanged is a free log retrieval operation binding the contract event 0xdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e6.
//
// Solidity: event MasterMinterChanged(address indexed newMasterMinter)
func (_Xbtc *XbtcFilterer) FilterMasterMinterChanged(opts *bind.FilterOpts, newMasterMinter []common.Address) (*XbtcMasterMinterChangedIterator, error) {

	var newMasterMinterRule []interface{}
	for _, newMasterMinterItem := range newMasterMinter {
		newMasterMinterRule = append(newMasterMinterRule, newMasterMinterItem)
	}

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "MasterMinterChanged", newMasterMinterRule)
	if err != nil {
		return nil, err
	}
	return &XbtcMasterMinterChangedIterator{contract: _Xbtc.contract, event: "MasterMinterChanged", logs: logs, sub: sub}, nil
}

// WatchMasterMinterChanged is a free log subscription operation binding the contract event 0xdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e6.
//
// Solidity: event MasterMinterChanged(address indexed newMasterMinter)
func (_Xbtc *XbtcFilterer) WatchMasterMinterChanged(opts *bind.WatchOpts, sink chan<- *XbtcMasterMinterChanged, newMasterMinter []common.Address) (event.Subscription, error) {

	var newMasterMinterRule []interface{}
	for _, newMasterMinterItem := range newMasterMinter {
		newMasterMinterRule = append(newMasterMinterRule, newMasterMinterItem)
	}

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "MasterMinterChanged", newMasterMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcMasterMinterChanged)
				if err := _Xbtc.contract.UnpackLog(event, "MasterMinterChanged", log); err != nil {
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

// ParseMasterMinterChanged is a log parse operation binding the contract event 0xdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e6.
//
// Solidity: event MasterMinterChanged(address indexed newMasterMinter)
func (_Xbtc *XbtcFilterer) ParseMasterMinterChanged(log types.Log) (*XbtcMasterMinterChanged, error) {
	event := new(XbtcMasterMinterChanged)
	if err := _Xbtc.contract.UnpackLog(event, "MasterMinterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Xbtc contract.
type XbtcMintIterator struct {
	Event *XbtcMint // Event containing the contract specifics and raw log

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
func (it *XbtcMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcMint)
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
		it.Event = new(XbtcMint)
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
func (it *XbtcMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcMint represents a Mint event raised by the Xbtc contract.
type XbtcMint struct {
	Minter common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed minter, address indexed to, uint256 amount)
func (_Xbtc *XbtcFilterer) FilterMint(opts *bind.FilterOpts, minter []common.Address, to []common.Address) (*XbtcMintIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "Mint", minterRule, toRule)
	if err != nil {
		return nil, err
	}
	return &XbtcMintIterator{contract: _Xbtc.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed minter, address indexed to, uint256 amount)
func (_Xbtc *XbtcFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *XbtcMint, minter []common.Address, to []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "Mint", minterRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcMint)
				if err := _Xbtc.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed minter, address indexed to, uint256 amount)
func (_Xbtc *XbtcFilterer) ParseMint(log types.Log) (*XbtcMint, error) {
	event := new(XbtcMint)
	if err := _Xbtc.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcMinterConfiguredIterator is returned from FilterMinterConfigured and is used to iterate over the raw logs and unpacked data for MinterConfigured events raised by the Xbtc contract.
type XbtcMinterConfiguredIterator struct {
	Event *XbtcMinterConfigured // Event containing the contract specifics and raw log

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
func (it *XbtcMinterConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcMinterConfigured)
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
		it.Event = new(XbtcMinterConfigured)
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
func (it *XbtcMinterConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcMinterConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcMinterConfigured represents a MinterConfigured event raised by the Xbtc contract.
type XbtcMinterConfigured struct {
	Minter              common.Address
	MinterAllowedAmount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMinterConfigured is a free log retrieval operation binding the contract event 0x46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20.
//
// Solidity: event MinterConfigured(address indexed minter, uint256 minterAllowedAmount)
func (_Xbtc *XbtcFilterer) FilterMinterConfigured(opts *bind.FilterOpts, minter []common.Address) (*XbtcMinterConfiguredIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "MinterConfigured", minterRule)
	if err != nil {
		return nil, err
	}
	return &XbtcMinterConfiguredIterator{contract: _Xbtc.contract, event: "MinterConfigured", logs: logs, sub: sub}, nil
}

// WatchMinterConfigured is a free log subscription operation binding the contract event 0x46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20.
//
// Solidity: event MinterConfigured(address indexed minter, uint256 minterAllowedAmount)
func (_Xbtc *XbtcFilterer) WatchMinterConfigured(opts *bind.WatchOpts, sink chan<- *XbtcMinterConfigured, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "MinterConfigured", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcMinterConfigured)
				if err := _Xbtc.contract.UnpackLog(event, "MinterConfigured", log); err != nil {
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

// ParseMinterConfigured is a log parse operation binding the contract event 0x46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20.
//
// Solidity: event MinterConfigured(address indexed minter, uint256 minterAllowedAmount)
func (_Xbtc *XbtcFilterer) ParseMinterConfigured(log types.Log) (*XbtcMinterConfigured, error) {
	event := new(XbtcMinterConfigured)
	if err := _Xbtc.contract.UnpackLog(event, "MinterConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the Xbtc contract.
type XbtcMinterRemovedIterator struct {
	Event *XbtcMinterRemoved // Event containing the contract specifics and raw log

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
func (it *XbtcMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcMinterRemoved)
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
		it.Event = new(XbtcMinterRemoved)
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
func (it *XbtcMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcMinterRemoved represents a MinterRemoved event raised by the Xbtc contract.
type XbtcMinterRemoved struct {
	OldMinter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed oldMinter)
func (_Xbtc *XbtcFilterer) FilterMinterRemoved(opts *bind.FilterOpts, oldMinter []common.Address) (*XbtcMinterRemovedIterator, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "MinterRemoved", oldMinterRule)
	if err != nil {
		return nil, err
	}
	return &XbtcMinterRemovedIterator{contract: _Xbtc.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed oldMinter)
func (_Xbtc *XbtcFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *XbtcMinterRemoved, oldMinter []common.Address) (event.Subscription, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "MinterRemoved", oldMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcMinterRemoved)
				if err := _Xbtc.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed oldMinter)
func (_Xbtc *XbtcFilterer) ParseMinterRemoved(log types.Log) (*XbtcMinterRemoved, error) {
	event := new(XbtcMinterRemoved)
	if err := _Xbtc.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Xbtc contract.
type XbtcOwnershipTransferredIterator struct {
	Event *XbtcOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *XbtcOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcOwnershipTransferred)
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
		it.Event = new(XbtcOwnershipTransferred)
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
func (it *XbtcOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcOwnershipTransferred represents a OwnershipTransferred event raised by the Xbtc contract.
type XbtcOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Xbtc *XbtcFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*XbtcOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &XbtcOwnershipTransferredIterator{contract: _Xbtc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Xbtc *XbtcFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *XbtcOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcOwnershipTransferred)
				if err := _Xbtc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Xbtc *XbtcFilterer) ParseOwnershipTransferred(log types.Log) (*XbtcOwnershipTransferred, error) {
	event := new(XbtcOwnershipTransferred)
	if err := _Xbtc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Xbtc contract.
type XbtcPausedIterator struct {
	Event *XbtcPaused // Event containing the contract specifics and raw log

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
func (it *XbtcPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcPaused)
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
		it.Event = new(XbtcPaused)
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
func (it *XbtcPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcPaused represents a Paused event raised by the Xbtc contract.
type XbtcPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Xbtc *XbtcFilterer) FilterPaused(opts *bind.FilterOpts) (*XbtcPausedIterator, error) {

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &XbtcPausedIterator{contract: _Xbtc.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Xbtc *XbtcFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *XbtcPaused) (event.Subscription, error) {

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcPaused)
				if err := _Xbtc.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Xbtc *XbtcFilterer) ParsePaused(log types.Log) (*XbtcPaused, error) {
	event := new(XbtcPaused)
	if err := _Xbtc.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Xbtc contract.
type XbtcTransferIterator struct {
	Event *XbtcTransfer // Event containing the contract specifics and raw log

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
func (it *XbtcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcTransfer)
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
		it.Event = new(XbtcTransfer)
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
func (it *XbtcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcTransfer represents a Transfer event raised by the Xbtc contract.
type XbtcTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Xbtc *XbtcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*XbtcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &XbtcTransferIterator{contract: _Xbtc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Xbtc *XbtcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *XbtcTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcTransfer)
				if err := _Xbtc.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Xbtc *XbtcFilterer) ParseTransfer(log types.Log) (*XbtcTransfer, error) {
	event := new(XbtcTransfer)
	if err := _Xbtc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcUnBlacklistedIterator is returned from FilterUnBlacklisted and is used to iterate over the raw logs and unpacked data for UnBlacklisted events raised by the Xbtc contract.
type XbtcUnBlacklistedIterator struct {
	Event *XbtcUnBlacklisted // Event containing the contract specifics and raw log

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
func (it *XbtcUnBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcUnBlacklisted)
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
		it.Event = new(XbtcUnBlacklisted)
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
func (it *XbtcUnBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcUnBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcUnBlacklisted represents a UnBlacklisted event raised by the Xbtc contract.
type XbtcUnBlacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnBlacklisted is a free log retrieval operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_Xbtc *XbtcFilterer) FilterUnBlacklisted(opts *bind.FilterOpts, _account []common.Address) (*XbtcUnBlacklistedIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "UnBlacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return &XbtcUnBlacklistedIterator{contract: _Xbtc.contract, event: "UnBlacklisted", logs: logs, sub: sub}, nil
}

// WatchUnBlacklisted is a free log subscription operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_Xbtc *XbtcFilterer) WatchUnBlacklisted(opts *bind.WatchOpts, sink chan<- *XbtcUnBlacklisted, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "UnBlacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcUnBlacklisted)
				if err := _Xbtc.contract.UnpackLog(event, "UnBlacklisted", log); err != nil {
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

// ParseUnBlacklisted is a log parse operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_Xbtc *XbtcFilterer) ParseUnBlacklisted(log types.Log) (*XbtcUnBlacklisted, error) {
	event := new(XbtcUnBlacklisted)
	if err := _Xbtc.contract.UnpackLog(event, "UnBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Xbtc contract.
type XbtcUnpausedIterator struct {
	Event *XbtcUnpaused // Event containing the contract specifics and raw log

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
func (it *XbtcUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcUnpaused)
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
		it.Event = new(XbtcUnpaused)
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
func (it *XbtcUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcUnpaused represents a Unpaused event raised by the Xbtc contract.
type XbtcUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Xbtc *XbtcFilterer) FilterUnpaused(opts *bind.FilterOpts) (*XbtcUnpausedIterator, error) {

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &XbtcUnpausedIterator{contract: _Xbtc.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Xbtc *XbtcFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *XbtcUnpaused) (event.Subscription, error) {

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcUnpaused)
				if err := _Xbtc.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Xbtc *XbtcFilterer) ParseUnpaused(log types.Log) (*XbtcUnpaused, error) {
	event := new(XbtcUnpaused)
	if err := _Xbtc.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XbtcUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Xbtc contract.
type XbtcUpgradedIterator struct {
	Event *XbtcUpgraded // Event containing the contract specifics and raw log

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
func (it *XbtcUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XbtcUpgraded)
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
		it.Event = new(XbtcUpgraded)
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
func (it *XbtcUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XbtcUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XbtcUpgraded represents a Upgraded event raised by the Xbtc contract.
type XbtcUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Xbtc *XbtcFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*XbtcUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Xbtc.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &XbtcUpgradedIterator{contract: _Xbtc.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Xbtc *XbtcFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *XbtcUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Xbtc.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XbtcUpgraded)
				if err := _Xbtc.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Xbtc *XbtcFilterer) ParseUpgraded(log types.Log) (*XbtcUpgraded, error) {
	event := new(XbtcUpgraded)
	if err := _Xbtc.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
