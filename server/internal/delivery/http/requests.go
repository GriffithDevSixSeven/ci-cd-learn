package http

type User struct {
	UserName string `validate:"required,max=40"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,max=100"`
}

type CreateUserRequest struct {
	UserName string `json:"UserName" validate:"required,max=40"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=100"`
}

type LoginUserRequest struct {
	UserName string `json:"UserName" validate:"required,max=40"`
	Password string `json:"password" validate:"required,max=100"`
}

type DeleteUserRequest struct {
	UserName string `json:"UserName" validate:"required,max=40"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=100"`
}