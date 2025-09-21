package main

import (
	"log"

	"github.com/hueypark/mcp/src/server"
)

func main() {
	server := server.New()

	err := server.ListenAndServe(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
