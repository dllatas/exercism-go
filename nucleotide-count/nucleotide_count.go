package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[byte]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	length := len(d)
	var h Histogram
	h = make(Histogram)
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0
	for i := 0; i < length; i++ {
		if _, ok := h[d[i]]; ok {
			h[d[i]] = h[d[i]] + 1
			continue
		}
		return h, errors.New("Strand with invalid nucleotides")
	}
	return h, nil
}
