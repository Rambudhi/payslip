package app

import (
	"fmt"
	"os"

	"github.com/Rambudhi/payslip/internal/middleware"
	"github.com/Rambudhi/payslip/internal/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Run() {
	LoadEnv()

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Starting Payslip Service...")

	db := util.InitPostgres()
	r := gin.Default()
	r.Use(cors.New(middleware.GetCorsConfig()))

	RegisterRoutes(r, db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	if err := r.Run(addr); err != nil {
		logrus.Fatal(err)
	}
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file found")
	}
}
