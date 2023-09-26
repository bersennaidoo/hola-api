package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bersennaidoo/hola-api/internal/service/translation"
	"github.com/bersennaidoo/hola-api/internal/transport/rest"
)

func TestTranslateHandler(t *testing.T) {

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/hello", nil)

	tr := translation.New()

	app := rest.New(tr)

	handler := http.HandlerFunc(app.TranslateHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("got %q want status '200'", rr.Code)
	}

	var resp rest.Resp
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp.Language != "english" {
		t.Errorf("got %q want 'english'", resp.Language)
	}

	if resp.Translation != "hello" {
		t.Errorf("got %q want 'hello'", resp.Translation)
	}
}
