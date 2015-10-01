package core

import "errors"

var (
	NotFound = errors.New("Not Found")
)

type errorInterface interface {
	GetErrors() []error
}

type Errors struct {
	errors []error
}

func (errs Errors) GetError() []error {
	return errs.errors
}
