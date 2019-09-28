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
	"github.com/assetsadapterstore/gxtglobal-adapter/gxtTransaction"
	"github.com/blocktree/go-owcrypt"
)

const (
	//币种
	Symbol    = "GXT"
	CurveType = owcrypt.ECC_CURVE_SECP256K1
	Decimals  = int32(8)
)

var (
	MainNetAddressPrefix = gxtTransaction.AddressPrefix{P2PKHPrefix: []byte{0x84}, P2WPKHPrefix: []byte{0x40}, P2SHPrefix: nil, Bech32Prefix: "hss"}
	TestNetAddressPrefix = gxtTransaction.AddressPrefix{P2PKHPrefix: []byte{0x89}, P2WPKHPrefix: []byte{0x13}, P2SHPrefix: nil, Bech32Prefix: "hss"}
)