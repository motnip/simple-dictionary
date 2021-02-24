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

