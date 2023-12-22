package cmd

import (
	"github.com/spf13/cobra"

  installFiles "templ-gen/cli/installFiles"
)


// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "This command will init a templ project",
	Long:  `This command will initialize a templ/GO project with Tailwind and TypeScript built in`,
	Run: func(cmd *cobra.Command, args []string) {
    installFiles.InstallFiles()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
