package collatzconjecture

import (
	"errors"
)

var err = errors.New("Bad input")

func CollatzConjecture(input int) (int, error) {
	if input < 1 {
		return 0, err
	}
	result := 0
	for input > 1 {
		if input%2 == 0 {
			input = input / 2
		} else {
			input = input*3 + 1
		}
		result++
	}
	return result, nil
}
