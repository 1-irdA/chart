package chart

import "strings"

// find longuest labels to make a good indendation
func findLonguest(labels []string) int {
	var longuest int
	for _, label := range labels {
		l := len(label)
		if l > longuest {
			longuest = l
		}
	}
	return longuest
}

// find the biggest number
func findMax(values []float64) float64 {
	var max float64
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

// align left string s with fill string repeat n times
// ex: "Test", "-", 4
// output: Test----
func alignLeft(s string, fill string, n int) string {
	return s + strings.Repeat(fill, n)
}

// align left string s with fill string repeat n times
// ex: "Test", "-", 4
// output: ----Test
func alignRight(s string, fill string, n int) string {
	return strings.Repeat(fill, n) + s
}

// center string s with fill string repeat n times
// ex: "Test", "-", 4
// output: ----Test----
func center(s string, fill string, n int) string {
	return strings.Repeat(fill, n) + s + strings.Repeat(fill, n)
}

// ajust left string s with fill string repeat n times
// ex: "Test", "-", 10
// output: Test------
func ajustLeft(s string, fill string, n int) string {
	var sLen = len(s)

	if sLen < n {
		return alignLeft(s, fill, n-sLen)
	}
	return s
}

// ajust right string s with fill string repeat n times
// ex: "Test", "-", 10
// output: ------Test
func ajustRight(s string, fill string, n int) string {
	var sLen = len(s)

	if sLen < n {
		return alignRight(s, fill, n-sLen)
	}
	return s
}
