package request

import (
	"fmt"
	"net/http"
)

type UserRequest struct {
	AuthToken
}

func (r *UserRequest) NewUserRequestByLogin(login string) *ClientHandle {
	return NewClientHandle(r.AuthToken.Token, fmt.Sprintf("/users/%s", login), http.MethodGet, nil)
}

//通过id获取用户信息
func (r *UserRequest) NewUserRequestById(id string) *ClientHandle {
	return NewClientHandle(r.AuthToken.Token, fmt.Sprintf("/users/%s", id), http.MethodGet, nil)
}
