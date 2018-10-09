package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(n int) int

func (s IntList) Foldl(f binFunc, initial int) int {
	acum := initial
	for j := 0; j < len(s); j++ {
		acum = f(acum, s[j])
	}
	return acum
}

func (s IntList) Foldr(f binFunc, initial int) int {
	acum := initial
	for j := len(s) - 1; j >= 0; j-- {
		acum = f(s[j], acum)
	}
	return acum
}

func (s IntList) Filter(f predFunc) IntList {
	var result IntList = []int{}
	for j := 0; j < len(s); j++ {
		if f(s[j]) {
			result = append(result, s[j])
		}
	}
	return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(f unaryFunc) IntList {
	var result IntList = []int{}
	for j := 0; j < len(s); j++ {
		result = append(result, f(s[j]))
	}
	return result
}

func (s IntList) Reverse() IntList {
	var result IntList = []int{}
	for j := len(s) - 1; j >= 0; j-- {
		result = append(result, s[j])
	}
	return result
}

func (s IntList) Append(a IntList) IntList {
	for j := 0; j < len(a); j++ {
		s = append(s, a[j])
	}
	return s
}

func (s IntList) Concat(a []IntList) IntList {
	for j := 0; j < len(a); j++ {
		s = s.Append(a[j])
	}
	return s
}
