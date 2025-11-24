package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.Pattern, r.RemoteAddr)
		w.Write([]byte("Hello from GoApp!"))
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.Pattern, r.RemoteAddr)
		w.Write([]byte("test"))
	})

	log.Println("GoApp server started")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
