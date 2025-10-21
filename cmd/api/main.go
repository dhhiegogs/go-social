package main

import (
	"go-social/internal/env"
	"log"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	log.Printf("starting server on %s", cfg.addr)

	err := app.run(app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
