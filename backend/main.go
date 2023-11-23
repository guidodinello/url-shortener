// hello world
package main

import (
	"fmt"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

func main() {
	http.HandleFunc("/", HelloWorldHandler)
	http.ListenAndServe(":8080", nil)
}
