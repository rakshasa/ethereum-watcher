package rpc

import (
	"github.com/rakshasa/ethwatcher/blockchain"
)

type IBlockChainRPC interface {
	GetCurrentBlockNum() (uint64, error)

	GetBlockByNum(uint64) (blockchain.Block, error)
	GetLiteBlockByNum(uint64) (blockchain.Block, error)
	GetTransactionReceipt(txHash string) (blockchain.TransactionReceipt, error)

	GetLogs(from, to uint64, address []string, topics []string) ([]blockchain.IReceiptLog, error)
}
