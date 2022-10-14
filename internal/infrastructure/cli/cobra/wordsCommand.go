package cli

import (
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"

	"github.com/atcheri/hexarch-go/internal/infrastructure/databases"
)

const (
	ColorDefault = "\x1b[39m"
	//ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
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
		printWordsInTable(words)
	},
}

func init() {
	wordsCmd.Flags().IntVarP(&offset, "offset", "o", 0, "The offset argument")
	wordsCmd.Flags().IntVarP(&limit, "limit", "i", 5, "The limit argument")
}

func printWordsInTable(words map[string]string) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: buildTableHeaders(),
	}

	table.Body = &simpletable.Body{
		Cells: buildTableCells(words),
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 3, Text: gray(fmt.Sprintf("Fetched %d words from DB", len(words)))},
		},
	}

	table.Println()
}

func buildTableHeaders() []*simpletable.Cell {
	return []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Text: "#"},
		{Align: simpletable.AlignCenter, Text: "KEY"},
		{Align: simpletable.AlignCenter, Text: "CONTENT"},
	}
}
func buildTableCells(words map[string]string) [][]*simpletable.Cell {
	cells := make([][]*simpletable.Cell, len(words))
	i := 0
	for k, word := range words {
		cell := []*simpletable.Cell{
			{Text: green(fmt.Sprintf("%d", i+1))},
			{Text: k},
			{Text: blue(word)},
		}
		cells[i] = cell
		i++
	}

	return cells
}

//func red(s string) string {
//	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
//}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
}
