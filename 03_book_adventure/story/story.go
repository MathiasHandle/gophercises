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

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type Story map[string]Chapter

// Parses JSON file to a struct
func ParseJSON(file *os.File) Story {
	d := json.NewDecoder(file)

	var jsonData Story
	err := d.Decode(&jsonData)
	if err != nil {
		fmt.Printf("Error parsing json data: %s", err)
	}

	return jsonData
}

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

type handler struct {
	s Story
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		tpl := template.Must(template.New("").Parse(defaultHandlerTmpl))
		err := tpl.Execute(w, chapter)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			log.Fatalf("Error executing template: %s\n", err)
		}
		return
	}

	http.Error(w, "chapter not found", http.StatusNotFound)
}
