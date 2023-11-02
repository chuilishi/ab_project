package wechat

import (
	"ab_project/model"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type TemplateData struct {
	ToUser     string `json:"touser"`
	TemplateID string `json:"template_id"`
	URL        string `json:"url"`
	Data       Data   `json:"data"`
}

type Data struct {
	Username Value `json:"username"`
	Message  Value `json:"message"`
	Status   Value `json:"status"`
}

type Value struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

/*
调用方式
resp := wechat.SendTemplateMessage(model.TemplateMessage{
		WxOpenId:  "",
		Name:      "",
		Message:   "",
		NowStatus: "",
		HTTP:      "",
	}, wechat.GetAccessToken(true))
*/

func SendTemplateMessage(Message model.TemplateMessage, AccessToken string) *http.Response {
	data := TemplateData{
		ToUser: Message.WxOpenId,
		//模板id
		TemplateID: "efZsl7ZUc1gA-JMz-CVYiv5wO4BLdSxemSYaSBohJI8",
		URL:        Message.HTTP,
		Data: Data{
			Username: Value{
				Value: Message.Name,
				Color: "#000000",
			},
			Message: Value{
				Value: Message.Message,
				Color: "#000000",
			},
			Status: Value{
				Value: Message.NowStatus,
				Color: "#FF0000",
			},
		},
	}
	marshalData, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	println(marshalData)
	resp, _ := http.Post("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+AccessToken, "application/json", bytes.NewBuffer(marshalData))
	content, _ := io.ReadAll(resp.Body)
	println(string(content))
	return resp
}
