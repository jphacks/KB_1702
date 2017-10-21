//go:generate goagen bootstrap -d github.com/pei0804/goa-docker-stater/design

package main

import (
	"flag"

	"app/app"
	"app/controller"
	"app/design"

	"github.com/deadcheat/goacors"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

// Server 実行に必要な値を保持している
type Server struct {
	service *goa.Service
}

// NewServer Server構造体を作成する
func NewServer(s *goa.Service) *Server {
	return &Server{
		service: s,
	}
}

func (s *Server) mountController() {
	// Mount "rooms" controller
	room := controller.NewRoomsController(s.service)
	app.MountRoomsController(s.service, room)
	// Mount "front" controller
	front := controller.NewFrontController(s.service)
	app.MountFrontController(s.service, front)
	// Mount "swagger" controller
	swagger := controller.NewSwaggerController(s.service)
	app.MountSwaggerController(s.service, swagger)
	// Mount "swaggerui" controller
	swaggerui := controller.NewSwaggeruiController(s.service)
	app.MountSwaggeruiController(s.service, swaggerui)
}

func (s *Server) mountMiddleware(env string) {
	s.service.Use(middleware.RequestID())
	s.service.Use(middleware.LogRequest(true))
	s.service.Use(middleware.ErrorHandler(s.service, true))
	s.service.Use(middleware.Recover())
	s.service.Use(goacors.WithConfig(s.service, design.CorsConfig[env]))
}

func main() {
	service := goa.New(design.REPO)
	var (
		port = flag.String("port", ":8080", "addr to bind")
		env  = flag.String("env", "production", "実行環境 (production, develop)")
	)
	flag.Parse()
	s := NewServer(service)
	s.mountMiddleware(*env)
	s.mountController()

	// Start service
	if err := service.ListenAndServe(*port); err != nil {
		service.LogError("startup", "err", err)
	}
}
