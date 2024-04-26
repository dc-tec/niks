package nikscli

import (
	"fmt"
	"github.com/dc-tec/niks-cli/pkg/niks-cli"
	"github.com/spf13/cobra"
)

var listGenerations = &cobra.Command{

	Use:   "list-generations",
	Short: "List all NixOS generations",
	Long:  `List all NixOS generations that are available on the system`,

	Run: func(cmd *cobra.Command, args []string) {

		err := nikscli.ListGenerations()

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listGenerations)
}
