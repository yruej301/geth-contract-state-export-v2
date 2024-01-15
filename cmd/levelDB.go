package cmd

import (
	"github.com/spf13/cobra"
	"main/tools"
)


// rootCmd represents the base command when called without any subcommands
var contractStateCmd = &cobra.Command{
	Use:   "contractState <LevelDB path> <contract address>",
	Short: "Return all contract state in the merkle-patricia tree part of LevelDB",
	Run: func(cmd *cobra.Command, args []string) {
		tools.ContractState(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(contractStateCmd)
}