package handlers

import (
	"errors"
)

func HandleErrorMethod(rMethod string, allowedMethod string) error {
	if rMethod != allowedMethod {
		return errors.New("этот метод не поддерживается")
	}
	return nil
}
