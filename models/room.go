package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Room model
type Room struct {
	gorm.Model
	AdminID int `gorm:"not null;" json:"admin_id"`
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

// GetAllRoom func
func (r *Room) GetAllRoom(db *gorm.DB) (*[]Room, error) {
	var err error
	rooms := []Room{}
	err = db.Debug().Model(&Room{}).Find(&rooms).Error
	if err != nil {
		return &[]Room{}, err
	}
	return &rooms, err
}

// GetOneRoom func
func (r *Room) GetOneRoom(db *gorm.DB, rid int) (*Room, error) {
	var err error
	err = db.Debug().Model(&Room{}).Where("id= ?", rid).Take(&r).Error
	if err != nil {
		return &Room{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Room{}, errors.New("Not found")
	}
	return r, err
}

// DeleteRoom func
func (r *Room) DeleteRoom(db *gorm.DB, rid int) (int64, error) {
	db = db.Debug().Model(&Room{}).Where("id = ?", rid).Take(&Room{}).Unscoped().Delete(&Room{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
