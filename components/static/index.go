package static

import (
	"html/template"
	"net/http"
)

type indexData struct {
	Greeting string
}

const indexHTML = `
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Base Go</title>
</head>

<body>
  <h1>{{ .Greeting }}</h1>
</body>

</html>
`

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("index").Parse(indexHTML))
	t.Execute(w, indexData{
		Greeting: "Hello world!",
	})
}

// NewStaticMux Does stuff
func NewStaticMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	return mux
}
