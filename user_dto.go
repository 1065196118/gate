package dto

type UserCreateDTO struct {
	Name      string `json:"name" binding:"required,min=3,max=100" validate:"is-cool"`
	Username  string `json:"username" binding:"required,min=3,max=100"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required,min=8,max=12"`
	PhoneCode string `json:"phone_code" binding:"required,min=2,max=5"`
	Password  string `json:"password" binding:"required,min=3,max=20"`
	Gender    string `json:"gender" binding:"required,min=3,max=20"`
}

type UserLoginDTO struct {
	ID       uint64 `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (UserLoginDTO) TableName() string {
	return "users"
}

func (UserCreateDTO) TableName() string {
	return "users"
}
