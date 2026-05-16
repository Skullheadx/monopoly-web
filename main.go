package main

import (

	"fmt"
	"io"
	"log"
	"net/http"

);

func main() {
	fmt.Println("ur mom")

	helloHandler:= func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello ur mom!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
