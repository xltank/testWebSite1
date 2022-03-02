package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Rtn    int         `json:"rtn"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

/*
res.SendOK(data)
res.SendParamErr(error.ErrUserLogin, "user id: xxxx")
res.SendParamErr(error.Errxxx, err)
*/

func response(rtn int, msg string) *Response {
	return &Response{
		Rtn:    rtn,
		ErrMsg: msg,
		Data:   nil,
	}
}

func (r *Response) Error(msg string) Response {
	if msg == "" {
		msg = "Param error"
	}
	return Response{
		Rtn:    r.Rtn,
		ErrMsg: msg,
		Data:   nil, //r.Data,
	}
}

func (r *Response) Json(data interface{}) Response {
	return Response{
		Rtn:    r.Rtn,
		ErrMsg: "",
		Data:   data,
	}
}

var (
	Ok        = response(0, "")               // 成功
	ServerErr = response(1000, "Serve Error") //服务器错误，请重新再试
	ParamErr  = response(2000, "Param Error")
	//UserPasswordErr = response(2002, "User Password Error")
	//TokenParseErr   = response(2004, "Token Parse Error")
	//AuthErr         = response(2006, "Auth Error")

)

func SendOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Ok.Json(data))
}

func SendParamError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, ParamErr.Error(msg))
}

func SendServerError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, ServerErr.Error(msg))
}
