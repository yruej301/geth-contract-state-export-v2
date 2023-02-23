package main

import (
	"main/tools"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "<LevelDB path> <contract address>",
	Short: "Return all contract state in the merkle-patricia tree part of LevelDB",
	Run: func(cmd *cobra.Command, args []string) {
		tools.ContractState(args[0], args[1])
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
