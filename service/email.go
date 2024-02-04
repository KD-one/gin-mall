package service

import (
	"gin-mall/conf"
	"gin-mall/dao"
	"gin-mall/pkg/util"
	"gin-mall/serializer"
	"github.com/gin-gonic/gin"
	"gopkg.in/mail.v2"
	"time"
)

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`             // 现阶段不去验证
	OperationType uint   `json:"operation_type" form:"operation_type"` // 1.绑定邮箱 2.解绑邮箱 3.修改密码
}

// Send 发送邮件
func (e *SendEmailService) Send(c *gin.Context) serializer.Response {
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
	}

	token, err := util.ReleaseEmailToken(uid, e.OperationType, e.Email, e.Password)
	if err != nil {
		return serializer.Response{
			Code: 400,
			Msg:  "发送失败",
			Data: nil,
		}
	}

	// 根据请求类型获取公告信息
	notice := dao.GetNoticeById(e.OperationType)

	address := conf.ValidEmail + token

	// 返回的邮件内容
	mailText := notice.Title + "<br><br>" + address

	// 设置邮件内容
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", e.Email)
	m.SetHeader("Subject", notice.Title)
	m.SetBody("text/html", mailText)

	// NewDialer 返回一个新的 SMTP 拨号器。给定的参数用于连接到 SMTP 服务器。
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)

	// StartTLSPolicy 表示与 SMTP 服务器通信所需的 TLS 安全级别。
	//默认为 OpportunisticStartTLS 以实现向后兼容性，但我们建议对所有现代 SMTP 服务器使用 MandatoryStartTLS。
	//如果 SSL 设置为 true，则此选项无效。
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// 通过 拨号器 d 的DialAndSend 方法连接到 SMTP 服务器并发送邮件内容 m 。
	if err := d.DialAndSend(m); err != nil {
		return serializer.Response{
			Code: 400,
			Msg:  "发送失败",
			Data: nil,
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "发送成功",
		Data: nil,
	}
}

type ValidEmailService struct {
}

// Valid 验证邮件
func (e *ValidEmailService) Valid(token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint

	// 验证token是否正确
	if token == "" {
		return serializer.Response{
			Code: 400,
			Msg:  "token为空",
			Data: nil,
		}
	}

	// 解析token 并验证token
	if _, claims, err := util.ParseEmailToken(token); err != nil {
		return serializer.Response{
			Code: 400,
			Msg:  "解析token失败",
			Data: nil,
		}
	} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
		return serializer.Response{
			Code: 400,
			Msg:  "token已过期",
			Data: nil,
		}
	} else {
		userId = claims.UserId
		email = claims.Email
		password = claims.Password
		operationType = claims.OperationType
	}

	// 获取用户信息
	user := dao.GetUserById(userId)

	// 根据不同的请求类型执行对应操作
	switch operationType {
	case 1:
		user.Email = email
	case 2:
		user.Email = ""
	case 3:
		user.Password = password
	}

	// 将用户信息更新到数据库
	if err := dao.UpdateUser(user); err != nil {
		return serializer.Response{
			Code: 400,
			Msg:  "更新用户信息失败",
			Data: nil,
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "验证成功",
		Data: serializer.BuildUser(user),
	}
}
