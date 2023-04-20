package chart

import "strings"

// findLonguest find longuest labels to make a good indendation
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

// findMax find the biggest number
func findMax(values []float64) float64 {
	var max float64
	for _, value := range values {
		if value > max {
			max = value
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

// duplicateHorizontal duplicate string s horizontaly n times
// ex: "Test", 3
// output: Test\nTest\nTest\n
func duplicateHorizontal(s string, n int) string {
	if n == 0 {
		return s
	}
	return s + "\n" + duplicateHorizontal(s, n-1)
}

// sum return the sum of the array
func sum(arr []float64) float64 {
	var sum float64

	for _, v := range arr {
		sum += v
	}
	return sum
}
