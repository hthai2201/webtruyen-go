// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/hthai2201/webtruyen-go/config"
	amqprpc "github.com/hthai2201/webtruyen-go/internal/controller/amqp_rpc"
	v1 "github.com/hthai2201/webtruyen-go/internal/controller/http/v1"
	"github.com/hthai2201/webtruyen-go/internal/usecase"
	"github.com/hthai2201/webtruyen-go/internal/usecase/repo"
	"github.com/hthai2201/webtruyen-go/internal/usecase/webapi"
	"github.com/hthai2201/webtruyen-go/pkg/appctx"
	"github.com/hthai2201/webtruyen-go/pkg/httpserver"
	"github.com/hthai2201/webtruyen-go/pkg/logger"
	"github.com/hthai2201/webtruyen-go/pkg/rabbitmq/rmq_rpc/server"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	dbUrl := os.Getenv("DATABASE_URL")
	fmt.Println(dbUrl)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	appCtx := appctx.NewAppContext(db.Debug())
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	// Use case
	translationUseCase := usecase.New(
		repo.New(appCtx),
		webapi.New(),
	)

	// RabbitMQ RPC Server
	rmqRouter := amqprpc.NewRouter(translationUseCase)

	rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	}

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, translationUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-rmqServer.Notify():
		l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = rmqServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}
}
