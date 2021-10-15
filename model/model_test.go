package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate_missingLabelField_ReturnError(t *testing.T) {

	expectedError := fmt.Errorf(MANDATORY_FIELD, "Label")
	word := Word{
		Label:    "",
		Meaning:  "meaning",
		Sentence: "sentence",
	}

	result, err := word.Validate()

	assert.False(t, result, "Router returned unexpected value: got %v want %v", result, "true")

	assert.EqualError(t, expectedError, err.Error(), "Router returned unexpected value: got %v want %v", err.Error(), expectedError)
}

func TestValidate_missingMeaningField_ReturnError(t *testing.T) {

	expectedError := fmt.Errorf(MANDATORY_FIELD, "Meaning")
	word := Word{
		Label:    "lable",
		Meaning:  "",
		Sentence: "sentence",
	}

	result, err := word.Validate()

	assert.False(t, result, "Router returned unexpected value: got %v want %v", result, "true")

	assert.EqualError(t, expectedError, err.Error(), "Router returned unexpected value: got %v want %v", err.Error(), expectedError)
}
