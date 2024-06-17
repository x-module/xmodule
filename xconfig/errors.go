/**
 * Created by Goland
 * @file   errors.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/17 10:33
 * @desc   errors.go
 */

package xconfig

import "github.com/x-module/xmodule/internal"

type ErrCode internal.ErrCode

//go:generate stringer -type ErrCode -linecomment -output errors_string.go

const (
	GetConfigErr ErrCode = 10000 + iota // 获取配置异常
)
