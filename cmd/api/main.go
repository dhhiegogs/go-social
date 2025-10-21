package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
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
