package dao

import "gin-mall/model"

// migration 数据库自动迁移
func migration() {
	err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
		AutoMigrate(
			&model.User{},       // 用户
			&model.Product{},    // 商品
			&model.ProductImg{}, // 商品图片
			&model.Admin{},      // 管理员
			&model.Cart{},       // 购物车
			&model.Order{},      // 订单
			&model.Address{},    // 地址
			&model.Carousel{},   // 轮播图
			&model.Category{},   // 分类
			&model.Favorite{},   // 收藏夹
			&model.Notice{},     // 公告
		)
	if err != nil {
		panic(err)
	}
}
