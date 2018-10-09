package proverb

import (
	"fmt"
)

var begin string = "For want of a %s the %s was lost."
var end string = "And all for the want of a %s."

func Proverb(rhyme []string) []string {
	length := len(rhyme)
	var result []string
	if length < 1 {
		return result
	}
	for i := 0; i < length-1; i++ {
		result = append(result, fmt.Sprintf(begin, rhyme[i], rhyme[i+1]))
	}
	return append(result, fmt.Sprintf(end, rhyme[0]))
}
