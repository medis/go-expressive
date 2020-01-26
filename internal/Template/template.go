package Template

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Template struct {
	templateCache map[string]*template.Template
}

func NewTemplate() (*Template, error) {
	dirs := make(map[string]string)
	items, err := ioutil.ReadDir("src")
	for _, item := range items {
		// Check if ConfigProvider.go exists.
		if item.IsDir() && fileExists(fmt.Sprintf("src/%s/ConfigProvider.go", item.Name())) {
			dirs[item.Name()] = fmt.Sprintf("src/%s/templates", item.Name())
		}
	}
	templateCache, err := newTemplateCache(dirs)
	if err != nil {
		return nil, err
	}
	return &Template{templateCache: templateCache}, nil
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// humanDate returns a nicely formatted string representation of a time.Time object.
func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dirs map[string]string) (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := make(map[string]*template.Template)

	for module, dir := range dirs {
		// Check if directory exists.
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			return nil, err
		}

		// Use the filepath.Glob function to get a slice of all filepaths with
		// the extension '.page.tmpl'. This essentially gives us a slice of all the
		// 'page' templates for the application.
		pages, err := filepath.Glob(filepath.Join(dir + "/pages", "*.page.gohtml"))
		if err != nil {
			return nil, err
		}

		// Loop through the pages one-by-one.
		for _, page := range pages {
			// Extract the file name (like 'home.page.tmpl') from the full file path
			// and assign it to the name variable.
			name := fmt.Sprintf("%s", filepath.Base(page))

			// The template.FuncMap must be registered with the template set before you
			// call the ParseFiles() method. This means we have to use template.New() to
			// create an empty template set, use the Funcs() method to register the
			// template.FuncMap, and then parse the file as normal.
			ts, err := template.New(name).Funcs(functions).ParseFiles(page)
			if err != nil {
				return nil, err
			}

			// Use the ParseGlob method to add any 'layout' templates to the
			// template set.
			ts, err = ts.ParseGlob(filepath.Join(dir + "/layouts", "*.layout.gohtml"))
			if err != nil {
				return nil, err
			}

			// Use the ParseGlob method to add any 'partial' templates to the
			// template set.
			ts, err = ts.ParseGlob(filepath.Join(dir + "/partials", "*.partial.gohtml"))
			if err != nil {
				return nil, err
			}

			// Add the template set to the cache, using the name of the page
			// (like 'home.page.tmpl') as the key.
			cache[module + "." + name] = ts
		}
	}

	return cache, nil
}

//func (t *Template) Render(name string, td *TemplateData) (*bytes.Buffer, error) {
func (t *Template) Render(name string) (*bytes.Buffer, error) {
	// Retrieve the appropriate template set from the cache based on the page name
	// (like 'home.page.tmpl'). If no entry exists in the cache with the
	// provided name, call the serverError helper method that we made earlier.
	ts, ok := t.templateCache[name]
	if !ok {
		return nil, fmt.Errorf("the template %s does not exist", name)
	}

	// Initialise new buffer.
	buf := new(bytes.Buffer)

	// Write the template to the buffer, instead of straight to the
	// http.ResponseWriter. Return if there's an error.
	//err := ts.Execute(buf, addDefaultData(td))
	err := ts.Execute(buf, nil)
	if err != nil {
		return nil, err
	}

	return buf, nil
}