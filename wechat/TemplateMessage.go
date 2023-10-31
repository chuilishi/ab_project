package wechat

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type TemplateData struct {
	ToUser     string `json:"touser"`
	TemplateID string `json:"template_id"`
	URL        string `json:"url"`
	Data       Data   `json:"data"`
}

type Data struct {
	Username  Value `json:"username"`
	IsSuccess Value `json:"isSuccess"`
}

type Value struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

func SendTemplateMessage(usropenid string, accessToken string) *http.Response {
	//username := "chuilishi"
	//templateId := ""
	//jsonData := `{
	//        touser:` + usropenid + `,
	//        template_id:` + templateId + `,
	//        url:` + `https://www.bilibili.com/video/BV1GJ411x7h7` + `,
	//        data:{
	//		   username:{
	//				value:` + username + `
	//		   },
	//		   isSuccess;{
	//				value:` + `成功` + `
	//		}
	//	}
	//}`
	data := TemplateData{
		ToUser:     usropenid,
		TemplateID: "bpv27yYOvMyjSy4s6aok9ebm3zfxG49IF7PpqYsMZ4o",
		URL:        "https://www.bilibili.com/video/BV1GJ411x7h7",
		Data: Data{
			Username: Value{
				Value: "chuilishi",
				Color: "#FF0000",
			},
			IsSuccess: Value{
				Value: "成功",
				Color: "#FF0000",
			},
		},
	}
	marshalData, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	println(string(marshalData))
	resp, _ := http.Post("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+accessToken, "application/json", bytes.NewBuffer(marshalData))
	return resp
}
