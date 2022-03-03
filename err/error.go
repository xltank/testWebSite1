package err

import (
	"fmt"
)

const (
	CodeServer    = 1000
	CodeParam     = 2000
	CodeAuthUser  = 2002
	CodeUserLogin = 2004
)

var (
	errCodeMap = map[int]string{
		CodeServer:    "Server Error",
		CodeParam:     "Param Error",
		CodeAuthUser:  "Auth User Error",
		CodeUserLogin: "User Login Error",
	}
)

type Error struct {
	Rtn    int
	ErrMsg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %v, %s", e.Rtn, e.ErrMsg)
}

func Err(code int, ext string) Error {
	if code == 0 {
		code = CodeParam
	}

	c := errCodeMap[code]
	msg := errCodeMap[CodeServer]

	if ext != "" {
		msg += ": " + ext
	}

	if c == "" {
		return Error{Rtn: CodeServer, ErrMsg: msg}
	}

	return Error{Rtn: code, ErrMsg: c}
}
