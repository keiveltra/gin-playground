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
	ContentTranslation  []ContentTranslation `gorm:"foreignKey:ReplyID"`
}

//
// TODO: to a new file
//
type ContentType string
const (
	ContentTypeReply  ContentType = "reply"
	ContentTypeReview ContentType = "review"
)

type ContentTranslation struct {
	Comment             string    `gorm:"type:varchar(1000)"`
	ReplyID             *uint     `gorm:"type:int unsigned" json:"reply_id"`
	ReviewContentID     *uint     `gorm:"type:int unsigned" json:"review_id"`
	LangID              int       `gorm:"type:int unsigned;index"`
	ContentType         string    `gorm:"type:enum('reply', 'review')" json:"type"`
}
