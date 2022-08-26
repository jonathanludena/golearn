package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	fmt.Println("Server listening on port 3000")
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
