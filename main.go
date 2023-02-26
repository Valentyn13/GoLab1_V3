package main

import (
	handler "lab1/handlers"
	lab1 "lab1/server"
	"log"
)

const port string = "8795"

func main() {
	handlers := new(handler.Handler)
	srv := new(lab1.Server)

	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatal("server error:", err.Error())
	}
}
