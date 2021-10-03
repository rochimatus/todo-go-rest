package response

type ToDoResponse struct {
	ID    int                `json:"id"`
	Title string             `json:"title"`
	User  CredentialResponse `json:"user"`
}
