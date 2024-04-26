package nikscli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "niks",
	Short: "A small utility wrapper to simplify working with NixOS and the Nix package manager.",
	Long:  `A small utility wrapper to simplify working with NixOS and the Nix package manager.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Yo, waddup?. We encounterd an error: %s\n", err)
		os.Exit(1)
	}

}
