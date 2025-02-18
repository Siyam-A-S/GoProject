package main

type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required_without=LastName"`
	LastName  string `json:"last_name" validate:"required_without=FirstName"`
	Email     string `json:"email" validate:"required,email"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name" validate:"required_without=LastName"`
	LastName  string `json:"last_name" validate:"required_without=FirstName"`
}

type UserResponse struct {
	ID          int             `json:"id"`
	FirstName   string          `json:"first_name"`
	LastName    string          `json:"last_name"`
	DisplayName string          `json:"display_name"`
	Emails      []EmailResponse `json:"emails"`
}

type EmailResponse struct {
	Address string `json:"address"`
	Primary bool   `json:"primary"`
}
