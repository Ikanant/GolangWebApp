// This file will hold the back controller functionality for the HOME page

package controllers

import (
	"net/http"      // Will allow us to work with the response writer and request object
	"text/template" // Will enable us to work with the VIEWs templates
	"viewmodels"    // Get the View model that provide the HOME page with its data
)

/*
Notice that both the HOME controller and its field begin with LOWER case letters
--> This means that they are scoped to the controller package. This will keep the public API of the controller package
	nice and light since nothing else really needs to know how the controller layer goes about organzing itself
*/
type homeController struct {
	template *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	// Since this controller method is only responsible for a single VIEW, its implementation becomes really simple

	// Grab a view model that is preconfigured with MOCK data using the VIEW model GETHOME
	vm := viewmodels.GetHome()

	// Tell the browser that we will be sending HTML by setting the content-type header on the response
	w.Header().Add("Content-Type", "text/html")

	/* We execute the controllers template field.
	This small/lite controlled funciton is important
	Due the controllers position in the MVC pattern it is aware of both of the other layer. This makes it easy to use the
	controller to handle things that really are the responsability of another portion of the application */
	this.template.Execute(w, vm)

}
