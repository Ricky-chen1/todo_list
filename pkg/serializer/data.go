package serializer

//带有token的data
type TokenData struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var StatusMap = map[int]string{
	-1: "事项已置为待办",
	1:  "事项已置为已完成",
}

type StatusData struct {
	StatusMsg string `json:"status_msg"`
}

type Datalist struct {
	Items interface{}
	Total uint
}
