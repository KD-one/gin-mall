package service

import (
	"fmt"
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

type FavoriteService struct {
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	ProductId  uint `json:"product_id" form:"product_id"`
	BoosId     uint `json:"boos_id" form:"boos_id"`
	model.BasePage
}

// Create 创建收藏夹
func (p *FavoriteService) Create(c *gin.Context) serializer.Response {
	// 获取用户
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
	}

	user := dao.GetUserById(uid)
	product, err := dao.GetProductById(p.ProductId)
	if err != nil {
		return serializer.Response{
			Code: e.Error,
			Msg:  e.GetMsg(e.Error),
			Data: err,
		}
	}
	boos := dao.GetUserById(product.BoosId)

	favorite := model.Favorite{
		UserId:    uid,
		User:      *user,
		ProductId: p.ProductId,
		Product:   *product,
		BoosId:    product.BoosId,
		Boos:      *boos,
	}

	// 判断当前商品是否在收藏夹中已经存在
	if dao.ExistFavorite(&favorite) {
		return serializer.Response{
			Code: e.Error,
			Msg:  "当前商品已在收藏夹中！",
		}
	}

	// 创建收藏夹
	if err = dao.CreateFavorite(&favorite); err != nil {
		return serializer.Response{
			Code: e.Error,
			Msg:  e.GetMsg(e.Error),
			Data: err,
		}
	}

	return serializer.Response{
		Code: e.Success,
		Msg:  e.GetMsg(e.Success),
		Data: serializer.BuildFavorite(&favorite),
	}

}

// Show 获取当前用户的收藏夹列表
func (p *FavoriteService) Show(c *gin.Context) serializer.Response {
	// 获取当前用户
	var uid uint
	if id, ok := c.Get("userId"); ok {
		uid = id.(uint)
	}

	favorites, err := dao.GetFavoritesByUid(uid)
	if err != nil {
		return serializer.Response{
			Code: e.NotFound,
			Msg:  e.GetMsg(e.NotFound),
			Data: err,
		}
	}
	// 因为GetFavorites函数内没有进行分页操作，所以可以直接使用len(favorites)获取总数
	return serializer.BuildListResponse(serializer.BuildFavorites(favorites), uint(len(favorites)))

}

// Delete 根据收藏夹id删除收藏夹
func (p *FavoriteService) Delete(c *gin.Context, id string) serializer.Response {
	fid, _ := strconv.Atoi(id)

	err := dao.DeleteFavorite(uint(fid))
	if err != nil {
		return serializer.Response{
			Code: e.NotFound,
			Msg:  e.GetMsg(e.NotFound),
			Data: err,
		}
	}

	return serializer.Response{
		Code: e.Success,
		Msg:  fmt.Sprintf("id为：%d 的收藏夹数据删除成功", fid),
	}
}
