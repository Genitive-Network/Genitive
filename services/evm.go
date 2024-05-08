package services

import (
	"Genitive/config"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"os"
	"strings"
)

var (
	client *ethclient.Client
)

func Runbevm() {
	client, err := ethclient.Dial("wss://testnet.bevm.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.GetConfig().Options.ContractAddress)
	contractAddress := common.HexToAddress(config.GetConfig().Options.ContractAddress)
	//contractAddress := common.HexToAddress("0x30A0e025BE2bbC80948f60647c48756815b78227")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	dir, err := os.Getwd()
	abiString, err := os.ReadFile(dir + "/config/abi.json")
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	contractAbi, err := abi.JSON(strings.NewReader(string(abiString)))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:

			fmt.Println("Log Block Number:", vLog.BlockNumber)
			fmt.Println("Log Index:", vLog.Index)
			fmt.Println("Log Data:", vLog.Data)
			fmt.Println("Log Topic:", vLog.Topics)
			fmt.Println("Log TxIndex:", vLog.TxIndex)
			fmt.Println("Log :", vLog)
			receipt, err := client.TransactionReceipt(context.Background(), vLog.TxHash)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("receipt :", receipt)
			fmt.Println("--------------------------------------------------------------------------------", receipt)
			var transferEvent LogTransfer
			err = contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Println(
				"event Transfer", transferEvent.From.Hex(), transferEvent.To.Hex(), transferEvent.Amount.String(),
			)
			fmt.Println(
				"Transaction Value:", vLog.Topics[0].Hex(), vLog.Topics[1].Hex(), decimal.NewFromBigInt(transferEvent.Amount, -18).String(),
			)

			if vLog.Topics[0].Hex() == "Transfer" && config.GetConfig().Options.UserAddress == vLog.Topics[2].Hex() {
				// TODO if vLog.Topics[0].Hex() == "Transfer" && evn.address == to
				// to 1
				// bevm-> fhevm 用户 address 1.xbtc (to 2xbtc) 跨链 bevm 转账到 特定地址 调用fhevm mint 接口传递 用户地址及 对应的token 给fhevm  mint
				// fhevm -> bevm 用户 address 1.xbtc (to 2xbtc) 跨链 fhevm 转账到 特定地址 调用bevm mint 接口传递 用户地址及 对应的token 给 bevm mint
				// 交易特定地址 调用fhevm mint 接口传递 发起人 对应的token 给fhevm  mint
				//mint 成功后 这边做 to address 的token 回收 1.xbtc
				transaction, _, err := client.TransactionByHash(context.Background(), vLog.TxHash)
				if err != nil {
					log.Fatal(err)
				}
				amount := transaction.Value()
				log.Println("Transaction Amount: ", amount)
			}
		}
	}
}

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
}

type mintArgs struct {
	_To     common.Address
	_Amount *big.Int
}
