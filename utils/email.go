package utils

import (
	"ab_project/global"
	mailer "gopkg.in/gomail.v2"
)

/**
to 主送 机构联系邮箱
cc 抄送  创建人邮箱
subject 标题
content 内容
result 1成功 2失败

host 邮件服务器
port 端口号
username 邮箱
password 授权码
*/

func SendMail(to, username, problem string) error {
	// 1. 首先构建一个 Message 对象，也就是邮件对象
	msg := mailer.NewMessage()
	// 2. 填充 From，注意第一个字母要大写
	msg.SetHeader("From", global.MailFrom)
	// 3. 填充 To
	msg.SetHeader("To", to)
	// 4. 如果需要可以填充 cc，也就是抄送
	//msg.SetHeader("Cc", "cc_address@example.com")
	// 5. 设置邮件标题
	msg.SetHeader("Subject", "异常情况，请处理！！！")
	// 6. 设置要发送的邮件正文
	// 第一个参数是类型，第二个参数是内容
	// 如果是 html，第一个参数则是 `text/html` 如果是文本则是"text/plain"
	msg.SetBody("text/plain", "用户"+username+"有异常情况，需要你处理！！！\n异常信息为"+problem)
	// 7. 添加附件，注意，这个附件是完整路径
	// msg.Attach("/Users/yufei/Downloads/1.jpg")
	// 到此，邮件消息构建完毕

	// 8. 创建 smtp 实例
	// 如果你的阿里云企业邮箱则是密码，否则一般情况下国内国外使用的都是授权码(例如，腾讯云企业邮箱)
	// 请注意 DialAndSend() 方法是一次性的，也就是连接邮件服务器，发送邮件，然后关闭连接。
	// dialer := mailer.NewDialer("smtp.mxhichina.com", 465, "阿里云企业邮箱账号", "阿里云企业邮箱密码")
	dialer := mailer.NewDialer("smtp.163.com", 465, "djy1349678887@163.com", "OAKNCZNXWJCWADWJ")
	// 9. 发送邮件，连接邮件服务器，发送完就关闭
	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}
