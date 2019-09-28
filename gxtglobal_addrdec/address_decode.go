package gxtglobal_addrdec

import (
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/openwallet/openwallet"
)

var (
	alphabet = addressEncoder.ZECAlphabet
)

var (

	GXT_mainnetAddressP2PKH         = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x84}, Suffix: nil}
	GXT_testnetAddressP2PKH         = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x89}, Suffix: nil}
	GXT_mainnetPrivateWIFCompressed = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "", HashLen: 32, Prefix: []byte{0x37}, Suffix: []byte{0x01}}
	GXT_testnetPrivateWIFCompressed = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "", HashLen: 32, Prefix: []byte{0xEF}, Suffix: []byte{0x01}}
	GXT_mainnetAddressP2SH          = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x40}, Suffix: nil}
	GXT_testnetAddressP2SH          = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x13}, Suffix: nil}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	*openwallet.AddressDecoderV2Base
	IsTestNet bool
}

//AddressDecode 地址解析
func (dec *AddressDecoderV2) AddressDecode(addr string, opts ...interface{}) ([]byte, error) {

	cfg := GXT_mainnetAddressP2PKH
	if dec.IsTestNet {
		cfg = GXT_testnetAddressP2PKH
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	return addressEncoder.AddressDecode(addr, cfg)
}

//AddressEncode 地址编码
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {

	cfg := GXT_mainnetAddressP2PKH
	if dec.IsTestNet {
		cfg = GXT_testnetAddressP2PKH
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	address := addressEncoder.AddressEncode(hash, cfg)
	return address, nil
}

// AddressVerify 地址校验
func (dec *AddressDecoderV2) AddressVerify(address string, opts ...interface{}) bool {

	return true
}