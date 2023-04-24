package chart

import "strings"

// findLonguest find longuest labels to make a good indendation
func findLonguest(series []Serie) int {
	var longuest int
	for _, serie := range series {
		l := len(serie.GetLabel())
		if l > longuest {
			longuest = l
		}
	}
	return longuest
}

// findMax find the biggest number
func findMax(series []Serie) float64 {
	var max float64
	for _, serie := range series {
		if serie.GetValue() > max {
			max = serie.GetValue()
		}
	}
	return max
}

// alignLeft align left string s with fill string repeat n times
// ex: "Test", "-", 4
// output: Test----
func alignLeft(s string, fill string, n int) string {
	return s + strings.Repeat(fill, n)
}

// alignRight align right string s with fill string repeat n times
// ex: "Test", "-", 4
// output: ----Test
func alignRight(s string, fill string, n int) string {
	return strings.Repeat(fill, n) + s
}

// center center string s with fill string repeat n times
// ex: "Test", "-", 4
// output: ----Test----
func center(s string, fill string, n int) string {
	return strings.Repeat(fill, n) + s + strings.Repeat(fill, n)
}

// ajustLeft ajust left string s with fill string repeat n times
// ex: "Test", "-", 10
// output: Test------
func ajustLeft(s string, fill string, n int) string {
	var sLen = len(s)

	if sLen < n {
		return alignLeft(s, fill, n-sLen)
	}
	return s
}

// ajustRight ajust right string s with fill string repeat n times
// ex: "Test", "-", 10
// output: ------Test
func ajustRight(s string, fill string, n int) string {
	var sLen = len(s)

	if sLen < n {
		return alignRight(s, fill, n-sLen)
	}
	return s
}
