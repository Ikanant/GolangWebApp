package controllers

import (
	"net/http"      // Will allow us to work with the response writer and request object
	"text/template" // Will enable us to work with the VIEWs templates
	"viewmodels"    // Get the View model that provide the HOME page with its data
	"controllers/util"
)

type aboutController struct {
	template *template.Template
}

func (this *aboutController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetAbout()

	w.Header().Add("Content-Type", "text/html")
	
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	this.template.Execute(responseWriter, vm)
}