package main

import (
	"datagps/internal/api/handlers"
	"datagps/internal/api/middlewares"
	"datagps/internal/api/routes"
	"datagps/internal/repository"
	"datagps/internal/service"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatal("No existe archivo de configuracion .env")
	}

	dns := os.Getenv("database_url")
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting generic database object: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Ping the database to verify the connection is alive
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	//db.AutoMigrate(&models.Data{})
	dataRepo, groupRepo := repository.NewRepository(db)
	ss := service.NewAppServiceStellar()

	srv := service.NewAppService(dataRepo, groupRepo, ss)
	h := handlers.NewHandler(srv)

	mode_debug := os.Getenv("gin_mode_debug")
	log.Printf("Modo Debug: %v", mode_debug)
	if mode_debug == "false" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()

	r.Use(middlewares.RequestLogger())

	routes.SetupRoutes(r, h)

	port := os.Getenv("server_port")
	srvServer := ":" + port
	r.Run(srvServer)
}
