package auth

type User struct {
	ID        *int64  `json:"id,omitempty"`
	FirstName *string `json:"first_name,omitempty" validate:"required"`
	LastName  *string `json:"last_name,omitempty" validate:"required"`
	Email     *string `json:"email,omitempty" validate:"required"`
	Password  *string `json:"password" validate:"required"`
}
type PublicUser struct {
	*User
	Password *struct{} `json:"password,omitempty"`
}
