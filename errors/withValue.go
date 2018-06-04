package errors

import (
	"fmt"
)

// errors with value
const (
	OrganizationNameAlreadyExist Type = baseWithValue + iota
	UserNameAlreadyExist
)

type withValue struct {
	typ    Type
	Values []interface{} `json:"values"`
}

var withValueStr = []string{
	"organization with name %s already exists",
	"user with name %s already exists",
}

func (e withValue) Error() string {
	return fmt.Sprintf(e.typ.Reference(), e.Values...)
}

func (e withValue) Code() Type {
	return e.typ
}

// NewOrganizationNameAlreadyExist wraps the OrganizationNameAlreadyExist with name.
func NewOrganizationNameAlreadyExist(name string) TypedError {
	return errWithValue(OrganizationNameAlreadyExist)(name)
}

// NewUserNameAlreadyExist wraps the UserNameAlreadyExist with name.
func NewUserNameAlreadyExist(name string) TypedError {
	return errWithValue(UserNameAlreadyExist)(name)
}

func errWithValue(typ Type) func(args ...interface{}) TypedError {
	return func(args ...interface{}) TypedError {
		return withValue{
			typ:    typ,
			Values: args,
		}
	}
}
