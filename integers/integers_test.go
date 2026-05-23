package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(10, 2)
	expected := 12

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
