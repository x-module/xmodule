/**
 * Created by Goland.
 * @file   account.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/8 19:32
 * @desc   account.go
 */

package api

import (
	"github.com/x-module/helper/debug"
	"github.com/x-module/xmodule/nakama/auth"
	"github.com/x-module/xmodule/xlog"
	"log"
	"testing"
)

func init() {
	xlog.InitLogger(xlog.LogConfig{
		LogModel: xlog.DebugMode,
		LogPath:  "./log",
		LogFile:  "auth.log",
	})
}

func getToken() auth.LoginToken {
	token, err := auth.NewAuth(auth.AuthData{
		Username: "admin",
		Password: "REDACTED",
		BaseUrl:  "http://192.168.1.187:7351",
		SignKey:  "defaultsigningkey",
		LogModel: xlog.DebugMode,
	}).GetToken(auth.LoginToken{})
	if err != nil {
		log.Printf("GetToken err:%s", err.Error())
		return auth.LoginToken{}
	}
	return token
}

func TestGetAccountList(T *testing.T) {
	token := getToken()
	accounts, err := NewAccount(token.Token).GetAccountList(GetAccountListParams{
		BaseUrl: "http://192.168.1.187:7351",
	})
	if err != nil {
		T.Error(err)
	}
	debug.DumpPrint(accounts)
	// T.Log(convertor.ToJsonString(accounts.Users))
}

func TestGetAccountDetail(T *testing.T) {
	token := getToken()
	accounts, err := NewAccount(token.Token).GetAccountDetail(GetAccountDetailParams{
		BaseUrl:  "http://192.168.1.187:7351",
		PlayerId: "001af836-70db-454b-a7b3-aee12d65c4c0",
	})
	if err != nil {
		T.Error(err)
	}
	debug.DumpPrint(accounts)
	// T.Log(convertor.ToJsonString(accounts.Users))
}
