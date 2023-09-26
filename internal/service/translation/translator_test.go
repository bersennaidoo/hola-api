package translation_test

import (
	"testing"

	"github.com/bersennaidoo/hola-api/internal/service/translation"
)

func TestTranslate(t *testing.T) {

	tt := []struct {
		Name        string
		Word        string
		Language    string
		Translation string
	}{
		{
			Name:        "hello to hello",
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Name:        "hello to hallo",
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Name:        "hello to hei",
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{
			Name:        `hello to ""`,
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
		{
			Name:        `bye to ""`,
			Word:        "bye",
			Language:    "dutch",
			Translation: "",
		},
		{
			Name:        `bye to ""`,
			Word:        "bye",
			Language:    "german",
			Translation: "",
		},
		{
			Name:        "hello to hallo",
			Word:        "hello",
			Language:    "German",
			Translation: "hallo",
		},
		{
			Name:        "hello to hallo",
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Name:        "hello to hallo",
			Word:        "hello ",
			Language:    "german",
			Translation: "hallo",
		},
	}

	for _, test := range tt {

		t.Run(test.Name, func(t *testing.T) {

			tr := translation.New()

			got := tr.Translate(test.Word, test.Language)

			want := test.Translation

			if got != want {
				t.Errorf("Translate(test.Word, test.Language) = %q want %q", got, want)
			}
		})
	}
}
