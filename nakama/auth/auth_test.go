/**
 * Created by Goland
 * @file   auth_test.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/21 16:39
 * @desc   auth_test.go
 */

package auth

import (
	"github.com/x-module/xmodule/xlog"
	"testing"
)

func init() {
	xlog.InitLogger(xlog.LogConfig{
		LogModel: xlog.DebugMode,
		LogPath:  "./log",
		LogFile:  "auth.log",
	})
}
func TestGetToken(t *testing.T) {
	token, err := NewAuth(AuthData{
		Username: "admin",
		Password: "REDACTED",
		BaseUrl:  "http://192.168.1.187:7351",
		SignKey:  "defaultsigningkey",
		LogModel: xlog.DebugMode,
	}).GetToken(LoginToken{})
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
