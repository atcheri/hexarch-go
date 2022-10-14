package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	list string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hexarch-go",
	Short: "A brief description of your application",
	Long:  `A longer description of the CLI would land here`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)

	rootCmd.PersistentFlags().StringVarP(&list, "list", "l", "", "List DTOs, for example \"words\" or \"sentences\" by specifying a name argument")
}
