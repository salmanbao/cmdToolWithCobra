package cmd

import (
	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
)

var RootCmd = &cobra.Command{
	Use: "app",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func init() {

	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	cmdPrint.AddCommand(cmdTimes)
	RootCmd.AddCommand(cmdPrint)
	RootCmd.AddCommand(cmdVersion)

	// you can just use -v flag instead of command
	RootCmd.Flags().BoolP("version", "v", false, "Print the software version")
}
