package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/teamlint/ardan/route"
)

func Home() route.Route {
	return route.Get("/", func(c *gin.Context) {
		c.String(200, "hello ardan")
	})
}
