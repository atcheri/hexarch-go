package adapters

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/atcheri/hexarch-go/internal/core/domain"
)

var (
	inMemoryLanguages []domain.LanguageCodeAndName
)

func init() {
	file, err := os.Open("./internal/core/adapters/right/repositories/inMemory/languages.json")
	if err != nil {
		panic("cannot load languages.json file data")
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(byteValue, &inMemoryLanguages)
	if err != nil {
		panic("cannot read languages.json file data")
	}
}
