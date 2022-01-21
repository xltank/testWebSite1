package res

type Response struct {
	Rtn    int         `json:"rtn"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

func response(rtn int, msg string) *Response {
	return &Response{
		Rtn:    rtn,
		ErrMsg: msg,
		Data:   nil,
	}
}

func (r *Response) Error(msg string) Response {
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
	MarshalJsonErr  = response(1000, "Marshal Json Error")
	Ok              = response(200, "")            // 成功
	Err             = response(500, "Serve Error") //服务器错误，请重新再试
	ServerErr       = response(500, "Serve Error") //服务器错误，请重新再试
	ParamErr        = response(3000, "Param Error")
	UserPasswordErr = response(3001, "User Password Error")
	TokenParseErr   = response(3002, "Token Parse Error")
	AuthErr         = response(4000, "Auth Error")
)
