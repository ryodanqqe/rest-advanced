package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"rest-api/internal/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("create router")
	router := httprouter.New()

	log.Println("register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	run(router)
}

func run(router *httprouter.Router) {
	log.Println("run application")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("server is listening port :1234")
	log.Fatalln(server.Serve(listener))
}
