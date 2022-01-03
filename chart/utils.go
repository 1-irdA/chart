package chart

import "strings"

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

func alignLeft(s string, fill string, n int) string {
	return s + strings.Repeat(fill, n)
}

func alignRight(s string, fill string, n int) string {
	return strings.Repeat(fill, n) + s
}

func center(s string, fill string, n int) string {
	return strings.Repeat(fill, n) + s + strings.Repeat(fill, n)
}

func ajustLeft(s string, fill string, n int) string {
	var sLen = len(s)

	if sLen < n {
		return alignLeft(s, fill, n-sLen)
	}
	return s
}

func ajustRight(s string, fill string, n int) string {
	var sLen = len(s)

	if sLen < n {
		return alignRight(s, fill, n-sLen)
	}
	return s
}
