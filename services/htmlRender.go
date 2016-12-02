package services

import (
	"html/template"
	"path/filepath"
	"strings"
	"github.com/gin-gonic/gin/render"
)

const (
	TemplatesDir = "src/github.com/stanisdev/indeed-clone/templates/" // TemplatesDir holds the location of the templates
	Layout = "layout" // Layout is the file name of the layout file
	Ext = ".html" // Ext is the file extension of the rendered templates
	Debug = false // Debug enables debug mode
)

type Render struct { // Render implements gin's HTMLRender and provides some sugar on top of it
	Templates    map[string]*template.Template
	Files        map[string][]string
	TemplatesDir string
	Layout       string
	Ext          string
	Debug        bool
}

func (r *Render) Instance(name string, data interface{}) render.Render { // Instance implements gin's HTML render interface
	var tpl *template.Template

	if r.Debug { // Check if gin is running in debug mode and load the templates accordingly
		tpl = r.loadTemplate(name)
	} else {
		tpl = r.Templates[name]
	}
	return render.HTML{
		Template: tpl,
		Data:     data,
	}
}

func (r *Render) loadTemplate(name string) *template.Template { // loadTemplate parses the specified template and returns it
	tpl, err := template.ParseFiles(r.Files[name]...)
	if err != nil {
		panic(name + " template name mismatch")
	}
	return template.Must(tpl, err)
}

func NewRender() Render { // New returns a fresh instance of Render
	return Render{
		Templates:    make(map[string]*template.Template),
		Files:        make(map[string][]string),
		TemplatesDir: TemplatesDir,
		Layout:       Layout,
		Ext:          Ext,
		Debug:        Debug,
	}
}

func (r *Render) Create() *Render { // Create goes through the `TemplatesDir` creating the template structure for rendering. Returns the Render instance.
	layout := r.TemplatesDir + r.Layout + r.Ext // templates/layouts/default.html

	tplRoot, err := filepath.Glob(r.TemplatesDir + "*" + r.Ext) // root dir
	// tplRoot => [templates/404.html]
	if err != nil {
		panic(err.Error())
	}

	tplSub, err := filepath.Glob(r.TemplatesDir + "**/*" + r.Ext) // sub dirs
	// tplSub => [templates/articles/index.html templates/layouts/default.html]
	if err != nil {
		panic(err.Error())
	}
	for _, tpl := range append(tplRoot, tplSub...) { // This check is to prevent `panic: template: redefinition of template "layout"`
		name := r.getTemplateName(tpl)
		// name => for example => articles/index || layouts/default || 404
		if name == r.Layout {
			continue
		}
		r.AddFromFiles(name, layout, tpl) // Попадет все то, что не относится к layout
	}
	return r
}

func (r *Render) AddFromFiles(name string, files ...string) *template.Template { // AddFromFiles parses the files and returns the result
	tmpl := template.Must(template.ParseFiles(files...))
	if r.Debug {
		r.Files[name] = files
	}
	r.Templates[name] = tmpl
	return tmpl
}

// getTemplateName returns the name of the template. For example, if the template path is `templates/articles/list.html`. getTemplateName would return `articles/list`
func (r *Render) getTemplateName(tpl string) string {
	dir, file := filepath.Split(tpl)
	dir = strings.Replace(dir, r.TemplatesDir, "", 1)
	file = strings.TrimSuffix(file, r.Ext)
	return dir + file
}
