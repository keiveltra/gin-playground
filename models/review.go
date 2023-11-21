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

type Status string
const (
	StatusNew       Status = "New"
	StatusPending   Status = "Pending"
	StatusPublished Status = "Published"
	StatusDeclined  Status = "Declined"
	StatusDeleted   Status = "Deleted"
)

type PtrStatus string
const (
	PtrStatusPending   Status = "Pending"
	PtrStatusPublished Status = "Published"
	PtrStatusDeclined  Status = "Declined"
)

type Review struct {
	gorm.Model
	ID                  uint       `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ServiceKey          ServiceKey `gorm:"type:enum('ac','ticket');index"`
	ProductID           uint64     `gorm:"type:int unsigned"`
	BookingID           uint64     `gorm:"type:int unsigned"`
	UserBasicID         uint64     `gorm:"type:int unsigned;index"`
	OrgReviewID         uint64     `gorm:"type:int unsigned"`
	LikeCount           uint64     `gorm:"type:int unsigned"`
	Status              string     `gorm:"type:enum('new','pending','published','declined','deleted');index"`
	UseFlag             uint8      `gorm:"type:tinyint unsigned;index"`
	MappingID           int64      `gorm:"type:int"`
	CdFlag              uint8      `gorm:"type:tinyint unsigned;default:0"`
	PostDate            *time.Time `gorm:"type:datetime"`
	StatusChangeDate    *time.Time `gorm:"type:datetime"`
	StatusChangeID      int        `gorm:"type:int"`
	MSiteID             int        `gorm:"type:int"`
	MOriginID           uint64     `gorm:"type:int unsigned"`
	Created             *time.Time `gorm:"type:datetime"`
	CreatedUserID       int        `gorm:"type:int"`
	CreatedURL          string     `gorm:"type:varchar(512)"`
	Updated             *time.Time `gorm:"type:datetime"`
	UpdatedUserID       int        `gorm:"type:int"`
	UpdatedURL          string     `gorm:"type:varchar(512)"`
	ACConversionFlag    uint8      `gorm:"type:tinyint unsigned;index;default:0"`

	Answer []Answer `gorm:"foreignKey:ReviewID"`
}
