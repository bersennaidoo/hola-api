package rest

import "net/http"

func (app *Application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", app.TranslateHandler)

	return mux
}
