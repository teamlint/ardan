package server

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/teamlint/ardan/config/section"
)

// GracefulShutdown 优雅关闭服务器
func GracefulShutdown(server *http.Server, quit <-chan os.Signal, done chan<- struct{}, conf ...*section.Config) {
	sig := <-quit
	log.Printf("Server received signal %s\n", sig)
	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), server.ReadTimeout)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	if len(conf) > 0 && !conf[0].App.Debug {
		select {
		case <-ctx.Done():
			log.Printf("timeout of %v\n", server.ReadTimeout)
		}
	}
	close(done)
}
