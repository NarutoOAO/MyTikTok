package model

import "github.com/jinzhu/gorm"

type Favorite struct {
	*gorm.Model
	Id      int64
	UserId  int64
	VideoId int64
}
