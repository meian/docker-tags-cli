package usecases

import (
	"errors"
	"fmt"
)

var (
	ErrNoImageName = errors.New("image name is not set.")
)

func ErrInvalidRequest(err error) error {
	return fmt.Errorf("invalid request: %w", err)
}

func ErrInvalidURL(url string) error {
	return fmt.Errorf("invalid url: %s", url)
}

func ErrInvalidStatus(code int) error {
	return fmt.Errorf("invalid status: %d", code)
}
