package model

import "fmt"

type Dictionary struct {
	Language string
	Words    []*Word
}

const MANDATORY_FIELD = "missing mandatory field %s"

type Word struct {
	Label    string
	Meaning  string
	Sentence string
}

func (w *Word) Validate() (bool, error) {

	if val, err := validateString(w.Label, "label"); val == false && err != nil {
		return val, err
	}

	if val, err := validateString(w.Meaning, "meaning"); val == false && err != nil {
		return val, err
	}

	return true, nil
}

func validateString(field string, fieldName string) (bool, error) {
	if len(field) == 0 {
		return false, fmt.Errorf(MANDATORY_FIELD, fieldName)
	}
	return true, nil
}
