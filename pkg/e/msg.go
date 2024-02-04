package e

var m = map[int]string{
	ParamErr:      "请求参数错误",
	Success:       "请求成功",
	Unauthorized:  "请求未授权",
	Error:         "请求失败",
	NotFound:      "请求资源不存在",
	ErrUserExist:  "用户已存在",
	ErrCreateUser: "创建用户失败",
}

func GetMsg(code int) string {
	if v, ok := m[code]; ok {
		return v
	} else {
		return "未知错误"
	}
}
