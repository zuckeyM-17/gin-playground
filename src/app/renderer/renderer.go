package renderer

import (
	"html/template"
	"io/ioutil"
	"path/filepath"

	"github.com/gin-gonic/gin/render"
)

type Renderer map[string]*template.Template

func CreateRenderer() Renderer {
	renderer := make(Renderer)
	files, _ := ioutil.ReadDir("templates")
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join("templates", file.Name())
		tmpl := template.Must(template.ParseFiles("templates/_layouts/base.html", filePath))
		renderer[file.Name()] = tmpl
	}

	return renderer
}

func (r Renderer) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r[name],
		Data:     data,
	}
}
