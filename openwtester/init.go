package openwtester

import (
	"github.com/Assetsadapter/ecoball-adapter/eco"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(eco.Symbol, eco.NewWalletManager())
}
