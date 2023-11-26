package wechat

import (
	"ab_project/global"
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
	Message  Value `json:"msg"`
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
		Msg:   "",
		NowStatus: "",
		HTTP:      "",
	}, wechat.GetAccessToken(true))
*/

func SendTemplateMessage(Message model.TemplateMessage) error {
	templateId := ""
	switch Message.Code {
	case 0:
		//日程安排
		templateId = "q9yRRqUZL9Mo_gkgfH3PVCmFNbCg1w7towmhE1rPl38"
	case 1:
		//面试结果
		templateId = "K2QX9SUT_XivtOD60DeTMDv3W9f4EwmtKEo2Y11B3W0"
	case 2:
		//面试会议
		templateId = "Hww8HsZFaOI0rU1Uaf0l9x6X9JgPOiD3TYM5AT5qdgs"
	}
	data := TemplateData{
		ToUser: Message.WxOpenId,
		//模板id
		TemplateID: templateId,
		URL:        Message.HTTP,
		Data: Data{
			Username: Value{
				Value: Message.Name,
				Color: "#FF0000",
			},
			Message: Value{
				Value: Message.Msg,
				Color: "#FF0000",
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
	resp, err := http.Post("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+global.AccessToken, "application/json", bytes.NewBuffer(marshalData))
	content, _ := io.ReadAll(resp.Body)
	println(string(content))
	return err
}
