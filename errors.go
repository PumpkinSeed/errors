package errors

import (
	stderrors "errors"
)

type err struct {
	original error
	wrapped  error
}

// New creates a new error by using the genesis error
func New(text string) error {
	return err{
		original: genesisErr{text},
		wrapped:  nil,
	}
}

func (e err) Error() string {
	var original string
	var wrapped string

	if e.wrapped == nil {
		return e.original.Error()
	}

	wrapped = e.wrapped.Error()
	if e.original != nil {
		original = e.original.Error()
	}
	return wrapped + ": " + original
}

func (e err) Is(target error) bool {
	if e == target {
		return true
	}
	if stderrors.Is(e.original, target) {
		return true
	}
	return stderrors.Is(e.wrapped, target)
}

func (e err) Unwrap() error {
	return e.original
}

// Wrap new error into an existing error
func Wrap(original error, wrapped error) error {
	return err{
		original: original,
		wrapped:  wrapped,
	}
}
