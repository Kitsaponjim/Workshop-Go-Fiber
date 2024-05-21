package models

import "gorm.io/gorm"

type Person struct {
	Email       string `json:"email" validate:"required,email"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required,min=6,max=20"`
	IdLine      string `json:"idLine"`
	Phone       string `json:"phone" validate:"required,numeric,max=10"`
	Type        string `json:"type" validate:"required"`
	NameWebsite string `json:"nameWebsite" validate:"required,min=2,max=30"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Count      int       ` json:"count"`
	Data       []DogsRes `json:"data"`
	Name       string    `json:"name"`
	CountRed   int       `json:"count_red"`
	CountGreen int       `json:"count_green"`
	CountPink  int       `json:"count_pink"`
	CountNo    int       `json:"count_nocolor"`
}

//7.0.1
type Company struct {
	gorm.Model
	CompanyName    string `json:"name"`
	CompanyAddress string `json:"address"`
	CompanyTel     string `json:"tel"`
	Company_Id     string `json:"id"`
}

// Project 2
type UserProfile struct {
	gorm.Model
	Employee_id int    `json:"employee_id"`
	Name        string `json:"name"`
	LastName    string `json:"lastName"`
	Birthday    string `json:"birthday"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	Tel         string `json:"tel"`
}
