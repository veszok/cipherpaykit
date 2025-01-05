package test

import (
	"bytes"
	"errors"
	"testing"
)

func checkValue(expected, actual interface{}) (bool, error) {
	switch expected := expected.(type) {
	case string:
		if actual_check, ok := actual.(string); ok {
			return expected == actual_check, nil
		}
		return false, errors.New("Argument types mismatch")
	case []byte:
		if actual_check, ok := actual.([]byte); ok {
			return bytes.Equal(expected, actual_check), nil
		}
		return false, errors.New("Argument types mismatch")
	default:
		return false, errors.New("Argument types mismatch")
	}
}

func Assert(t *testing.T, expected, actual interface{}) {
	res, err := checkValue(expected, actual)
	if err != nil {
		t.Errorf("Error occur: %v", err)
	}
	if !res {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
