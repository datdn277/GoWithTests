package slice

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum([]int{1, 2, 3})
	expected := 6

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func TestSumAll(t *testing.T) {
	checkSum := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !slices.Equal(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	}

	t.Run("SumAllTails with empty", func(t *testing.T) {
		got := SumAllTails()
		expected := []int{}
		checkSum(t, got, expected)
	})

	t.Run("SumAllTails with 1 slice one item", func(t *testing.T) {
		got := SumAllTails([]int{1})
		expected := []int{0}
		checkSum(t, got, expected)
	})

	t.Run("SumAllTails with 1 slice multiple items", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3})
		expected := []int{5}
		checkSum(t, got, expected)
	})

	t.Run("SumAllTails with 2 slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1})
		expected := []int{5, 0}
		checkSum(t, got, expected)
	})
}
