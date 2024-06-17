/**
 * Created by goland.
 * @file   log.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/1 19:10
 * @desc   log.go
 */

package xlog

import (
	"github.com/druidcaesa/gotool"
	"github.com/druidcaesa/gotool/openfile"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/x-module/helper/fileutil"
	"github.com/x-module/helper/xerror"
	"github.com/x-module/helper/xlog"
	"os"
	"path"
	"time"
)

// Logger 日志句柄
var Logger *Xlogger

type LogMode string

const (
	// DebugMode indicates gin mode is debug.
	DebugMode LogMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode LogMode = "release"
	// TestMode indicates gin mode is test.
	TestMode LogMode = "test"
)

type XLogInter interface {
	xlog.LogInter
	// GetLoggerSource 获取日志资源
	GetLoggerSource() *logrus.Logger
}

// Xlogger 系统日志
type Xlogger struct {
	logger *logrus.Logger
}

// WithField 添加字段
func (x *Xlogger) WithField(key string, value any) *logrus.Entry {
	return x.logger.WithField(key, value)
}

// Debug 调试日志
func (x *Xlogger) Debug(args ...any) {
	x.logger.Debug(args...)
}

// Info 信息日志
func (x *Xlogger) Info(args ...any) {
	x.logger.Info(args...)
}

// Warn 警告日志
func (x *Xlogger) Warn(args ...any) {
	x.logger.Warn(args...)
}

// Error 错误日志
func (x *Xlogger) Error(args ...any) {
	x.logger.Error(args...)
}

// Fatal 致命错误日志
func (x *Xlogger) Fatal(args ...any) {
	x.logger.Fatal(args...)
}

// GetLoggerSource 获取日志资源
func (x *Xlogger) GetLoggerSource() *logrus.Logger {
	return x.logger
}

type LogConfig struct {
	LogPath  string
	LogFile  string
	LogModel LogMode
}

// InitLogger 日志模块初始化
func InitLogger(config LogConfig) *Xlogger {
	if !fileutil.IsExist(config.LogFile) {
		err := os.MkdirAll(config.LogFile, os.ModePerm)
		xerror.PanicErr(err, "init system error. make log data err.path:"+config.LogFile)
	}
	// 日志文件
	fileName := path.Join(config.LogFile, config.LogFile)
	if !gotool.FileUtils.Exists(fileName) {
		openfile.Create(fileName)
		if !gotool.FileUtils.Exists(fileName) {
			panic("init system error. create log file err. log file:" + fileName)
		}
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	xerror.PanicErr(err, "open log file error")
	// 实例化
	logger := logrus.New()
	// 设置日志级别
	switch config.LogModel {
	case ReleaseMode:
		logger.SetLevel(logrus.WarnLevel)
	case DebugMode:
		logger.SetLevel(logrus.DebugLevel)
	case TestMode:
		logger.SetLevel(logrus.InfoLevel)
	}

	// 设置输出
	logger.Out = src
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	logger.SetOutput(os.Stdout)

	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(30*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))
	// 是否显示文件位置
	// Logger.SetReportCaller(true)
	Logger = &Xlogger{logger: logger}
	return Logger
}
