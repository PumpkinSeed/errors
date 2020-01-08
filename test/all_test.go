package test

import (
	"github.com/PumpkinSeed/errors"
	"testing"
)

func checkResult(result int) error {
	if result != 4 {
		errors.New("result not equals with 4")
	}
	return nil
}

func countResult(a, b int) error {
	if err := checkResult(a+b); err != nil {
		return errors.Wrap(err, errors.New("Result is the sum of a + b"))
	}
	return nil
}

func TestCountResult(t *testing.T) {
	if err := countResult(2, 5); err == nil {
		t.Error("error shouldn't be nil")
	}
}
