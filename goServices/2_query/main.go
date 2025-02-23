package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/user", queryHandler)
	fmt.Println("Server is started on :8082")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера")
	}
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(400)
		w.Write([]byte("Параметр name не получен"))
		return
	}
	var out string
	out = "Hello, " + name + "!"
	fmt.Fprint(w, out)
}
