package dklserrors

import "fmt"

func newError(message string) error {
	return fmt.Errorf("%s", message)
}

func NotRelizedError() error {
	return newError("Not realized yet")
}

func NotFoundError() error {
	return newError("Not found")
}

func Wrap(op string, err error) error {
	return fmt.Errorf("%s: %w", op, err)
}
