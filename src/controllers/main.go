package controllers

// We will need all of the imports form the old main.go file
import (
	"bufio" 		// Allows us to buffer the output
	"net/http"
	"os"            // Access to file system opeartion like reading the templates
	"strings"       // We will use it to compare the requested file extensions
	"text/template" // Gives the templating functionallity we need
	"viewmodels"    //access new functionality package
)

// Will serve as the point where the controller layer configures itself and prepares to handle requests
// Receive the template CACHE that we build in the main package
// We will grab the three http functions form the main package
func Register(templates *template.Template) {
	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {

			requestFile := req.URL.Path[1:]
			template := templates.Lookup(requestFile + ".html")

			/*
				Since we are using this function to handle each of the pages, we are going to have to handle several types of context objects
				To manage that we will set the context to be on type INTERFACE, so that anything can be assigned to it later
			*/
			var context interface{} = nil

			switch requestFile {
			case "home":
				context = viewmodels.GetHome()
			case "about":
				context = viewmodels.GetAbout()
			}

			if template != nil {
				template.Execute(w, context)
			} else {
				w.WriteHeader(404)
			}
		})

	http.HandleFunc("/images/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/fonts/", serveResource)
	http.HandleFunc("/js/", serveResource)
}

// Responsible for processing the resource request and returning them to the requester
func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
