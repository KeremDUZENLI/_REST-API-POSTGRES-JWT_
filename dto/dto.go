package dto

// SignUp
type DtoSignUp struct {
	Password string `gorm:"column:password"`
	Token    string `gorm:"column:token"`

	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	UserType  string `gorm:"column:user_type"`
}

// LogIn
type DtoLogIn struct {
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

// GetUserById

// GetUsers
