package nikscli

import (
	"fmt"

	"github.com/dc-tec/niks-cli/pkg/niks-cli"
	"github.com/spf13/cobra"
)

func updateCmd() *cobra.Command {
	var config string
	var dryRun bool

	cmd := &cobra.Command{

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

	cmd.Flags().StringVarP(&config, "config", "c", "", "Which configuration to update")
	cmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Dry run the update")
	cmd.MarkFlagRequired("config")

	return cmd
}

func main() {
	rootCmd.AddCommand(updateCmd())
	rootCmd.Execute()
}
