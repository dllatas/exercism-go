package accumulate

type fn func(string) string

func Accumulate(a []string, f fn) []string {
	length := len(a)
	b := make([]string, length)
	for i := 0; i < length; i++ {
		b[i] = f(a[i])
	}
	return b
}
