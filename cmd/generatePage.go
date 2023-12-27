package cmd

import (
	"github.com/spf13/cobra"

  generatePage "templ-gen/cli/generatePage"
)


var genPageCmd = &cobra.Command{
	Use:   "p [NAME]",
	Short: "This will generate a page",
	Long:  `This command will generate a new templ page`,
  Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
    generatePage.GeneratePage(args[0])
	},
}

func init() {
	rootCmd.AddCommand(genPageCmd)
}
