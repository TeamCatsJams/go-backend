package errors

import (
	"errors"
)

var ErrSomethingWentWrong = errors.New("something went wrong")
var ErrRecordNotFound = errors.New("record not found")
