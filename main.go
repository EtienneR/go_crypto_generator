package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"hash"
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Message, Secret, Md5, Sha1, Sha256, Sha512, HmacMd5, HmacSha1, HmacSha256, HmacSha512 string
}

func Encrypt(format string, message string, secret string) string {
	var h hash.Hash
	var key []byte

	if format == "hmacmd5" || format == "hmacsha1" || format == "hmacsha256" || format == "hmacsha512" {
		key = []byte(secret)
	}

	if format == "md5" {
		h = md5.New()
	} else if format == "sha1" {
		h = sha1.New()
	} else if format == "sha256" {
		h = sha256.New()
	} else if format == "sha512" {
		h = sha512.New()
	} else if format == "hmacmd5" {
		h = hmac.New(md5.New, key)
	} else if format == "hmacsha1" {
		h = hmac.New(sha1.New, key)
	} else if format == "hmacsha256" {
		h = hmac.New(sha256.New, key)
	} else if format == "hmacsha512" {
		h = hmac.New(sha512.New, key)
	}

	h.Write([]byte(message))

	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	port := ":3000"
	log.Println("Listen http://127.0.0.1" + port)

	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("home.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf(err.Error())
	}

	message := r.FormValue("message")

	var data interface{}

	if message != "" {
		secret := r.FormValue("secret")

		if secret != "" {
			data = Data{
				message,
				secret,
				Encrypt("md5", message, ""),
				Encrypt("sha1", message, ""),
				Encrypt("sha256", message, ""),
				Encrypt("sha512", message, ""),
				Encrypt("hmacmd5", message, secret),
				Encrypt("hmacsha1", message, secret),
				Encrypt("hmacsha256", message, secret),
				Encrypt("hmacsha512", message, secret),
			}
		} else {
			data = Data{
				message,
				"",
				Encrypt("md5", message, ""),
				Encrypt("sha1", message, ""),
				Encrypt("sha256", message, ""),
				Encrypt("sha512", message, ""),
				"",
				"",
				"",
				"",
			}
		}
	}

	if message != "" {
		json, _ := json.Marshal(data)
		log.Println(string(json))
	}

	tmpl.Execute(w, data)
}
