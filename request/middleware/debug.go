/**
 * Created by Goland
 * @file   debug.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/25 13:38
 * @desc   debug.go
 */

package middleware

import (
	"bytes"
	"fmt"
	"github.com/x-module/helper/convertor"
	"gopkg.in/h2non/gentleman.v2/context"
	"io"
	"log"
	"strconv"
	"strings"
)

// getBody 获取请求体
func getBody(data io.ReadCloser) string {
	temp := strings.Split(fmt.Sprintf("%s", data), "[")
	if len(temp) > 1 {
		temp = strings.Split(temp[1], "]")
		if len(temp) > 1 {
			bits := strings.Split(temp[0], " ")
			var d []byte
			for _, v := range bits {
				if v != "" {
					b, _ := strconv.Atoi(v)
					d = append(d, byte(b))
				}
			}
			return string(d)
		} else {
			return "12"
		}
	}
	return "1"
}

// DebugMiddleware 调试中间件
func DebugMiddleware() context.HandlerFunc {
	return func(ctx *context.Context, h context.Handler) {
		log.Println("\n\n*--------------------------------------------------[Request Debug]--------------------------------------------------*")
		log.Printf("Request URL:%s \t\n", ctx.Request.URL)
		log.Printf("Request Headers:%s \t\n", convertor.ToJsonString(ctx.Request.Header))
		// ctx.Request.ContentLength
		log.Printf("Request ContentLength:%d \t\n", ctx.Request.ContentLength)
		log.Printf("Request Form:%s \t\n", ctx.Request.Form)
		if ctx.Request.Body != nil {
			log.Printf("Request Body:%s \t\n", getBody(ctx.Request.Body))
		}
		h.Next(ctx)
		if ctx.Response != nil {
			log.Printf("Response Status :%d \t\n", ctx.Response.StatusCode)
			log.Printf("Response Headers:%s \t\n", convertor.ToJsonString(ctx.Response.Header))
			if bodyBytes, err := io.ReadAll(ctx.Response.Body); err == nil {
				log.Printf("Response Body:%s \t\n", string(bodyBytes))
				ctx.Response.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Restore the request body
			}
		} else if ctx.Error != nil {
			log.Printf("Request Error:%s \t\n", ctx.Error.Error())
		}
		log.Printf("*--------------------------------------------------[Request Debug]--------------------------------------------------*\n\n")
	}
}
