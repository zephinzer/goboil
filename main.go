package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, GetUniversalMessage())
	})
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port " + port)
	http.ListenAndServe("0.0.0.0:"+port, nil)
}
