package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var number string
var noAnsi bool

func init() {
	// Register command
	rootCmd.AddCommand(scanCmd)

	// Register flags
	scanCmd.PersistentFlags().StringVarP(&number, "number", "n", "", "The phone number to scan (E164 or international format)")
	scanCmd.PersistentFlags().StringVarP(&number, "input", "i", "", "Text file containing a list of phone numbers to scan (one per line)")
	scanCmd.PersistentFlags().StringVarP(&number, "output", "o", "", "Output to save scan results")
	scanCmd.PersistentFlags().StringVarP(&number, "scanner", "s", "", "Scanner to use")
	scanCmd.PersistentFlags().StringVar(&number, "recon", "", "Launch custom format reconnaissance")
	scanCmd.PersistentFlags().BoolVar(&noAnsi, "no-ansi", false, "Disable colored output")
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a phone number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scan")
	},
}
