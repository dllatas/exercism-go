package strand

var comp = map[byte]string{
	'C': "G",
	'G': "C",
	'T': "A",
	'A': "U",
}

func ToRNA(dna string) string {
	rna := ""
	for i := 0; i < len(dna); i++ {
		rna = rna + comp[dna[i]]
	}
	return rna
}
