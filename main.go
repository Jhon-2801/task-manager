package main

import (
	"log"
	"v3/bd"
	"v3/handlers"
)

func main() {
	if bd.ChequeoConnectioDb() == 0 {
		log.Fatal("Sin conexion a la DB")
	}
	handlers.Manejadores()
}
