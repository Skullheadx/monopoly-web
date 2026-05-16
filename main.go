package main

import (

	"fmt"
	"io"
	"log"
	"net/http"

);

func main() {
	fmt.Println("monopoly-web")

	healthHandler:= func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Status: healthy\n")
	}

	http.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
