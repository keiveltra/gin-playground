package models

import (
	"gorm.io/gorm"
)

type ImageStatus string
const (
	ImageStatusActive  ImageStatus = "Active"
	ImageStatusDeleted ImageStatus = "Deleted"
)

type ReviewImage struct {
	gorm.Model
	ID                 uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ReviewID           uint       `gorm:"type:int unsigned" json:"review_id"`
	Filename           string     `gorm:"type:varchar(128)"`
	FilenameBase       string     `gorm:"type:varchar(128)"`
	Status             string     `gorm:"type:enum('active','deleted');index"`
	Width              uint64     `gorm:"type:int unsigned"`
	Height             uint64     `gorm:"type:int unsigned"`
	Size               uint64     `gorm:"type:int unsigned"`
	Comment            string     `gorm:"type:varchar(1000)"`
	CreatedUserID      uint64     `gorm:"type:int unsigned"`
	CreatedURL         string     `gorm:"type:varchar(512)"`
	UpdatedUserID      int        `gorm:"type:int"`
	UpdatedURL         string     `gorm:"type:varchar(512)"`
	ContentTranslation []*ContentTranslation `gorm:"foreignKey:ReviewImageID"`

        ReviewContents     []ReviewContentImage `gorm:"foreignKey:ReviewImageID"`
}
