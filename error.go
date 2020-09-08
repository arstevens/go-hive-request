package hrequest

import "fmt"

// Shared attributes between module error types
type simpleErr struct {
	err error
	msg string
}

// Prints error message
func (se *simpleErr) Error() string {
	return fmt.Errorf(se.msg+": %v", se.err).Error()
}

// Returns the underlying error
func (se *simpleErr) Unwrap() error {
	return se.err
}

// MarshalErr returned when fail to deserialize data
type MarshalErr struct {
	simpleErr
}
