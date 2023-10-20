package service

import "encoding/json"

// Marshal 序列化json+json自动换行（解决/n问题）
func Marshal(v any) []byte {
	ret, _ := json.Marshal(v)
	return json.RawMessage(string(ret))
}
