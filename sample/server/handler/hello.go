package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/teamlint/ardan/container"
	"github.com/teamlint/ardan/route"
	"github.com/teamlint/ardan/sample/app/model"
	"github.com/teamlint/ardan/sample/app/service"
)

func User() route.Route {
	return route.Get("/user", func(c *gin.Context) {
		var result *model.User
		var err error
		// service locator 模式
		container.Invoke(func(svc service.UserService) {
			log.Printf("[handler] container svc: %v", svc)
			id := c.DefaultQuery("id", "1213")
			result, err = svc.Get(id)
			if err != nil {
				c.JSON(200, err.Error())
				return
			}
			log.Printf("[handler] Hello() load service.Get: %+v", *result)
		})
		c.JSON(200, result)
	})
}

func World(path string) route.Route {
	return route.GetAndPost(path, func(c *gin.Context) {
		c.String(200, "[Handler] World()")
	})
}
