package cli

import (
	"github.com/spf13/cobra"
	"log"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List entities (e.g: words, sentences), given offset and limits",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	// listCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the DTO")
	//
	// if err := listCmd.MarkFlagRequired("name"); err != nil {
	//	fmt.Println(err)
	// }

	listCmd.AddCommand(wordsCmd)
}
