package controller

import (
	"FirstWebApplication/webapp/viewmodel"
	"html/template"
	"net/http"
)

type home struct {
	hometemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewBase()
	h.hometemplate.Execute(w, vm)
}
