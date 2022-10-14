package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/atcheri/hexarch-go/internal/infrastructure/databases"
)

var (
	offset int
	limit  int
)

// wordsCmd represents the words command
var wordsCmd = &cobra.Command{
	Use:   "words",
	Short: "List the words using offset and limit",
	Long:  `The default value of offset is 0, and limit is 5`,
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		db := databases.NewInMemoryDB()
		words := db.GetWords(offset, limit)
		for i, w := range words {
			fmt.Printf("%d. %s\n", i, w)
		}
	},
}

func init() {
	wordsCmd.Flags().IntVarP(&offset, "offset", "o", 0, "The offset argument")
	wordsCmd.Flags().IntVarP(&limit, "limit", "i", 5, "The limit argument")
}
