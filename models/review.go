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

type Review struct {
	gorm.Model
	ID                  uint       `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ServiceKey          ServiceKey `gorm:"type:enum('ac','ticket');index"`
	ProductID           uint64     `gorm:"type:int unsigned"`
	QuestionID          uint      `gorm:"type:int unsigned" json:"question_id"`
	BookingID           uint64     `gorm:"type:int unsigned"`
	UserBasicID         uint64     `gorm:"type:int unsigned;index"`
	LikeCount           uint64     `gorm:"type:int unsigned"`
	UseFlag             uint8      `gorm:"type:tinyint unsigned;index"`
	MappingID           int64      `gorm:"type:int"`
	CdFlag              uint8      `gorm:"type:tinyint unsigned;default:0"`
	Hash                string     `gorm:"type:varchar(512)"`
	PostDate            *time.Time `gorm:"type:datetime"`
	StatusChangeDate    *time.Time `gorm:"type:datetime"`
	StatusChangeID      int        `gorm:"type:int"`
	MSiteID             int        `gorm:"type:int"`
	MOriginID           uint64     `gorm:"type:int unsigned"`
	CreatedUserID       int        `gorm:"type:int"`
	CreatedURL          string     `gorm:"type:varchar(512)"`
	UpdatedUserID       int        `gorm:"type:int"`
	UpdatedURL          string     `gorm:"type:varchar(512)"`

	Answer              []Answer             `gorm:"foreignKey:ReviewID"`
	ReviewImage         []ReviewImage        `gorm:"foreignKey:ReviewID"`
}
