/**
 * Created by Goland.
 * @file   leaderboard.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2022/4/13 17:56
 * @desc   leaderboard.go
 */

package api

import (
	"github.com/x-module/helper/debug"
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

// TestGetLeaderboardList 测试获取排行榜列表
func TestGetLeaderboardList(T *testing.T) {
	token := getToken()
	list, err := NewLeaderboard(token.Token).GetLeaderboardList(GetLeaderboardListParams{
		BaseUrl: "http://192.168.1.187:7351",
	})
	if err != nil {
		T.Error(err)
	}
	debug.DumpPrint(list)
}

// TestGetLeaderboardDetail 测试获取排行榜详情
func TestGetLeaderboardDetail(T *testing.T) {
	token := getToken()
	list, err := NewLeaderboard(token.Token).GetLeaderboardDetail(GetLeaderboardDetailParams{
		BaseUrl:     "http://192.168.1.187:7351",
		Leaderboard: "airstrikeCount",
	})
	if err != nil {
		T.Error(err)
	}
	debug.DumpPrint(list)
}

// TestGetLeaderboardRecord 测试获取排行榜记录
func TestGetLeaderboardRecord(T *testing.T) {
	token := getToken()
	list, err := NewLeaderboard(token.Token).GetLeaderboardRecord(GetLeaderboardRecordParams{
		BaseUrl:     "http://192.168.1.187:7351",
		Leaderboard: "airstrikeCount",
	})
	if err != nil {
		T.Error(err)
	}
	debug.DumpPrint(list)
}
