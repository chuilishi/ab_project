package wechat

import (
	"ab_project/service/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// 这个access_token是移动端的那个专用access_token,与wechat中那个不同
var mobileAccessToken = ""

// 移动端登录的时候给前端返回的信息
type Userinfo struct {
	Openid string `json:"openid"`
	//移动端是可以获取微信相关信息的,头像,名称等等
	//Nickname   string `json:"nickname"`
	//Sex        string `json:"sex"`
	//Country    string `json:"country"`
	//Headimgurl string `json:"headimgurl"`
}

func WechatLogin(c *gin.Context) {
	code := c.Query("code")
	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", appid, secret, code))
	if err != nil {
		return
	}
	respjson := make(map[string]interface{})
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	json.Unmarshal(all, &respjson)

	//上面说的那个access_token
	mobileAccessToken = respjson["access_token"].(string)
	openid := respjson["openid"].(string)
	marshal, err := json.Marshal(Userinfo{
		Openid: openid,
	})
	if err != nil {
		return
	}
	response.OkWithData(marshal, c)
}

/*
https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx807d86fb6b3d4fd2&redirect_uri=http%3A%2F%2Fdevelopers.weixin.qq.com&response_type=code&scope=snsapi_userinfo#wechat_redirect
*/

/*
{
  "access_token":"ACCESS_TOKEN",
  "expires_in":7200,
  "refresh_token":"REFRESH_TOKEN",
  "openid":"OPENID",
  "scope":"SCOPE",
  "is_snapshotuser": 1,
  "unionid": "UNIONID"
}
*/
