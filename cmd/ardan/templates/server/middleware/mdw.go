package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teamlint/ardan/server"
)

type DemoMDW struct{}

func NewDemoMDW() *DemoMDW {
	return &DemoMDW{}
}

// server module register interface
func (m *DemoMDW) RegisterModule(s *server.Server) {
	s.Use(m.Log)
}
func (m *DemoMDW) Log(ctx *gin.Context) {
	log.Printf("[DemoMDW] log %v\n", time.Now().Unix())
	ctx.Next()
}
