package handler

import (
	"net/http"

	"github.com/Rambudhi/payslip/internal/request"
	"github.com/Rambudhi/payslip/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AttendanceHandler struct {
	service service.AttendanceService
}

func NewAttendanceHandler(s service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{service: s}
}

// SubmitAttendance
// @Summary Submit attendance
// @Description Submit attendance for the current day. Attendance is not allowed on weekends. Duplicate same-day attendance is ignored.
// @Security BearerAuth
// @Tags Attendance
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "attendance submitted successfully"
// @Failure 400 {object} map[string]string "bad request / validation error"
// @Failure 500 {object} map[string]string "internal server error"
// @Router /api/attendance [post]
func (h *AttendanceHandler) SubmitAttendance(c *gin.Context) {
	ipAddr := c.ClientIP()

	var userID uint
	if val, exists := c.Get("userID"); exists {
		if id, ok := val.(uint); ok {
			userID = id
		}
	}

	req := request.SubmitAttendanceRequest{
		UserID: userID,
	}

	requestID := c.GetHeader("X-Request-ID")
	if requestID == "" {
		requestID = uuid.New().String()
	}

	if err := h.service.SubmitAttendance(req, ipAddr, requestID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "attendance submitted successfully"})
}
