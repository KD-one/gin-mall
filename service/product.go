package service

import (
	"fmt"
	"gin-mall/conf"
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/serializer"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strconv"
)

type ProductService struct {
	Id             uint   `json:"id" form:"id"`
	Name           string `json:"name" form:"name"`
	Title          string `json:"title" form:"title"`
	Info           string `json:"info" form:"info"`
	Price          string `json:"price" form:"price"`
	Category       string `json:"category" form:"category"`
	ImgPath        string `json:"img_path" form:"img_path"`
	DiscountPrice  string `json:"discount_price" form:"discount_price"`
	OnSale         bool   `json:"on_sale" form:"on_sale"` // 是否上架
	Num            int    `json:"num" form:"num"`
	model.BasePage        // 继承分页功能
}

// Create 创建商品（以第一张上传的图片作为商品默认图片）
func (p *ProductService) Create(c *gin.Context, files []*multipart.FileHeader) serializer.Response {
	// 获取用户
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
	}
	user := dao.GetUserById(uid)

	// 将首张图片作为商品图片
	path, err := FileUpload(c, files[0], uid, conf.Product, "product")
	if err != nil {
		return serializer.Response{
			Code: 400,
			Msg:  "上传失败",
		}
	}

	// 将商品信息存入数据库
	product := model.Product{
		Name:          p.Name,
		Category:      p.Category,
		Title:         p.Title,
		Info:          p.Info,
		ImgPath:       path,
		Price:         p.Price,
		DiscountPrice: p.DiscountPrice,
		OnSale:        true,
		Num:           p.Num,
		BoosId:        uid, // 商家id
		BoosName:      user.Name,
		BoosAvatar:    user.Avatar,
	}
	if err = dao.CreateProduct(&product); err != nil {
		return serializer.Response{
			Code: 400,
			Msg:  "创建失败",
		}
	}

	// 处理除首张外的其余图片
	for i, file := range files {

		// 将图片上传到服务器静态资源路径下
		path, err = FileUpload(c, file, uid, conf.Product, "product")
		if err != nil {
			return serializer.Response{
				Code: 400,
				Msg:  fmt.Sprintf("上传失败，第%d张图片: %s", i+1, file.Filename),
			}
		}

		// 将图片信息存入数据库
		productImg := model.ProductImg{
			// product没有初始化id，默认使用当前用户id
			ProductId: product.ID,
			ImgPath:   path,
		}
		if err = dao.CreateProductImg(&productImg); err != nil {
			return serializer.Response{
				Code: 400,
				Msg:  fmt.Sprintf("文件存入数据库失败！，文件名: %s", file.Filename),
			}
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "创建成功",
		Data: serializer.BuildProduct(&product),
	}
}

// Show 根据商品分类获取商品列表（没填写分类则展示所有商品）
func (p *ProductService) Show(c *gin.Context) serializer.Response {
	var products []*model.Product

	// 默认每页展示15条数据
	if p.PageSize == 0 {
		p.PageSize = 15
	}

	//  将商品根据分类查询
	condition := make(map[string]interface{})

	// 如果category为空，where中传入的map为空map，则此时条件为空，将查询所有商品
	if p.Category != "" {
		condition["category"] = p.Category
	}

	// 根据condition条件查询商品总数
	total, err := dao.GetProductCountByCondition(condition)
	if err != nil {
		return serializer.Response{
			Code: 400,
			Msg:  "查询商品数据库失败",
		}
	}

	// 根据condition条件分页查询商品
	products, err = dao.GetProductsByCondition(condition, p.BasePage)
	if err != nil {
		return serializer.Response{
			Code: 400,
			Msg:  "查询商品失败",
		}
	}

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))

}

// Search 目前是根据商品名查询相关的所有商品
func (p *ProductService) Search(c *gin.Context) serializer.Response {

	if p.PageSize == 0 {
		p.PageSize = 15
	}

	// 根据商品名查询商品总数
	total, err := dao.GetProductCountByName(p.Name)
	if err != nil {
		return serializer.Response{
			Code: 400,
			Msg:  "查询商品名查询商品数量失败！",
		}
	}

	// 根据商品名查询商品
	products, err := dao.GetProductsByName(p.Name, p.BasePage)
	if err != nil {
		return serializer.Response{
			Code: 400,
			Msg:  "查询商品失败！",
		}
	}

	// 最后返回的total为什么不能使用len(products)？
	// 1、并发问题： 如果在多线程或多请求情况下，其他请求同时修改了数据库中的商品记录数，而这里的查询和统计不是在一个事务内完成的，那么返回的商品总数可能与实际数据库中的商品数量有差异。
	// 2、分页查询与总数不一致： 如果SearchProductByName函数内部实现了分页查询逻辑（例如：使用LIMIT和OFFSET关键字），并且没有同时查询出总记录数，那么返回的uint(len(products))只是当前分页内的商品数量，而不是数据库中的商品总数。
	// 3、数据更新导致的不一致： 在执行查询和计算长度之间，如果有其他进程或事务删除或新增了商品记录，也可能导致返回的商品总数与查询时的实际记录数不符。
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

// ProductInfo 根据商品id查询商品信息
func (p *ProductService) ProductInfo(c *gin.Context, id string) serializer.Response {
	// 获取商品id
	pid, _ := strconv.Atoi(id)

	// 根据商品id查询商品
	product, err := dao.GetProductById(uint(pid))
	if err != nil {
		return serializer.Response{
			Code: 400,
			Msg:  "查询商品失败",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "获取用户信息成功",
		Data: serializer.BuildProduct(product),
	}
}
