package ethwatcher

import (
	"context"
	"fmt"
	"testing"

	"github.com/rakshasa/ethwatcher/plugin"
	"github.com/rakshasa/ethwatcher/structs"
	"github.com/rakshasa/ethwatcher/utils"
	"github.com/sirupsen/logrus"
)

func TestReceiptLogsPlugin(t *testing.T) {
	utils.SetCategoryLogLevel(logrus.DebugLevel)

	api := "https://kovan.infura.io/v3/19d753b2600445e292d54b1ef58d4df4"
	w := NewHttpBasedEthWatcher(api)

	contract := "0x63bB8a255a8c045122EFf28B3093Cc225B711F6D"
	// Match
	topics := []string{"0x6bf96fcc2cec9e08b082506ebbc10114578a497ff1ea436628ba8996b750677c"}

	w.RegisterReceiptLogPlugin(plugin.NewReceiptLogPlugin(contract, topics, func(receipt *structs.RemovableReceiptLog) {
		if receipt.IsRemoved {
			utils.Infof("Removed >> %+v", receipt)
		} else {
			utils.Infof("Adding >> %+v, tx: %s, logIdx: %d", receipt, receipt.IReceiptLog.GetTransactionHash(), receipt.IReceiptLog.GetLogIndex())
		}
	}))

	//startBlock := 12304546
	startBlock := 12101723
	err := w.RunTillExitFromBlock(context.Background(), uint64(startBlock))

	fmt.Println("err:", err)
}
