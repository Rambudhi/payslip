package request

type AttendanceRequest struct {
	Date string `json:"date" binding:"required"` // format: YYYY-MM-DD
}
