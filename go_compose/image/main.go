package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Image Service")
	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API : /image/upload, /image/crop")
	})
	http.HandleFunc("/image/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Upload image")
	})

	http.HandleFunc("/image/crop", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Crop image")
	})

	log.Fatal(http.ListenAndServe(":8088", nil))

}
