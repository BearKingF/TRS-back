package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Sex      string `json:"sex"`
	PhoneNum string `json:"phone_num"`
	Email    string `json:"email"`
	Major    string `json:"major"`
	Password string `json:"-"` //返回给前端时忽略
}
