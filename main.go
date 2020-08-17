package main

import (
	"log"

	"github.com/IanDex/twitter/db"
	"github.com/IanDex/twitter/handlers"
)

func main() {
	if db.CheckConection() == 0 {
		log.Fatal("efdfdf")
		return
	}

	handlers.Manejadores()
}
