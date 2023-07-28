package main

import (
	"net"
	"net/http"
	"time"

	"rest-api/internal/user"
	"rest-api/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()

	logger.Info("create router")
	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	run(router)
}

func run(router *httprouter.Router) {
	logger := logging.GetLogger()

	logger.Info("run application")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server is listening port :1234")
	logger.Fatalln(server.Serve(listener))
}
