package chart

import "testing"

func TestCenter(t *testing.T) {
	got := center("Test", " ", 4)
	want := "    Test    "

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAlignLeft(t *testing.T) {
	got := alignLeft("Test", "-", 4)
	want := "Test----"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAlignRight(t *testing.T) {
	got := alignRight("Test", "*", 4)
	want := "****Test"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAjustLeft(t *testing.T) {
	got := ajustLeft("Test", "-", 10)
	want := "Test------"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAjustRight(t *testing.T) {
	got := ajustRight("Test", "/", 10)
	want := "//////Test"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
