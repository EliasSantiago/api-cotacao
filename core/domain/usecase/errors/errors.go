package errors

import "errors"

var (
	CheckDataExists = func() error {
		err := errors.New("Failed to check if data exists")
		return err
	}
	InvalidParameters = func() error {
		err := errors.New("Invalid parameters")
		return err
	}
)
