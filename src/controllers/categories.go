package controllers

import (
	"net/http"      // Will allow us to work with the response writer and request object
	"text/template" // Will enable us to work with the VIEWs templates
	"viewmodels"    // Get the View model that provide the HOME page with its 
)

type categoriesController struct {
	template *template.Template
}
func (this *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetCategory()

	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(w, vm)
}

type categoryController struct {
	template *template.Template
}
func (this *categoryController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetProducts(1)

	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(w, vm)
}