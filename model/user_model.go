package model

type User struct {
	BaseModel
	FullName string `faker:"name"`
	Email    string `faker:"email"`
	Password string `faker:"-"`
	Role     string `faker:"oneof: admin, guest"`
}
