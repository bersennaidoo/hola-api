package rest

import "github.com/bersennaidoo/hola-api/internal/service/translation"

type Application struct {
	Translator *translation.Translator
}

func New(tr *translation.Translator) *Application {
	return &Application{
		Translator: tr,
	}
}
