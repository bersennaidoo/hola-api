package main

import (
	"github.com/bersennaidoo/hola-api/internal/service/translation"
	"github.com/bersennaidoo/hola-api/internal/transport/rest"
)

func main() {

	tr := translation.New()

	app := rest.New(tr)

	app.ServerRun(":3000")
}
