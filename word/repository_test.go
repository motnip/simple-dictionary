package word

import "testing"

func TestAddWord(t *testing.T) {

    repo := NewRepository()

    newWord := Word{
        Label:   "hello",
        Meaning: "ciao",
    }

    result := repo.AddWord(&newWord)

    if len(result.words) < 1 {
        t.Error("word has not been persisted", result)
    }
}

func TestListWord(t *testing.T) {

    repo := NewRepository()

    newWord := Word{
        Label:   "hello",
        Meaning: "ciao",
    }

    _ = repo.AddWord(&newWord)

    result := repo.ListWords()

    if len(result) < 1 {
        t.Error("no list of words have been returned", result)
    }

    if result[0].Label != newWord.Label {
        t.Errorf("expected %s, got %s ", newWord.Label, result[0].Label)
    }
}

