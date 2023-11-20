package models

import (
	"gorm.io/gorm"
	"time"
)

type Reply struct {
	gorm.Model
	ID                  uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ReviewID            uint       `gorm:"type:int unsigned" json:"review_id"`
	PtrBasicID          int        `gorm:"type:int unsigned;index"`
	PtrComment          string     `gorm:"type:varchar(1000)"`
	PtrStatus           string     `gorm:"type:enum('pending','published','declined')"`
	PtrStatusChangeDate *time.Time `gorm:"type:datetime"`
	PtrStatusChangeID   int        `gorm:"type:int"`

	Review              Review     `gorm:"foreignKey:ReviewID"`
}
