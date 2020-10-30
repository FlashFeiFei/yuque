package request

import "net/http"

//选项
type Option func(client *Client)

//post、put请求的时候设置Content-Type 是 application/json 类型
func contentTypeOp() Option {

	return func(client *Client) {
		client.request.Header.Add("Content-Type", "application/json")
	}

}

//baseUri设置
func NewBaseUriOp(base_uri string) Option {

	return func(client *Client) {
		client.baseUri = base_uri
	}
}

//这是token
func NewTokenOp(token string) Option {
	return func(client *Client) {
		client.token = token
	}
}

//设置用户代理
func NewUserAgentOp(user_agent string) Option {
	return func(client *Client) {
		client.request.Header.Add("User-Agent", user_agent)
	}
}

//请求客户端
type Client struct {
	request *http.Request //http请求
	Op      []Option      //选项
	baseUri string        //baseuri
	token   string        //用户token
}

func (c *Client) Before() {

	switch c.request.Method {
	case http.MethodPut:
		c.Op = append(c.Op, contentTypeOp())
		break
	case http.MethodPost:
		c.Op = append(c.Op, contentTypeOp())
		break
	}

	//回调请求前回调下选项设计
	for _, op := range c.Op {
		op(c)
	}
	
}

func (c *Client) Request() {

}
