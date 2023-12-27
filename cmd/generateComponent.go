package cmd

import (
	"github.com/spf13/cobra"

  generateComponent "templ-gen/cli/generateComponent"
)


var genComponentCmd = &cobra.Command{
	Use:   "c [NAME]",
	Short: "This will generate a component",
	Long:  `This command will generate a new templ component`,
  Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
    generateComponent.GenerateComponent(args[0])
	},
}

func init() {
	rootCmd.AddCommand(genComponentCmd)
}
