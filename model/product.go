package model

import (
	"context"
	"gin-mall/cache"
	"gorm.io/gorm"
	"strconv"
)

// Product 商品
type Product struct {
	gorm.Model
	Name          string `json:"name" gorm:"type:varchar(20) not null"`
	Category      string `json:"category" gorm:"type:varchar(20) not null"`
	Title         string `json:"title" gorm:"type:varchar(20) not null"`
	Info          string `json:"info" gorm:"type:varchar(225) not null"`
	ImgPath       string `json:"img_path" gorm:"type:varchar(225) not null"`
	Price         string `json:"price" gorm:"type:varchar(20) not null"`
	DiscountPrice string `json:"discount_price" gorm:"type:varchar(20) not null"`
	OnSale        bool   `json:"on_sale" gorm:"type:bool not null"`             // 是否在售
	Num           int    `json:"num" gorm:"type:int(11) not null"`              // 数量
	BoosId        uint   `json:"boos_id" gorm:"type:int(11) not null"`          // 商家id
	BoosName      string `json:"boos_name" gorm:"type:varchar(20) not null"`    // 商家名字
	BoosAvatar    string `json:"boos_avatar" gorm:"type:varchar(225) not null"` // 商家头像
}

// View 通过Redis获取商品浏览量
func (p *Product) View() uint64 {
	result, err := cache.RedisClient.Get(context.TODO(), cache.FormatProductViewKey(p.ID)).Result()
	if err != nil {
		return 0
	}
	count, _ := strconv.ParseUint(result, 10, 64)
	return count
}

// AddView 增加商品浏览量
func (p *Product) AddView() {
	// Incr 命令：增加商品浏览量
	// 1、如果key不存在，那么Redis会先创建该键，并将其值初始化为0，然后执行自增操作，结果就是1。
	// 2、如果key存在并且其值可以被解释为十进制整数，那么它的值将增加1。
	// 3、如果key存储的不是整数值或者无法表示为数字，INCR命令会返回一个错误。
	cache.RedisClient.Incr(context.TODO(), cache.FormatProductViewKey(p.ID))

	// 商品浏览量排名增加
	// ZINCRBY key increment member
	// -key：表示有序集合的键名。
	// -increment：要增加的浮点数值。
	// -member：有序集合中的成员。
	// 此命令执行后：
	// 1、如果key不存在，则创建一个新的有序集合，并添加member作为其成员，其分数为给定的increment值。
	// 2、如果member已经存在于有序集合中，则其现有的分数会增加指定的increment值。
	cache.RedisClient.ZIncrBy(context.TODO(), cache.RANKKEY, 1, strconv.Itoa(int(p.ID)))
}
