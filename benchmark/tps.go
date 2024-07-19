package benchmark

import (
	"context"
	"decipher.com/tps/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type blockTPSInfo struct {
	blockDelay           int
	pendingTransaction   int
	confirmedTransaction int
	tps                  uint64
}

func CheckTpsByBlock(total int, filename string) {
	config.WaitSubscribeBlockHead.Add(1)
	defer config.WaitSubscribeBlockHead.Done()

	client, err := ethclient.Dial(config.Host2)
	if err != nil {
		log.Fatal("Fail to Dial", err)
	}
	defer client.Close()
	ctx := context.Background()
	client2 := client.Client()

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(ctx, headers)
	if err != nil {
		log.Fatal("Fail to subscribe", err)
	}
	totalTransactions := 0

	recordAvgTPS := make(map[int]blockTPSInfo)

	startTime := <-config.ChStart
	startConsensusTime := startTime
	blockNumber := 0
	failCount := -1
	sendFinish := false
	writeFile := false
	for {
		select {
		case err = <-sub.Err():
			log.Fatal("Failed to subscribe", err)
		case failCount = <-config.ChFailedCount:
			log.Println("failed to count:", failCount)
			if totalTransactions >= total-failCount {
				if !sendFinish {
					sendFinish = true
					config.ChFinish <- totalTransactions
				}
				if !writeFile {
					writeFile = true
					go StoreDataOnFile(recordAvgTPS, filename)
					log.Printf("max tps = %v\n", config.MaxTPS)
					log.Printf("max delay = %v\n", config.MaxBlockTime)
					log.Printf("total delay = %v\n\n", config.TotalDelay)
				}
			}

		case header := <-headers:
			currentDelay := time.Since(startConsensusTime).Seconds()
			startConsensusTime = time.Now()
			config.TotalDelay = time.Since(startTime).Seconds()
			block, err := client.BlockByNumber(ctx, header.Number)
			if err != nil {
				log.Fatal("getBlock ", err)
			}

			transactions := len(block.Transactions())
			totalTransactions += transactions
			currentTps := float64(transactions) / currentDelay
			tps := float64(totalTransactions) / config.TotalDelay

			client2.CallContext(ctx, &config.Result, "txpool_status")
			pendingTransaction, _ := strconv.ParseInt(config.Result["pending"], 0, 64)
			queuedTransaction, _ := strconv.ParseInt(config.Result["queued"], 0, 64)

			log.Printf("===== block no. %v =====\n", block.Number().Uint64())
			log.Printf("confrimed_transactions:%v\n", transactions)
			log.Printf("total_confirmed_transactions:%v\n", totalTransactions)
			log.Printf("pending_transactions:%v\n", pendingTransaction)
			log.Printf("queued_transactions:%v\n", queuedTransaction)
			log.Printf("block_latency:  %v\n", currentDelay)
			log.Printf("current_tps:%v\n", currentTps)
			log.Printf("total_tps:%v\n\n", tps)

			blockNumber = int(block.NumberU64())
			recordAvgTPS[blockNumber] = blockTPSInfo{int(currentDelay), int(pendingTransaction), transactions, uint64(tps)}

			if tps > config.MaxTPS {
				config.MaxTPS = tps
			}
			if failCount >= 0 {
				// TODO: Should be add Timeout Logic?
				if totalTransactions >= total-failCount {
					log.Printf("max tps = %v\n", config.MaxTPS)
					log.Printf("max delay = %v\n\n", config.MaxBlockTime)
					log.Printf("total delay = %v\n\n", config.TotalDelay)
					if !writeFile {
						writeFile = true
						go StoreDataOnFile(recordAvgTPS, filename)
					}
					if !sendFinish {
						sendFinish = true
						config.ChFinish <- totalTransactions
					}
				}
			}

		case <-config.ChFileWriteFinish:
			log.Println("file write finished")
			return
		}
	}

}

func StoreDataOnFile(data map[int]blockTPSInfo, filename string) {
	file, err := os.Create(filepath.Join(".", "result", filename))
	if err != nil {
		log.Println("file:", err)
		return
	}
	defer file.Close()

	keys := make([]int, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	for _, k := range keys {
		fmt.Fprintf(file, "%d	%d    %d	%d   %d\n", k, data[k].blockDelay, data[k].pendingTransaction, data[k].confirmedTransaction, data[k].tps)
	}
	config.ChFileWriteFinish <- true
}
