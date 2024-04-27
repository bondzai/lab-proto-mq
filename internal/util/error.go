package util

import (
	"errors"
)

func NewError(msg string) error {
	return errors.New(msg)
}

var ErrCommon = NewError("common error")
