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



