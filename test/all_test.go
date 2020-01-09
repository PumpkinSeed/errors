package test

import (
	"testing"

	"github.com/PumpkinSeed/errors"
)

var ErrGlobal = errors.New("global error")

func checkResult(result int) error {
	if result != 4 {
		return errors.New("result not equals with 4")
	}
	return nil
}

func countResult(a, b int) error {
	if err := checkResult(a + b); err != nil {
		return errors.Wrap(err, errors.New("Result is the sum of a + b"))
	}
	return nil
}

func globlaTest() error {
	if err := countResult(12, 4); err != nil {
		return errors.Wrap(err, ErrGlobal)
	}
	return nil
}

func TestCountResult(t *testing.T) {
	if err := countResult(2, 5); err == nil {
		t.Error("error shouldn't be nil")
	}
}

func BenchmarkCountResult(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := countResult(2, 5)
		keep(err)
	}
}

func TestGloblaTest(t *testing.T) {
	if err := globlaTest(); err == nil {
		t.Error("error shouldn't be nil")
	} else {
		var expected = "global error: Result is the sum of a + b: result not equals with 4"
		if err.Error() != expected {
			t.Errorf("Error should be %s, instead of %s", expected, err)
		}
	}
}

func BenchmarkGloblaTest(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := globlaTest()
		keep(err)
	}
}

func keep(err error) {}
