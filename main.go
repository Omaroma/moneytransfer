package main

import (
	"github.com/Omaroma/moneytransfer/middleware"
	"log"
)

func main() {
	r := middleware.Router()
	log.Println("starting server on port", middleware.PORT)
	err := r.Run(middleware.PORT)
	if err != nil {
		log.Fatal("failed to run te server, ", err)
	}
}
