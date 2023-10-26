package wechat

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

var myOpenId string = "oQQWt53FZT7A8pmqb7KVhLt68AOo"
var usrOpenId string = ""
var appid string = "wx54539a45cc39782b"
var secret string = "a5591b49cef41cbe80a4197a6b5ef6cb"

type SubscribeBody struct {
	ToUserName string `xml:"ToUserName"`
	Openid     string `xml:"FromUserName"` //关注者
	MsgType    string `xml:"MsgType"`      //关注事件的话这一条的值是event
	Event      string `xml:"Event"`        //可能的值 subscribe unsubscribe 等等
	EventKey   string `xml:"EventKey"`     //一般为空?
	CreateTime string `xml:"CreateTime"`
	/*
		timestamp := int64(CreateTime)
		t := time.Unix(timestamp, 0)
		t 为 2023-10-18 08:14:50 UTC
	*/
}

// 回复用户的数据
type responseXML struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	MsgType      string `xml:"MsgType"`
	CreateTime   int64  `xml:"CreateTime"`
	Content      string `xml:"Content"`
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

func main() {
	//GetAccess_token()
	r := gin.Default()
	r.GET("/")
	r.POST("/", func(c *gin.Context) {
		body, err1 := io.ReadAll(c.Request.Body)
		if err1 != nil {
			println("Body读取错误")
		} else {
			println("#####body####" + string(body) + "#####body####")
		}
		var subscribe SubscribeBody
		//myOpenId = subscribe.ToUserName
		err := xml.Unmarshal(body, &subscribe)
		if err != nil {
			println("绑定错误")
			return
		} else {
			println("绑定成功")
		}
		switch subscribe.MsgType {
		case "event":
			switch subscribe.Event {
			case "subscribe":
				SubscribeHandler(c)
			case "unsubscribe":

			}

			break
		case "text":

		}
	})
	err := r.Run(":80")
	if err != nil {
		fmt.Println("error")
		return
	}
}

func SubscribeHandler(c *gin.Context) {
	rxml := responseXML{
		ToUserName:   myOpenId,
		FromUserName: usrOpenId,
		MsgType:      "text",
		CreateTime:   time.Now().Unix(),
		Content:      "关注福利:关上屏幕可以看到帅哥/美女",
	}
	c.XML(http.StatusOK, rxml)
}

//
//func TextHandler(c *gin.Context) {
//	rxml := responseXML{
//		ToUserName:   myOpenId,
//		FromUserName: usrOpenId,
//		MsgType:      "text",
//		CreateTime:   time.Now().Unix(),
//		Content:      ,
//	}
//	c.XML(http.StatusOK, rxml)
//}

func GetAccess_token() {
	var access_token string
	var expires_in int64 = 0 //access_token 凭证有效时间
	var checkTime int64      //开始检查时的时间
	checkTime = time.Now().Unix()
	for {
		if time.Now().Unix()-checkTime < expires_in-60 {
			time.Sleep(60 * time.Second) //60秒检测一次
			continue
		}
		get, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appid, secret))
		if err != nil {
			break
		}
		body, _ := io.ReadAll(get.Body)
		var data map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			break
		}
		access_token = data["access_token"].(string)
		println("凭证是" + string(access_token))
		expires_in = data["expires_in"].(int64)
		println("有效时间是" + string(expires_in))
		checkTime = time.Now().Unix()
	}
}
