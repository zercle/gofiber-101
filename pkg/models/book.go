package models

import (
	"time"

	"gorm.io/gorm"
)

// Book model
// https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go
// https://gorm.io/docs/models.html
type Book struct {
	ID        uint           `from:"id" json:"id" gorm:"primarykey"`
	Title     string         `from:"title" json:"title" gorm:"index"`
	Author    string         `from:"author" json:"author" gorm:"index"`
	Rating    float64        `from:"rating" json:"rating" gorm:""`
	Price     float64        `from:"price" json:"price" gorm:"index"`
	CreatedAt time.Time      `from:"created_at" json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `from:"updated_at" json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `from:"deleted_at" json:"deleted_at" gorm:"index"`
}
