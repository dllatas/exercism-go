package luhn

import (
	"regexp"
	"strings"
)

var reg, _ = regexp.Compile("[^0-9]+")

func Double(index int, digit byte) int {
	num := int(digit - '0')
	num = 2 * num
	if num > 9 {
		num = num - 9
	}
	return num
}

func Valid(input string) bool {

	noSpaceInput := strings.Replace(input, " ", "", -1)
	if len(noSpaceInput) < 2 {
		return false
	}

	filterInput := reg.ReplaceAllString(noSpaceInput, "")
	if len(noSpaceInput) != len(filterInput) {
		return false
	}

	acum := 0
	for i := len(filterInput) - 2; i >= 0; i = i - 2 {
		digit := Double(i, filterInput[i])
		acum = acum + (digit + int(filterInput[i+1]-'0'))
	}

	if acum%10 == 0 {
		return true
	}

	return false
}
