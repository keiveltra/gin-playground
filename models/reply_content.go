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

type ReplyContent struct {
	gorm.Model
	ID                  uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ReplyID             uint64     `gorm:"type:int unsigned" json:"reply_id"`
	ReviewContentID     uint       `gorm:"type:int unsigned" json:"review_content_id"`
	PtrComment          string     `gorm:"type:varchar(1000)" comment: "comment by the Ptr"`
	PtrStatus           string     `gorm:"type:enum('pending','published','declined')" comment: "status by the Ptr"`
	PtrStatusChangeDate *time.Time `gorm:"type:datetime" comment: "status change date by Ptr"`
	PtrStatusChangeID   int        `gorm:"type:int" comment: "status change ID by Ptr"`

	Reply               Reply      `gorm:"foreignKey:ReplyID"`
	ContentTranslations []ContentTranslation `gorm:"polymorphic:Content;"`
}
