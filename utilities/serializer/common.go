package serializer

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Err    error       `json:"err"`
}

func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: Datalist{
			Items: items,
			Total: total,
		},
		Msg: "ok",
	}
}
