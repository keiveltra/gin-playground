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
	// ContentTranslation  []ContentTranslation `gorm:"foreignKey:ReplyID"`
}

//
// TODO: to a new file
//
type ContentType string
const (
	ContentTypeReply  ContentType = "reply"
	ContentTypeReview ContentType = "review"
)

//type ContentTranslation struct {
//	ID                  uint    `gorm:"type:int unsigned;primaryKey;autoIncrement"`
//	TranslatedContent   string  `gorm:"type:varchar(1000)"`
//	ContentType         string  `gorm:"type:enum('reply', 'review', 'image')" json:"type"`
//	ContentID           uint    `gorm:"type:int unsigned" json:"content_id"` // foreign_key
//	LangID              int     `gorm:"type:int unsigned;index"`
//	// + Translator tool [Google | DeepL] [Enum]
//	// + Human Approval ID [userID VTR or PTR or TR]
//}
