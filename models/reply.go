package models

import (
	"gorm.io/gorm"
	"time"
)

type PtrStatus string
const (
	PtrStatusPending   Status = "Pending"
	PtrStatusPublished Status = "Published"
	PtrStatusDeclined  Status = "Declined"
)

type Reply struct {
	gorm.Model
	ID                  uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ReviewID            uint       `gorm:"type:int unsigned" json:"review_id"`
	PtrBasicID          int        `gorm:"type:int unsigned;index"`

	Review              Review     `gorm:"foreignKey:ReviewID"`
}

type ReplyContent struct {
	gorm.Model
	ID                  uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ReplyID             uint64     `gorm:"type:int unsigned" json:"reply_id"`
	PtrComment          string     `gorm:"type:varchar(1000)"`
	PtrStatus           string     `gorm:"type:enum('pending','published','declined')"`
	PtrStatusChangeDate *time.Time `gorm:"type:datetime"`
	PtrStatusChangeID   int        `gorm:"type:int"`

	Reply               Reply      `gorm:"foreignKey:ReplyID"`
	ContentTranslation  []ContentTranslation `gorm:"polymorphic:Content;"`
}
