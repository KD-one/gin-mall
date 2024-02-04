package service

import (
	"gin-mall/conf"
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util"
	"gin-mall/serializer"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type UserService struct {
	NickName  string `json:"nick_name" form:"nick_name"`
	UserName  string `json:"user_name" form:"user_name"`
	Password  string `json:"password" form:"password"`
	SecretKey string `json:"secret_key" form:"secret_key"` // 现阶段不去验证
}

// Register 注册
func (u *UserService) Register() serializer.Response {
	var user model.User
	//if u.SecretKey == "" || len(u.SecretKey) < 16 {
	//	code := e.Error
	//	return serializer.Response{
	//		Code: code,
	//		Msg:  e.GetMsg(code),
	//		Data: "密钥不合法",
	//	}
	//}

	// 验证用户名是否存在
	_, exist, err := dao.CheckUserName(u.UserName)
	if err != nil {
		return serializer.Response{
			Code: e.Error,
			Msg:  e.GetMsg(e.Error),
			Data: err,
		}
	}
	// 用户名存在
	if exist {
		return serializer.Response{
			Code: e.ErrUserExist,
			Msg:  e.GetMsg(e.ErrUserExist),
		}
	}

	// 加密密码
	encryptPassword, err := util.EncryptAES(u.Password)
	if err != nil {
		return serializer.Response{
			Code: e.Error,
			Msg:  e.GetMsg(e.Error),
			Data: err,
		}
	}

	// 填写用户信息
	user = model.User{
		NickName: u.NickName,
		Name:     u.UserName,
		Password: encryptPassword,
		Avatar:   "avatar.jpg",
		Status:   1,
		Money:    "100",
	}

	// 创建用户
	err = dao.CreateUser(&user)
	if err != nil {
		return serializer.Response{
			Code: e.ErrCreateUser,
			Msg:  e.GetMsg(e.ErrCreateUser),
			Data: err,
		}
	}

	return serializer.Response{
		Code: e.Success,
		Msg:  e.GetMsg(e.Success),
		Data: "注册成功",
	}
}

// Login 登录
func (u *UserService) Login() serializer.Response {
	// 验证用户名是否存在
	user, exist, err := dao.CheckUserName(u.UserName)
	if err != nil {
		return serializer.Response{
			Code: e.Error,
			Msg:  e.GetMsg(e.Error),
			Data: err,
		}
	}

	// 用户名不存在
	if !exist {
		return serializer.Response{
			Code: e.Error,
			Msg:  e.GetMsg(e.Error),
			Data: "用户不存在",
		}
	}

	// 加密密码
	encryptPassword, err := util.EncryptAES(u.Password)
	if err != nil {
		return serializer.Response{
			Code: e.Error,
			Msg:  e.GetMsg(e.Error),
			Data: err,
		}
	}

	// 验证密码
	_, exist, err = dao.CheckUserPassword(u.UserName, encryptPassword)
	if err != nil {
		return serializer.Response{
			Code: e.Error,
			Msg:  e.GetMsg(e.Error),
			Data: err,
		}
	}

	// 密码错误
	if !exist {
		return serializer.Response{
			Code: e.Error,
			Msg:  e.GetMsg(e.Error),
			Data: "密码错误",
		}
	}

	// 发放token
	tokenString, err := util.ReleaseToken(user.ID)
	if err != nil {
		return serializer.Response{
			Code: e.Error,
			Msg:  e.GetMsg(e.Error),
			Data: "token发放失败！！！",
		}
	}
	// 登录成功
	// 返回给前端其需要的信息
	return serializer.Response{
		Code: e.Success,
		Msg:  e.GetMsg(e.Success),
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: tokenString,
		},
	}
}

// Update 修改用户信息
func (u *UserService) Update(c *gin.Context) serializer.Response {
	// 获取当前登录用户（登陆时在上下文中设置了("userId", user.ID)键值对）
	var user *model.User
	if uid, ok := c.Get("userId"); ok {
		user = dao.GetUserById(uid.(uint))
	}
	// TODO: 修改用户目前只能修改nickName昵称
	if u.NickName != "" {
		user.NickName = u.NickName
	}
	// 在数据库中修改用户修改了的内容
	if err := dao.UpdateUser(user); err != nil {
		return serializer.Response{
			Code: e.Error,
			Msg:  e.GetMsg(e.Error),
			Data: err,
		}
	}
	return serializer.Response{
		Code: e.Success,
		Msg:  e.GetMsg(e.Success),
		Data: "修改成功",
	}
}

// AvatarUpdate 修改用户头像
func (u *UserService) AvatarUpdate(c *gin.Context, avatar *multipart.FileHeader) serializer.Response {
	// 获取当前登录用户（登陆时在上下文中设置了("userId", user.ID)键值对）
	var user *model.User
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
		user = dao.GetUserById(uid)
	}

	// 将上传的头像保存本地并返回路径
	path, err := FileUpload(c, avatar, uid, conf.Avatar, "user")
	if err != nil {
		return serializer.Response{
			Code: e.Error,
			Msg:  e.GetMsg(e.Error),
			Data: err,
		}
	}
	// 修改用户头像
	user.Avatar = path

	// 在数据库中修改用户修改了的内容
	if err := dao.UpdateUser(user); err != nil {
		return serializer.Response{
			Code: e.Error,
			Msg:  e.GetMsg(e.Error),
			Data: err,
		}
	}
	return serializer.Response{
		Code: e.Success,
		Msg:  e.GetMsg(e.Success),
		Data: "修改成功",
	}
}
