package request

type OvertimeRequest struct {
	Date  string `json:"date" binding:"required"` // format: YYYY-MM-DD
	Hours int    `json:"hours" binding:"required,min=1,max=3"`
}
