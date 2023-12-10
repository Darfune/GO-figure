package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string)  {

	tc, err := CreateTemplateCashe()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// CreateTemplateCashe creates a template cache as a map
func CreateTemplateCashe() (map[string]*template.Template, error)  {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil{
		return myCache, err
	}
	for _, page := range pages{
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil{
			return myCache, err
		}
		matches, err := filepath.Glob("./template/*.layout.tmpl")
		if err != nil{
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./template/*.layout.tmpl")

			if err != nil{
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
