package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from html</h1>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("server started running at port 3000...")
	http.ListenAndServe(":3000", nil)
}
