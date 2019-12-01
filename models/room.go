package models

import "github.com/jinzhu/gorm"

// Room model
type Room struct {
	gorm.Model
	AdminID int `gorm:"not null"; json:"admin_id"`
}

// CreateRoom func
func (r *Room) CreateRoom(db *gorm.DB) (*Room, error) {
	var err error
	err = db.Debug().Model(&Room{}).Create(&r).Error
	if err != nil {
		return &Room{}, err
	}
	return r, nil
}
