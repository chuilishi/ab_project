package wechat

import (
	"ab_project/middle"
	"ab_project/model"
	"ab_project/service/response"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
	"time"
)

var myOpenId string = "gh_e7a9f5071ab9"
var usrOpenId string = "oQQWt53FZT7A8pmqb7KVhLt68AOo"

var appid string = "wx8ba1b60caf51ed26"

//var appid string = "wx54539a45cc39782b"

var secret string = "e1a9d9c66d49b7425ab0ec7d90635f4c"

// var secret string = "4cff84381199dda3c57badb5d5fa8f26"
var access_token string = "74_Z7RyL4Tpd1L1LOyKOj8mDS8zlECboNUGG-ki5l9c7oLwXphwbTYyn4sEIlzy9tQARSBz81Ph_4PqX3BdW8G_K768n8SEUNund9O10J76f2kdyeVAqE-q_6QFJLcWJYiAEAMZO"

// EventBody 微信所有事件(关注,消息等)的结构体
type EventBody struct {
	FromUserName string `xml:"ToUserName"` //用户
	Openid       string `xml:"FromUserName"`
	MsgType      string `xml:"MsgType"` //关注事件的话这一条的值是event
	Content      string `xml:"Msg"`
	Event        string `xml:"Event"`    //可能的值 subscribe unsubscribe 等等
	EventKey     string `xml:"EventKey"` //二维码携带的场景值
	CreateTime   string `xml:"CreateTime"`
	Ticket       string `xml:"Ticket"` //二维码对应ticket
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
	Content      string `xml:"Msg"`
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

// 带着ticket参数来获取openid
func openidHandler(c *gin.Context) {
	openidChan, exist := ticketToOpenId[c.Query("ticket")]
	if exist {
		openid := <-openidChan
		c.String(http.StatusOK, openid)
		println("#####成功#####")
		delete(ticketToOpenId, c.Query("ticket"))
	} else {
		code := c.Query("code")
		if code == "" {
			response.FailWithMessage("无效的ticket或者code", c)
		} else {
			mobileOpenidHandler(c)
		}
	}
}

// 移动端来获取openid
func mobileOpenidHandler(c *gin.Context) {
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

	//上面说的那个moblieAccesstoken
	mobileAccessToken = respjson["access_token"].(string)
	openid := respjson["openid"].(string)
	fmt.Println("$$$$$$$$" + openid + "$$$$$$$$")

	if err != nil {
		return
	}
	response.OkWithData(openid, c)
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
	//一直更新一个access_token (不然会过期)
	go GetAccessToken(false)
	r := gin.Default()
	r.Use(middle.Cors())
	r.GET("/", func(context *gin.Context) {
		response.OkWithMessage("Test", context)
	})
	r.GET("/qrcode", qrcodeHandler)
	r.GET("/openid", openidHandler)
	r.GET("/message", TemplateMessageHandler)
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
		if eventBody.Openid == "" {
			return
		} //不是微信方发来的消息退出
		switch eventBody.MsgType {
		case "event":
			switch eventBody.Event {
			case "subscribe":
				SendTemplateMessage(model.TemplateMessage{
					WxOpenId:  eventBody.Openid,
					Name:      "从数据库拿过来",
					Msg:       "test",
					NowStatus: "test状态",
					HTTP:      "http://123.207.73.185:100?id=" + eventBody.Openid,
				})
				//Reply(c, "achobeta,启动!", eventBody.Openid)
				//if eventBody.Ticket != "" { //登录事件
				//	ch, exist := ticketToOpenId[eventBody.Ticket]
				//	if !exist {
				//		ticketToOpenId[eventBody.Ticket] = make(chan string)
				//	}
				//	ch <- eventBody.Openid
				//}
				break
			case "unsubscribe":
				break
			case "SCAN":
				SendTemplateMessage(model.TemplateMessage{
					WxOpenId:  eventBody.Openid,
					Name:      "从数据库拿",
					Msg:       "test",
					NowStatus: "test状态",
					HTTP:      "http://123.207.73.185:100?id=" + eventBody.Openid,
				})
				//if eventBody.Ticket != "" { //登录事件
				//	ch, exist := ticketToOpenId[eventBody.Ticket]
				//	if !exist {
				//		ticketToOpenId[eventBody.Ticket] = make(chan string)
				//	}
				//	ch <- eventBody.Openid
				//}
				break
			}
			break
		case "text":
			println("发送的是消息")
			Reply(c, eventBody.Content+"当然是正确的,"+"你说的对，但是一小时有60分钟，一分钟有60秒，3600个一秒可以组成一小时。这些你都知道。你甚至知道，地球是圆的，太阳不是宇宙的中心，银河系也不是宇宙唯一的星系。如此高深的知识，你都知道。可是你不知道，每一秒，每一分钟我都在想着你。我打开手机，打开ipad，第一眼看到的就是你。可你不知道。你上知天文，下知地理，通晓时空。可你不知道我对你的心。一年有31536000秒，这你一定知道。你不知道的是，每一秒我心里都有你的位置。你知道宇宙的中心不是太阳，却又万万不知道在我心里，宇宙的中心是你。你知道整个世界，可你不知道我对你的心意!!!", eventBody.Openid)
			break
		}
	})
	err := r.Run(":80")
	if err != nil {
		fmt.Println("error")
		return
	}
}

// 回复消息
func Reply(c *gin.Context, message string, openid string) {
	rxml := responseXML{
		ToUserName:   openid,
		FromUserName: myOpenId,
		MsgType:      "text",
		CreateTime:   time.Now().Unix(),
		Content:      message,
	}
	c.XML(http.StatusOK, rxml)
}

func GetAccessToken(once bool) string {
	if once {
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
		return access_token
	}

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

// 模板消息的handler
func TemplateMessageHandler(c *gin.Context) {
	TemplateData := model.TemplateMessage{}
	err := c.ShouldBind(&TemplateData)
	if err != nil {
		fmt.Sprintf("#####%v#####", err)
		return
	}
	SendTemplateMessage(TemplateData)
	response.OkWithMessage("已接收", c)
}
