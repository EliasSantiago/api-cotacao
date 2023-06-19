package errors

import "errors"

var (
	InvalidParameters = func() error {
		err := errors.New("Invalid parameters")
		return err
	}
)
