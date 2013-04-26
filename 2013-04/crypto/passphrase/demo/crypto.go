package main

import (
	"github.com/kisom/gocrypto/chapter2/symmetric"
	"github.com/kisom/gocrypto/chapter4/hash"
	"net/http"
)

func decrypt(w http.ResponseWriter, upload *Upload) {
	salt := upload.Data[:hash.SaltLength]
	enc := upload.Data[hash.SaltLength:]
	key := hash.DeriveKeyWithSalt(upload.Password, salt)
	if key == nil {
		serverError(w, "failed to generate key")
		return
	}
	dec, err := symmetric.Decrypt(key.Key, enc)
	if err != nil {
		serverError(w, "encryption failure: "+err.Error())
		return
	}
	w.Header().Add("content-type", "application/octet-stream")
	w.Header().Add("content-disposition", "attachment; filename="+upload.Name)
	w.Write(dec)
}

func encrypt(w http.ResponseWriter, upload *Upload) {
	key := hash.DeriveKey(upload.Password)
	if key == nil {
		serverError(w, "failed to generate key")
		return
	}

	out := key.Salt
	enc, err := symmetric.Encrypt(key.Key, upload.Data)
	if err != nil {
		serverError(w, "encryption failure: "+err.Error())
		return
	}
	out = append(out, enc...)
	w.Header().Add("content-type", "application/octet-stream")
	w.Header().Add("content-disposition", "attachment; filename="+upload.Name)
	w.Write(out)
}
