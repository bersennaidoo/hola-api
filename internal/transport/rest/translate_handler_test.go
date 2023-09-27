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

	tt := []struct {
		Name                string
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Name:                "returns 'hello'",
			Endpoint:            "/hello",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Name:                "returns 'hallo'",
			Endpoint:            "/hello?language=german",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			Name:                "returns empty string",
			Endpoint:            "/hello?language=dutch",
			StatusCode:          http.StatusNotFound,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}

	for _, test := range tt {

		t.Run(test.Name, func(t *testing.T) {
			tr := translation.New()

			app := rest.New(tr)

			handler := http.HandlerFunc(app.TranslateHandler)

			rr := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", test.Endpoint, nil)

			handler.ServeHTTP(rr, req)

			if rr.Code != test.StatusCode {
				t.Errorf("expected status %d but received %d", test.StatusCode, rr.Code)
			}

			var resp rest.Resp
			json.Unmarshal(rr.Body.Bytes(), &resp)

			if resp.Language != test.ExpectedLanguage {
				t.Errorf("expected language '%s' but received %s", test.ExpectedLanguage, resp.Language)
			}

			if resp.Translation != test.ExpectedTranslation {
				t.Errorf("expected Translation '%s' but received %s",
					test.ExpectedTranslation, resp.Translation)
			}
		})
	}

}
