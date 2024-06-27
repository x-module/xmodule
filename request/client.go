/**
 * Created by Goland
 * @file   request.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/24 18:12
 * @desc   request.go
 */

package request

import (
	middleware2 "github.com/x-module/xmodule/request/middleware"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/middleware"
)

type Client struct {
	gentleman.Client
	isJson bool
}

func NewClient() *Client {
	request := &Client{}
	request.Context = context.New()
	request.Middleware = middleware.New()
	return request
}

// Debug 开启调试模式
func (r *Client) Debug(debug ...bool) *Client {
	debugMode := true
	if len(debug) > 0 {
		debugMode = debug[0]
	}
	if debugMode {
		r.UseResponse(middleware2.DebugMiddleware())
	}
	return r
}
