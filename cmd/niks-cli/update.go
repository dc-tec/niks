package nikscli

import (
	"fmt"

	"github.com/dc-tec/niks-cli/pkg/niks-cli"
	"github.com/spf13/cobra"
)

var config string
var path string

var updateCmd = &cobra.Command{

	Use:   "update",
	Short: "Update the system configuration",
	Long:  `Update the system configuration`,

	Run: func(cmd *cobra.Command, args []string) {

		err := nikscli.Update(path, config)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	updateCmd.Flags().StringVarP(&config, "path", "p", ".", "Which path to use")
	updateCmd.Flags().StringVarP(&config, "config", "c", "", "Which configuration to update")
	updateCmd.MarkFlagRequired("config")

	rootCmd.AddCommand(updateCmd)
}
