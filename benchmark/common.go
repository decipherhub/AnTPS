package benchmark

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"decipher.com/tps/config"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func UpdateConfig(config string) {
	filePath := "config/config.go"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)

	if config == "ava" {
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "var ChainID") {
				line = fmt.Sprintf("var ChainID = big.NewInt(%s)", big.NewInt(43112))
			} else if strings.Contains(line, "var Host1") {
				line = fmt.Sprintf("var Host1 = \"%s\"", "ws://127.0.0.1:9650/ext/bc/C/ws")
			} else if strings.Contains(line, "var Host2") {
				line = fmt.Sprintf("var Host2 = \"%s\"", "ws://127.0.0.1:9651/ext/bc/C/ws")
			} else if strings.Contains(line, "var Network") {
				line = fmt.Sprintf("var Network = \"%s\"", config)
			}
			lines = append(lines, line)
		}
	} else if config == "klay" {
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "var ChainID") {
				line = fmt.Sprintf("var ChainID = big.NewInt(%s)", big.NewInt(8216))
			} else if strings.Contains(line, "var Host1") {
				line = fmt.Sprintf("var Host1 = \"%s\"", "ws://127.0.0.1:9551")
			} else if strings.Contains(line, "var Host2") {
				line = fmt.Sprintf("var Host2 = \"%s\"", "ws://127.0.0.1:9551")
			} else if strings.Contains(line, "var Network") {
				line = fmt.Sprintf("var Network = \"%s\"", config)
			}
			lines = append(lines, line)
		}
	} else if config == "eth" {
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "var ChainID") {
				line = fmt.Sprintf("var ChainID = big.NewInt(%s)", big.NewInt(32382))
			} else if strings.Contains(line, "var Host1") {
				line = fmt.Sprintf("var Host1 = \"%s\"", "ws://127.0.0.1:8546")
			} else if strings.Contains(line, "var Host2") {
				line = fmt.Sprintf("var Host2 = \"%s\"", "ws://127.0.0.1:9546")
			} else if strings.Contains(line, "var Network") {
				line = fmt.Sprintf("var Network = \"%s\"", config)
			}
			lines = append(lines, line)
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("Failed to scan file: %s", err)
	}

	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed to open file for writing: %s", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
}

func UpdateAddress(ERC20, ERC721, ERC1155 common.Address) {
	filePath := filepath.Join("config", "config.yml")
	existingContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("failed to read file: %v", err)
	}

	updatedContent := updateAddresses(string(existingContent), ERC20, ERC721, ERC1155)
	err = os.WriteFile(filePath, []byte(updatedContent), 0644)
	if err != nil {
		log.Printf("failed to write file: %v", err)
	}
}

func updateAddresses(content string, ERC20, ERC721, ERC1155 common.Address) string {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		switch {
		case strings.Contains(line, "erc20:"):
			lines[i+1] = fmt.Sprintf("    address: \"%s\"", ERC20.Hex())
		case strings.Contains(line, "erc721:"):
			lines[i+1] = fmt.Sprintf("    address: \"%s\"", ERC721.Hex())
		case strings.Contains(line, "erc1155:"):
			lines[i+1] = fmt.Sprintf("    address: \"%s\"", ERC1155.Hex())
		}
	}
	return strings.Join(lines, "\n")
}

func InitAccount(count int) {
	file, err := os.Open("./account/privateKey_100k")
	if err != nil {
		log.Fatalf("Failed to open account file: %v", err)
	}
	if config.Err != nil {
		log.Fatalf("Failed to load config: %v", config.Err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "\"", "", -1)
		line = strings.Replace(line, ",", "", -1)
		line = strings.Replace(line, "0x", "", -1)

		config.PrivateKeyHex = append(config.PrivateKeyHex, line)
		if len(config.PrivateKeyHex) >= count {
			break
		}
	}

	for _, hex := range config.PrivateKeyHex {
		key, _ := crypto.HexToECDSA(hex)
		config.PrivateKey = append(config.PrivateKey, key)
	}

	for i := 0; i < len(config.PrivateKey); i++ {
		config.PrivateKey[i], _ = crypto.HexToECDSA(config.PrivateKeyHex[i])
	}
}

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

func InitContract() (common.Address, common.Address, common.Address) {
	client, err := ethclient.Dial(config.Host1)
	if err != nil {
		log.Fatalf("Error connect to client: %v", err)
	}
	defer client.Close()
	ERC20, ERC721, ERC1155 := DeployContract(client, config.PrivateKey[0])

	return ERC20, ERC721, ERC1155
}

func updateNonce(client *ethclient.Client, owner common.Address, chain *bind.TransactOpts) {
	currentNonce, err := client.PendingNonceAt(context.Background(), owner)
	if err != nil {
		log.Fatalf("Failed to retrieve current nonce: %v", err)
	}
	chain.Nonce = big.NewInt(int64(currentNonce))
}

func initialize(client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*ethclient.Client, *bind.TransactOpts, common.Address) {
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	Trader := crypto.PubkeyToAddress(*publicKeyECDSA)

	chain, err := bind.NewKeyedTransactorWithChainID(privateKey, config.ChainID)
	if err != nil {
		log.Fatalf("NewKeyedTransactorWithChainID: %s", err)
	}
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	chain.GasPrice = gasPrice
	chain.GasLimit = config.GasLimit

	nonce, err := client.PendingNonceAt(context.Background(), Trader)
	if err != nil {
		log.Fatalf("Nonce: %v", err)
	}
	chain.Nonce = big.NewInt(int64(nonce))

	return client, chain, Trader
}

func waitMined(ctx context.Context, client *ethclient.Client, tx *types.Transaction, msg ...string) bool {
	queryTicker := time.NewTicker(time.Millisecond * 100)
	defer queryTicker.Stop()

	count := 0
	for {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			if receipt.Status == 0 {
				log.Println("failed transaction:", msg, receipt)
				return false
			}
			return true
		}
		select {
		case <-ctx.Done():
			log.Println("failed transaction error:", ctx.Err())
			return false
		case <-queryTicker.C:
			count++
			if count >= 600 {
				var result map[string]string
				err = client.Client().CallContext(ctx, &result, "txpool_status")
				if err != nil {
					return false
				}
				pendingTransaction, _ := strconv.ParseInt(result["pending"], 0, 64)
				if pendingTransaction == 0 {
					log.Println("Fail to mine:", count, tx.Hash())
					return false
				}
			}
		}
	}
}
