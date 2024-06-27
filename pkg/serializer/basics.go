package serializer

//基本响应信息
type Response struct {
	StatusCode uint64 `json:"status_code"`          //状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` //返回状态描述
}
