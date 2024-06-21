package model

import "time"

//Basics 基础信息模型
type Basics struct {
	ID        uint64    `json:"id" gorm:"primarykey"` //uuid唯一
	CreatedAt time.Time `json:"created_at"`           //创建时间
	UpdatedAt time.Time `json:"updated_at"`           //更新时间
}
