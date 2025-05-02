package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errNotFound = errors.New("not found")

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
	var multErr *MultiError
	assert.True(t, errors.As(err, &multErr))

	expectedMessage = "3 errors occured:\n\t* error 1\t* error 2\t* not found\n"
	err = Append(err, errNotFound)
	assert.EqualError(t, err, expectedMessage)
	assert.True(t, errors.Is(err, errNotFound))
}
