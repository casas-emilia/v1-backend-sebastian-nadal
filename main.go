package main

import (
	"log"
	"v1-backend-sebastian-nadal/configs"
	"v1-backend-sebastian-nadal/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env, asegurarse de que las variables de entorno estén configuradas")
	}

	// Conectar a la base de datos
	configs.ConnectToDB()
}

func main() {
	log.Println("Iniciando el servidor...")
	router := routers.SetupRouter()

	// Ruta de health check
	log.Println("Configurando ruta de health check...")
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	log.Println("Servidor corriendo en el puerto 8080")
	router.Run(":8080")
}
