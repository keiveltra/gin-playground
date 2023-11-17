package models

import (
    "time"
)

type ReviewKeys struct {
	ID                 uint      `gorm:"column:id;primaryKey"`
	BookingID          uint      `gorm:"column:booking_id;index"`
	TrUserBasicID      uint      `gorm:"column:tr_user_basic_id"`
	Hash               string    `gorm:"column:hash;unique"`
	Created            time.Time `gorm:"column:created"`
	CreatedUserID      uint      `gorm:"column:created_user_id"`
	CreatedURL         string    `gorm:"column:created_url"`
	Updated            time.Time `gorm:"column:updated"`
	UpdatedUserID      uint      `gorm:"column:updated_user_id"`
	UpdatedURL         string    `gorm:"column:updated_url"`
}
