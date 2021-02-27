package word

type Word struct {
	Label    string
	Meaning  string
	Sentence string
}

type Dictionary struct {
	Language string
	Words    []*Word
}
