package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var cmdPrint = &cobra.Command{
	Use:   "print [string to print]",
	Short: "print anything to screen.",
	Long:  "print is for printing anything back to the screen.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("print: " + strings.Join(args, " "))
	},
}
