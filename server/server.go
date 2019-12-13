package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teamlint/ardan/config/section"
	"github.com/teamlint/ardan/route"
	"go.uber.org/dig"
)

const (
	// DebugMode indicates server mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates server mode is release.
	ReleaseMode = "release"
	// TestMode indicates server mode is test.
	TestMode = "test"
)

type Server struct {
	*gin.Engine
}
type RouteParams struct {
	dig.In
	Routes []route.RouteInfo `group:"route"`
}

// Configurator 服务器配置
type Configurator func(*Server)

// SetMode 设置服务模式
func SetMode(mode string) {
	gin.SetMode(mode)
}

// Mode returns currently server mode.
func Mode() string {
	return gin.Mode()
}

// New 初始化服务器
func New(rp RouteParams, cs ...Configurator) *Server {
	e := gin.New()
	s := &Server{e}
	s.registerRouteParams(rp)
	s.Configure(cs...)
	return s
}

// Default 默认服务器
func Default(rp RouteParams, cs ...Configurator) *Server {
	e := gin.Default()
	s := &Server{e}
	s.registerRouteParams(rp)
	s.Configure(cs...)
	return s
}

// registerRouteInfo 注册路由留牌
func (s *Server) registerRouteInfo(r route.RouteInfo) *Server {
	switch r.Method {
	case "ANY":
		s.Any(r.Path, r.Handlers...)
		return s
	case "GET_POST":
		s.GET(r.Path, r.Handlers...)
		s.POST(r.Path, r.Handlers...)
		return s
	default:
		s.Handle(r.Method, r.Path, r.Handlers...)
		return s
	}
}

func (s *Server) registerRouteParams(rp RouteParams) *Server {
	for _, r := range rp.Routes {
		s.registerRouteInfo(r)
	}
	return s
}

// RegisterRoutes 注册路由
func (s *Server) RegisterRoute(rs ...route.Route) *Server {
	for _, r := range rs {
		return s.RegisterRouteInfo(r.RouteInfo)
	}
	return s
}

// RegisterRoutes 注册路由信息
func (s *Server) RegisterRouteInfo(rs ...route.RouteInfo) *Server {
	for _, r := range rs {
		s.registerRouteInfo(r)
	}
	return s
}

// Configure 配置
func (s *Server) Configure(cs ...Configurator) *Server {
	for _, c := range cs {
		c(s)
	}
	return s
}

// HttpServer 标准Http服务器
func (s *Server) HttpServer(conf *section.Config) *http.Server {
	readTimeout, _ := time.ParseDuration(conf.Server.ReadTimeout)
	writeTimeout, _ := time.ParseDuration(conf.Server.WriteTimeout)
	idleTimeout, _ := time.ParseDuration(conf.Server.IdleTimeout)
	srv := &http.Server{
		Addr:         conf.Server.HttpAddr,
		Handler:      s,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
	return srv
}
