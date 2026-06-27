package domain

type User struct {
	ID       int
	UserName string
	Email    string
	Password string
}

type Credentials struct {
	UserName string
	Password string
}
