package main // Define el paquete principal que ejecuta el programa (requerido en Go para que funcione como ejecutable)

import (
	"database/sql"  // Proporciona funciones para trabajar con bases de datos relacionales (como PostgreSQL)
	"encoding/json" // Permite codificar y decodificar datos JSON
	"fmt"           // Proporciona funciones de formato, como imprimir en consola
	"log"           // Manejo de logs (mensajes de error o depuración)
	"net/http"      // Proporciona funciones para crear servidores HTTP

	_ "github.com/lib/pq" // Importa el driver de PostgreSQL. El guion bajo indica que se importa solo por sus efectos secundarios (registro del driver)
)

// Video representa la estructura de un video que obtendremos desde la base de datos
type Video struct {
	ID    int    `json:"id"`    // ID único del video. El tag `json:"id"` indica cómo se verá en el JSON
	Title string `json:"title"` // Título del video
	URL   string `json:"url"`   // Enlace del video en YouTube
}

func main() {
	// Cadena de conexión a la base de datos
	// Cambia TU_CONTRASEÑA por tu contraseña real
	connStr := "user=postgres password=12345 dbname=youtube_links sslmode=disable"

	// Abre la conexión con la base de datos usando el driver "postgres"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err) // Si hay un error, lo imprime y termina el programa
	}

	// Define el manejador para la ruta /api/videos
	http.HandleFunc("/api/videos", func(w http.ResponseWriter, r *http.Request) {
		// Ejecuta la consulta SQL para obtener los videos
		rows, err := db.Query("SELECT id, title, url FROM videos")
		if err != nil {
			// Si hay error al consultar, devolver error HTTP 500
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close() // Asegura que se cierren los resultados cuando ya no se usen

		var videos []Video // Arreglo para almacenar los videos

		// Recorre cada fila del resultado
		for rows.Next() {
			var v Video
			// Lee los valores de la fila y los asigna a v
			if err := rows.Scan(&v.ID, &v.Title, &v.URL); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Agrega el video al arreglo
			videos = append(videos, v)
		}

		// Configura la cabecera de respuesta como JSON
		w.Header().Set("Content-Type", "application/json")

		// Convierte el arreglo de videos a JSON y lo escribe en la respuesta
		json.NewEncoder(w).Encode(videos)
	})

	// Muestra en consola que el servidor está funcionando
	fmt.Println("Servidor en http://localhost:8080")

	// Inicia el servidor en el puerto 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// El servidor escuchará en http://localhost:8080/api/videos para responder con los videos en formato JSON
