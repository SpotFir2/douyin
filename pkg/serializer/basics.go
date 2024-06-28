package serializer

const (
	CodeSuccess = 0
)

// 基本响应信息
type Response struct {
	StatusCode int    `json:"status_code"`          //状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` //返回状态描述
}

func NewResponse(statusCode int, statusMsg string) Response {
	return Response{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	}
}
