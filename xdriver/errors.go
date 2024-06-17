/**
 * Created by Goland
 * @file   errors.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/17 10:33
 * @desc   errors.go
 */

package dirver

import "github.com/x-module/xmodule/internal"

//go:generate stringer -type ErrCode -linecomment

const (
	InitRedisErr    internal.ErrCode = 11000 + iota // 初始化系统-连接Redis数据库异常
	ConnectMysqlErr                                 // 连接数据库异常

)
