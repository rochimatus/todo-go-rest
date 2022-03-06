package response

import "time"

type AttachmentResponse struct {
	ID         int       `json:"id"`
	Url        string    `json:"url"`
	Caption    string    `json:"caption"`
	ToDoListId int       `json:"todolist_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
