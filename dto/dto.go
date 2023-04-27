package dto

// SignUp
type DtoSignUp struct {
	Password string `json:"password" binding:"required"`

	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	UserType  string `json:"usertype" binding:"required"`
}

// LogIn
type DtoLogIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// GetUserById

// GetUsers
