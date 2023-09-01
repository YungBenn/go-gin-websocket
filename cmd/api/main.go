package main

import (
	"log"
	"net/http"
	"workshop-web-golang/config"
	_ "workshop-web-golang/docs/statik"
	"workshop-web-golang/internal/controller"
	"workshop-web-golang/internal/db"
	"workshop-web-golang/internal/router"
	ws "workshop-web-golang/internal/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
)

func main() {
	env, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	configCORS := cors.DefaultConfig()
	configCORS.AllowOrigins = []string{"http://localhost:4000", "http://localhost:3000"}
	configCORS.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH"}

	r.Use(cors.New(configCORS))
	// r.Use(cors.Default())

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	swaggerHandler := http.StripPrefix("/docs/", http.FileServer(statikFS))
	r.GET("/docs/*any", func(c *gin.Context) {
		swaggerHandler.ServeHTTP(c.Writer, c.Request)
	})

	config := &db.ConfigDB{
		Host:     env.DB_HOST,
		User:     env.DB_USER,
		Password: env.DB_PASS,
		DBName:   env.DB_NAME,
		Port:     env.DB_PORT,
		SSLMode:  env.DB_SSLMODE,
	}

	db := db.ConnectDB(config)

	router := router.NewRouter(controller.NewStudentsController(db))
	router.StudentRoutes(&r.RouterGroup)

	// websocket endpoint
	r.GET("/ws", ws.HandleWebSocket)

	go ws.StartWebSocketServer(env.WS_PORT)

	r.Run(":" + env.PORT)
}
