package request

type ToDoRequest struct {
	Title string `json:"title" binding:"required"`
}
