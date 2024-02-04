package serializer

import "gin-mall/model"

// Carousel 轮播图序列化器
type CarouselSerializer struct {
	Id        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductId uint   `json:"product_id"`
}

func BuildCarousel(item *model.Carousel) *CarouselSerializer {
	return &CarouselSerializer{
		Id:        item.ID,
		ImgPath:   item.ImgPath,
		ProductId: item.ProductId,
	}
}

func BuildCarousels(items []*model.Carousel) (carousels []*CarouselSerializer) {
	for _, item := range items {
		carousels = append(carousels, BuildCarousel(item))
	}
	return carousels
}
