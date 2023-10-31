package wechat

import (
	"net/http"
	"strings"
)

func SendTemplateMessage(openid string, accessToken string) *http.Response {
	username := "chuilishi"
	templateId := "osNMd68bPwCn7FRP-NISWGwg0Ybk"
	jsonData := `{
           touser:` + openid + `,
           template_id:` + templateId + `,
           url:` + `https://www.bilibili.com/video/BV1GJ411x7h7` + `,
           data:{
			   username:{
					value:` + username + `
			   },
			   isSuccess;{
					value:` + `成功` + `
			}
		}
   }`
	resp, _ := http.Post("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+accessToken, "application/json", strings.NewReader(jsonData))
	return resp
}
