package models

import "github.com/jinzhu/gorm"

// Member model
type Member struct {
	gorm.Model
	RoomID int `gorm:"not null;" json:"room_id"`
	UserID int `gorm:"not null;" json:"user_id"`
}

// AddMember to room
func (m *Member) AddMember(db *gorm.DB) (*Member, error) {
	var err error
	err = db.Debug().Create(&m).Error
	if err != nil {
		return &Member{}, err
	}
	return m, nil
}

// GetMember in room
func (m *Member) GetMember(db *gorm.DB, rid int) (*[]Member, error) {
	var err error
	members := []Member{}
	err = db.Debug().Model(&Member{}).Where("room_id = ?", rid).Find(&members).Error
	if err != nil {
		return &[]Member{}, err
	}
	return &members, err
}
