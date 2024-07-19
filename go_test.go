package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"testing"
	"time"

	"decipher.com/tps/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

var (
	gasPrice     = big.NewInt(int64(0))
	gasLimit     = uint64(3_500_000)
	chainId      = big.NewInt(43112)
	oneEther     = big.NewInt(params.Ether)
	client       *ethclient.Client
	chain        *bind.TransactOpts
	owner        common.Address
	ERC20Address common.Address
)

func GetKeyAndAddress(pk string) (*ecdsa.PrivateKey, common.Address) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatalf("private key hex to ECDSA conversion failed: %v", err)
	}

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	return privateKey, crypto.PubkeyToAddress(*publicKeyECDSA)
}

func initialize() (*ethclient.Client, *bind.TransactOpts, common.Address) {
	pk, user1 := GetKeyAndAddress("2e0834786285daccd064ca17f1654f67b4aef298acbb82cef9ec422fb4975622")

	chain, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		log.Fatalf("key: %s", err)
	}
	gasPrice, _ = client.SuggestGasPrice(context.Background())
	chain.GasPrice = gasPrice
	chain.GasLimit = gasLimit

	nonce, err := client.PendingNonceAt(context.Background(), owner)
	if err != nil {
		log.Fatalf("nonce: %v", err)
	}
	chain.Nonce = big.NewInt(int64(nonce))

	return client, chain, user1
}

func WaitMined(ctx context.Context, client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Millisecond * 100)
	defer queryTicker.Stop()

	count := 0
	for {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			return receipt, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
			count++
			log.Printf("Waiting for transaction %s to be confirmed... (%d/600)", tx.Hash().Hex(), count)
			if count >= 600 {
				return nil, fmt.Errorf("transaction timed out: %v", tx.Hash().Hex())
			}
		}
	}
}

func checkWaitMined(ctx context.Context, client *ethclient.Client, tx *types.Transaction, msg ...string) bool {
	receipt, err := WaitMined(ctx, client, tx)
	if err != nil {
		log.Println("failed transaction error:", err)
		return false
	}
	if receipt.Status == 0 {
		log.Println("failed transaction:", msg, receipt)
		return false
	}
	return true
}

func waitForNodeToBeReady(client *ethclient.Client) error {
	ctx := context.Background()
	for {
		_, err := client.BlockByNumber(ctx, nil)
		if err == nil {
			log.Println("Node is ready.")
			break
		}
		if strings.Contains(err.Error(), "chain is not done bootstrapping") {
			log.Println("Node is bootstrapping. Waiting...")
			time.Sleep(5 * time.Second)
			continue
		}
		return err
	}
	return nil
}

func TestDeployContract(t *testing.T) {
	var err error
	client, err = ethclient.Dial("http://127.0.0.1:9650/ext/bc/C/rpc")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	log.Println("Waiting for node to be ready...")
	err = waitForNodeToBeReady(client)
	if err != nil {
		t.Fatalf("Node is not ready: %v", err)
	}

	client, chain, owner = initialize()
	ERC20Address, tx, _, err := abi.DeployERC20(chain, client, oneEther)
	if err != nil {
		t.Fatalf("failed to deploy %v", err)
	}
	t.Logf("ERC20 Contract address: %s", ERC20Address)
	t.Logf("Transaction hash: %s", tx.Hash().Hex())

	ctx := context.Background()
	if !checkWaitMined(ctx, client, tx) {
		t.Fatalf("failed to mine transaction")
	}
	receipt, err := client.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		t.Fatalf("failed to get transaction receipt: %v", err)
	}
	if receipt.Status == 1 {
		t.Logf("Transaction status: success")
	} else {
		t.Logf("Transaction status: failed")
	}
}

func updateNonce(client *ethclient.Client, chain *bind.TransactOpts, address common.Address) error {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return err
	}
	chain.Nonce = big.NewInt(int64(nonce))
	return nil
}

func TestMintAndTransfer(t *testing.T) {
	var err error
	client, err = ethclient.Dial("http://127.0.0.1:9650/ext/bc/C/rpc")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	log.Println("Waiting for node to be ready...")
	err = waitForNodeToBeReady(client)
	if err != nil {
		t.Fatalf("Node is not ready: %v", err)
	}

	client, chain, owner = initialize()

	erc20, err := abi.NewERC20(ERC20Address, client)
	if err != nil {
		t.Fatalf("failed to instantiate ERC20 contract: %v", err)
	}

	ctx := context.Background()

	mintAmount := new(big.Int).Mul(oneEther, big.NewInt(2000000000))
	err = updateNonce(client, chain, owner)
	if err != nil {
		t.Fatalf("failed to update nonce: %v", err)
	}
	tx, err := erc20.Mint(chain, owner, mintAmount)
	if err != nil {
		t.Fatalf("failed to mint tokens: %v", err)
	}
	log.Printf("Minting transaction hash: %s", tx.Hash().Hex())
	if !checkWaitMined(ctx, client, tx, "minting") {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			log.Printf("Failed Minting transaction receipt: %+v\n", receipt)
		}
		t.Fatalf("failed to mint tokens")
	}

	toAddress := common.HexToAddress("0x25dBeC20C5d60f405F4daA2B6008e03eC1ec6095")
	err = updateNonce(client, chain, owner)
	if err != nil {
		t.Fatalf("failed to update nonce: %v", err)
	}
	tx, err = erc20.Transfer(chain, toAddress, mintAmount.Div(mintAmount, big.NewInt(2)))
	if err != nil {
		t.Fatalf("failed to transfer tokens: %v", err)
	}
	log.Printf("Transfer transaction hash: %s", tx.Hash().Hex())
	if !checkWaitMined(ctx, client, tx, "transfer") {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			log.Printf("Failed Transfer transaction receipt: %+v\n", receipt)
		}
		t.Fatalf("failed to transfer tokens")
	}

	t.Logf("Mint and transfer test completed successfully")
}
