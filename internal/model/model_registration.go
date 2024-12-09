package model

// UserLogin represents a user login entity with Login and Password fields.
type UserLogin struct {
	Login    string
	Password string
}

// UserRegistration represents a user registration entity with Login and Password fields.
type UserRegistration struct {
	Login           string
	Password        string
	ConfirmPassword string
}
