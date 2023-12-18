package models

import (
	"gorm.io/gorm"
)

type Vote struct {
	gorm.Model
	ID                  uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ReviewID            uint       `gorm:"type:int unsigned" json:"review_id"`
	UserID              uint64     `gorm:"type:int unsigned;index" comment:"ID of Tr posts the vote. Former tr_user_basic_id."`

	Review              Review     `gorm:"foreignKey:ReviewID"`
}
