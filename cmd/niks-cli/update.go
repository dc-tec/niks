package nikscli

import (
	"fmt"

	"github.com/dc-tec/niks-cli/pkg/niks-cli"
	"github.com/spf13/cobra"
)

var config string
var dryRun bool

var updateCmd = &cobra.Command{

	Use:   "update",
	Short: "Update the system configuration",
	Long:  `Update the system configuration`,

	Run: func(cmd *cobra.Command, args []string) {

		err := nikscli.Update(config, dryRun)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	updateCmd.Flags().StringVarP(&config, "config", "c", "", "Which configuration to update")
	updateCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Dry run the update")
	updateCmd.MarkFlagRequired("config")

	rootCmd.AddCommand(updateCmd)
}
