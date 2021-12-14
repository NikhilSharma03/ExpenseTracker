package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/NikhilSharma03/expensetracker/server/expensepb"
	"github.com/spf13/cobra"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "display the current available balance",
	Long: `The balance subcommand is used to display the current available balance.

Example:
	expensetracker balance`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := expenseClient.GetBalance(context.Background(), &expensepb.Empty{})
		if err != nil {
			log.Fatal("Something went wrong", err.Error())
		}
		fmt.Println("Current Balance:", res.GetBalance())
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)
}
