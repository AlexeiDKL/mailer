package dklserrors

import "fmt"

func newError(message string) error {
	return fmt.Errorf("%s", message)
}

func NotRealizedError() error {
	return newError("Not realized yet")
}

func NotFoundError() error {
	return newError("Not found")
}

func Wrap(op string, err error) error {
	return fmt.Errorf("%s: %w", op, err)
}

func UnsupportedType(unType string) error {
	return newError("unsupported type: " + unType)
}
