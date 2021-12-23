package model

import (
	"github.com/ethereum/go-ethereum/common"
)

type Account struct {
	name    string
	alias   string
	address common.Address
}

func NewAccount(name, alias string, address string) *Account {
	return &Account{
		name:    name,
		alias:   alias,
		address: common.HexToAddress(address),
	}
}

func (m *Account) Name() string {
	return m.name
}

func (m *Account) Alias() string {
	return m.alias
}

func (m *Account) Address() common.Address {
	return m.address
}
