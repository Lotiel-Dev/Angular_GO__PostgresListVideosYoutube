# PASO 1
## Angular_GO__PostgresListVideosYoutube
para crear el proyecto en Angular promero debemos crear la aplicaicon en angular:
```bash
    ng new youtube-video-list --routing --style=scss
```
y lo iniciamos con:
```bash 
ng serve
```
esto iniciara un servidor local en
```bash
http://localhost:4200
```
## GO
Iniciamos el modulo de GO para permitirnos manejar dependencias
```bash
go mod init github.com/tuusuario/youtube-video-api
```
ejecutamos el backend con:
```bash
go run main.go
```
y los abrimos con:
```bash
http://localhost:8080/api/videos
```
# PASO 2: Conectar Golang con PostgreSQL
## Instalamos el controlador de PostgresSQL para Go
desde la carpeta backend
```bash
go get github.com/lib/pq
```
esto descarga la libreria y la deja lista para usar

## Crear la base de datos y tabla en PostgreSQL
Abre tu consola de PostgreSQL:
```bash
psql -U postgres
```
Crea una base de datos:
```bash
CREATE DATABASE youtube_links;
```
Entra en la base de datos:
```bash
\c youtube_links
```
Crea una tabla videos:
```bash
CREATE TABLE videos (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  url TEXT NOT NULL
);
```
Inserta un ejemplo:
```bash
INSERT INTO videos (title, url) VALUES
('Video 1', 'https://www.youtube.com/watch?v=dQw4w9WgXcQ');
```
## Escribir código en Go para conectar y leer
Corre el servidor:
```bash
go run main.go
```
Abre tu navegador o usa Postman para ir a:
```bash
http://localhost:8080/api/videos
```
deberiamos ver:
```bash
[
  {
    "id": 1,
    "title": "Video 1",
    "url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
  }
]
```

# PASO 3: Consumir la API de Golang desde Angular
## A. Crear un servicio Angular para llamar al backend
desde la carepeta de frontend ejecutaremos el siguiente comando para generar el servicio:
```bash
ng generate service services/video
```
y creara dos archivos:
- src/app/services/video.service.ts
- src/app/services/video.service.spec.ts (para pruebas, lo usaremos más adelante)