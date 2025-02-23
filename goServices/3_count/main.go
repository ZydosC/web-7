package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var countRes int = 0

func countHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != "GET" && r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
	}
	if r.Method == "GET" {
		fmt.Fprint(w, countRes)
	}
	if r.Method == "POST" {
		count, err := strconv.Atoi(r.FormValue("count"))
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("Введено не число"))
			return
		}
		countRes += count
		w.Write([]byte("Добавлено"))
	}

}

func main() {
	http.HandleFunc("/count", countHandler)
	fmt.Println("starting server...")
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
