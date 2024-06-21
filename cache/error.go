/**
 * Created by Goland
 * @file   error.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/17 10:41
 * @desc   error.go
 */

package cache

import (
	"errors"
	"fmt"
	"github.com/x-module/xmodule/internal"
	"github.com/x-module/xmodule/xlog"
	"runtime/debug"
)

func Error(msg string) error {
	return errors.New(msg)
}

func HasErr(err error, errCode fmt.Stringer) bool {
	if err != nil {
		xlog.Logger.WithField(internal.ErrField, err).Error(errCode.String())
		return true
	}
	return false
}

func HasWar(err error, errCode fmt.Stringer) bool {
	if err != nil {
		xlog.Logger.WithField(internal.ErrField, err).Warn(errCode.String())
		return true
	}
	return false
}

func Success(status int) bool {
	return internal.ErrCode(status) == internal.Success
}

func PanicErr(err error, log string) {
	if err != nil {
		xlog.Logger.Fatal("error:%s,message:%s,trace:%s", err.Error(), log, string(debug.Stack()))
	}
}
