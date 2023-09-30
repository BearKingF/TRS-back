package models

type Team struct {
	TeamID    uint   `json:"team_id" ,gorm:"primary_key"` //主键
	TeamName  string `json:"team_name"`                   //团队名称不可重复
	CaptainID uint   `json:"captain_id"`
	Password  string `json:"-"`
	Total     uint   `json:"total"`
}
