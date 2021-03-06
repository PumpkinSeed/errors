package errors

import (
	stderrors "errors"
	"testing"
)

func TestNew(t *testing.T) {
	testText := "test text"
	err := New(testText)
	if err == nil {
		t.Error("Error shouldn't be nil")
	}
	if err != nil && err.Error() != "test text" {
		t.Errorf("Error should be '%s', instead of %s", testText, err.Error())
	}
}

func TestUnwrap(t *testing.T) {
	originalErr := New("test")
	wrappedErr := stderrors.New("text")
	err := Wrap(originalErr, wrappedErr)
	unwrappedErr := stderrors.Unwrap(err)

	if unwrappedErr == nil || unwrappedErr.Error() != "test" {
		t.Error("Unwrapped error should be 'test'")
	}
}

func TestIs(t *testing.T) {
	testErr1 := New("test err 1")
	testErr2 := New("test err 2")
	testErr3 := New("test err 3")
	testErr4 := New("test err 4")
	testFail := New("test fail")

	err := Wrap(New("origin"), testErr1)
	err = Wrap(err, testErr2)
	err = Wrap(err, testErr3)
	err = Wrap(err, testErr4)

	testtrue(t, stderrors.Is(err, testErr1), "testErr1")
	testtrue(t, stderrors.Is(err, testErr2), "testErr2")
	testtrue(t, stderrors.Is(err, testErr3), "testErr3")
	testtrue(t, stderrors.Is(err, testErr4), "testErr4")
	testtrue(t, stderrors.Is(err, err), "own")
	testfalse(t, stderrors.Is(err, testFail), "testFail")
}

func TestWrappedError(t *testing.T) {
	testErr1 := New("test err 1")
	testErr2 := New("test err 2")
	testErr3 := New("test err 3")
	testErr4 := New("test err 4")

	err := Wrap(New("origin"), testErr1)
	err = Wrap(err, testErr2)
	err = Wrap(err, testErr3)
	err = Wrap(err, testErr4)
	expected := "test err 4: test err 3: test err 2: test err 1: origin"
	if err == nil || err.Error() != expected {
		t.Errorf("Error should be %s, instead of %s", expected, err.Error())
	}
}

//func BenchmarkNew(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		err := New("test err 1")
//		keep(err)
//	}
//}

func BenchmarkWrap(b *testing.B) {
	testErr1 := New("test err 1")
	err := New("origin")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ne := Wrap(err, testErr1)
		keep(ne)
	}
}

func keep(err error) {}

func testtrue(t *testing.T, value bool, msg ...string) {
	if !value {
		t.Errorf("Should be true, %v", msg)
	}
}

func testfalse(t *testing.T, value bool, msg ...string) {
	if value {
		t.Errorf("Should be false, %v", msg)
	}
}
