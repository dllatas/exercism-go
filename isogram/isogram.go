package isogram

import (
	"regexp"
	"strings"
)

//var dummy int = 1
var reg, _ = regexp.Compile("[^a-zA-Z]+")

func IsIsogram(w string) bool {
	regWord := reg.ReplaceAllString(w, "")
	regWord = strings.ToLower(regWord)
	wordDist := make(map[rune]int)
	for i, c := range regWord {
		wordDist[c] = i 
	}
	if len(wordDist) == len(regWord) {
		return true
	}
	return false
}
