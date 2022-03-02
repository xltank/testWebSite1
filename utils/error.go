package utils

import (
	"fmt"
)

var (
	ErrServer = Error{Rtn: 1000, ErrMsg: "Server Error"}

	ErrParam     = Error{Rtn: 2000, ErrMsg: "Param Error"}
	ErrAuthUser  = Error{Rtn: 2002, ErrMsg: "Auth User Error"}
	ErrUserLogin = Error{Rtn: 2002, ErrMsg: "User Login Error"}
)

type Error struct {
	Rtn    int
	ErrMsg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %v, %s", e.Rtn, e.ErrMsg)
}
