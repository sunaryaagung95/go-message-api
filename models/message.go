package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Message model
type Message struct {
	ID        int       `gorm:"primary_key;auto_increment;" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	SenderID  int       `gorm:"not null;" json:"sender_id"`
	RoomID    int       `gormL:"not null;" json:"room_id"`
	Content   string    `gorm:"size:300;" json:"content"`
}

// Prepare user model
func (m *Message) Prepare() {
	m.ID = 0
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
}

// CreateMessage func
func (m *Message) CreateMessage(db *gorm.DB) (*Message, error) {
	var err error
	err = db.Debug().Create(&m).Error
	if err != nil {
		return &Message{}, err
	}
	return m, nil
}

// GetMessage on room func
func (m *Message) GetMessage(db *gorm.DB, rid int) (*[]Message, error) {
	var err error
	messages := []Message{}
	err = db.Debug().Model(&Message{}).Where("room_id = ?", rid).Find(&messages).Error
	if err != nil {
		return &[]Message{}, err
	}
	return &messages, err
}

// DeleteMessage func
func (m *Message) DeleteMessage(db *gorm.DB, mid int) (int64, error) {
	db = db.Debug().Model(&Message{}).Where("id = ?", mid).Take(&Message{}).Unscoped().Delete(&Message{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
