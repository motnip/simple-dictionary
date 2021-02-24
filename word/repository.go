package word

type repository struct {
    language string
    words []*Word
}

func NewRepository() *repository {
    return &repository{
        language: "EN",
    }
}

func (r *repository) AddWord(word *Word) *repository {
    r.words = append(r.words, word)
    return r
}
