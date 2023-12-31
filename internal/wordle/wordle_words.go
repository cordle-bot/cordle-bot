package wordle

import (
	"encoding/json"
	"os"

	"cordle/internal/config"
	"cordle/internal/pkg/util"
)

// Calls LoadWords when the module is imported
func init() {
	LoadWords(config.Config.Game.AnswersPath, config.Config.Game.GuessesPath)
}

// loadWordsFromFile reads the given json files containing allowed words and answers and returns them (answers, guesses)
func loadWordsFromFile(aPath string, gPath string) ([]string, []string) {
	afile, err := os.ReadFile(aPath)
	util.CheckErrMsg(err, "Failed to load answers")
	gfile, err := os.ReadFile(gPath)
	util.CheckErrMsg(err, "Failed to load guesses")

	// Decode JSON
	answers := []string{}
	err = json.Unmarshal(afile, &answers)
	util.CheckErrMsg(err, "Failed to decode answers")

	guesses := []string{}
	err = json.Unmarshal(gfile, &guesses)
	util.CheckErrMsg(err, "Failed to decode guesses")

	return answers, guesses
}
