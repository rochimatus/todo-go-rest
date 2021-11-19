package response

type ToDoListResponse struct {
	ID     int          `json:"id"`
	ToDo   ToDoResponse `json:"todo"`
	Task   string       `json:"task"`
	Status string       `json:"status"`
}
