package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/NikhilSharma03/expensetracker/server/expensepb"
	"github.com/jedib0t/go-pretty/table"
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
		stream, err := expenseClient.GetExpenseHistory(context.Background(), &expensepb.Empty{})
		if err != nil {
			log.Fatal("Something went wrong", err.Error())
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Amount", "Type"})
		i := 1
		for {
			val, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("Something went wrong", err.Error())
			}
			amount := strings.Trim(fmt.Sprintf("%v", val.GetAmount()), "-")
			t.AppendRow([]interface{}{i, amount, strings.ToUpper(val.GetType())})
			i += 1
		}
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
