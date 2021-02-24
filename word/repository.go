package word

type repository struct {
    language string
    words []Word
}

//TOMAS at the moment instantiate a new repo with language English
func NewRepository() *repository {
    return &repository{
        language: "EN",
    }
}

func (r *repository) AddWord(word *Word) bool {
    return false
}
