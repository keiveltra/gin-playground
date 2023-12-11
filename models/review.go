package models

import (
	"gorm.io/gorm"
	"time"
)

type ServiceKey string
const (
	Activity ServiceKey = "ac"
	Ticket   ServiceKey = "ticket"
	Store    ServiceKey = "store"
)

//
// CategoryID is (in activity context) parent of product_id(=ac_id)
// When querying review count/ave.score then we need dedicated view/table
// for memory usage reduction.
//
type Review struct {
	gorm.Model
	ID                  uint       `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ServiceKey          ServiceKey `gorm:"type:enum('ac','ticket');index" comment:"The key to infer the external service(ES)"`
	ProductID           uint64     `gorm:"type:int unsigned" comment:"The ID to identify the product id in the ES (i.e. activity id)"`
	CategoryID          uint64     `gorm:"type:int unsigned" comment:"If the ES is ac, then this field is used."`
	QuestionID          uint       `gorm:"type:int unsigned" json:"question_id" comment:"The ID of question (= customizable review)"`
	UserBasicID         uint64     `gorm:"type:int unsigned;index" comment:"ID of Tr posts the review"`
	VoteCount           uint64     `gorm:"type:int unsigned" comment:"When people clicked likes button in the review, this column is incremented."`
	Hash                string     `gorm:"type:varchar(512)" comment:"For allowing user to redirect to review page with review_id agnostic."`
	AttendedAsID        uint16     `gorm:"type:smallint unsigned" comment: "former go_with_id"`
	ActivityDate        *time.Time `gorm:"type:date" comment: "when Tr joined the activity"`
	PostDate            *time.Time `gorm:"type:datetime" comment:"Datetime of posting Review"`
	StatusChangeDate    *time.Time `gorm:"type:datetime" comment:"Datetime of updating the status of Review"`
	StatusChangeID      int        `gorm:"type:int" comment:"ID of updating the status of review"`
	MSiteID             int        `gorm:"type:int" comment:"Site ID(=CS)"`
	MOriginID           uint64     `gorm:"type:int unsigned" comment:"Location ID"`
	LangID              int        `gorm:"type:int unsigned;index" comment:"language ID"`
	CreatedUserID       int        `gorm:"type:int" comment:"ID of Tr posted the review"`
	CreatedURL          string     `gorm:"type:varchar(512)" comment: "source URL of the review creation"`
	UpdatedUserID       int        `gorm:"type:int" comment:"ID of Tr updated the review"`
	UpdatedURL          string     `gorm:"type:varchar(512)" comment: "source URL of the review update"`

	Answer              []Answer             `gorm:"foreignKey:ReviewID"`
	ReviewImage         []ReviewImage        `gorm:"foreignKey:ReviewID"`
	Vote                []Vote               `gorm:"foreignKey:ReviewID"`
}
