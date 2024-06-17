/**
 * Created by Goland
 * @file   storage.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/17 12:58
 * @desc   storage.go
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

type Storage struct {
	Base
}

func NewStorage() *Storage {
	return &Storage{}
}

type StorageData struct {
	Collection      string    `json:"collection"`
	Key             string    `json:"key"`
	UserId          string    `json:"user_id"`
	Value           string    `json:"value"`
	Version         string    `json:"version"`
	PermissionRead  int       `json:"permission_read"`
	PermissionWrite int       `json:"permission_write"`
	CreateTime      time.Time `json:"create_time"`
	UpdateTime      time.Time `json:"update_time"`
}

type StorageList struct {
	Objects    []StorageData `json:"objects"`
	TotalCount int           `json:"total_count"`
	NextCursor string        `json:"next_cursor"`
}

type GetStorageListParams struct {
	ApiUrl  string
	Cursor  string
	Mode    xlog.LogMode
	Timeout time.Duration
}

// GetStorageList 获取存储列表
func (s *Storage) GetStorageList(params GetStorageListParams) (StorageList, error) {
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
	response, err := xconfig.NewRequest().Debug(params.Mode == xlog.DebugMode).SetHeaders(GeHeader(s.Token, domain)).SetTimeout(params.Timeout).Get(params.ApiUrl)
	if xerror.HasErr(err, GetAccountListErr) {
		xlog.Logger.Error("request api[accounts-list] error:", err)
		return StorageList{}, err
	}
	defer response.Close()
	if !xerror.Success(response.StatusCode()) {
		content, _ := response.Content()
		xlog.Logger.Error("request api[accounts-list] error,result:", content)
		return StorageList{}, errors.New("request nakama server error")
	}
	var storageList StorageList
	err = response.Json(&storageList)
	if xerror.HasErr(err, ParseJsonDataErr) {
		return StorageList{}, err
	}
	return storageList, nil
}
