package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!\n")
	fmt.Println(ioutil.ReadAll(r.Body))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Simple Helloworld server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

