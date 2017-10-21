//go:generate goagen bootstrap -d github.com/pei0804/goa-docker-stater/design

package main

import (
	"flag"
	"log"

	"app/app"
	"app/config"
	"app/controller"
	"app/design"

	"path/filepath"

	"github.com/deadcheat/goacors"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	mgo "gopkg.in/mgo.v2"
)

// Server 実行に必要な値を保持している
type Server struct {
	service *goa.Service
	mongodb *mgo.Session
}

// NewServer Server構造体を作成する
func NewServer(s *goa.Service) *Server {
	return &Server{
		service: s,
	}
}

// loadConfig 設定ファイルの読み込み
func (s *Server) loadConfig(settingFolder string, env string) {
	cs, err := config.NewMongodbConfigsFromFile(filepath.Join(settingFolder, "mongodb.yml"))
	if err != nil {
		log.Fatalf("cannot open mongodb configuration. exit. %s", err)
	}
	s.mongodb, err = cs.MongodbOpen(env)
	if err != nil {
		log.Fatalf("mongodb initialization failed: %s", err)
	}
}

func (s *Server) mountController() {
	// Mount "rooms" controller
	room := controller.NewRoomsController(s.service, s.mongodb)
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
		port    = flag.String("port", ":8080", "addr to bind")
		env     = flag.String("env", "develop", "実行環境 (production, develop)")
		confDir = flag.String("confDir", "./config", "設定ファイルの場所")
	)
	flag.Parse()
	s := NewServer(service)
	s.loadConfig(*confDir, *env)
	s.mountMiddleware(*env)
	s.mountController()

	// Start service
	if err := service.ListenAndServe(*port); err != nil {
		service.LogError("startup", "err", err)
	}
}
