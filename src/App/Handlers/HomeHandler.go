package Handlers

import (
	"github.com/medis/go-expressive/internal/Response"
	"github.com/medis/go-expressive/internal/Template"
	"net/http"
)

type homeHandlerStruct struct {
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
		Response.ServerError(w, err)
	}

	headers := make(map[string]string)
	Response.HttpResponse(output, w, http.StatusOK, headers)
}
