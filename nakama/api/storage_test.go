/**
 * Created by Goland
 * @file   storage.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/21 19:07
 * @desc   storage.go
 */

package api

import (
	"github.com/x-module/helper/debug"
	"testing"
)

// TestGetStorageList 测试获取存储列表
func TestGetStorageList(T *testing.T) {
	token := getToken()
	list, err := NewStorage(token.Token, "").GetStorageList(GetStorageListParams{
		BaseUrl:    "http://192.168.1.187:7351",
		Collection: "PLAYER_DATA",
		Key:        "DOG_TAG",
	})
	if err != nil {
		T.Error(err)
	}
	debug.DumpPrint(list)
}

// TestGetStorageDetail 测试获取存储详情
func TestGetStorageDetail(T *testing.T) {
	token := getToken()
	list, err := NewStorage(token.Token, map[string]ValueData{}).GetStorageDetail(GetStorageDetailParams{
		BaseUrl:    "http://192.168.1.187:7351",
		Collection: "PLAYER_DATA",
		Key:        "DOG_TAG",
		PlayerId:   "45e7a83a-d4e8-494e-99bd-5f54ab29197a",
	})
	if err != nil {
		T.Error(err)
	}
	debug.DumpPrint(list)
}

type ValueData struct {
	Count     int `json:"count"`
	CreatedAt int `json:"createdAt"`
}
