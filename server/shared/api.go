package shared

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseCode string

const (
	ResponseCodeError   = "01"
	ResponseCodeOk      = "00"
	ResponseCodePending = "02"
)

func GetResponse(code string, msg string, data interface{}) *Response {
	return &Response{Code: code, Message: msg, Data: data}
}
