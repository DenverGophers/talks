package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

type Upload struct {
	Data     []byte
	Name     string
	Password string
	Mode     string
}

func (upload *Upload) Valid() bool {
	if len(upload.Data) == 0 {
		log.Println("invalid file")
		return false
	} else if upload.Name == "" {
		log.Print("invalid filename")
		return false
	} else if upload.Password == "false" {
		log.Println("invalid password")
		return false
	} else if (upload.Mode != "encrypt") && (upload.Mode != "decrypt") {
		log.Println("invalid mode", upload.Mode)
		return false
	}

	return true
}

func serverError(w http.ResponseWriter, msg string) {
	log.Println(msg)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(msg))
}

func parseForm(w http.ResponseWriter, r *http.Request) *Upload {
	var upload Upload

	mp_rdr, err := r.MultipartReader()
	if err != nil {
		serverError(w, "error reading multipart: "+err.Error())
		return nil
	}

	for {
		part, err := mp_rdr.NextPart()

		if err == io.EOF {
			break
		}

		switch part.FormName() {
		case "file":
			upload.Data = readPart(part)
		case "filename":
			upload.Name = string(readPart(part))
		case "password":
			upload.Password = string(readPart(part))
		case "mode":
			upload.Mode = string(readPart(part))
		default:
			serverError(w, "invalid form part: "+part.FormName())
			return nil
		}
	}

	return &upload
}

func readPart(part *multipart.Part) []byte {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, part)
	if err != nil {
		return []byte{}
	}
	return buf.Bytes()
}

func dispatch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		rootHandler(w, r)
		return
	} else if r.Method == "POST" {
		upload := parseForm(w, r)
		if !upload.Valid() {
			serverError(w, "invalid form")
			return
		} else if upload.Mode == "encrypt" {
			encrypt(w, upload)
		} else if upload.Mode == "decrypt" {
			decrypt(w, upload)
		}
	}
}

func main() {
	port := flag.String("p", "8080", "port to listen on")
	flag.Parse()

	http.HandleFunc("/", dispatch)
	address := "127.0.0.1:" + *port
	log.Print("listening on http://", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
