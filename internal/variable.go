/**
 * Created by GoLand
 * @file   variable.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/6/26 22:16
 * @desc   variable.go
 */

package internal

type ErrCode int64

const (
	Success  ErrCode = 200 // Success
	ErrField         = "Err"
)
