package errors

import (
	stdlib "errors"
	"fmt"
)

type err struct {
	original error
	wrapped  error
}

func New(text string) error {
	return &err{
		original: stdlib.New(text),
		wrapped:  nil,
	}
}

func (e *err) Error() string {
	var original string
	var wrapped string

	if e.wrapped == nil {
		return e.original.Error()
	}

	wrapped = e.wrapped.Error()
	if e.original != nil {
		original = e.original.Error()
	}
	return fmt.Sprintf("%s: %s", wrapped, original)
}

func (e *err) Is(target error) bool {
	if e == target {
		return true
	}
	if stdlib.Is(e.original, target) {
		return true
	}
	return stdlib.Is(e.wrapped, target)
}

func (e *err) Unwrap() error {
	return e.original
}

func (e *err) Cause() error {
	return e.original
}

func Wrap(original error, wrapped error) error {
	return &err{
		original: original,
		wrapped: wrapped,
	}
}
