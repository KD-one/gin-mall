package serializer

// serializer 序列化器

// Response 通用Response结构体
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

// TokenData 通用tokenData结构体
type TokenData struct {
	User  interface{} `json:"user_id"`
	Token string      `json:"token"`
}

type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

func BuildListResponse(item interface{}, total uint) Response {
	return Response{
		Code: 200,
		Data: DataList{
			Items: item,
			Total: total,
		},
		Msg: "ok",
	}
}
