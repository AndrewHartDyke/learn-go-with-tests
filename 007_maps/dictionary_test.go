package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("word is present in dict", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})
	t.Run("word is not present in dict", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		if got == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("adding a new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefintion(t, dictionary, word, definition)
	})
	t.Run("adding a word that already exists", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new")

		assertError(t, err, ErrWordExists)
		assertDefintion(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertNoError(t, err)
		assertDefintion(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	t.Run("delete an existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "this is a test"}
		err := dictionary.Delete(word)
		assertNoError(t, err)
		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})
	t.Run("delete a non existant word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertDefintion(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}
