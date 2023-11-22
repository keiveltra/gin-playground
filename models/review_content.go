package models

import (
	"gorm.io/gorm"
	"time"
)

type ReviewContent struct {
	gorm.Model
	ID                  uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	VersionID           uint       `gorm:"type:int unsigned" json:"review_id"`
	ReviewID            uint       `gorm:"type:int unsigned" json:"review_id"`
	LatestContent       bool       `gorm:"column:latest_content"`
	DisplayUserName     string     `gorm:"type:varchar(64)"`
	Title               string     `gorm:"type:varchar(256)"`
	Rate                uint8      `gorm:"type:tinyint unsigned;index:idx_rate;default:5"`
	Status              string     `gorm:"type:enum('new','pending','published','declined','deleted');index"`
	ContentEn           string     `gorm:"type:varchar(4000)"`
	ContentJp           string     `gorm:"type:varchar(4000)"`
	Advice              string     `gorm:"type:varchar(4000)"`
	GoWithID            uint16     `gorm:"type:smallint unsigned"`
	LangID              int        `gorm:"type:int unsigned;index"`
	CommentDate         *time.Time `gorm:"type:datetime"`
	ActivityDate        *time.Time `gorm:"type:date"`
	PointCurrency       string     `gorm:"type:varchar(10)"`

	CreatedUserID       int        `gorm:"type:int"`
	CreatedURL          string     `gorm:"type:varchar(512)"`
	UpdatedUserID       int        `gorm:"type:int"`
	UpdatedURL          string     `gorm:"type:varchar(512)"`

	Review              Review     `gorm:"foreignKey:ReviewID"`
}
