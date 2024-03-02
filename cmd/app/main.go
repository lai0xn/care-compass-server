package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lai0xn/hackiwna-backend/internal/storage"
	"github.com/lai0xn/hackiwna-backend/internal/storage/migrations"
	"github.com/lai0xn/hackiwna-backend/internal/transport/routes"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	storage.Connect()
	migrations.Migrate()
	routes.Setup(r)
	r.Run()
}
