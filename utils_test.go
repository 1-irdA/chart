package chart

import "testing"

func TestCenter(t *testing.T) {
	got := center("Test", " ", 4)
	want := "    Test    "

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAlignLeft(t *testing.T) {
	got := alignLeft("Test", "-", 4)
	want := "Test----"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAlignRight(t *testing.T) {
	got := alignRight("Test", "*", 4)
	want := "****Test"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAjustLeft(t *testing.T) {
	got := ajustLeft("Test", "-", 10)
	want := "Test------"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAjustLeftSmaller(t *testing.T) {
	got := ajustLeft("Test", "-", 1)
	want := "Test"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAjustRight(t *testing.T) {
	got := ajustRight("Test", "/", 10)
	want := "//////Test"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAjustRightSmaller(t *testing.T) {
	got := ajustRight("Test", "/", 3)
	want := "Test"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFindLonguest(t *testing.T) {
	got := findLonguest([]string{"this", "is", "a", "test"})
	want := 4

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFindMax(t *testing.T) {
	got := findMax([]float64{458.3, -78.5, 5, 1875.48, 0.0, 56.8})
	want := 1875.48

	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}
