package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//
	Username     string `gorm:"unique;size:6"`
	HashPassword string `gorm:"size:128"`
}

func (User) TableName() string {
	return "t_user"
}

type Star struct {
	gorm.Model
	//
	AuthorID uint
	Public   bool
	Title    string `gorm:"size:36"`
	Content  string
}

func (Star) TableName() string {
	return "t_star"
}

type Media struct {
	gorm.Model
	//
	Name        string `gorm:"uniqueIndex;size:56"`
	UrlPath     string `gorm:"size:128"`
	OriName     string `gorm:"size:56"`
	ContentType string `gorm:"size:36"`
	Size        int64  // unit: bytes
	StoragePath string `gorm:"size:128"`
}

func (Media) TableName() string {
	return "t_media"
}
