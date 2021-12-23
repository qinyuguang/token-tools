package model

import (
	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	symbol  string
	address common.Address
	network string
}

func NewToken(symbol, address, network string) *Token {
	return &Token{
		symbol:  symbol,
		address: common.HexToAddress(address),
		network: network,
	}
}

func (m *Token) Symbol() string {
	return m.symbol
}

func (m *Token) Address() common.Address {
	return m.address
}
