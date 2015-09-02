package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandlerFunc)
	mux.HandleFunc("/basic", BasicHandlerFunc)
	mux.HandleFunc("/no-cache", NoCacheHandlerFunc)
	mux.HandleFunc("/no-store", NoStoreHandlerFunc)
	mux.HandleFunc("/ultimate-no-store", UltimateNoStoreHandlerFunc)
	http.ListenAndServe(":8888", mux)
}

func RootHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	addData(w)
}

func BasicHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	addData(w)
}

func NoCacheHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Type", "text/html")
	addData(w)
}

func NoStoreHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "text/html")
	addData(w)
}

func UltimateNoStoreHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "text/html")
	curr_time, _ := time.Now().MarshalText()
	io.WriteString(w, "<body onunload='' onbeforeunload=''>")
	io.WriteString(w, fmt.Sprintf("%x", md5.Sum(curr_time)))
	io.WriteString(w, "<br><a href='./'>Root</a>")
	io.WriteString(w, "<br><a href='./basic'>Basic</a>")
	io.WriteString(w, "<br><a href='./no-cache'>No Cache</a>")
	io.WriteString(w, "<br><a href='./no-store'>No Store</a>")
	io.WriteString(w, "<br><a href='./ultimate-no-store'>Ultimate No Store</a>")
	io.WriteString(w, "<br><a href='https://www.google.co.in'>Google</a>")
	io.WriteString(w, "</body>")
}

func addData(w http.ResponseWriter) {
	curr_time, _ := time.Now().MarshalText()
	io.WriteString(w, "<body>")
	io.WriteString(w, fmt.Sprintf("%x", md5.Sum(curr_time)))
	io.WriteString(w, "<br><a href='./'>Root</a>")
	io.WriteString(w, "<br><a href='./basic'>Basic</a>")
	io.WriteString(w, "<br><a href='./no-cache'>No Cache</a>")
	io.WriteString(w, "<br><a href='./no-store'>No Store</a>")
	io.WriteString(w, "<br><a href='./ultimate-no-store'>Ultimate No Store</a>")
	io.WriteString(w, "<br><a href='https://www.google.co.in'>Google</a>")
	io.WriteString(w, "</body>")
}
