package charlatan

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type charlatanErrors struct {
	Messages []string
}

func (ce *charlatanErrors) CheckType(name string, fieldType reflect.Type, expected string) bool {
	if fieldType.String() != expected {
		ce.Messages = append(ce.Messages, fmt.Sprintf("field %s is of type %s, expected `%s`", name, fieldType, expected))
		return false
	}

	return true
}

func (ce *charlatanErrors) Errors() error {
	if len(ce.Messages) > 0 {
		return errors.New(strings.Join(ce.Messages, "\n"))
	}

	return nil
}
