package model

// request
type (
	CreateUserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UpdateUserRequest struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

// response
type (
	DataUserResponse struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}
)
