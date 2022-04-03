package Exceptions

import (
	"errors"
)

func IndexOutOfBoundsException() error {
	return errors.New("index out of bounds")
}
