package model

import (
	"fmt"
	"testing"
)

func TestValidate_missingLabelField_ReturnError(t *testing.T) {

	expectedError := fmt.Sprintf(MANDATORY_FIELD, "Label")
	word := Word{
		Label:    "",
		Meaning:  "meaning",
		Sentence: "sentence",
	}

	result, err := word.Validate()

	if result != false {
		t.Errorf("Router returned unexpected value: got %v want %v", result, "true")
	}
	if err.Error() != expectedError {
		t.Errorf("Router returned unexpected value: got %v want %v", err.Error(), expectedError)
	}
}

func TestValidate_missingMeaningField_ReturnError(t *testing.T) {

	expectedError := fmt.Sprintf(MANDATORY_FIELD, "Meaning")
	word := Word{
		Label:    "lable",
		Meaning:  "",
		Sentence: "sentence",
	}

	result, err := word.Validate()

	if result != false {
		t.Errorf("Router returned unexpected value: got %v want %v", result, "true")
	}
	if err.Error() != expectedError {
		t.Errorf("Router returned unexpected value: got %v want %v", err.Error(), expectedError)
	}
}
