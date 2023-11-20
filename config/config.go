package config

import "github.com/Sellsuki/sellsuki-go-logger/level"

type Config struct {
	LogLevel      level.Level
	AppName       string
	Version       string
	MaxBodySize   int
	Readable      bool
	HardCodedTime string
}
