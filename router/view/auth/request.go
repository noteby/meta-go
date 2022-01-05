package auth

type RegisterRequest struct {
	Username  string `validate:"required,min=2,max=6"`
	Password  string `validate:"required,min=2,max=16"`
	Password2 string `validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
