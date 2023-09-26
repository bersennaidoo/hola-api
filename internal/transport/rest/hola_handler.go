package rest

import (
	"encoding/json"
	"net/http"
)

func (app *Application) holaHandler(w http.ResponseWriter, r *http.Request) {

	enc := json.NewEncoder(w)
	w.Header().
		Set("Content-Type", "application/json; charset=utf-8")

	resp := Resp{
		Language:    "Spanish",
		Translation: "Hola",
	}

	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
