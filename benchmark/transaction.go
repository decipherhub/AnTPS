package benchmark

import (
	"context"
	"crypto/ecdsa"
	"decipher.com/tps/abi"
	"decipher.com/tps/config"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BenchmarkContext struct {
	Client          *ethclient.Client
	Chain           *bind.TransactOpts
	Owner           common.Address
	ContractAddress common.Address
	Total           int
	SendRate        int
	Wait            sync.WaitGroup
	FailCount       int
	FailCountMutex  *sync.Mutex
	NonceMutex      *sync.Mutex
	MaxElapsed      float64
	TotalElapsed    float64
	TotalMutex      *sync.Mutex
	Ctx             context.Context
}

func initializeBenchmark(total int, sendRate int, operationType string, contractAddress common.Address) (*BenchmarkContext, string) {
	client, err := ethclient.Dial(config.Host1)
	if err != nil {
		log.Fatalf("client: %v", err)
	}

	filename := fmt.Sprintf("%v.%v.%v.%v.%v.txt", config.Network, time.Now().Format("20060102_150405"), total, sendRate, operationType)
	go CheckTpsByBlock(total, filename)
	config.ChStart <- time.Now()

	_, chain, owner := initialize(client, config.PrivateKey[0])

	return &BenchmarkContext{
		Client:          client,
		Chain:           chain,
		Owner:           owner,
		ContractAddress: contractAddress,
		Total:           total,
		SendRate:        sendRate,
		FailCountMutex:  new(sync.Mutex),
		NonceMutex:      new(sync.Mutex),
		TotalMutex:      new(sync.Mutex),
		Ctx:             context.Background(),
	}, filename
}

func (bc *BenchmarkContext) Benchmark(txFunc func(int) (*types.Transaction, error)) {
	for i := 1; i <= bc.Total; i++ {
		bc.Wait.Add(1)
		go func(id int) {
			defer bc.Wait.Done()
			for {
				bc.NonceMutex.Lock()
				updateNonce(bc.Client, bc.Owner, bc.Chain)
				bc.NonceMutex.Unlock()
				config.Start = time.Now()
				tx, err := txFunc(id)
				if err != nil {
					/*
						Have to handle the Nonce Error case
					*/
					if err.Error() == "replacement transaction underpriced" || err.Error() == "already known" || err.Error() == "there is another tx which has the same nonce in the tx pool" {
						updateNonce(bc.Client, bc.Owner, bc.Chain)
						time.Sleep(time.Second)
						continue
					} else {
						bc.FailCountMutex.Lock()
						bc.FailCount++
						bc.FailCountMutex.Unlock()
						return
					}
				}
				if !waitMined(bc.Ctx, bc.Client, tx) {
					bc.FailCountMutex.Lock()
					bc.FailCount++
					bc.FailCountMutex.Unlock()
					return
				}
				elapsed := time.Since(config.Start).Seconds()
				bc.TotalMutex.Lock()
				bc.TotalElapsed += elapsed
				if elapsed > bc.MaxElapsed {
					bc.MaxElapsed = elapsed
				}
				bc.TotalMutex.Unlock()
				break
			}
		}(i)

		if i%bc.SendRate == 0 {
			log.Println("send ", i)
			time.Sleep(time.Second)
		}
	}
	bc.Wait.Wait()
	log.Println("max latency", bc.MaxElapsed)
	config.ChFailedCount <- bc.FailCount
	total := <-config.ChFinish
	log.Println("avg latency:", bc.TotalElapsed/float64(total))
}

func DeployContract(client *ethclient.Client, privateKey *ecdsa.PrivateKey) (common.Address, common.Address, common.Address) {
	client, chain, owner := initialize(client, privateKey)
	totalSupply := new(big.Int).Mul(config.OneEther, big.NewInt(1000000000))
	ERC20Address, _, _, err := abi.DeployERC20(chain, client, totalSupply)
	updateNonce(client, owner, chain)
	ERC721Address, _, _, err := abi.DeployERC721(chain, client)
	updateNonce(client, owner, chain)
	ERC1155Address, _, _, err := abi.DeployERC1155(chain, client)
	if err != nil {
		log.Fatalln("failed to deploy", err)
	}
	log.Printf("ERC20 Contract address: %s", ERC20Address)
	log.Printf("ERC721 Contract address: %s", ERC721Address)
	log.Printf("ERC1155 Contract address: %s", ERC1155Address)
	return ERC20Address, ERC721Address, ERC1155Address
}

func ERC20Mint(total int, sendRate int, contractAddress common.Address) {
	bc, _ := initializeBenchmark(total, sendRate, "mint_erc20", contractAddress)
	token, _ := abi.NewERC20(contractAddress, bc.Client)
	mintAmount := new(big.Int).Mul(config.OneEther, big.NewInt(1))

	txFunc := func(id int) (*types.Transaction, error) {
		_, toAddress := GetKeyAndAddress(config.PrivateKeyHex[id-1])
		return token.Mint(bc.Chain, toAddress, mintAmount)
	}

	bc.Benchmark(txFunc)
	config.WaitSubscribeBlockHead.Wait()
}

func ERC20Transfer(total int, sendRate int, contractAddress common.Address) {
	bc, _ := initializeBenchmark(total, sendRate, "transfer_erc20", contractAddress)
	token, _ := abi.NewERC20(contractAddress, bc.Client)
	Amount := new(big.Int).Mul(config.OneEther, big.NewInt(1))

	txFunc := func(id int) (*types.Transaction, error) {
		_, toAddress := GetKeyAndAddress(config.PrivateKeyHex[id-1])
		return token.Transfer(bc.Chain, toAddress, Amount)
	}

	bc.Benchmark(txFunc)
	config.WaitSubscribeBlockHead.Wait()
}

func ERC721Mint(total int, sendRate int, contractAddress common.Address) {
	bc, _ := initializeBenchmark(total, sendRate, "mint_erc721", contractAddress)
	token, _ := abi.NewERC721(contractAddress, bc.Client)

	txFunc := func(id int) (*types.Transaction, error) {
		return token.Mint(bc.Chain, bc.Owner)
	}

	bc.Benchmark(txFunc)
	config.WaitSubscribeBlockHead.Wait()
}

func ERC721Transfer(total int, sendRate int, contractAddress common.Address) {
	bc, _ := initializeBenchmark(total, sendRate, "transfer_erc721", contractAddress)
	token, _ := abi.NewERC721(contractAddress, bc.Client)

	txFunc := func(id int) (*types.Transaction, error) {
		_, toAddress := GetKeyAndAddress(config.PrivateKeyHex[id-1])
		return token.TransferFrom(bc.Chain, bc.Owner, toAddress, big.NewInt(int64(id)))
	}

	bc.Benchmark(txFunc)
	config.WaitSubscribeBlockHead.Wait()
}

func ERC1155Mint(total int, sendRate int, contractAddress common.Address) {
	bc, _ := initializeBenchmark(total, sendRate, "mint_erc1155", contractAddress)
	token, _ := abi.NewERC1155(contractAddress, bc.Client)
	Amount := new(big.Int).Mul(config.OneEther, big.NewInt(1))

	txFunc := func(id int) (*types.Transaction, error) {
		return token.Mint(bc.Chain, bc.Owner, Amount)
	}

	bc.Benchmark(txFunc)
	config.WaitSubscribeBlockHead.Wait()
}

func ERC1155Transfer(total int, sendRate int, contractAddress common.Address) {
	bc, _ := initializeBenchmark(total, sendRate, "transfer_erc1155", contractAddress)
	token, _ := abi.NewERC1155(contractAddress, bc.Client)
	Amount := new(big.Int).Mul(config.OneEther, big.NewInt(1))

	txFunc := func(id int) (*types.Transaction, error) {
		_, toAddress := GetKeyAndAddress(config.PrivateKeyHex[id-1])
		return token.SafeTransferFrom(bc.Chain, bc.Owner, toAddress, big.NewInt(int64(id)), Amount, nil)
	}

	bc.Benchmark(txFunc)
	config.WaitSubscribeBlockHead.Wait()
}

func NativeTransfer(total int, sendRate int) {
	bc, _ := initializeBenchmark(total, sendRate, "transfer_native", common.Address{})
	transferAmount := new(big.Int).Mul(config.OneEther, big.NewInt(1))

	txFunc := func(id int) (*types.Transaction, error) {
		_, toAddress := GetKeyAndAddress(config.PrivateKeyHex[id-1])

		nonce, _ := bc.Client.PendingNonceAt(bc.Ctx, bc.Owner)
		gasPrice, _ := bc.Client.SuggestGasPrice(bc.Ctx)

		tx := types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			To:       &toAddress,
			Value:    transferAmount,
			Gas:      config.GasLimit,
			GasPrice: gasPrice,
		})

		signedTx, err := bc.Chain.Signer(bc.Owner, tx)
		if err != nil {
			return nil, err
		}

		return signedTx, bc.Client.SendTransaction(bc.Ctx, signedTx)
	}

	bc.Benchmark(txFunc)
	config.WaitSubscribeBlockHead.Wait()
}

