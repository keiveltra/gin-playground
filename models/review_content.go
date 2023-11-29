package models

import (
	"gorm.io/gorm"
	"time"
)

type Status string
const (
	StatusNew       Status = "New"
	StatusPending   Status = "Pending"
	StatusPublished Status = "Published"
	StatusDeclined  Status = "Declined"
	StatusDeleted   Status = "Deleted"
)

type ReviewContent struct {
	gorm.Model
	ID                  uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ReviewID            uint       `gorm:"type:int unsigned" json:"review_id"`
	DisplayUserName     string     `gorm:"type:varchar(64)"`
	Title               string     `gorm:"type:varchar(256)"`
	Rate                uint8      `gorm:"type:tinyint unsigned;index:idx_rate;default:5"`
	Status              string     `gorm:"type:enum('new','pending','published','declined','deleted');index"`
	Content             string     `gorm:"type:varchar(4000)"`
	Advice              string     `gorm:"type:varchar(4000)"`
	AttendedAsID        uint16     `gorm:"type:smallint unsigned"` // Former go_with_id
	LangID              int        `gorm:"type:int unsigned;index"`
	CommentDate         *time.Time `gorm:"type:datetime"`
	ActivityDate        *time.Time `gorm:"type:date"`
	PointCurrency       string     `gorm:"type:varchar(10)"`

	CreatedUserID       int        `gorm:"type:int"`
	CreatedURL          string     `gorm:"type:varchar(512)"`
	UpdatedUserID       int        `gorm:"type:int"`
	UpdatedURL          string     `gorm:"type:varchar(512)"`

	Review              Review     `gorm:"foreignKey:ReviewID"`
	ContentTranslation  []*ContentTranslation `gorm:"foreignKey:ReviewContentID"`
}
