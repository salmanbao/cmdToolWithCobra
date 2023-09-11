package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Printer",
	Long:  "All software has versions",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	fmt.Printf("Software Version: %s\n", version)
}
