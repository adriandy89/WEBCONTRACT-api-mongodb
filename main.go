package main

import (
	"log"

	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/handlers"
)

func main() {

	if !db.TestConnection() {
		log.Fatal("Sin conexion")
		return
	}
	handlers.Handlers()
}
