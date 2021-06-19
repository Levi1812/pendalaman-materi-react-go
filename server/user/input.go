package user

type UserRegisterInput struct {
	Name      string `json:"name" binding:"required"`
	Address   string `json:"address" binding:"required"`
	DateBirth string `json:"date_birth" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
