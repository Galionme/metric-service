package handlers

import (
	"github.com/Galionme/metric-service.git/internal/storage"
	"html/template"
	"net/http"
)

func HomeMetrics(res http.ResponseWriter, req *http.Request) {

	tmpl := `
  <!DOCTYPE html>
  <html lang="en">
  <head>
   <meta charset="UTF-8">
   <title>service-template</title>
  </head>
  <body>
   <h1>List of Names and Values:</h1>
   <ul>
    {{range $key, $value := .}}
     <li>{{$key}}: {{$value}}</li>
    {{end}}
   </ul>
  </body>
  </html>
 `
	data := storage.GlobalMemStorage.GetAll()

	t := template.Must(template.New("list").Parse(tmpl))

	t.Execute(res, data)

	res.WriteHeader(http.StatusOK)
}
