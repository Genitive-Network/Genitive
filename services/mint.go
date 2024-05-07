package services

import (
	"Genitive/config"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"math/big"
	"os"
)

func Mint(address string, amount int64) error {
	dir, err := os.Getwd()
	fmt.Println(dir + "/config/example.env")
	if err := godotenv.Load(dir + "/example.env"); err != nil {
		fmt.Errorf(err.Error())
		return err
	}
	privateKeyString := config.GetConfig().Options.PrivateKey

	client, err := ethclient.Dial(config.GetConfig().Options.BevmRpc)
	if err != nil {
		return err
	}
	contractAddress := common.HexToAddress(config.GetConfig().Options.ContractAddress)
	//contractABI, err := abi.JSON(strings.NewReader(abiString))
	privateKey, _ := crypto.HexToECDSA(privateKeyString)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Errorf("error:", err)
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	amountData := decimal.NewFromBigInt(big.NewInt(amount), 18).BigInt()

	to := common.HexToAddress(address)

	contractInstance, err := NewXbtc(contractAddress, client)
	tx, err := contractInstance.Mint(auth, to, amountData)
	if err != nil {
		return ErrFailedExtractToken
	} else {
		fmt.Println("tx sent: ", tx.Hash().Hex())
	}

	return nil
}
