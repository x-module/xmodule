/**
 * Created by Goland
 * @file   request_test.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/27 20:12
 * @desc   request_test.go
 */

package request

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/middleware"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/plugins/cookies"
	url2 "gopkg.in/h2non/gentleman.v2/plugins/url"
	"net/url"
	"testing"
)

func TestOne(T *testing.T) {
	// Create a new client
	cli := gentleman.New()
	// Define cookies

	// Perform the request
	res, err := cli.Use(cookies.Set("showdown_console", "user_cookie")).Use(cookies.Jar()).Request().URL("http://ec2-34-201-82-238.compute-1.amazonaws.com:9000").Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}
	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s", res.String())
}
func TestTwo(T *testing.T) {
	cli := NewClient()
	cli.Debug(true)
	// cli.Use(logger.New(os.Stdout))
	// cli.Use(cookies.Set("showdown_console", "user_cookie")).Use(cookies.Jar())
	req := cli.Request().URL("https://showdown.us-east1.nakamacloud.io/v2/console/authenticate").Method("POST")
	formData := url.Values{}
	formData.Set("name", "bar")
	formData.Set("age", "qux")
	req.Use(body.String("name=bar&age=111"))
	req.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	// req.Use(body.JSON(map[string]string{"name": "showdow333n"}))
	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}
	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s", res.String())
}

func TestThree(T *testing.T) {
	// 创建一个新的客户端实例
	client := gentleman.New()
	client.UseRequest(func(ctx *context.Context, h context.Handler) {
		fmt.Println("Request URL:", ctx.Request.URL)
		fmt.Println("Request Headers:", ctx.Request.Header)
		if ctx.Request.Body != nil {
			fmt.Println("Request Body:", ctx.Request.Body)
		}
		h.Next(ctx)
		if ctx.Response != nil {
			fmt.Println("Response Status:", ctx.Response.StatusCode)
			fmt.Println("Response Headers:", ctx.Response.Header)
			fmt.Println("Response Body:", ctx.Response.Body)
		} else if ctx.Error != nil {
			fmt.Println("Request Error:", ctx.Error)
		}
	})
	mid := middleware.New()
	mid.GetStack()
	// 定义请求 URL
	req := client.Request()
	req.Use(url2.URL("http://localhost:9000"))

	// 使用 net/url 包来构建 form 数据
	params := url.Values{}
	params.Add("name", "333")

	// 设置请求体为 application/x-www-form-urlencoded 格式
	req.Use(body.String(params.Encode()))

	// 设置 Content-Type 请求头
	req.SetHeader("Content-Type", "application/x-www-form-urlencoded")

	// Add the debug plugin to the request
	req.Method("POST")
	// .Use(DebugPlugin())

	// 发送 POST 请求
	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request failed: %s\n", err)
		return
	}

	// 读取并输出响应的状态码和主体内容
	fmt.Printf("Status: %d\n", res.StatusCode)
	if res.Ok {
		fmt.Println("Response received:", res.String())
	} else {
		fmt.Println("Failed to fetch response:", res.String())
	}
}
