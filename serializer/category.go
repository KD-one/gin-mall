package serializer

import "gin-mall/model"

// CategorySerializer 轮播图序列化器
type CategorySerializer struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func BuildCategory(item *model.Category) *CategorySerializer {
	return &CategorySerializer{
		Id:   item.ID,
		Name: item.Name,
	}
}

func BuildCategorys(items []*model.Category) (categorys []*CategorySerializer) {
	for _, item := range items {
		categorys = append(categorys, BuildCategory(item))
	}
	return categorys
}