func MultiTransfer(total int) {
	client, err := ethclient.Dial(config.Host1)
	if err != nil {
		log.Fatalf("client: %v", err)
	}
	privateKeys := config.PrivateKey[:config.Multi]
	filename := fmt.Sprintf("%v.%v.%v.%v.txt", config.Network, time.Now().Format("20060102_150405"), total, "transfer_multi")
	go CheckTpsByBlock(total, filename)
	config.ChStart <- time.Now()

	transferAmount := new(big.Int).SetInt64(100000000000000000) // 0.1ETH
	var Wait sync.WaitGroup
	failCount := 0
	failCountMutex := new(sync.Mutex)
	nonceMutex := new(sync.Mutex)
	maxElapsed := 0.0
	totalElapsed := 0.0
	totalMutex := new(sync.Mutex)
	ctx := context.Background()
	txsPerAccount := total / len(privateKeys)

	for i, privateKey := range privateKeys {
		Wait.Add(1)
		go func(pk *ecdsa.PrivateKey, id int) {
			defer Wait.Done()
			_, chain, owner := initialize(client, pk)
			_, toAddress := GetKeyAndAddress(config.PrivateKeyHex[id])

			for j := 0; j < txsPerAccount; j++ {
				nonceMutex.Lock()
				updateNonce(client, owner, chain)
				nonce, _ := client.PendingNonceAt(ctx, owner)
				nonceMutex.Unlock()

				gasPrice, _ := client.SuggestGasPrice(ctx)
				nativeTx := types.NewTx(&types.LegacyTx{
					Nonce:    nonce,
					To:       &toAddress,
					Value:    transferAmount,
					Gas:      config.GasLimit,
					GasPrice: gasPrice,
				})

				signedTx, _ := types.SignTx(nativeTx, types.NewEIP155Signer(config.ChainID), pk)
				config.Err = client.SendTransaction(ctx, signedTx)
				if config.Err != nil {
					if config.Err.Error() == "replacement transaction underpriced" {
						time.Sleep(time.Second)
						updateNonce(client, owner, chain)
						continue
					} else {
						failCountMutex.Lock()
						failCount++
						failCountMutex.Unlock()
						return
					}
				}
				config.Start = time.Now()
				if !waitMined(ctx, client, signedTx) {
					failCountMutex.Lock()
					failCount++
					failCountMutex.Unlock()
					return
				}
				elapsed := time.Since(config.Start).Seconds()
				totalMutex.Lock()
				totalElapsed += elapsed
				totalMutex.Unlock()
				if elapsed > maxElapsed {
					maxElapsed = elapsed
				}
			}
		}(privateKey, i)
	}
	Wait.Wait()
	log.Println("max latency", maxElapsed)
	config.ChFailedCount <- failCount
	total2 := <-config.ChFinish
	log.Println("avg latency:", totalElapsed/float64(total2))
	config.WaitSubscribeBlockHead.Wait()
}
