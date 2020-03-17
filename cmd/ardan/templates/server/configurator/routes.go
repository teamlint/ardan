package configurator

import (
	"github.com/teamlint/ardan/sample/server/handler"
	"github.com/teamlint/ardan/server"
)

func WithRoutes() server.Configurator {
	return func(s *server.Server) {
		// s.GET("/c", func(c *gin.Context) {
		// 	c.String(200, "configurator")
		// })
		s.RegisterRoute(handler.Home())
		s.RegisterRoute(handler.User())
		s.RegisterRoute(handler.World("/w"))
	}
}
