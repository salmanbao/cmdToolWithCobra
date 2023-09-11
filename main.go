package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
)

func main() {
	var echoTimes int

	var cmdPrint = &cobra.Command{
		Use:   "print [string to print]",
		Short: "print anything to screen.",
		Long:  "print is for printing anything back to the screen.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("print: " + strings.Join(args, " "))
		},
	}

	var cmdTimes = &cobra.Command{
		Use:   "times [string to echo]",
		Short: "Echo anything to the screen more times",
		Long:  "echo things multiple times back to the user by providing a count and a string.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}

	var cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Printer",
		Long:  "All software has versions",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
		},
	}

	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
		},
	}
	cmdPrint.AddCommand(cmdTimes)
	rootCmd.AddCommand(cmdPrint)
	rootCmd.AddCommand(cmdVersion)

	// you can just use -v flag instead of command
	rootCmd.Flags().BoolP("version", "v", false, "Print the software version")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printVersion() {
	fmt.Printf("Software Version: %s\n", version)
}

// Examples:
// 1- ./binary-file version
// 2- ./binary-file -v
// 3- ./binary-file --version
// 4- ./binary-file print hello
// 4- ./binary-file print hello
// 5- ./binary-file print times hello --times 5
