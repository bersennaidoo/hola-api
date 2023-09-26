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

	language := defaultLanguage
	word := strings.ReplaceAll(r.URL.Path, "/", "")
	translation := app.Translator.Translate(word, language)

	resp := Resp{
		Language:    language,
		Translation: translation,
	}

	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
