package config

import (
	"log"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/reader"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/config/source/file"
	"github.com/teamlint/ardan/config/section"
	"github.com/teamlint/ardan/pkg"
)

var (
	conf *section.Config
)

const (
	ConfigFile                = "./config.yml"
	DefaultTitle              = "teamlint"
	DefaultCopyright          = "teamlint.com"
	DefaultTimeFormat         = "2006-01-02 15:04:05"
	DefaultCharset            = "UTF-8"
	DefaultServerHTTPAddr     = ":1234"
	DefaultServerReadTimeout  = "5s"
	DefaultServerWriteTimeout = "10s"
	DefaultServerIdleTimeout  = "15s"
)

type Option func(conf *section.Config)

func init() {
	conf = &section.Config{
		App:    &section.App{},
		Server: &section.Server{},
	}
	defaultOption(conf)
	sources := []source.Source{env.NewSource()}
	if pkg.Exists(ConfigFile) {
		sources = append(sources, file.NewSource(file.WithPath(ConfigFile)))
	}
	err := config.Load(sources...)
	if err != nil {
		panic(err)
	}
}

// LoadFile 加载配置文件
func LoadFile(path string) error {
	return config.LoadFile(path)
}

// Load config sources
func Load(source ...source.Source) error {
	return config.Load(source...)
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
		conf.App.Title = DefaultTitle
	}
	if conf.App.Copyright == "" {
		conf.App.Copyright = DefaultCopyright
	}
	if conf.App.TimeFormat == "" {
		conf.App.TimeFormat = DefaultTimeFormat
	}
	if conf.App.Charset == "" {
		conf.App.Charset = DefaultCharset
	}
	// Server
	if conf.Server.HttpAddr == "" {
		conf.Server.HttpAddr = DefaultServerHTTPAddr
	}
	if conf.Server.ReadTimeout == "" {
		conf.Server.ReadTimeout = DefaultServerReadTimeout
	}
	if conf.Server.WriteTimeout == "" {
		conf.Server.WriteTimeout = DefaultServerWriteTimeout
	}
	if conf.Server.IdleTimeout == "" {
		conf.Server.IdleTimeout = DefaultServerIdleTimeout
	}
}

// WithHTTPAddr Http服务地址
func WithHTTPAddr(addr string) Option {
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
