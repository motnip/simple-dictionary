package word

type repository struct {
	Dictionary *Dictionary
}

func NewRepository() *repository {
	return &repository{
		Dictionary: &Dictionary{},
	}
}

func (r *repository) CreateDictionary(language string) *Dictionary {
	return &Dictionary{
		Language: language,
	}
}

func (r *repository) AddWord(word *Word) *repository {
	r.Dictionary.Words = append(r.Dictionary.Words, word)
	return r
}

func (r *repository) ListWords() []*Word {
	return r.Dictionary.Words
}
