package request

import (
	"fmt"
	"net/http"
)


//通过login获取用户信息
func NewUserRquestByLogin(token,login string) *ClientHandle{
	return  NewClientHandle(token,fmt.Sprintf("/users/%s",login),http.MethodGet,nil)
}

//通过id获取用户信息
func NewUserRequestById(token,id string) *ClientHandle{
	return  NewClientHandle(token,fmt.Sprintf("/users/%s",id),http.MethodGet,nil)
}


