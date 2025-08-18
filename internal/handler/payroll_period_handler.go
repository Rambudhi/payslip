package handler

import (
	"net/http"

	"github.com/Rambudhi/payslip/internal/request"
	"github.com/Rambudhi/payslip/internal/service"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type PayrollPeriodHandler struct {
	service service.PayrollPeriodService
}

func NewPayrollPeriodHandler(s service.PayrollPeriodService) *PayrollPeriodHandler {
	return &PayrollPeriodHandler{service: s}
}

// Create Payroll Period
// @Summary Create payroll period
// @Description Create a new payroll period and enqueue for processing
// @Security BearerAuth
// @Tags PayrollPeriod
// @Accept json
// @Produce json
// @Param payroll_period body request.CreatePayrollPeriodRequest true "Payroll Period"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/payroll-period [post]
func (h *PayrollPeriodHandler) Create(c *gin.Context) {
	var req request.CreatePayrollPeriodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ipAddr := c.ClientIP()
	var userID *uint
	if val, exists := c.Get("userID"); exists {
		if id, ok := val.(uint); ok {
			userID = &id
		}
	}

	requestID := c.GetHeader("X-Request-ID")
	if requestID == "" {
		requestID = uuid.New().String()
	}

	if err := h.service.Create(req, userID, ipAddr, requestID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payroll period queued successfully"})
}
