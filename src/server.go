package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/lgranade/minesweeperapi/dao"
)

var server *http.Server

func start() error {
	var err error

	err = dao.Start(&dao.DBConfig{
		Host:     params.DBHost,
		Port:     params.DBPort,
		User:     params.DBUser,
		Password: params.DBPassword,
		Name:     params.DBName,
		SSLMode:  params.DBSSLMode,
	})
	if err != nil {
		return err
	}

	handler, err := createRoutes()
	if err != nil {
		return err
	}

	walkFunc := func(method string,
		route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler) error {
		log.Println(method, " ", route)
		return nil
	}
	if err := chi.Walk(handler, walkFunc); err != nil {
		log.Fatal("Can't walk through routes, err: ", err)
	}

	server = &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	go func() {
		err := server.Serve(listener)
		if err != http.ErrServerClosed {
			log.Println("Error starting api server")
		}
	}()

	return nil
}

func stop() {
	log.Println("Shutting down api server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err := server.Shutdown(ctx)
	if err != nil {
		log.Println("Graceful shutdown of api server failed, err: ", err)
	}
	cancel()
	dao.Stop()
}
