package main

import (
	"controllers"	// Will allow us to drop into the main function and add a call to the controllers register function
	"net/http"
	"os"            // Access to file system opeartion like reading the templates
	"text/template" // Gives the templating functionallity we need
)

func main() {
	templates := populateTemplates()

	controllers.Register(templates)

	http.ListenAndServe(":8000", nil)
}

// Will return a pointer to a template object (A tempalte cache that I am creating)
func populateTemplates() *template.Template {

	// Creates a new template called "templates" that will server as our template cache
	result := template.New("templates")
	
	// Open the folder templates and close it when done
	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	// Read EVERYTHING inside the Folder opened above
	templatePathsRaw, _ := templateFolder.Readdir(-1)
	
	// Create empty slice object of type string
	templatePaths := new([]string)

	// Parse through elements. If something is not a directory, add it to our templatePaths slice
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}

	// Insert all found templates inside our BIG template cache
	result.ParseFiles(*templatePaths...)
	
	return result
}
