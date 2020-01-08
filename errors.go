package errors

import (
	stdlib "errors"
)

type err struct {
	original error
	wrapped error
}

func New(text string) error {
	return &err{
		original: stdlib.New(text),
		wrapped: nil,
	}
}

func (e *err) Error() string {

}
