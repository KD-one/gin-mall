package serializer

import "gin-mall/model"

type MoneySerializer struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(user *model.User, key string) *MoneySerializer {
	return &MoneySerializer{
		UserId:    user.ID,
		UserName:  user.Name,
		UserMoney: user.Money,
	}
}
