/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

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
			fmt.Println("Credit called", credit)
			return
		}

		if debit != "" {
			fmt.Println("Dredit called", debit)
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
