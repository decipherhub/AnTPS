package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v2"
	"log"
	"math/big"
	"os"
	"path/filepath"
)

var ChainID = big.NewInt(8216)
var Host1 = "ws://127.0.0.1:9551"
var Host2 = "ws://127.0.0.1:9552"
var Network = "klay"
var config Config

type Config struct {
	Contracts struct {
		ERC20 struct {
			Address string `yaml:"address"`
		} `yaml:"erc20"`
		ERC721 struct {
			Address string `yaml:"address"`
		} `yaml:"erc721"`
		ERC1155 struct {
			Address string `yaml:"address"`
		} `yaml:"erc1155"`
	} `yaml:"contracts"`
	Condition struct {
		Rate struct {
			Value int `yaml:"value"`
		} `yaml:"rate"`
		Total struct {
			Value int `yaml:"value"`
		} `yaml:"total"`
		GasLimit struct {
			Value uint64 `yaml:"value"`
		} `yaml:"gasLimit"`
	} `yaml:"condition"`
	Multi struct {
		Value int `yaml:"value"`
	}
}

func LoadAddresses(filename string) {
	filePath := filepath.Join(filename)
	content := fmt.Sprintf(`contracts:
  erc20:
    address: "0x0000000000000000000000000000000000000000"
  erc721:
    address: "0x0000000000000000000000000000000000000000"
  erc1155:
    address: "0x0000000000000000000000000000000000000000"
condition:
  rate:
    value: 50
  total:
    value: 500
  gasLimit:
    value: 21000
multi:
  value: 50
`)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		err = os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			log.Printf("failed to create file: %v", err)
		}
	}
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("failed to unmarshal YAML: %v", err)
	}

	ERC20ADDRESS = common.HexToAddress(config.Contracts.ERC20.Address)
	ERC721ADDRESS = common.HexToAddress(config.Contracts.ERC721.Address)
	ERC1155ADDRESS = common.HexToAddress(config.Contracts.ERC1155.Address)
	Rate = config.Condition.Rate.Value
	Total = config.Condition.Total.Value
	GasLimit = config.Condition.GasLimit.Value
	Multi = config.Multi.Value
}
