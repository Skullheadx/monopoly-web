package main

import (

	"fmt"
	"io"
	"log"
	"net/http"

);


func main() {
	fmt.Println("monopoly-web backend")

	// register routes
	http.HandleFunc("/health", healthHandler)


	// listen and serve
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Status: healthy\n")
}
