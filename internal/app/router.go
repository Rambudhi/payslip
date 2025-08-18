package app

import (
	"runtime"

	"github.com/Rambudhi/payslip/internal/handler"
	"github.com/Rambudhi/payslip/internal/middleware"
	"github.com/Rambudhi/payslip/internal/queue"
	"github.com/Rambudhi/payslip/internal/repository"
	"github.com/Rambudhi/payslip/internal/service"
	"github.com/Rambudhi/payslip/internal/worker"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// === Auth ===
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)
	r.POST("/login", authHandler.Login)

	// === Queue & Worker ===
	bufferSize := runtime.NumCPU() * 100
	workerCount := runtime.NumCPU() * 2

	q := queue.NewQueue(bufferSize)
	w := worker.NewWorker(q)

	payrollRepo := repository.NewPayrollPeriodRepository(db)
	worker.RegisterPayrollPeriodWorker(w, payrollRepo)

	logRepo := repository.NewLogRepository(db)
	worker.RegisterLogActivityWorker(w, logRepo)

	w.Start(workerCount)

	payrollService := service.NewPayrollPeriodService(q)
	payrollHandler := handler.NewPayrollPeriodHandler(payrollService)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/me", authHandler.Me)

		api.Use(middleware.RoleMiddleware("admin"))
		{
			api.POST("/payroll-period", payrollHandler.Create)
		}
	}
}
