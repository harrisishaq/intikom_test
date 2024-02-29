package model

// request
type (
	CreateUserRequest struct {
		Name     string `json:"name" binding:"required,min=4"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	UpdateUserRequest struct {
		ID       string `json:"id"`
		Name     string `json:"name" binding:"required,min=4"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
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
