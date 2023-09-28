package rest

import "net/http"

func (app *Application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/translate/hello", app.TranslateHandler)
	mux.HandleFunc("/health", app.HealthCheckHandler)

	return mux
}
