package model

import "time"

// 通用头
// todo ： 考虑使用 gorm.Model

type Universal struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id" mapstructure:"id"`
	CreatedAt time.Time `json:"created_at" mapstructure:"-"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"-"`
}
