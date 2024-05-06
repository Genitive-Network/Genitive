package services

import (
	"Genitive/config"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"os"
)

func Mint() error {
	dir, err := os.Getwd()
	fmt.Println(dir + "/config/example.env")
	if err := godotenv.Load(dir + "/example.env"); err != nil {
		log.Fatal(err)
		return err
	}
	privateKeyString := os.Getenv("BEVM_DEV_PRIVATE_KEY")

	client, err := ethclient.Dial("wss://testnet.bevm.io/ws")
	if err != nil {
		return err
	}
	contractAddress := common.HexToAddress(config.GetConfig().Options.ContractAddress)
	//contractABI, err := abi.JSON(strings.NewReader(abiString))
	privateKey, _ := crypto.HexToECDSA(privateKeyString)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
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

	amount := decimal.NewFromBigInt(big.NewInt(10), 18).BigInt()
	to := common.HexToAddress(config.GetConfig().Options.UserAddress)

	contractInstance, err := NewXbtc(contractAddress, client)
	tx, err := contractInstance.Mint(auth, to, amount)
	if err != nil {
		return ErrFailedExtractToken
	} else {
		fmt.Println("tx sent: ", tx.Hash().Hex())
	}

	return nil
}
