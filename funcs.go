package luluo

import (
	"errors"
)

var ErrNotFound = errors.New("value not found")

var notFound = func(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func SetNotFound(cb func(error) bool) {
	notFound = cb
}

func IsErrNotFound(err error) bool {
	return notFound(err)
}

type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

var errorWrapFunc func(err error, msg string) error

func SetErrorWrap(cb func(err error, msg string) error) {
	errorWrapFunc = cb
}

func errorWrap(err error, msg string) error {
	if errorWrapFunc != nil {
		return errorWrapFunc(err, msg)
	}

	return &wrapError{
		err: err,
		msg: msg + ": " + err.Error(),
	}
}

func newArgumentError(name string, msg string) error {
	return errors.New(msg)
}
