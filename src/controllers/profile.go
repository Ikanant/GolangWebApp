package controllers

import (
	"controllers/util"
	"net/http"
	"viewmodels"
	"text/template"
)

type profileController struct {
	template *template.Template
	
	
}

func (this *profileController) handle(w http.ResponseWriter, req *http.Request){
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	vm := viewmodels.GetProfile()
	if req.Method == "POST" {
		vm.User.First = req.FormValue("firstName")
		vm.User.Last = req.FormValue("lastName")
		vm.User.Email = req.FormValue("email")
	}
	
	responseWriter.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
}