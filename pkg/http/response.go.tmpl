package http

const (
	// 正常
	OK int = 0
	// 参数错误
	ParamsError int = 4001
	// 数据错误
	DataError int = 4002
)

// Response Http响应
type Response struct {
	// 是否成功
	Success bool `json:"success"`
	// 错误代码
	ErrorCode int `json:"error_code"`
	// 消息
	Message string `json:"message"`
	// 数据项
	Data interface{} `json:"data"`
}

// NewResponse 创建响应
func NewResponse(errCode int, msg string, data ...interface{}) *Response {
	resp := Response{ErrorCode: errCode, Message: msg}
	if resp.ErrorCode == 0 {
		resp.Success = true
	}
	if len(data) > 0 {
		resp.Data = data[0]
	}
	return &resp
}

// NewResponseOK 创建成功响应
func NewResponseOK(msg string, data ...interface{}) *Response {
	resp := Response{Success: true, ErrorCode: 0, Message: msg}
	if len(data) > 0 {
		resp.Data = data[0]
	}
	return &resp
}

// NewResponseError 创建错误响应, 显示原始错误
func NewResponseError(err interface{}, data ...interface{}) *Response {
	var msg string
	switch e := err.(type) {
	case error:
		msg = e.Error()
	case string:
		msg = e
	default:
		msg = ""
	}
	resp := Response{Success: false, ErrorCode: DataError, Message: msg}
	if len(data) > 0 {
		resp.Data = data[0]
	}
	return &resp
}
