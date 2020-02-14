package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sundowndev/phoneinfoga/scanners"
	"github.com/sundowndev/phoneinfoga/utils"
)

func init() {
	// Register command
	rootCmd.AddCommand(reconCmd)

	// Register flags
	reconCmd.PersistentFlags().StringVarP(&number, "number", "n", "", "The phone number to scan (E164 or international format)")
	reconCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Text file containing a list of phone numbers to scan (one per line)")
	reconCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output to save scan results")
	reconCmd.PersistentFlags().BoolVar(&NoAnsi, "no-ansi", false, "Disable colored output")
}

var reconCmd = &cobra.Command{
	Use:   "recon",
	Short: "Launch custom format reconnaissance",
	Run: func(cmd *cobra.Command, args []string) {
		utils.LoggerService.Infoln("Custom recon for phone number", number)

		scanners.ScanCLI(number)

		utils.LoggerService.Infoln("Job finished.")
	},
}
