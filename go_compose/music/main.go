package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Music Service")
	http.HandleFunc("/music", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API : /music/play, /music/stop")
	})
	http.HandleFunc("/music/play", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Play music")
	})

	http.HandleFunc("/music/stop", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "stop music")
	})

	log.Fatal(http.ListenAndServe(":8088", nil))

}
