// Package cfg /*
// 用于读写 ini 配置文件。
// Used to read and write ini configuration files.
package cfg

import (
	"github.com/go-ini/ini"
	"os"
)

func GetConf(path string, section string, item string, def string) string {
	//var iniFile *ini.File
	_, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return def
	}
	// 载入配置文件
	cfg, err := ini.Load(path)
	if err != nil {
		return def
	}

	// 取项目: [项目名]
	se, err := cfg.GetSection(section)
	if err != nil {
		return def
	}

	// 取子项
	value := se.Key(item).Value()
	return value
}

func SetConf(path string, section string, item string, value string) bool {
	_, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}

	cfg, err := ini.Load(path)
	if err != nil {
		return false
	}
	cfg.Section(section).Key(item).SetValue(value)
	err = cfg.SaveTo(path)
	if err != nil {
		return false
	}
	return true
}
