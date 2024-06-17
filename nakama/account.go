/**
 * Created by Goland.
 * @file   account.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/8 19:32
 * @desc   account.go
 */

package nakama

import (
	"errors"
	"fmt"
	"github.com/x-module/xmodule/xerror"
	"github.com/x-module/xmodule/xlog"
	xconfig "github.com/x-module/xmodule/xrequest"
	"net/url"
	"time"
)

type Account struct {
	Token string
}

type Accounts struct {
	Users      []User `json:"users"`
	TotalCount int    `json:"total_count"`
	NextCursor string `json:"next_cursor"`
}

type User struct {
	ID                    string    `json:"id"`
	Username              string    `json:"username"`
	DisplayName           string    `json:"display_name"`
	AvatarURL             string    `json:"avatar_url"`
	LangTag               string    `json:"lang_tag"`
	Location              string    `json:"location"`
	Timezone              string    `json:"timezone"`
	Metadata              string    `json:"metadata"`
	FacebookID            string    `json:"facebook_id"`
	GoogleID              string    `json:"google_id"`
	GameCenterID          string    `json:"gamecenter_id"`
	SteamID               string    `json:"steam_id"`
	Online                bool      `json:"online"`
	EdgeCount             int       `json:"edge_count"`
	CreateTime            time.Time `json:"create_time"`
	UpdateTime            time.Time `json:"update_time"`
	FacebookInstantGameID string    `json:"facebook_instant_game_id"`
	AppleID               string    `json:"apple_id"`
}

func NewAccount(token string) *Account {
	return &Account{Token: token}
}

type GetAccountListParams struct {
	ApiUrl  string
	Filter  string
	Cursor  string
	Mode    xlog.LogMode
	Timeout time.Duration
}

// GetAccountList 获取用户列表
func (a *Account) GetAccountList(params GetAccountListParams) (Accounts, error) {
	params.ApiUrl = params.ApiUrl + "?a=a"
	if params.Filter != "" {
		params.Filter = url.QueryEscape(params.Filter)
		params.ApiUrl = fmt.Sprintf("%s&filter=%s", params.ApiUrl, params.Filter)
	}
	if params.Cursor != "" {
		params.ApiUrl = fmt.Sprintf("%s&cursor=%s", params.ApiUrl, params.Cursor)
	}
	if params.Mode == "" {
		params.Mode = xlog.DebugMode
	}

	if params.Timeout == 0 {
		params.Timeout = 10
	}

	xlog.Logger.Info("当前运行模式为:", params.Mode)
	t, _ := url.Parse(params.ApiUrl)
	domain := fmt.Sprintf("%s://%s", t.Scheme, t.Host)
	response, err := xconfig.NewRequest().Debug(params.Mode == xlog.DebugMode).SetHeaders(GeHeader(a.Token, domain)).SetTimeout(params.Timeout).Get(params.ApiUrl)
	if xerror.HasErr(err, GetAccountListErr) {
		xlog.Logger.Error("request api[accounts-list] error:", err)
		return Accounts{}, err
	}
	defer response.Close()
	if !xerror.Success(response.StatusCode()) {
		content, _ := response.Content()
		xlog.Logger.Error("request api[accounts-list] error,result:", content)
		return Accounts{}, errors.New("request nakama server error")
	}
	var accounts Accounts
	err = response.Json(&accounts)
	if xerror.HasErr(err, ParseJsonDataErr) {
		return Accounts{}, err
	}
	return accounts, nil
}

type AccountInfo struct {
	Account     AccountData `json:"account"`
	DisableTime any         `json:"disable_time"`
}
type AccountData struct {
	User        User   `json:"user"`
	Wallet      string `json:"wallet"`
	Email       string `json:"email"`
	Devices     []any  `json:"devices"`
	CustomID    string `json:"custom_id"`
	VerifyTime  any    `json:"verify_time"`
	DisableTime string `json:"disable_time"`
}

type GetAccountDetailParams struct {
	ApiUrl  string
	Mode    xlog.LogMode
	Timeout time.Duration
}

// GetAccountDetail 获取用户详情
func (a *Account) GetAccountDetail(params GetAccountDetailParams) (AccountInfo, error) {
	t, _ := url.Parse(params.ApiUrl)
	domain := fmt.Sprintf("%s://%s", t.Scheme, t.Host)
	if params.Mode == "" {
		params.Mode = xlog.DebugMode
	}
	if params.Timeout == 0 {
		params.Timeout = 10
	}
	response, err := xconfig.NewRequest().Debug(params.Mode == xlog.DebugMode).SetHeaders(GeHeader(a.Token, domain)).SetTimeout(params.Timeout).Get(params.ApiUrl)
	if xerror.HasErr(err, GetAccountDetailErr) {
		return AccountInfo{}, err
	}
	defer response.Close()
	if !xerror.Success(response.StatusCode()) {
		xlog.Logger.Error("request api[accounts-detail] error,code:", response.StatusCode())
		return AccountInfo{}, errors.New("request nakama server error")
	}
	var accountInfo AccountInfo
	err = response.Json(&accountInfo)
	if xerror.HasErr(err, ParseJsonDataErr) {
		return AccountInfo{}, err
	}
	return accountInfo, nil
}
