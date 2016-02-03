package controllers

import (
	"controllers/util"
	"converter"
	"models"
	"net/http"      // Will allow us to work with the response writer and request object
	"text/template" // Will enable us to work with the VIEWs templates
	"viewmodels"    // Get the View model that provide the HOME page with its data
)

type aboutController struct {
	template *template.Template
}

func (this *aboutController) get(w http.ResponseWriter, req *http.Request) {
	skills := models.GetSkills()
	skillsVM := []viewmodels.Skill{}

	for _, modelSkills := range skills {
		skillsVM = append(skillsVM, converter.ConvertSkillToViewModel(modelSkills))
	}

	vm := viewmodels.GetAbout()
	
	vm.Skills = skillsVM
	
	w.Header().Add("Content-Type", "text/html")

	responseWriter := util.GetResponseWriter(w, req)
	
	_, err := req.Cookie("goSessionId")
	if err == nil {
		vm.LoggedIn = true
	} else {
		vm.LoggedIn = false
	}
	
	this.template.Execute(responseWriter, vm)
	defer responseWriter.Close()
}
