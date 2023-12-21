package cmd

import (
	"github.com/spf13/cobra"

  fns "templ-gen/fns"
)


// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "This command will init a templ project",
	Long:  `This command will initialize a templ/GO project with Tailwind and TypeScript built in`,
	Run: func(cmd *cobra.Command, args []string) {
		/*var gopherName = "dr-who"

		if len(args) >= 1 && args[0] != "" {
			gopherName = args[0]
		}*/

    fns.InitProject()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
