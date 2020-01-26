package Handlers

import (
	"github.com/medis/go-expressive/internal/Template"
	"net/http"
)

type homeHandlerStruct struct{
	template *Template.Template
}

func NewHomeHandler(template *Template.Template) *homeHandlerStruct {
	return &homeHandlerStruct{
		template: template,
	}
}

func (handler *homeHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	output, err := handler.template.Render("App.home.page.gohtml")
	if err != nil {
		//Server.ServerError(a.errorLog, w, err)
		w.Write([]byte("error"))
	}

	_, err = output.WriteTo(w)
	if err != nil {
		w.Write([]byte("error"))
	}
}
