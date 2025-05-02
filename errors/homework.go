package main

import (
	"errors"
	"fmt"
	"strings"
)

// go test -v homework_test.go

type MultiError struct {
	errs []error
}

func (e *MultiError) Error() string {
	if e == nil || e.errs == nil {
		return ""
	}

	errLen := len(e.errs)
	if errLen == 0 {
		return ""
	}

	if errLen == 1 {
		return fmt.Sprintf("1 error occured:\n\t* %s\n", e.errs[0])
	}

	points := make([]string, len(e.errs))
	for i, err := range e.errs {
		points[i] = fmt.Sprintf("* %s", err)
	}

	return fmt.Sprintf("%d errors occured:\n\t%s\n", errLen, strings.Join(points, "\t"))
}

func (e *MultiError) Unwrap() []error {
	if e == nil || len(e.errs) == 0 {
		return nil
	}

	return e.errs
}

func Append(err error, errs ...error) *MultiError {
	if len(errs) == 0 {
		return nil
	}

	var multiErr *MultiError
	switch {
	case errors.As(err, &multiErr):
		if err == nil {
			err = new(MultiError)
		}

		for _, e := range errs {
			var mulErr *MultiError
			switch {
			case errors.As(e, &mulErr):
				if e != nil {
					multiErr.errs = append(multiErr.errs, mulErr.errs...)
				}
			default:
				if e != nil {
					multiErr.errs = append(multiErr.errs, e)
				}
			}
		}

		return multiErr
	default:
		newErrs := make([]error, 0, len(errs)+1)
		if err != nil {
			newErrs = append(newErrs, err)
		}

		newErrs = append(newErrs, errs...)

		return Append(&MultiError{}, newErrs...)
	}
}
