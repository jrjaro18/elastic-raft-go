package proxy

import (
	"context"
	"log"
	"os"
	"os/signal"
	"slices"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jrjaro18/elastic-raft-go/proxy/internal/middleware"
)

type Proxy struct {
	app     *fiber.App
	port    string
	apiKey  string
	servers map[string]struct{}
}

func NewProxy(port string, apiKey string) *Proxy {
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	})
	return &Proxy{
		app:     app,
		port:    port,
		apiKey:  apiKey,
		servers: map[string]struct{}{},
	}
}

func (p *Proxy) Start() {
	go func() {
		log.Println("Starting proxy @ port", p.port)
		if err := p.app.Listen(p.port); err != nil {
			log.Fatalln("Proxy couldn't be started ---", err)
		}
	}()
	p.initServerRoutes()
	p.closeProxy()
}

func (p *Proxy) closeProxy() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Starting to close proxy")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := p.app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Proxy forced to shutdown: %v\n", err)
	}

	log.Println("Proxy exited gracefully.")
}

func (p *Proxy) GetListedServer(exception []string) []string {
	a := []string{}
	for x := range p.servers {
		if !slices.Contains(exception, x) {
			a = append(a, x)
		}
	}
	return a
}

func (p *Proxy) registerServer(server string) {
	p.servers[server] = struct{}{}
}

func (p *Proxy) initServerRoutes() {
	s := p.app.Group("server")
	api := s.Group("v1")

	api.Use(middleware.ApiKeyMiddleware(p.apiKey))

	api.Get("/", p.serverStatus)
	api.Get("/register", p.registerServerHandler)
}
