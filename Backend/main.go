package main

import (
	"fmt"
	"log/slog"

	"Github.com/Hams/abdi/InventoryManagmentSystem/infra"
	"Github.com/Hams/abdi/InventoryManagmentSystem/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("Initialing the env variables")
	infra.InitEnv()
	config := infra.Configurations

	slog.Info("Connecting to database")
	infra.ConnectDB()
	slog.Info("✅ Connected to database successfully")

	r := gin.Default()

	
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"http://localhost:5003","http://localhost:5173",},
	// 	AllowMethods: []string{"GET","POST","PUT","DELETE","OPTIONS",},
	// 	AllowHeaders: []string{"Origin","Content-Type","Authorization",},
	// }))

	routes.RegisterRoutes(r)

	slog.Info("✅ Application is running on port :5000")

	if err := r.Run(fmt.Sprintf(":%s", config.Port)); err != nil {
		slog.Error("Server failed", "error", err)
	}
}