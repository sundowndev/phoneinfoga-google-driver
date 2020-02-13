package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sundowndev/phoneinfoga/scanners"
	"github.com/sundowndev/phoneinfoga/utils"
)

// NoAnsi is an option to disable colors for CLI logs
var NoAnsi bool // TODO

var number string
var input string   // TODO
var output string  // TODO
var scanner string // TODO
var recon bool

func init() {
	// Register command
	rootCmd.AddCommand(scanCmd)

	// Register flags
	scanCmd.PersistentFlags().StringVarP(&number, "number", "n", "", "The phone number to scan (E164 or international format)")
	scanCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Text file containing a list of phone numbers to scan (one per line)")
	scanCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output to save scan results")
	scanCmd.PersistentFlags().StringVarP(&scanner, "scanner", "s", "all", "Scanner to use")
	scanCmd.PersistentFlags().BoolVar(&recon, "recon", false, "Launch custom format reconnaissance")
	scanCmd.PersistentFlags().BoolVar(&NoAnsi, "no-ansi", false, "Disable colored output")
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan a phone number",
	Run: func(cmd *cobra.Command, args []string) {
		utils.LoggerService.Infoln("Scanning phone number", number)

		scanners.ScanCLI(number)

		utils.LoggerService.Infoln("Scan finished.")
	},
}
