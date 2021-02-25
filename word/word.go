package word

type Word struct {
	Label    string
	Meaning  string
	Sentence string
}

func CreateWord(label string, meaning string, sentence string) *Word {
	return &Word{
		Label:    label,
		Meaning:  meaning,
		Sentence: sentence,
	}
}
