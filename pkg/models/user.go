package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// Book model
// https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go
// https://gorm.io/docs/models.html
type User struct {
	Username  string         `from:"username" json:"username" gorm:"primarykey"`
	Password  string         `from:"password" json:"password" gorm:""`
	Name      string         `from:"name" json:"name" gorm:"index"`
	Credit    json.Number    `from:"credit" json:"rating" gorm:""`
	CreatedAt time.Time      `from:"created_at" json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `from:"updated_at" json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `from:"deleted_at" json:"deleted_at" gorm:"index"`
}
