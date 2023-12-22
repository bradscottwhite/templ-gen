package cmd

import (
	"github.com/spf13/cobra"

  generateComponent "templ-gen/cli/generateComponent"
)


var genCmd = &cobra.Command{
	Use:   "g [NAME]",
	Short: "This will generate a templ component",
	Long:  `This command will generate a templ component`,
  Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
    generateComponent.GenerateComponent(args[0])
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
