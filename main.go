package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/ffauzann/simpleapi/docs"
	"github.com/ffauzann/simpleapi/internal/config"
	customMiddleware "github.com/ffauzann/simpleapi/internal/middleware"
	"github.com/ffauzann/simpleapi/internal/model/entity"
	authRepo "github.com/ffauzann/simpleapi/internal/module/authentication/repository"
	authRouter "github.com/ffauzann/simpleapi/internal/module/authentication/router"
	authService "github.com/ffauzann/simpleapi/internal/module/authentication/service"
	trxRepo "github.com/ffauzann/simpleapi/internal/module/transaction/repository"
	trxRouter "github.com/ffauzann/simpleapi/internal/module/transaction/router"
	trxService "github.com/ffauzann/simpleapi/internal/module/transaction/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// --------------------------//
//        SWAGGER SPECS      //
// --------------------------//
// @title Playground
// @version 1.0
// @description This is a playground server.

// @host localhost:3000
// @BasePath /api/v1
// @schemes http

var conf entity.Config

func init() {
	conf, _ = config.Setup()
}

func main() {
	// Create echo's instance
	e := echo.New()

	// Setup middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Groupping router
	g := e.Group("/api/v1")

	// Initialize repositories
	authr := authRepo.Init(conf.Connection.MySQL)
	trxr := trxRepo.Init(conf.Connection.MySQL)

	// Initialize services
	auths := authService.Init(authr, conf.App)
	trxs := trxService.Init(trxr, conf.App)

	// Initalize public routes
	authRouter.Init(g, auths, conf.App)
	// Swagger's endpoint could be restricted with server configuration
	// In this case, let's keep it open
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Initialize private routes
	g.Use(customMiddleware.JWT(conf.App.JWT.Secret))
	trxRouter.Init(g, trxs, conf.App)

	// Start server
	startServer(e)
}

func startServer(e *echo.Echo) {
	/*
		gracefully shuts down the server.
		reference:
			- https://medium.com/honestbee-tw-engineer/gracefully-shutdown-in-go-http-server-5f5e6b83da5a
			- https://echo.labstack.com/cookbook/graceful-shutdown/
	*/

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", conf.App.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		// Close MySQL connection
		db, _ := conf.Connection.MySQL.DB()
		db.Close()

		cancel()
	}()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
