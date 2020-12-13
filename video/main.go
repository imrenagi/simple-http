package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/video", video)
	http.ListenAndServe("0.0.0.0:80", nil)
}

func video(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	w.Write([]byte(fmt.Sprintf("data video service: %s", name)))
}

func home(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	w.Write([]byte(fmt.Sprintf("hc video service: %s", name)))
}
