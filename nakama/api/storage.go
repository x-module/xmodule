/**
 * Created by Goland
 * @file   storage.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/21 19:07
 * @desc   storage.go
 */

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/x-module/helper/debug"
	"github.com/x-module/xmodule/cache"
	"github.com/x-module/xmodule/internal"
	"github.com/x-module/xmodule/nakama/common"
	"github.com/x-module/xmodule/request"
	"github.com/x-module/xmodule/xlog"
	"time"
)

type Storage[T any] struct {
	common.NakamaApi
}

func NewStorage[T any](token string, _ T) *Storage[T] {
	storage := new(Storage[T])
	storage.Token = token
	return storage
}

type GetStorageListParams struct {
	BaseUrl    string       `json:"base_url"`
	Key        string       `json:"key"`
	Collection string       `json:"collection"`
	Cursor     string       `json:"cursor"`
	LogMode    xlog.LogMode `json:"log_mode"`
}

// GetStorageListResponse 获取存储列表
type GetStorageListResponse struct {
	Objects    []StorageData `json:"objects"`
	TotalCount int           `json:"total_count"`
	NextCursor string        `json:"next_cursor"`
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

// https://showdown.us-east1.nakamacloud.io/v2/console/storage?key=DOG_TAG&collection=PLAYER_DATA&cursor=Tv-PAwEBFGNvbnNvbGVTdG9yYWdlQ3Vyc29yAf-QAAEEAQNLZXkBDAABBlVzZXJJRAH_hgABCkNvbGxlY3Rpb24BDAABBFJlYWQBBAAAABD_hQYBAQRVVUlEAf-GAAAALf-QAQdET0dfVEFHARAALcvHPedOiZDtuxcy7PFSAQtQTEFZRVJfREFUQQEEAA

func (s *Storage[T]) GetStorageList(params GetStorageListParams) (list GetStorageListResponse, err error) {
	apiUrl := fmt.Sprintf("%s%s?key=%s&collection=%s&cursor=%s", params.BaseUrl, common.StorageApi, params.Key, params.Collection, params.Cursor)
	xlog.Logger.Info("当前运行模式为:", params.LogMode)
	response, err := request.NewRequest().Debug(params.LogMode == xlog.DebugMode).SetHeaders(s.GetNakamaHeader(s.Token)).SetTimeout(10).Get(apiUrl)
	if cache.HasErr(err, internal.GetLeaderboardListErr) {
		return GetStorageListResponse{}, err
	}
	defer response.Close()
	if !cache.Success(response.StatusCode()) {
		content, _ := response.Content()
		xlog.Logger.Error("request api[leaderboard-list] error,result:", content)
		return GetStorageListResponse{}, errors.New("request nakama server error")
	}
	var storageListResponse GetStorageListResponse
	err = response.Json(&storageListResponse)
	if cache.HasErr(err, internal.ParseJsonDataErr) {
		return GetStorageListResponse{}, err
	}
	return storageListResponse, nil
}

type GetStorageDetailParams struct {
	BaseUrl    string       `json:"base_url"`
	Key        string       `json:"key"`
	Collection string       `json:"collection"`
	LogMode    xlog.LogMode `json:"logMode"`
	PlayerId   string       `json:"playerId"`
}

type GetStorageDetailResponse struct {
	Collection      string    `json:"collection"`
	Key             string    `json:"key"`
	UserId          string    `json:"user_id"`
	Value           any       `json:"value"`
	Version         string    `json:"version"`
	PermissionRead  int       `json:"permission_read"`
	PermissionWrite int       `json:"permission_write"`
	CreateTime      time.Time `json:"create_time"`
	UpdateTime      time.Time `json:"update_time"`
}

func (s *Storage[T]) GetStorageDetail(params GetStorageDetailParams) (detail GetStorageDetailResponse, err error) {
	apiUrl := fmt.Sprintf("%s%s/%s/%s/%s", params.BaseUrl, common.StorageApi, params.Collection, params.Key, params.PlayerId)
	xlog.Logger.Info("当前运行模式为:", params.LogMode)
	response, err := request.NewRequest().Debug(params.LogMode == xlog.DebugMode).SetHeaders(s.GetNakamaHeader(s.Token)).SetTimeout(10).Get(apiUrl)
	if cache.HasErr(err, internal.GetLeaderboardListErr) {
		return GetStorageDetailResponse{}, err
	}
	defer response.Close()
	if !cache.Success(response.StatusCode()) {
		content, _ := response.Content()
		xlog.Logger.Error("request api[leaderboard-list] error,result:", content)
		return GetStorageDetailResponse{}, errors.New("request nakama server error")
	}
	var storageDetailResponse GetStorageDetailResponse
	err = response.Json(&storageDetailResponse)
	if cache.HasErr(err, internal.ParseJsonDataErr) {
		return GetStorageDetailResponse{}, err
	}
	debug.DumpPrint(storageDetailResponse.Value)
	var value T
	_ = json.Unmarshal([]byte(fmt.Sprint(storageDetailResponse.Value)), &value)
	storageDetailResponse.Value = value
	return storageDetailResponse, nil
}
