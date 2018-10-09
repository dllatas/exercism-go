package scale

import (
	"strings"
)

var m = map[string]string{
	"C":  "s",
	"a":  "s",
	"G":  "s",
	"D":  "s",
	"A":  "s",
	"E":  "s",
	"B":  "s",
	"F#": "s",
	"e":  "s",
	"b":  "s",
	"f#": "s",
	"c#": "s",
	"g#": "s",
	"d#": "s",
	"F":  "f",
	"Bb": "f",
	"Eb": "f",
	"Ab": "f",
	"Db": "f",
	"Gb": "f",
	"d":  "f",
	"g":  "f",
	"c":  "f",
	"f":  "f",
	"bb": "f",
	"eb": "f",
}

var tone = map[rune]int{
	'M': 2,
	'm': 1,
	'A': 3,
}

var sharps = [12]string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

var flats = [12]string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

func TonicScale(tonic string, scale [12]string) []string {
	var tonicScale []string
	for i, note := range scale {
		if tonic == note {
			tonicScale = append(scale[i:], scale[0:i]...)
			break
		}
	}
	return tonicScale
}

func Scale(tonic string, interval string) []string {
	var tonicScale []string
	scale := m[tonic]
	tonic = strings.ToUpper(tonic[0:1]) + tonic[1:]
	if scale == "s" {
		tonicScale = TonicScale(tonic, sharps)
	} else {
		tonicScale = TonicScale(tonic, flats)
	}
	if interval == "" {
		return tonicScale
	}
	resultScale := append(tonicScale[0:1])
	index := 0
	maxIndex := len(tonicScale)
	for _, i := range interval {
		index = index + tone[i]
		if maxIndex <= index {
			break
		}
		resultScale = append(resultScale, tonicScale[index:index+1]...)
	}
	return resultScale
}
