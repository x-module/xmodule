/**
 * Created by Goland
 * @file   common.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/17 12:06
 * @desc   common.go
 */

package nakama

type Base struct {
	Token string
}

// SetToken 设置token
func (b *Base) SetToken(token string) {
	b.Token = token
}

// GeHeader 获取请求header
func GeHeader(token string, url string) map[string]string {
	Authorization := "Bearer " + token
	return map[string]string{
		"Authority":          "contractors.us-east1.nakamacloud.io",
		"Accept":             "application/json, text/plain, */*",
		"Accept-Language":    "zh-CN,zh;q=0.9",
		"Cache-Control":      "no-cache",
		"Authorization":      Authorization,
		"Pragma":             "no-cache",
		"Referer":            url,
		"Sec-Ch-Ua":          "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"100\", \"Google Chrome\";v=\"100\"",
		"Sec-Ch-Ua-Mobile":   "?0",
		"Sec-Ch-Ua-Platform": "\"macOS\"",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-origin",
		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Safari/537.36",
	}
}
