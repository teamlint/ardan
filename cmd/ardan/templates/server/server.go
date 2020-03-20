package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/teamlint/ardan/config"
	"github.com/teamlint/container"

	// "github.com/teamlint/ardan/container"
	"github.com/teamlint/ardan/sample/app/repository"
	"github.com/teamlint/ardan/sample/app/service"
	"github.com/teamlint/ardan/sample/server/configurator"
	"github.com/teamlint/ardan/sample/server/controller"
	"github.com/teamlint/ardan/sample/server/middleware"
	"github.com/teamlint/ardan/server"
)

// Start use teamlint container
func Start() {
	server.SetMode(server.DebugMode)
	container.Build(
		container.Provide(repository.NewDB),
		container.Provide(repository.NewUserRepository),
		container.Provide(service.NewUserService),
		// modules
		container.Provide(controller.NewUserController, container.As(new(server.Module))),
		container.Provide(middleware.NewDemoMDW, container.As(new(server.Module))),
		// with modules server
		container.Provide(server.NewWithModules),
	)
	var srv *server.Server
	container.MustExtract(&srv)

	srv.Use(middleware.Demo())
	srv.Configure(configurator.WithRoutes())
	log.Printf("[server] mode=%v", server.Mode())
	log.Printf("[server] config.App=%+v\n", config.App())
	// http server
	conf := config.Config()
	s := srv.HttpServer(conf)
	// greaceful shutdown
	done := make(chan struct{}, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	go server.GracefulShutdown(s, quit, done, conf)
	log.Println("Server is ready to handle requests at", conf.Server.HttpAddr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", conf.Server.HttpAddr, err)
	}
	<-done
	log.Println("Server stopped")
}
