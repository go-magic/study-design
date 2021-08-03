package factory

import "errors"

type ErrorFactory struct {
}

func (f *ErrorFactory) Check(subTask string) (string, error) {
	return "", errors.New("error")
}
