package charlatan

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

var wordList []string

// Generates a random word
func word() interface{} {
	if len(wordList) == 0 {
		loadDictionary()
	}

	idx := rand.Intn(len(wordList))

	return wordList[idx]
}

func words() interface{} {
	n := rand.Intn(10)

	var words []string
	for i := 0; i < n; i++ {
		words = append(words, word().(string))
	}

	return words
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
