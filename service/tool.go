package service

import "encoding/json"

// Marshal 序列化json+json自动换行（解决/n问题）
func Marshal(v any) json.RawMessage {
	ret, _ := json.Marshal(v)
	re := string(ret)

	return json.RawMessage(re)
}

// Response 返回信息
type Response struct {
	Code int    `json:"code"` //相应码，错误为0，正确为1
	Msg  string `json:"msg"`  //信息，错误为具体信息，正确为空
}
