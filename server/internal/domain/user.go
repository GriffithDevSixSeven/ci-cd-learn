package domain

type User struct {
	ID       int    
	UserName string `json:"UserName"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Credentials struct {
	UserName string
	Password string
}
