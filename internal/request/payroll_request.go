package request

type CreatePayrollPeriodRequest struct {
	StartDate string `json:"start_date" binding:"required" example:"2025-08-01"`
	EndDate   string `json:"end_date" binding:"required" example:"2025-08-31"`
}
