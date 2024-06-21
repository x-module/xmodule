/**
 * Created by Goland
 * @file   config.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/17 10:30
 * @desc   mysql.go
 */

package dirver

import (
	"fmt"
	"github.com/x-module/helper/xerror"
	"github.com/x-module/xmodule/internal"
	"github.com/x-module/xmodule/xlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var Db *gorm.DB

type LinkParams struct {
	Host        string
	Port        int
	UserName    string
	DbName      string
	Password    string
	MaxOpenConn int
	MaxIdleConn int
	LogLevel    int
}

// InitializeDB 初始化管理后台数据库
func InitializeDB(params LinkParams) (*gorm.DB, error) {
	linkParams := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(linkParams, params.UserName, params.Password, params.Host, params.Port, params.DbName)
	newLogger := logger.New(
		xlog.New(NewMysqlLogger(), "\r\n", xlog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,                      // Slow SQL threshold
			LogLevel:                  logger.LogLevel(params.LogLevel), // Log level
			IgnoreRecordNotFoundError: true,                             // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,                            // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	xerror.PanicErr(err, internal.ConnectMysqlErr.String())
	// 链接池设置
	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(params.MaxOpenConn)
	sqlDb.SetMaxIdleConns(params.MaxIdleConn)
	Db = db
	return db, nil
}

type MysqlLogger struct {
}

func NewMysqlLogger() *MysqlLogger {
	return new(MysqlLogger)
}

// Write 实现Write接口，用于写入
func (l *MysqlLogger) Write(p []byte) (n int, err error) {
	xlog.Logger.Debug(string(p))
	return 1, nil
}
