package dto

type SignUpDto struct {
	Firstname string `json:"firstname" validate:"required,max=50"`
	Lastname  string `json:"lastname" validate:"required,max=50"`
	Username  string `json:"username" validate:"required,min=3,max=32"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,password_strong"`
}

type SignInDto struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=3,max=32"`
}

type UserDto struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}
