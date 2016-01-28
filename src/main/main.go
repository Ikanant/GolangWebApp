package main

import (
	"controllers"
	"net/http"
	"os"            // Access to file system opeartion like reading the templates
	"text/template" // Gives the templating functionallity we need
)

func main() {
	templates := populateTemplates()

	controllers.Register(templates)

	http.ListenAndServe(":8000", nil)
}

// Will return a pointer to a template object (A tempalte cache that we are creating)
func populateTemplates() *template.Template {

	result := template.New("templates")
	basePath := "templates"
	templateFolder, _ := os.Open(basePath)

	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)
	templatePaths := new([]string)

	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)
	return result
}
