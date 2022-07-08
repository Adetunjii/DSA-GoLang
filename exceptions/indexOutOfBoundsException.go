package exceptions

import (
	"errors"
)

func IndexOutOfBoundsException() error {
	return errors.New("index out of bounds")
}
