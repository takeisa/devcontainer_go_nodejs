package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// CORSポリシーの設定
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3080")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		fmt.Fprintf(w, "Hello World")
	})

	fmt.Println("Server listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
