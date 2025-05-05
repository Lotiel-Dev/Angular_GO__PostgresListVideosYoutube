package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Video struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

func main() {
	connStr := "user=postgres password=12345 dbname=youtube_links sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/videos", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, title, url FROM videos")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var videos []Video
		for rows.Next() {
			var v Video
			if err := rows.Scan(&v.ID, &v.Title, &v.URL); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			videos = append(videos, v)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(videos)
	})

	fmt.Println("Servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
