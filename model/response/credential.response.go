package response

type CredentialResponse struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
}
