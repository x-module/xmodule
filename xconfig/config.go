/**
 * Created by Goland
 * @file   config.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/17 10:30
 * @desc   config.go
 */

package xconfig

import (
	"github.com/x-module/helper/fileutil"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
)

var (
	// ConfigPatch 配置路径
	ConfigPatch = "config"
	// ENVKey 环境变量
	ENVKey = "ENV"
	// DefaultENV 默认环境
	DefaultENV = "dev"
)

// ParseConfig 解析配置
func ParseConfig[T any](configPath string, config T) T {
	err := GetConfig(GetConfigFile(configPath), &config)
	if err != nil {
		log.Fatal(err, GetConfigErr.String())
	}
	return config
}

// GetConfig 获取配置
func GetConfig(path string, config any) error {
	if content, err := fileutil.ReadFileToString(path); err != nil {
		return err
	} else if err = yaml.Unmarshal([]byte(content), config); err != nil {
		return err
	}
	return nil
}

// GetConfigFile 获取当前运行环境下的配置
func GetConfigFile(configFile string) string {
	mode := os.Getenv(ENVKey)
	if mode == "" {
		mode = DefaultENV
	}
	return path.Join(ConfigPatch, mode, configFile)
}
