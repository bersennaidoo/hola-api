package main

import "github.com/bersennaidoo/hola-api/internal/transport/rest"

func main() {

	app := rest.New()

	app.ServerRun(":3000")
}
