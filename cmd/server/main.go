package main

import (
	"log"

	"github.com/dnjooiopa/tcp-server/server"
)

func main() {
	srv := server.New()
	log.Println("server started at :8080")
	if err := srv.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
