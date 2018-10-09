package grains

import (
	"errors"
	"math"
)

func Square(x int) (uint64, error) {
	if x < 1 || x > 64 {
		return 0, errors.New("Square limited between 1 and 64")
	}
	return uint64(math.Pow(2, float64(x-1))), nil
}

func Total() uint64 {
        prod := uint64(math.Pow(2, 63))
        return prod<<2 - 1
}
