package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "displays the history of transactions",
	Long: `The history subcommand is used to display the history of transactions.

Example:
	expensetracker history
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("history called")
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
