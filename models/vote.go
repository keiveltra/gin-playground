package models

import (
	"gorm.io/gorm"
)

type Vote struct {
	gorm.Model
	ID                  uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ReviewID            uint       `gorm:"type:int unsigned" json:"review_id"`
	TrUserBasicID       uint64     `gorm:"type:int unsigned;index"`

	Review              Review     `gorm:"foreignKey:ReviewID"`
}
