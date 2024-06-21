/**
 * Created by Goland.
 * @file   leaderboard.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/13 17:56
 * @desc   leaderboard.go
 */

package api

import (
	"errors"
	"fmt"
	"github.com/x-module/xmodule/cache"
	"github.com/x-module/xmodule/internal"
	"github.com/x-module/xmodule/nakama/common"
	"github.com/x-module/xmodule/request"
	"github.com/x-module/xmodule/xlog"
	"time"
)

type Leaderboard struct {
	common.NakamaApi
}
type LeaderboardList struct {
	Total        int               `json:"total"`
	Cursor       string            `json:"cursor"`
	Leaderboards []LeaderboardInfo `json:"leaderboards"`
}
type LeaderboardInfo struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Category      int    `json:"category"`
	SortOrder     int    `json:"sort_order"`
	Size          int    `json:"size"`
	MaxSize       int    `json:"max_size"`
	MaxNumScore   int    `json:"max_num_score"`
	Operator      int    `json:"operator"`
	EndActive     int    `json:"end_active"`
	ResetSchedule string `json:"reset_schedule"`
	Metadata      string `json:"metadata"`
	CreateTime    any    `json:"create_time"`
	StartTime     any    `json:"start_time"`
	EndTime       any    `json:"end_time"`
	Duration      int    `json:"duration"`
	StartActive   int    `json:"start_active"`
	JoinRequired  bool   `json:"join_required"`
	Authoritative bool   `json:"authoritative"`
	Tournament    bool   `json:"tournament"`
}

type LeaderboardRecord struct {
	Records      []Records `json:"records"`
	OwnerRecords []any     `json:"owner_records"`
	NextCursor   string    `json:"next_cursor"`
	PrevCursor   string    `json:"prev_cursor"`
}

type Records struct {
	LeaderboardID string    `json:"leaderboard_id"`
	OwnerID       string    `json:"owner_id"`
	Username      string    `json:"username"`
	Score         string    `json:"score"`
	Subscore      string    `json:"subscore"`
	NumScore      int       `json:"num_score"`
	Metadata      string    `json:"metadata"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
	ExpiryTime    any       `json:"expiry_time"`
	Rank          string    `json:"rank"`
	MaxNumScore   int       `json:"max_num_score"`
}

func NewLeaderboard(token string) *Leaderboard {
	leaderboard := new(Leaderboard)
	leaderboard.Token = token
	return leaderboard
}

type GetLeaderboardListParams struct {
	LogMode xlog.LogMode
	BaseUrl string
	Cursor  string
}

// GetLeaderboardList 获取排行榜列表
func (a *Leaderboard) GetLeaderboardList(params GetLeaderboardListParams) (LeaderboardList, error) {
	apiUrl := fmt.Sprintf("%s%s", params.BaseUrl, common.LeaderboardApi)
	if params.Cursor != "" {
		apiUrl = fmt.Sprintf("%s&cursor=%s", apiUrl, params.Cursor)
	}
	xlog.Logger.Info("当前运行模式为:", params.LogMode)
	response, err := request.NewRequest().Debug(params.LogMode == xlog.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Get(apiUrl)
	if cache.HasErr(err, internal.GetLeaderboardListErr) {
		return LeaderboardList{}, err
	}
	defer response.Close()
	if !cache.Success(response.StatusCode()) {
		content, _ := response.Content()
		xlog.Logger.Error("request api[leaderboard-list] error,result:", content)
		return LeaderboardList{}, errors.New("request nakama server error")
	}
	var leaderboardList LeaderboardList
	err = response.Json(&leaderboardList)
	if cache.HasErr(err, internal.ParseJsonDataErr) {
		return LeaderboardList{}, err
	}
	return leaderboardList, nil
}

type GetLeaderboardDetailParams struct {
	LogMode     xlog.LogMode
	BaseUrl     string
	Leaderboard string
}

// GetLeaderboardDetail 获取排行榜详情
func (a *Leaderboard) GetLeaderboardDetail(params GetLeaderboardDetailParams) (LeaderboardInfo, error) {
	apiUrl := fmt.Sprintf("%s%s/%s", params.BaseUrl, common.LeaderboardApi, params.Leaderboard)
	xlog.Logger.Info("当前运行模式为:", params.LogMode)
	response, err := new(request.Request).Debug(params.LogMode == xlog.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Get(apiUrl)
	if cache.HasErr(err, internal.GetLeaderboardDetailErr) {
		return LeaderboardInfo{}, err
	}
	defer response.Close()
	if !cache.Success(response.StatusCode()) {
		xlog.Logger.Error("request api[leaderboard-detail] error,code:", response.StatusCode())
		return LeaderboardInfo{}, errors.New(internal.GetLeaderboardDetailErr.String())
	}
	var leaderboardInfo LeaderboardInfo
	err = response.Json(&leaderboardInfo)
	if cache.HasErr(err, internal.ParseJsonDataErr) {
		return LeaderboardInfo{}, err
	}
	return leaderboardInfo, nil
}

type GetLeaderboardRecordParams struct {
	LogMode     xlog.LogMode
	BaseUrl     string
	Leaderboard string
	Cursor      string
}

// https://showdown.us-east1.nakamacloud.io/v2/console/leaderboard/trainingSurvivalTime_2024-06-21/records?limit=100&cursor=_4H_jQMBARtsZWFkZXJib2FyZFJlY29yZExpc3RDdXJzb3IB_44AAQcBBklzTmV4dAECAAENTGVhZGVyYm9hcmRJZAEMAAEKRXhwaXJ5VGltZQEEAAEFU2NvcmUBBAABCFN1YnNjb3JlAQQAAQdPd25lcklkAQwAAQRSYW5rAQQAAABR_44BAQEfdHJhaW5pbmdTdXJ2aXZhbFRpbWVfMjAyNC0wNi0yMQIeAiQ0ZjgwZTE0YS02MzY1LTQ1NmUtODM3Ny1iZmU5ZTZjNDZjN2YB_8gA

// GetLeaderboardRecord 获取排行榜记录
func (a *Leaderboard) GetLeaderboardRecord(params GetLeaderboardRecordParams) (LeaderboardRecord, error) {
	apiUrl := fmt.Sprintf("%s%s/%s/records?limit=100", params.BaseUrl, common.LeaderboardApi, params.Leaderboard)
	xlog.Logger.Info("当前运行模式为:", params.LogMode)
	if params.Cursor != "" {
		apiUrl = fmt.Sprintf("%s&cursor=%s", apiUrl, params.Cursor)
	}
	response, err := new(request.Request).Debug(params.LogMode == xlog.DebugMode).SetHeaders(a.GetNakamaHeader(a.Token)).SetTimeout(10).Get(apiUrl)
	if cache.HasErr(err, internal.GetAccountListErr) {
		return LeaderboardRecord{}, err
	}
	defer response.Close()
	if !cache.Success(response.StatusCode()) {
		content, _ := response.Content()
		xlog.Logger.Error("request api[accounts-list] error,result:", content)
		return LeaderboardRecord{}, errors.New(internal.GetAccountListErr.String())
	}
	var leaderboardRecord LeaderboardRecord
	err = response.Json(&leaderboardRecord)
	if cache.HasErr(err, internal.ParseJsonDataErr) {
		return LeaderboardRecord{}, err
	}
	return leaderboardRecord, nil
}
