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
	DisplayUserName     string     `gorm:"type:varchar(64)" comment: "Nick name of the Tr who posts the review"`
	Title               string     `gorm:"type:varchar(256)" comment: "title of the review"`
	Rate                uint8      `gorm:"type:tinyint unsigned;index:idx_rate;default:5" comment: "rate score of the review (1-5)"`
	Status              string     `gorm:"type:enum('new','pending','published','declined','deleted');index" comment: "status of the review"`
	Content             string     `gorm:"type:varchar(4000)" comment: "review text itself. here Tr can describe his/her experience of the activity"`
	Advice              string     `gorm:"type:varchar(4000)" comment: "advice to the other Tr written by Reviewer(Tr)"`
	ActivityDate        *time.Time `gorm:"type:date" comment: "when Tr joined the activity"`
	CommentDate         *time.Time `gorm:"type:datetime"`    

	CreatedUserID       int        `gorm:"type:int"          comment:"ID of Tr posted the review"`
	CreatedURL          string     `gorm:"type:varchar(512)" comment: "source URL of the review creation"`
	UpdatedUserID       int        `gorm:"type:int"          comment:"ID of Tr updated the review"`
	UpdatedURL          string     `gorm:"type:varchar(512)" comment: "source URL of the review update"`

	Review              Review     `gorm:"foreignKey:ReviewID"`
	ContentTranslations []ContentTranslation `gorm:"polymorphic:Content;"`
}
