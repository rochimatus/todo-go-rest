package request

type StatusRequest struct {
	Name string `json:"name" binding:"required"`
}
