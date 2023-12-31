package rest

import (
	"encoding/json"
	"net/http"
	"strings"
)

const defaultLanguage = "english"

func (app *Application) TranslateHandler(w http.ResponseWriter, r *http.Request) {

	enc := json.NewEncoder(w)
	w.Header().
		Set("Content-Type", "application/json; charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = defaultLanguage
	}

	word := strings.ReplaceAll(r.URL.Path, "/", "")
	translation := app.Translator.Translate(word, language)
	if translation == "" {
		language = ""
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp := Resp{
		Language:    language,
		Translation: translation,
	}

	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}

func (app *Application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	enc := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json; charset=utf8")

	resp := map[string]string{"status": "up"}
	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
