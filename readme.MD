# Desafio W

## Detalles tecnicos

### Backend 🔧

- Lenguaje: Go
- Base de datos: MongoDB

#### Estructura del backend

- 📁 db ➡️ Conexión y consultas a la base de datos
- 📁 handler ➡️ Punto de entrada http
- 📁 middleware ➡️ Capas entrante para las rutas.
- 📁 models ➡️ Representación de la data entrante e información saliente.
- 📁 repositories ➡️ Enlace entre un controlador y el modelo, reutilizable en otras partes.
- 📁 routes ➡️ Acciones asociadas a las rutas del handler
- 📁 seeder ➡️ Collecciones que son cargadas cuando se inicia el proyecto
- 📁 static ➡️ Archivos estaticos
- 📁 resources ➡️ Vistas

#### Dependencias del lenguaje 🔌

- github.com/gorilla/mux
- github.com/rs/cors
- go.mongodb.org/mongo-driver

### Frontend

- Lenguaje: JavaScript ES6
- Framework: ReactJs
- Repositorio: https://github.com/fenriz07/frw

#### Dependencias del front

- react
- react-bootstrap
- react-dom
- react-scripts
- bootstrap
- cra-template
- lodash
- nodesass

## Instrucciones

1. Clone este repositorio o descargelo
2. Entre al directorio raíz del repositorio
3. Ejecute `sudo docker-compose up`
4. Abra su explorador favorito y escriba la dirección: `http://localhost:8080`

## Docker

Proyecto creado con 2 contenedores, vinculados en una red.

## Tests

1.Ejecutar tests: `go test $(pwd)/tests` dentro del directorio raíz del proyecto

## Parametros para buscar

![imagen](https://user-images.githubusercontent.com/9199380/87319056-1fefd900-c4f7-11ea-81b3-aaf9cb2b47ca.png)

Puede buscar por id del producto: rango(1,3000), por su marca o descripción. *sjlzxeo*

## Notas

1. Reconstruir `sudo docker-compose up --build -d`
