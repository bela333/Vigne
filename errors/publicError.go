package errors

import "errors"

type PublicError struct {
	PublicPart string
	error
}


//Returns a new error
func New(private string, public string) error {
	e := PublicError{}
	e.error = errors.New(private)
	e.PublicPart = public
	return &e
}