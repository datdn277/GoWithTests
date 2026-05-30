package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("unknow word", func(t *testing.T) {
		_, err := dict.Search("unknow")

		assertNotNil(t, err)
		assertError(t, err, NotFoundError)
	})

	t.Run("Add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "word"
		definition := "new word"
		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Add word existed", func(t *testing.T) {
		err := dict.Add("test", "updated value")
		assertNotNil(t, err)
		assertString(t, err.Error(), ErrWordExists.Error())
	})

	t.Run("Update new word", func(t *testing.T) {
		word := "word"
		definition := "update word"
		err := dict.Update(word, definition)
		assertNotNil(t, err)
		assertError(t, err, NotFoundError)
	})

	t.Run("Update existed word", func(t *testing.T) {
		word := "test"
		definition := "update word"
		dict.Update(word, definition)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("Delete new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		err := dictionary.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})

	t.Run("Delete existed word", func(t *testing.T) {
		word := "test"
		dict.Delete(word)
		_, err := dict.Search(word)
		assertError(t, err, NotFoundError)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNotNil(t testing.TB, got error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected to get an error.")
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertString(t, got, definition)
}
