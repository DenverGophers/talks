package main

import (
	"io"
	"net/http"
	"path/filepath"
)

const someHtml = `<!DOCTYPE html>
<html>

<head>
	<link rel="stylesheet" type="text/css" href="/css/screen.css">
</head>

<body>
	<h2>Hello</h2>
</body>

</html>
`

func main() {

	var (
		uri = "/css/"
		fs  = http.FileServer(http.Dir(filepath.Join("./assets", uri)))
	)

	http.Handle(uri, http.StripPrefix(uri, fs))

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, someHtml)
	})

	if err := http.ListenAndServe(":4002", nil); err != nil {
		panic(err)
	}

}
