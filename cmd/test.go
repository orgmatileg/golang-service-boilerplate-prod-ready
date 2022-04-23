package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testCMD)
}

var testCMD = &cobra.Command{
	Use:   "test",
	Short: "Command for testing code",
	Long:  `Command for testing code`,
	Run: func(cmd *cobra.Command, args []string) {
		test()
	},
}

func test() {

}
