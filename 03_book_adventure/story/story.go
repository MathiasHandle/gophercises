package story

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>CYA</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Story}}
			<p>{{.}}</p>
		{{ end }}
    <ul>
    {{range .Options}}
      <li><a href="/{{.Arc}}">{{.Text}}</a></li>
    {{end}}
    </ul>
	</body>
</html>
`

//  Template to be used for story
var tpl *template.Template

// Setting up default value for template
func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

// Chapter represents a CYOA story chapter. Each
// chapter includes its title, the story,
// and options available for the reader to take.
// If the options are empty it is
// assumed that you have reached the end of that particular
// story path. Arc is a link to a new story chapter and the Text
// is a text that user will see as option to choose from.
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

// Story represents a Choose Your Own Adventure story.
// Each key is the name of a story chapter (aka "arc"), and
// each value is a Chapter.
type Story map[string]Chapter

// ParseJSON will decode a story using the incoming reader
// and the encoding/json package. It is assumed that the
// provided reader has the story stored in JSON.
func ParseJSON(file *os.File) Story {
	d := json.NewDecoder(file)

	var jsonData Story
	err := d.Decode(&jsonData)
	if err != nil {
		fmt.Printf("Error parsing json data: %s", err)
	}

	return jsonData
}

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			log.Fatalf("Error executing template: %s\n", err)
		}
		return
	}

	http.Error(w, "chapter not found", http.StatusNotFound)
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	return path
}

// NewHandler will construct a http.Handler that will render
// the story provided.
// The default handler will use the full path (minus the / prefix)
// as the chapter name, defaulting to "intro" if the path is
// empty. The default template creates option links that follow
// this pattern.
func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	// default value for story and template
	h := handler{s, tpl, defaultPathFn}

	// applying optional values
	for _, opt := range opts {
		opt(&h)
	}

	return h
}

// HandlerOption are used with the NewHandler function to
// configure the http.Handler returned.
type HandlerOption func(h *handler)

// WithTemplate is an option to provide a custom template to
// be used when rendering stories.
func WithTemplate(t *template.Template) func(h *handler) {
	return func(h *handler) {
		h.t = t
	}
}

func WithPath(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}
