package main

import (
	"FirstWebApplication/webapp/controller"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	templates := populateTemplate()
	controller.Startup(templates)
	log.Fatal(http.ListenAndServe(":8082", nil))

}

func populateTemplate() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dirr, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory " + err.Error())
	}
	fis, err := dirr.Readdir(-1)
	if err != nil {
		panic("Failed to read the content of the content directory " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		defer f.Close()
		if err != nil {
			panic("Failed to open template " + fi.Name() + err.Error())
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content of file " + fi.Name() + err.Error())
		}
		tmp1 := template.Must(layout.Clone())
		_, err = tmp1.Parse(string(content))
		if err != nil {
			panic("Error  parsing  the content of file " + fi.Name() + " as template")
		}
		result[fi.Name()] = tmp1
	}
	return result
}
