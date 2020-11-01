package request

//请求流程
type RequestHandle interface {
	Request(response interface{}) error  //请求
}


func Request(request_handle RequestHandle,response interface{}) {
	request_handle.Request(response)
}
