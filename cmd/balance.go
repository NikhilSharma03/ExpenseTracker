package cmd

import (
	"fmt"

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
		fmt.Println("balance called")
	},
}

func init() {
	rootCmd.AddCommand(balanceCmd)
}
