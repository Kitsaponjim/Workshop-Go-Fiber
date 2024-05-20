package models

type Person struct {
	Email       string `json:"email" validate:"required,email"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required,min=6,max=20"`
	IdLine      string `json:"idLine"`
	Phone       string `json:"phone" validate:"required,numeric,max=10"`
	Type        string `json:"type" validate:"required"`
	NameWebsite string `json:"nameWebsite" validate:"required,min=2,max=30"`
}
