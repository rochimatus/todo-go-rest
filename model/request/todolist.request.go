package request

type ToDoListRequest struct {
	Task     string `json:"task" binding:"required"`
	ToDoID   int    `json:"todo" binding:"required"`
	StatusID int    `json:"status" binding:"required"`
}
