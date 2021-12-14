package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/NikhilSharma03/expensetracker/server/expensepb"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "credit/debit money into expenses",
	Long: `The new command is used to credit/debit money into expenses
	
To credit money into expenses, add --credit or -C flag
Example:
	expensetracker new --credit 1000

To debit money from expenses, add --debit or -D flag
Example:
	expensetracker new --debit 1000

`,
	Run: func(cmd *cobra.Command, args []string) {
		credit, _ := cmd.Flags().GetString("credit")
		debit, _ := cmd.Flags().GetString("debit")

		if credit != "" && debit != "" {
			fmt.Println("Please use one flag at one time.")
			return
		}

		if credit != "" {
			cfVal, erro := strconv.ParseFloat(credit, 64)
			if erro != nil {
				log.Fatal("Something went wrong", erro.Error())
			}

			res, err := expenseClient.AddExpense(context.Background(), &expensepb.Transaction{Type: "credit", Amount: cfVal})
			if err != nil {
				log.Fatal("Something went wrong", err.Error())
			}
			fmt.Println("Successful", "Amount:", res.GetAmount(), "Type:", strings.ToUpper(res.GetType()))
			return
		}

		if debit != "" {
			dfVal, erro := strconv.ParseFloat(debit, 64)
			if erro != nil {
				log.Fatal("Something went wrong", erro.Error())
			}
			res, err := expenseClient.AddExpense(context.Background(), &expensepb.Transaction{Type: "debit", Amount: dfVal})
			if err != nil {
				log.Fatal("Something went wrong", err.Error())
			}
			fmt.Println("Successful", "Amount:", res.GetAmount(), "Type:", strings.ToUpper(res.GetType()))
			return
		}

		fmt.Println(`
The new command is used to credit/debit money into expenses
	
To credit money into expenses, add --credit or -C flag
Example:
	expensetracker new --credit 1000

To debit money from expenses, add --debit or -D flag
Example:
	expensetracker new --debit 1000

		`)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("credit", "C", "", "The credit flag is used to add money in expenses")
	newCmd.Flags().StringP("debit", "D", "", "The debit flag is used to deduct money from expenses")
}
