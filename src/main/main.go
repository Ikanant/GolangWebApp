package main

import (
	"net/http"
	"os"            // Access to file system opeartion like reading the templates
	"text/template" // Gives the templating functionallity we need

	"bufio"   // Allows us to buffer the output
	"strings" // We will use it to compare the requested file extensions
)

func main() {
	// Define a templates variable to get the results of that function
	templates := populateTemplates()

	// Handler that will listen to page requests and return the correct template when found.
	// "/" Will listen at the root
	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {

			// Get the request file path from the URL of the request.
			// This path is always prefix with the / which don't want
			// [1:] by taking a slice of the path string, we can remove that first char
			requestFile := req.URL.Path[1:]

			// Match the request file with the template
			template := templates.Lookup(requestFile + ".html")

			if template != nil {
				// Execute template and for now, a nil data context (future used for the data injection)
				template.Execute(w, nil)
			} else {
				w.WriteHeader(404)
			}
		})
	
	// use HandleFunc to register routes that will use the serveResource function
	// Since this route is more specific than the handler above (listening to ROOT), it will take presedence
	
	http.HandleFunc("/images/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/fonts/", serveResource)
	http.HandleFunc("/js/", serveResource)

	// Change second parameter to nil so that the default server mux is used
	http.ListenAndServe(":8000", nil)
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

	// Will receive the response on the call to open a file with the requested path
	f, err := os.Open(path)
	
	if err == nil{
		// Defer call to close the file handler to ensure to free up the resources that are no longer required
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		
		// Create new buffer reader to wrap the file handle
		br := bufio.NewReader(f)
		// With the above line defined, we can simply write the results of the buffer reader to the response writter and be all set
		br.WriteTo(w)
	} else {
		// NOT found response to them
		w.WriteHeader(404)
	}
}

// Will return a pointer to a template object (A tempalte cache that we are creating)
func populateTemplates() *template.Template {

	// equal a new template. This one name will not matter yet
	result := template.New("templates")

	// Gives the location of the tempalte folder relative to the project ROOT
	basePath := "templates"

	// Open up that folder. We can ignore the error since we know that this function exists
	templateFolder, _ := os.Open(basePath)

	// Defer keyboard will allow us to tell the app to close the handle but only after this function has finished executing
	defer templateFolder.Close()

	// We need to grab its content and find any template inside
	// Will read the entries from the folder every time is called and returned the assiciated file info object
	// By passing a negative number we are telling it we want all of the content inside the folder
	// Readdir will ensure that we get not only the template file names, but also everything else
	templatePathsRaw, _ := templateFolder.Readdir(-1)
	templatePaths := new([]string)

	// LOOP through the file info objects and, if it ISNT a directory, added to our templatePath slices
	// This will work on any path delimeter. GO will convert BACKSLASHES if necessary
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}

	// Load these in as child templates. We can pass in an arbitrary number of characters into the method and they will be processed one by one
	// The type of the parameters must me strings
	// ... after the name of the slice. Allows the compiler to recognize that is getting a slice as an input and it will handle it accordingly
	result.ParseFiles(*templatePaths...)

	return result
}
