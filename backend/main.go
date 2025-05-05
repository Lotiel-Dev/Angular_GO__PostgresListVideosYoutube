package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/videos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Aquí estará nuestra lista de videos")
	})

	fmt.Println("Servidor en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
