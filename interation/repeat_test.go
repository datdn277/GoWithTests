package interation

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat string 'a' five times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q but expected %q", got, want)
		}
	})

	t.Run("Repeat string 'b' five times", func(t *testing.T) {
		got := Repeat("b", 5)
		want := "bbbbb"

		if got != want {
			t.Errorf("got %q but expected %q", got, want)
		}
	})

	t.Run("Repeat empty string five times", func(t *testing.T) {
		got := Repeat("", 5)
		want := ""

		if got != want {
			t.Errorf("got %q but expected %q", got, want)
		}
	})

	t.Run("Repeat 'a' ten times", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"

		if got != want {
			t.Errorf("got %q but expected %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
