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
	ServiceKey          ServiceKey `gorm:"type:enum('ac','ticket');index" comment: "The key to infer the external service(ES)"`
	ProductID           uint64     `gorm:"type:int unsigned" comment: "The ID to identify the product id in the ES (i.e. activity id)"`
	CategoryID          uint64     `gorm:"type:int unsigned" comment: "If the ES is ac, then this field is used."`
	QuestionID          uint       `gorm:"type:int unsigned" json:"question_id" comment: "The ID of question (= customizable review)"`
	BookingID           uint64     `gorm:"type:int unsigned" comment: "Booking ID in CS"`
	UserBasicID         uint64     `gorm:"type:int unsigned;index" comment: "ID of traveller who posts the review"`
	VoteCount           uint64     `gorm:"type:int unsigned" comment: "When people clicked likes botton in the review, this column is incremented."`
	UseFlag             uint8      `gorm:"type:tinyint unsigned;index" comment: ""`
	MappingID           int64      `gorm:"type:int"`
	CdFlag              uint8      `gorm:"type:tinyint unsigned;default:0"`
	Hash                string     `gorm:"type:varchar(512)"`
	PostDate            *time.Time `gorm:"type:datetime"`
	StatusChangeDate    *time.Time `gorm:"type:datetime"`
	StatusChangeID      int        `gorm:"type:int"`
	MSiteID             int        `gorm:"type:int"`
	MOriginID           uint64     `gorm:"type:int unsigned"`
	LangID              int        `gorm:"type:int unsigned;index" comment: "language ID"`
	CreatedUserID       int        `gorm:"type:int"`
	CreatedURL          string     `gorm:"type:varchar(512)"`
	UpdatedUserID       int        `gorm:"type:int"`
	UpdatedURL          string     `gorm:"type:varchar(512)"`

	Answer              []Answer             `gorm:"foreignKey:ReviewID"`
	ReviewImage         []ReviewImage        `gorm:"foreignKey:ReviewID"`
	Vote                []Vote               `gorm:"foreignKey:ReviewID"`
}
