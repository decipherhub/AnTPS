package config

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"sync"
	"time"
)

var (
	PrivateKeyHex          []string
	PrivateKey             []*ecdsa.PrivateKey
	Result                 map[string]string
	ChStart                = make(chan time.Time)
	ChFinish               = make(chan int)
	ChFailedCount          = make(chan int)
	ChFileWriteFinish      = make(chan bool)
	WaitSubscribeBlockHead sync.WaitGroup

	ERC20ADDRESS   common.Address
	ERC721ADDRESS  common.Address
	ERC1155ADDRESS common.Address
	Rate           int
	Total          int
	GasLimit       uint64

	OneEther     = big.NewInt(params.Ether)
	Start        time.Time
	Err          error
	Multi        int
	TotalDelay   float64
	MaxTPS       float64
	MaxBlockTime = 0
	NonceMutex   sync.Mutex
	LastNonce    uint64
)
