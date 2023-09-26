package rest

import (
	"log"
	"net/http"
)

func (app *Application) ServerRun(addr string) {

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, app.routes()))
}
