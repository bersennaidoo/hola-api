package rest

import "net/http"

func (app *Application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/hola", app.holaHandler)

	return mux
}
