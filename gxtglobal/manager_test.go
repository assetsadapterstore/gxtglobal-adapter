/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package gxtglobal

import (
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/log"
	"github.com/shopspring/decimal"
	"path/filepath"
	"testing"
)

var (
	tw *WalletManager
)

func init() {

	tw = testNewWalletManager()
}

func testNewWalletManager() *WalletManager {
	wm := NewWalletManager()

	//读取配置
	absFile := filepath.Join("conf", "conf.ini")
	//log.Debug("absFile:", absFile)
	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return nil
	}
	wm.LoadAssetsConfig(c)
	//wm.ExplorerClient.Debug = false
	wm.WalletClient.Debug = true
	return wm
}

func TestGetCoreWalletinfo(t *testing.T) {
	tw.GetCoreWalletinfo()
}

func TestGetBlockChainInfo(t *testing.T) {
	b, err := tw.GetBlockChainInfo()
	if err != nil {
		t.Errorf("GetBlockChainInfo failed unexpected error: %v\n", err)
	} else {
		t.Logf("GetBlockChainInfo info: %v\n", b)
	}
}

func TestListUnspent(t *testing.T) {
	utxos, err := tw.ListUnspent(0, "vLZxJVV3wajNzqJUP3qxRVEDitdWiqEbr4")
	if err != nil {
		t.Errorf("ListUnspent failed unexpected error: %v\n", err)
		return
	}
	totalBalance := decimal.Zero
	for _, u := range utxos {
		t.Logf("ListUnspent %s: %s = %s\n", u.Address, u.AccountID, u.Amount)
		amount, _ := decimal.NewFromString(u.Amount)
		totalBalance = totalBalance.Add(amount)
	}

	t.Logf("totalBalance: %s \n", totalBalance.String())
}

func TestEstimateFee(t *testing.T) {
	feeRate, _ := tw.EstimateFeeRate()
	t.Logf("EstimateFee feeRate = %s\n", feeRate.StringFixed(8))
	fees, _ := tw.EstimateFee(10, 2, feeRate)
	t.Logf("EstimateFee fees = %s\n", fees.StringFixed(8))
}

func TestGetNetworkInfo(t *testing.T) {
	tw.GetNetworkInfo()
}

func TestPrintConfig(t *testing.T) {
	tw.Config.PrintConfig()
}

func TestWalletManager_ImportAddress(t *testing.T) {
	addr := "H9wJq1HLkY3fCRJQa7JXXRdQ4zkbDzpq5n"
	err := tw.ImportAddress(addr, "")
	if err != nil {
		t.Errorf("RestoreWallet failed unexpected error: %v\n", err)
		return
	}
	log.Info("imported success")
}

func TestWalletManager_ListAddresses(t *testing.T) {
	addresses, err := tw.ListAddresses()
	if err != nil {
		t.Errorf("GetAddressesByAccount failed unexpected error: %v\n", err)
		return
	}

	for i, a := range addresses {
		t.Logf("ListAddresses address[%d] = %s\n", i, a)
	}
}

func TestWalletManager_GetInfo(t *testing.T) {
	tw.GetInfo()
}
