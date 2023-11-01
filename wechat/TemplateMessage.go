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

func SendTemplateMessage(message model.TemplateMessage, accessToken string) *http.Response {
	data := TemplateData{
		ToUser: message.WxOpenId,
		//模板id
		TemplateID: "wjDMuDWaEcbc3Woq2igqPAwgO4lnJlrRES7CcJFml8g",
		URL:        message.HTTP,
		Data: Data{
			Username: Value{
				Value: message.Name,
				Color: "#000000",
			},
			Message: Value{
				Value: message.Message,
				Color: "#000000",
			},
			Status: Value{
				Value: message.NowStatus,
				Color: "#FF0000",
			},
		},
	}
	marshalData, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	resp, _ := http.Post("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+accessToken, "application/json", bytes.NewBuffer(marshalData))
	content, _ := io.ReadAll(resp.Body)
	println(string(content))
	return resp
}
