package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/NikhilSharma03/expensetracker/server/expensepb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/spf13/viper"
)

var cfgFile string

var gRPCClient *grpc.ClientConn
var expenseClient expensepb.ExpenseServiceClient

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ExpenseTracker",
	Short: "A expense tracker cli.",
	Long: `ExpenseTracker is an cli application build with Cobra.
	
It can do many things including

- Managing your expenses
- Credit to expenses
- Debit from expenses
- Keep track of total amount present
- History of transactions
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	var err error
	gRPCClient, err = grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Something went wrong", err.Error())
	}
	expenseClient = expensepb.NewExpenseServiceClient(gRPCClient)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ExpenseTracker" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ExpenseTracker")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
