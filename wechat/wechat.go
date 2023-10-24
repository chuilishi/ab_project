package wechat

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
	"time"
)

var myOpenId string = "oQQWt53FZT7A8pmqb7KVhLt68AOo"
var usrOpenId string = "oQQWt53FZT7A8pmqb7KVhLt68AOo"
var appid string = "wx8ba1b60caf51ed26"
var secret string = "e1a9d9c66d49b7425ab0ec7d90635f4c"
var access_token string = "73_povW6cXWc6rITF10Pk53AfYAUj09bXvNcJdcqoy01jOgv5TddZjPXODmWXIYkEwA00M7Bl-zcgRvZGAGX-VY-KYCbHvN5toXv5icmaWvS2lybBCU879jd_RYAxIFFDaABATTR"

// EventBody 微信所有事件(关注,消息等)的结构体
type EventBody struct {
	ToUserName string `xml:"ToUserName"`
	Openid     string `xml:"FromUserName"` //关注者
	MsgType    string `xml:"MsgType"`      //关注事件的话这一条的值是event
	Content    string `xml:"Content"`
	Event      string `xml:"Event"`    //可能的值 subscribe unsubscribe 等等
	EventKey   string `xml:"EventKey"` //二维码携带的场景值
	CreateTime string `xml:"CreateTime"`
	Ticket     string `xml:"Ticket"` //二维码对应ticket
	/*
		timestamp := int64(CreateTime)
		t := time.Unix(timestamp, 0)
		t 为 2023-10-18 08:14:50 UTC
	*/
}

// ticket对应的openid管道
var ticketToOpenId = make(map[string]chan string)

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

// 带着ticket参数来获取openid
func openidHandler(c *gin.Context) {
	openidChan, exist := ticketToOpenId[c.Query("ticket")]
	if exist {
		openid := <-openidChan
		c.String(http.StatusOK, openid)
		delete(ticketToOpenId, c.Query("ticket"))
	} else {
		c.String(http.StatusOK, "") //一直都没扫返回空字符串
	}
}

// 请求qrcode之后返回一个json
func qrcodeHandler(c *gin.Context) {
	url := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + access_token
	jsondata := `{
        "action_name": "QR_LIMIT_SCENE",
        "action_info": {
            "scene": {"scene_id": 123}
        }
    }`
	resp, err := http.Post(url, "application/json", strings.NewReader(jsondata))
	if err != nil {
		return
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	respjson := make(map[string]interface{})
	json.Unmarshal(all, &respjson)
	ticketToOpenId[respjson["ticket"].(string)] = make(chan string)
	c.JSON(http.StatusOK, map[string]string{
		"url":    "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + respjson["ticket"].(string),
		"ticket": respjson["ticket"].(string),
	})
}
func Wechat() {
	go GetAccessToken()
	r := gin.Default()
	r.GET("/")
	r.GET("/qrcode", qrcodeHandler)
	r.GET("/openid", openidHandler)
	r.POST("/", func(c *gin.Context) {
		body, err1 := io.ReadAll(c.Request.Body)
		if err1 != nil {
			println("Body读取错误")
		} else {
			println("#####body####" + string(body) + "#####body####")
		}
		var eventBody EventBody
		//myOpenId = eventBody.ToUserName
		err := xml.Unmarshal(body, &eventBody)
		if err != nil {
			println("绑定错误")
			return
		} else {
			println("绑定成功")
		}
		if eventBody.ToUserName == "" {
			return
		} //不是微信方发来的消息退出
		switch eventBody.MsgType {
		case "event":
			switch eventBody.Event {
			case "subscribe":
				Reply(c, "achobeta,启动!")
				if eventBody.Ticket != "" { //登录事件
					ch, exist := ticketToOpenId[eventBody.Ticket]
					if !exist {
						ticketToOpenId[eventBody.Ticket] = make(chan string)
					}
					ch <- eventBody.Openid
				}
				break
			case "unsubscribe":
				break
			case "SCAN":
				Reply(c, "欢迎回来")
				if eventBody.Ticket != "" { //登录事件
					ch, exist := ticketToOpenId[eventBody.Ticket]
					if !exist {
						ticketToOpenId[eventBody.Ticket] = make(chan string)
					}
					ch <- eventBody.Openid
				}
				break
			}
			break
		case "text":
			Reply(c, "为什么你要发送"+eventBody.Content)
			break
		}
	})
	err := r.Run(":80")
	if err != nil {
		fmt.Println("error")
		return
	}
}

// 回复
func Reply(c *gin.Context, message string) {
	rxml := responseXML{
		ToUserName:   myOpenId,
		FromUserName: usrOpenId,
		MsgType:      "text",
		CreateTime:   time.Now().Unix(),
		Content:      message,
	}
	c.XML(http.StatusOK, rxml)
}

//	func TextHandler(c *gin.Context) {
//		rxml := responseXML{
//			ToUserName:   myOpenId,
//			FromUserName: usrOpenId,
//			MsgType:      "text",
//			CreateTime:   time.Now().Unix(),
//			Content:      ,
//		}
//		c.XML(http.StatusOK, rxml)
//	}
func GetAccessToken() {
	for {
		get, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appid, secret))
		if err != nil {
			println("获取Access_token 异常")
		}
		body, _ := io.ReadAll(get.Body)
		var data map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			println("获取Access_token 异常")
		}
		access_token = data["access_token"].(string)
		println("凭证是" + string(access_token))
		time.Sleep(7000 * time.Second)
	}
}
