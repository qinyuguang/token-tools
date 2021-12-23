package service

import (
	"errors"

	"github.com/gogf/gf/util/gconv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"token-tools/internal/model"
)

type Account interface {
	Account() *model.Account
}

type account struct {
	account *model.Account
}

func NewAccount(cmd *cobra.Command) (*account, error) {
	alias, err := cmd.Flags().GetString("account")
	if err != nil {
		return nil, err
	}

	var configs []*model.AccountConfig
	if err := gconv.Structs(viper.Get("accounts"), &configs); err != nil {
		return nil, err
	}

	for i := range configs {
		if configs[i].Alias == alias {
			return &account{
				account: model.NewAccount(configs[i].Name, alias, configs[i].Address),
			}, nil
		}
	}

	return nil, errors.New("not found the account in config yaml")
}

func (s *account) Account() *model.Account {
	return s.account
}
