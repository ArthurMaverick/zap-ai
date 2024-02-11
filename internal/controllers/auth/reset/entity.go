package reset

type InputReset struct {
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,gte=8"`
	ChangePassword string `json:"change_password" validate:"required,gte=8"`
	Active         bool   `json:"active"`
}
