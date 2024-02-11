package register

type InputRegister struct {
	FullName string `json:"full_name" validate:"required,lowercase"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}
