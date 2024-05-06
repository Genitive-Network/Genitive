package services

import (
	"Genitive/config"
	"crypto/ecdsa"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"math/big"
	"os"

	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"

	"github.com/joho/godotenv"
)

func Burn() error {
	dir, err := os.Getwd()
	if err := godotenv.Load(dir + "/config/example.env"); err != nil {
		log.Error(err)
		return err
	}
	privateKeyString := os.Getenv("BEVM_DEV_PRIVATE_KEY")

	client, err := ethclient.Dial("wss://testnet.bevm.io/ws")
	if err != nil {
		log.Error(err)
		return err
	}
	contractAddress := common.HexToAddress(config.GetConfig().Options.ContractAddress)
	//contractABI, err := abi.JSON(strings.NewReader(abiString))
	privateKey, _ := crypto.HexToECDSA(privateKeyString)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Error(err)
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error(err)
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
		fmt.Println(err)
		return err
	} else {
		fmt.Println("tx sent: ", tx.Hash().Hex())
	}

	return nil
}
