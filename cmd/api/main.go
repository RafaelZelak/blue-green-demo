package main

import (
	"log"
	"net/http"

	"github.com/RafaelZelak/blue-green-demo/internal/routes"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's alive!"))
	})

	http.HandleFunc("/users", routes.UsersHandler)
	http.HandleFunc("/now", routes.NowHandler)

	http.HandleFunc("/users", routes.UsersHandler)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
