package tests

import "testing"

func TestRepeater(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

}

const repeatCount = 5

func Repeat(character string) string {
	var repated string
	for i := 0; i < repeatCount; i++ {
		repated += character
	}
	return repated
}
