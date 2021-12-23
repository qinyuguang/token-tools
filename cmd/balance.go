/*
Copyright © 2021 Qin Yuguang <qinyuguang@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"token-tools/internal/service"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "query account balance",
	RunE: func(cmd *cobra.Command, args []string) error {
		token, err := service.NewToken(cmd)
		if err != nil {
			return err
		}

		account, err := service.NewAccount(cmd)
		if err != nil {
			return err
		}

		balance, err := token.BalanceOf(account)
		if err != nil {
			return err
		}

		fmt.Printf("%s balance of %s is %s\n",
			token.Token().Symbol(),
			account.Account().Name(),
			balance.String(),
		)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	balanceCmd.Flags().StringP("account", "a", "", "account alias")
	_ = balanceCmd.MarkFlagRequired("account")
}
