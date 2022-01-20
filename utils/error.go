package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

const CodeServerError int = 1000
const CodeParamError int = 2000
const CodeAuthError int = 3000
const CodeLoginError int = 3001

type Error struct {
	Rtn    int
	ErrMsg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %v, %s", e.Rtn, e.ErrMsg)
}

func create(rtn int, err interface{}) Error {
	var e Error
	switch v := err.(type) {
	case string:
		e = Error{
			Rtn:    rtn,
			ErrMsg: v,
		}
	case error:
		e = Error{
			Rtn:    rtn,
			ErrMsg: v.Error(),
		}
	default:
		s, err := json.Marshal(e)
		if err == nil {
			e = Error{
				Rtn:    rtn,
				ErrMsg: string(s),
			}
		} else {
			panic("Invalid param in error.New()")
		}
	}

	log.Println(e)
	return e
}

func NewServerError(e interface{}) Error {
	if e == nil {
		e = "Server Error"
	}
	return create(CodeServerError, e)
}

func NewParamError(e interface{}) Error {
	if e == nil {
		e = "Param Error"
	}
	return create(CodeParamError, e)
}

func NewLoginError(e interface{}) Error {
	if e == nil {
		e = "Login Error"
	}
	return create(CodeLoginError, e)
}

func NewAuthError(e interface{}) Error {
	if e == nil {
		e = "Auth Error"
	}
	return create(CodeAuthError, e)
}
