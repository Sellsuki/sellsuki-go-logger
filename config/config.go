package config

import "github.com/Sellsuki/sellsuki-go-logger/v2/level"

type Config struct {
	LogLevel      level.Level
	AppName       string
	Version       string
	MaxBodySize   int
	Readable      bool
	HardCodedTime string
}
