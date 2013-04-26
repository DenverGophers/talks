package main

import "log"
import "html/template"
import "net/http"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	err := rootTemplate.Execute(w, nil)
	if err != nil {
		log.Fatal("[!] serving root: ", err.Error())
	}
}

var rootTemplate = template.Must(template.New("root").Parse(`
<!doctype.html>
<html>
  <head>
    <meta charset="utf8" />
    <title>April Golang Meetup Demo</title>
    <link href="//netdna.bootstrapcdn.com/twitter-bootstrap/2.3.1/css/bootstrap-combined.no-icons.min.css"
          rel="stylesheet">
    <style type="text/css">
        html,
        body {
            height: 100%;
        }
        #wrap {
            min-height: 100%;
            height: auto !important;
            height: 100%;
            margin: 0 auto -60px;
        }
        #push, #footer {
            height: 60px;
        }
        #footer {
            background-color: #f5f5f5;
        }
        @media (max-width: 767px) {
            #footer {
                margin-left: -20px;
                margin-right: -20px;
                padding-left: 20px;
                padding-right: 20px;
            }
        }
        .container {
            width: auto;
            max-width: 680px;
        }
        .container .credit {
            margin: 20px 0;
        }
    </style>
  </head>

  <body>
    <div id="wrap">
      <div class="container">
        <div class="row">
          <div class="span4"></div>
          <div class="span4">
          <h1>April Demo</h1>
          <p>This is a local webapp to encrypt and decrypt files using a
          password.</p>
          <form action="/" name="demo" method="POST" enctype="multipart/form-data">
            <label>File: </label>
            <input type="file" name="file" size="45"><br>
            <label>Output Filename: </label>
            <input type="text" name="filename" style="height: 2em"><br>
            <label>Mode: </label>
            <input type="radio" name="mode" value="encrypt" checked> Encrypt</input> 
            <input type="radio" name="mode" value="decrypt"> Decrypt</input>
            <br>
            <label>Password: </label>
            <input type="password" name="password">
            <input type="submit" value="Submit">
          </form>
          </div>
          <div class="span4"></div>
        </div>
      </div>
    </div>
  </body>
</html>`))
