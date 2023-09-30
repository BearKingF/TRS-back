package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"` //主键
	Username string `json:"username"`
	Sex      string `json:"sex"`
	PhoneNum string `json:"phone_num"`
	Email    string `json:"email"`
	Major    string `json:"major"`
	Password string `json:"-"` //返回给前端时忽略
	//new!!
	Type      uint `json:"type"`    //1 表示普通用户 2 表示管理员用户
	TeamID    int  `json:"team_id"` //未加入团队时记为-1
	IsCaptain bool `json:"-"`       //前端感觉不会用到
}
