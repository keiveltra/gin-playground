package main

import (
	"gorm.io/gorm"
	"time"
)

type ReviewImage struct {
	gorm.Model
	ID                 uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	Filename           string     `gorm:"type:varchar(128)"`
	FilenameBase       string     `gorm:"type:varchar(128)"`
	Width              uint64     `gorm:"type:int unsigned"`
	Height             uint64     `gorm:"type:int unsigned"`
	Size               uint64     `gorm:"type:int unsigned"`
	Comment            string     `gorm:"type:varchar(1000)"`
	Created            *time.Time `gorm:"type:datetime"`
	CreatedUserID      uint64     `gorm:"type:int unsigned"`
	CreatedURL         string     `gorm:"type:varchar(512)"`
	Updated            *time.Time `gorm:"type:datetime"`
	UpdatedUserID      int        `gorm:"type:int"`
	UpdatedURL         string     `gorm:"type:varchar(512)"`
	ACConversionFlag   uint8      `gorm:"type:tinyint unsigned;index;default:0"`
}


