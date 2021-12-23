package service

import (
	"errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/util/gconv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"token-tools/internal/model"
	"token-tools/pkg/erc20"
)

type Token interface {
	BalanceOf(account Account) (uint, error)
	Decimals() uint8
}

type token struct {
	token    *model.Token
	client   *ethclient.Client
	instance *erc20.Token
}

func NewToken(cmd *cobra.Command) (*token, error) {
	symbol, err := cmd.Flags().GetString("token")
	if err != nil {
		return nil, err
	}

	var configs []*model.TokenConfig
	if err := gconv.Structs(viper.Get("tokens"), &configs); err != nil {
		return nil, err
	}

	for i := range configs {
		if configs[i].Symbol == symbol {
			t := model.NewToken(symbol, configs[i].Address, configs[i].Network)

			client, err := ethclient.Dial(configs[i].Network)
			if err != nil {
				return nil, err
			}

			instance, err := erc20.NewToken(t.Address(), client)
			if err != nil {
				return nil, err
			}

			return &token{
				token:    t,
				client:   client,
				instance: instance,
			}, nil
		}
	}

	return nil, errors.New("not found the token in config yaml")
}

func (s *token) Token() *model.Token {
	return s.token
}

func (s *token) BalanceOf(account Account) (*big.Float, error) {
	balance, err := s.instance.BalanceOf(&bind.CallOpts{}, account.Account().Address())
	if err != nil {
		return nil, err
	}

	if balance == nil {
		return nil, errors.New("balance is nil")
	}

	return new(big.Float).Quo(
		new(big.Float).SetInt(balance),
		big.NewFloat(math.Pow10(int(s.Decimals()))),
	), nil
}

func (s *token) Decimals() uint8 {
	decimals, err := s.instance.Decimals(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return decimals
}
