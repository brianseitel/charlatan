package charlatan

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

var wordList []string

// Generates a random word
func word() string {
	if len(wordList) == 0 {
		loadDictionary()
	}

	idx := rand.Intn(len(wordList))

	return wordList[idx]
}

func loadDictionary() {
	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	wordList = strings.Split(string(bytes), "\n")
	return
}
