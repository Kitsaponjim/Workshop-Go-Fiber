package models

type Person struct {
	Email       string `json:"email" validate:"required,email"`
	Username    string `json:"username" validate:"required,min=5,max=20"`
	Password    string `json:"password" validate:"required"`
	IdLine      string `json:"idLine" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	Type        string `json:"type" validate:"required"`
	NameWebsite string `json:"nameWebsite" validate:"required"`
}
