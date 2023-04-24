package chart

import "strings"

// findLongest find the longest labels to make a good spaces
func findLongest(series []Series) int {
	var longest int
	for _, s := range series {
		l := len(s.GetLabel())
		if l > longest {
			longest = l
		}
	}
	return longest
}

// findMax find the biggest number
func findMax(series []Series) float64 {
	var max float64
	for _, s := range series {
		if s.GetValue() > max {
			max = s.GetValue()
		}
	}
	return max
}

// left align to the left string s with fill string repeat n times
// ex: "Test", "-", 4
// output: Test----
func left(s string, fill string, n int) string {
	return s + strings.Repeat(fill, n)
}

// right align to the right string s with fill string repeat n times
// ex: "Test", "-", 4
// output: ----Test
func right(s string, fill string, n int) string {
	return strings.Repeat(fill, n) + s
}

// center string s with fill string repeat n times
// ex: "Test", "-", 4
// output: ----Test----
func center(s string, fill string, n int) string {
	return strings.Repeat(fill, n) + s + strings.Repeat(fill, n)
}

// alignLeft align to the left string s with fill string repeat n times
// ex: "Test", "-", 10
// output: Test------
func alignLeft(s string, fill string, n int) string {
	var sLen = len(s)

	if sLen < n {
		return left(s, fill, n-sLen)
	}
	return s
}

// alignRight align to the right string s with fill string repeat n times
// ex: "Test", "-", 10
// output: ------Test
func alignRight(s string, fill string, n int) string {
	var sLen = len(s)

	if sLen < n {
		return right(s, fill, n-sLen)
	}
	return s
}
