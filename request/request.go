package request

//请求流程
type RequestHandle interface {
	Before()  //请求之前
	Request() //请求
}

func Request(request_handle RequestHandle) {
	request_handle.Before()
	request_handle.Request()
}
