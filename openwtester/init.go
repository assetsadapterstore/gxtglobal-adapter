package openwtester

import (
	"github.com/assetsadapterstore/gxtglobal-adapter/gxtglobal"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(gxtglobal.Symbol, gxtglobal.NewWalletManager())
}
