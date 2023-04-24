package chart

import (
	"testing"
)

var testSeries = []Series{
	NewSeries("go", "=", 25.0),
	NewSeries("cpp", "=", 10.0),
	NewSeries("java", "=", 2.0),
	NewSeries("ts", "=", 30.5),
	NewSeries("c#", "=", 0.5),
}

func TestCenter(t *testing.T) {
	got := center("Test", " ", 4)
	want := "    Test    "

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestLeft(t *testing.T) {
	got := left("Test", "-", 4)
	want := "Test----"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRight(t *testing.T) {
	got := right("Test", "*", 4)
	want := "****Test"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAlignLeft(t *testing.T) {
	got := alignLeft("Test", "-", 10)
	want := "Test------"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAlignLeftSmaller(t *testing.T) {
	got := alignLeft("Test", "-", 1)
	want := "Test"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAlignRight(t *testing.T) {
	got := alignRight("Test", "/", 10)
	want := "//////Test"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAlignRightSmaller(t *testing.T) {
	got := alignRight("Test", "/", 3)
	want := "Test"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFindLongest(t *testing.T) {
	got := findLongest(testSeries)
	want := 4

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFindMax(t *testing.T) {
	got := findMax(testSeries)
	want := 30.5

	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}
