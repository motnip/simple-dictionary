package word

type Word struct {
	Label   string
	Meaning string
}

func CreateWord(label string, meaning string) *Word {
	return &Word{
		Label:   label,
		Meaning: meaning,
	}
}
