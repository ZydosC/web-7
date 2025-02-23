package main

import (
	"fmt"
	"net/http"
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, "Hello, web!")
}

func main() {
	http.HandleFunc("/get", countHandler)
	fmt.Println("starting server...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
