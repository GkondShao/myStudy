package models

type User struct {
	Id       string `orm:"pk"`
	Nick     string
	Info     string `orm:"null"`
	Email    string `orm:"null"`
	Password string
}
