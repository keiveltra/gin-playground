package main

import (
	"gorm.io/gorm"
	"time"
)

type ServiceKey string
const (
    Activity = iota
    Ticket
    Store
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
	ID                  uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ServiceKey          string     `gorm:"type:enum('ac','ticket');index"`
	ServiceCategoryID   uint64     `gorm:"type:int unsigned"`
	BookingID           uint64     `gorm:"type:int unsigned"`
	UserBasicID         uint64     `gorm:"type:int unsigned;index"`
	Rate                uint8      `gorm:"type:tinyint unsigned;index:idx_rate;default:5"`
	DisplayUserName     string     `gorm:"type:varchar(64)"`
	Title               string     `gorm:"type:varchar(256)"`
	Review              string     `gorm:"type:varchar(4000)"`
	Advice              string     `gorm:"type:varchar(4000)"`
	GoWithID            uint16     `gorm:"type:smallint unsigned"`
	FirstReviewID       uint64     `gorm:"type:int unsigned;index"`
	OrgReviewID         uint64     `gorm:"type:int unsigned"`
	PtrComment          string     `gorm:"type:varchar(1000)"`
	LikeCount           uint64     `gorm:"type:int unsigned"`
	Status              string     `gorm:"type:enum('new','pending','published','declined','deleted');index"`
	PtrStatus           string     `gorm:"type:enum('pending','published','declined')"`
	UseFlag             uint8      `gorm:"type:tinyint unsigned;index"`
	MappingID           int64      `gorm:"type:int"`
	CdFlag              uint8      `gorm:"type:tinyint unsigned;default:0"`
	PostDate            *time.Time `gorm:"type:datetime"`
	CommentDate         *time.Time `gorm:"type:datetime"`
	StatusChangeDate    *time.Time `gorm:"type:datetime"`
	StatusChangeID      int        `gorm:"type:int"`
	PtrStatusChangeDate *time.Time `gorm:"type:datetime"`
	PtrStatusChangeID   int        `gorm:"type:int"`
	MSiteID             int        `gorm:"type:int"`
	LangID              int        `gorm:"type:int unsigned;index"`
	MOriginID           uint64     `gorm:"type:int unsigned"`
	ActivityDate        *time.Time `gorm:"type:date"`
	PtrBasicID          int        `gorm:"type:int unsigned;index"`
	PointCurrency       string     `gorm:"type:varchar(10)"`
	Created             *time.Time `gorm:"type:datetime"`
	CreatedUserID       int        `gorm:"type:int"`
	CreatedURL          string     `gorm:"type:varchar(512)"`
	Updated             *time.Time `gorm:"type:datetime"`
	UpdatedUserID       int        `gorm:"type:int"`
	UpdatedURL          string     `gorm:"type:varchar(512)"`
	ACConversionFlag    uint8      `gorm:"type:tinyint unsigned;index;default:0"`

	Answers []Answer `gorm:"foreignKey:ReviewID"`
}

// func (r *Review) UpdateMyself(newLabel string) {
//     r.ServiceCategoryID = 12
//     r.UpdatedAt = time.Now()
// }
