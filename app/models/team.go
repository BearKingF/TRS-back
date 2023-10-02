package models

type Team struct {
	TeamID    uint   `json:"team_id" gorm:"primaryKey"` //主键 （中间不加逗号！！！）
	TeamName  string `json:"team_name"`                 //团队名称不可重复
	CaptainID uint   `json:"captain_id"`
	Password  string `json:"-"`
	Total     uint   `json:"total"`
	Status    uint   `json:"status"` //状态：2表示未提交报名，1表示已提交
}
