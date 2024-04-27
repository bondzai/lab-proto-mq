package util

import (
	"errors"
)

func NewError(msg string) error {
	return errors.New(msg)
}

var ErrCommon = NewError("common error")
var ErrDecodeMsg = NewError("error decoding message")
var ErrCreatePublisher = NewError("error create rabbitmq publisher")
var ErrCreateConsumer = NewError("error create rabbitmq consumer")
