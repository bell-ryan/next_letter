package next_letter

import "testing"

func TestNextLetterSingle(t *testing.T) {
	c := "a"

	want := "b"
	got := FindNext(c)

	if got != want {
		t.Errorf("next letter is incorrect. expected %s got %s", want, got)
	}
}
func TestNextLetterWrapAround(t *testing.T) {
	c := "z"

	want := "aa"
	got := FindNext(c)

	if got != want {
		t.Errorf("next letter is incorrect. expected %s got %s", want, got)
	}
}
func TestNextLetterMultiple(t *testing.T) {
	c := "bbb"

	want := "ccc"
	got := FindNext(c)

	if got != want {
		t.Errorf("next letter is incorrect. expected %s got %s", want, got)
	}
}
