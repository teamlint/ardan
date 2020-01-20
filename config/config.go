package config

import (
	"log"
	"os"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/reader"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/file"
	"github.com/teamlint/ardan/config/section"
)

var (
	conf *section.Config
)

const (
	CONFIG_FILE                  = "./config.yml"
	DEFAULT_TITLE                = "teamlint"
	DEFAULT_COPYRIGHT            = "teamlint.com"
	DEFAULT_TIMEFORMAT           = "2006-01-02 15:04:05"
	DEFAULT_CHARSET              = "UTF-8"
	DEFAULT_SERVER_HTTP_ADDR     = ":1234"
	DEFAULT_SERVER_READ_TIMEOUT  = "5s"
	DEFAULT_SERVER_WRITE_TIMEOUT = "10s"
	DEFAULT_SERVER_IDLE_TIMEOUT  = "15s"
)

type Option func(conf *section.Config)

func init() {
	conf = &section.Config{
		App:    &section.App{},
		Server: &section.Server{},
	}
	defaultOption(conf)
	sources := []source.Source{env.NewSource()}
	if Exists(CONFIG_FILE) {
		sources = append(sources, file.NewSource(file.WithPath(CONFIG_FILE)))
	}
	err := config.Load(sources...)
	if err != nil {
		panic(err)
	}
}

// Exists 判断指定文件/目录是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// Config 配置
func Config(opts ...Option) *section.Config {
	err := config.Get().Scan(conf)
	if err != nil {
		log.Fatalf("config file read err: %v\n", err)
	}
	for _, opt := range opts {
		opt(conf)
	}
	return conf
}

// App 应用程序配置
func App() *section.App {
	return Config().App
}

// Server 服务器配置
func Server() *section.Server {
	return Config().Server
}

// Databases 数据库配置
func Databases(name string) *section.Database {
	if v, ok := Config().Databases[name]; ok {
		return v
	}
	return nil
}

// Caches 缓存配置
func Caches(name string) *section.Cache {
	if v, ok := Config().Caches[name]; ok {
		return v
	}
	return nil
}

// Map 获取配置字典
func Map() map[string]interface{} {
	return config.Map()
}

// Get 获取指定路径配置
func Get(path ...string) reader.Value {
	return config.Get(path...)
}

// Option
// -------------------------------------------------------------
// defaultOption 默认值
func defaultOption(conf *section.Config) {
	// App
	conf.App.Debug = true
	if conf.App.Title == "" {
		conf.App.Title = DEFAULT_TITLE
	}
	if conf.App.Copyright == "" {
		conf.App.Copyright = DEFAULT_COPYRIGHT
	}
	if conf.App.TimeFormat == "" {
		conf.App.TimeFormat = DEFAULT_TIMEFORMAT
	}
	if conf.App.Charset == "" {
		conf.App.Charset = DEFAULT_CHARSET
	}
	// Server
	if conf.Server.HttpAddr == "" {
		conf.Server.HttpAddr = DEFAULT_SERVER_HTTP_ADDR
	}
	if conf.Server.ReadTimeout == "" {
		conf.Server.ReadTimeout = DEFAULT_SERVER_READ_TIMEOUT
	}
	if conf.Server.WriteTimeout == "" {
		conf.Server.WriteTimeout = DEFAULT_SERVER_WRITE_TIMEOUT
	}
	if conf.Server.IdleTimeout == "" {
		conf.Server.IdleTimeout = DEFAULT_SERVER_IDLE_TIMEOUT
	}
}

// Http服务地址
func WithHttpAddr(addr string) Option {
	return func(conf *section.Config) {
		conf.Server.HttpAddr = addr
	}
}

// WithTitle 应用标题
func WithTitle(title string) Option {
	return func(conf *section.Config) {
		conf.App.Title = title
	}
}

// WithCharset 字符集
func WithCharset(charset string) Option {
	return func(conf *section.Config) {
		conf.App.Charset = charset
	}
}
