package reverse

func String(a string) string {
	b := ""
	for _, runeValue := range a {
		b = string(runeValue) + b
	}
	return b
}
