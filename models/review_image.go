package models

import (
	"gorm.io/gorm"
)

type ReviewImage struct {
	gorm.Model
	ID                 uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ReviewID           uint      `gorm:"type:int unsigned" json:"review_id"`
	Filename           string     `gorm:"type:varchar(128)"`
	FilenameBase       string     `gorm:"type:varchar(128)"`
	Width              uint64     `gorm:"type:int unsigned"`
	Height             uint64     `gorm:"type:int unsigned"`
	Size               uint64     `gorm:"type:int unsigned"`
	Comment            string     `gorm:"type:varchar(1000)"`
	CreatedUserID      uint64     `gorm:"type:int unsigned"`
	CreatedURL         string     `gorm:"type:varchar(512)"`
	UpdatedUserID      int        `gorm:"type:int"`
	UpdatedURL         string     `gorm:"type:varchar(512)"`
	ACConversionFlag   uint8      `gorm:"type:tinyint unsigned;index;default:0"` // TODO: ask KL whether still used/not
	ContentTranslation  []*ContentTranslation `gorm:"foreignKey:ReviewImageID"`
}
