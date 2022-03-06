package request

type ToDoListRequest struct {
	Task     string `json:"task" binding:"required"`
	ToDoID   int    `json:"todo_id" binding:"required"`
	StatusID int    `json:"status_id" binding:"required"`
}
