package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r * http.Request)  {
	response := fmt.Sprintf("This is reponse to: %s%s", r.Host, r.URL.Path)

	_, err := w.Write([]byte(response))
	if err != nil {
		fmt.Printf("Could not send response because of an error: %s", err.Error())
	}
}

func main() {
	http.HandleFunc("/echo", Handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server returned with error: %s", err.Error())
	}
}
