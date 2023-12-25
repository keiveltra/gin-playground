package models

import (
	"gorm.io/gorm"
)

type Reply struct {
	gorm.Model
	ID         uint64 `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ReviewID   uint   `gorm:"type:int unsigned" json:"review_id"`
        PtrBasicID int    `gorm:"type:int unsigned;index" comment: "Ptr's ID in CS DB"`

	Review     Review `gorm:"foreignKey:ReviewID"`
}
