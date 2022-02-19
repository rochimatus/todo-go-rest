package request

type UserRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	RoleID   int    `json:"role_id" binding:"required"`
}
