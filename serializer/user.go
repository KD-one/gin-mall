package serializer

import (
	"gin-mall/conf"
	"gin-mall/model"
)

// UserSerializer 定义给前端需要展示的user数据 VO(view object)
type UserSerializer struct {
	Id       uint   `json:"id"`
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
}

// BuildUser 创建并返回前端需要的User数据
func BuildUser(u *model.User) *UserSerializer {
	return &UserSerializer{
		Id:       u.ID,
		NickName: u.NickName,
		UserName: u.Name,
		Avatar:   conf.Host + conf.Port + conf.Avatar + u.Avatar,
		Email:    u.Email,
		Status:   u.Status,
	}
}
