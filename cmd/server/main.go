package main

import (
	"net/http"

	"github.com/luiszkm/api/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBHost)
	

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8080", nil)
}