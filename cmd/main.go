package main

import (
	_ "github.com/Rambudhi/payslip/docs" // 👈 penting buat swagger docs
	"github.com/Rambudhi/payslip/internal/app"
)

// @title Payslip API
// @version 1.0
// @description API service for Payslip.
// @host localhost:8009
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app.Run()
}
