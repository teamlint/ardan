package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teamlint/ardan/server"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// server module register interface
func (c *UserController) RegisterModule(s *server.Server) {
	g := s.Group("/user")
	g.GET("/list", c.List)
	g.GET("/info/:id", c.Info)
}

func (c *UserController) List(ctx *gin.Context) {
	ctx.String(http.StatusOK, "[UserController] get user list")
}
func (c *UserController) Info(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "[UserController] get user[%v] info", id)
}
