package main

import (
	"context"
	"log"

	"github.com/raihansuwanto/go-boilerplate/app"
)

func main() {
	ctx := context.Background()

	app, err := app.ProvideApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	app.Run(ctx)
}
