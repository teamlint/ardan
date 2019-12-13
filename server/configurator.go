package server

import "github.com/gin-gonic/gin"

// WithRelease release模式
func WithRelease() Configurator {
	return func(s *Server) {
		gin.SetMode(gin.ReleaseMode)
	}
}
