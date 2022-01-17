package error

import (
	"encoding/json"
	"fmt"
	"log"
)

const CODE_SERVER_ERROR int = 1000
const CODE_PARAM_ERROR int = 2000
const CODE_AUTH_ERROR int = 3000
const CODE_LOGIN_ERROR int = 3001

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
	return create(CODE_SERVER_ERROR, e)
}

func NewParamError(e interface{}) Error {
	if e == nil {
		e = "Param Error"
	}
	return create(CODE_PARAM_ERROR, e)
}

func NewLoginError(e interface{}) Error {
	if e == nil {
		e = "Login Error"
	}
	return create(CODE_LOGIN_ERROR, e)
}

func NewAuthError(e interface{}) Error {
	if e == nil {
		e = "Auth Error"
	}
	return create(CODE_AUTH_ERROR, e)
}
