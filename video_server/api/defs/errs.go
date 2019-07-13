package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErroResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErroResponse{HttpSC: 400, Error{Error: "Request body is not correct", ErrorCode: "001"}}
	ErrorNotAuthUser            = ErroResponse{HttpSC: 401, Error{Error: "User autoentication failed", ErrorCode: "002"}}
)
