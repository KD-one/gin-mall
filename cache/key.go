package cache

import "strconv"

const (
	RANKKEY = "rank"
)

// FormatProductViewKey 将商品id转换成"view:product:pid"特殊字符串
func FormatProductViewKey(pid uint) string {
	return "view:product:" + strconv.FormatUint(uint64(pid), 10)
}
