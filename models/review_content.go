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
	AttendedAsID        uint16     `gorm:"type:smallint unsigned" comment: "former go_with_id. who Tr accompanied with in Ac."`
	LangID              int        `gorm:"type:int unsigned;index" comment: "language ID"`
	CommentDate         *time.Time `gorm:"type:datetime"`
	ActivityDate        *time.Time `gorm:"type:date" comment: "when Tr joined the activity"`
	PointCurrency       string     `gorm:"type:varchar(10)" comment: "i.e. JPY | Dollar etc"`

	CreatedUserID       int        `gorm:"type:int"`
	CreatedURL          string     `gorm:"type:varchar(512)"`
	UpdatedUserID       int        `gorm:"type:int"`
	UpdatedURL          string     `gorm:"type:varchar(512)"`

	Review              Review     `gorm:"foreignKey:ReviewID"`
	ContentTranslation  []ContentTranslation `gorm:"polymorphic:Content;"`

        ReviewImages        []ReviewContentImage `gorm:"foreignKey:ReviewContentID"`
}

//
// since review history (=review content) should have reference to
// review image, review_content and review_image's relationship
// has to be many-to-many.
// plus, in gorm we need immediate table to realize this
// many-to-many relationships.
//
type ReviewContentImage struct {
    gorm.Model
    ReviewContentID uint64 `gorm:"type:int unsigned"`
    ReviewImageID   uint64 `gorm:"type:int unsigned"`
}
