package routes

import (
	"gin-mall/api/v1"
	"gin-mall/conf"
	"gin-mall/middleware"
	"github.com/gin-gonic/gin"
)

func NweRouter() {
	// 初始化路由
	r := gin.Default()

	// 跨域中间件
	r.Use(middleware.Cors())

	// 静态文件
	r.Static("/static", "../static")

	// 邮箱操作
	// 验证邮箱链接
	r.GET("user/validEmail/:token", api.ValidEmail)

	// v1路由分组
	v1 := r.Group("/api/v1")
	{
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 轮播图
		v1.GET("carousels", api.ShowCarousels)

		// 商品操作
		// 分页显示所有商品
		v1.GET("products", api.ShowProducts)
		// 根据商品名搜索商品
		v1.GET("product/search", api.SearchProduct)
		// 根据商品id显示商品
		v1.GET("product/:id", api.ShowProduct)
		// 根据商品id显示所有商品图片
		v1.GET("imgs/:id", api.ShowProductImg)
		// 展示所有商品分类
		v1.GET("categories", api.ShowCategories)

		autho := v1.Group("/")
		autho.Use(middleware.CheckToken)
		{
			// 用户操作
			autho.PUT("user/update", api.UserUpdate)
			autho.POST("avatar", api.AvatarUpdate)

			// 邮件操作
			// 发送邮件
			autho.POST("user/sendEmail", api.SendEmail)
			//autho.POST("user/validEmail/:token", api.ValidEmail)

			// 金钱操作
			// 显示用户金额
			autho.GET("money", api.ShowMoney)

			// 商品操作
			// 创建商品
			autho.POST("product/create", api.CreateProduct)

			// 收藏夹操作
			// 创建收藏夹
			autho.POST("favorite/create", api.CreateFavorite)
			autho.GET("favorites", api.ShowFavorites)
			autho.DELETE("favorite/:id", api.DeleteFavorite)

			// 地址操作
			// 创建地址
			autho.POST("address/create", api.CreateAddress)
			// 显示某个地址
			autho.GET("address/:id", api.ShowAddress)
			// 显示所有地址
			autho.GET("addresses", api.ShowAddresses)
			// 修改地址
			autho.PUT("address/:id", api.UpdateAddress)
			// 删除地址
			autho.DELETE("address/:id", api.DeleteAddress)

			// 购物车操作
			// 创建购物车记录
			autho.POST("cart/create", api.CreateCart)
			// 显示某个购物车记录
			autho.GET("cart/:id", api.ShowCart)
			// 显示所有购物车记录
			autho.GET("carts", api.ShowCarts)
			// 修改购物车记录
			autho.PUT("cart/:id", api.UpdateCart)
			// 删除购物车记录
			autho.DELETE("cart/:id", api.DeleteCart)

			// 订单操作
			// 创建订单
			autho.POST("order/create", api.CreateOrder)
			// 显示某个订单
			autho.GET("order/:id", api.ShowOrder)
			// 显示所有订单
			autho.GET("orders", api.ShowOrders)
			// 删除订单
			autho.DELETE("order/:id", api.DeleteOrder)

			// 订单支付
			autho.POST("order/pay", api.PayOrder)
		}
	}

	// 启动服务
	_ = r.Run(conf.Port)

}
