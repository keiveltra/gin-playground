package models

import (
	"gorm.io/gorm"
)

type ImageStatus string
const (
	ImageStatusActive  ImageStatus = "Active"
	ImageStatusDeleted ImageStatus = "Deleted"
)

type ReviewImage struct {
	gorm.Model
	ID                 uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ReviewID           uint       `gorm:"type:int unsigned" json:"review_id"`
	ReviewContentID    uint       `gorm:"type:int unsigned" json:"review_content_id"`
	Filename           string     `gorm:"type:varchar(128)" comment: "file_path of review image in S3 or CDN"`
	FilenameBase       string     `gorm:"type:varchar(128)" comment: "pure file_name of review image"`
	Status             string     `gorm:"type:enum('active','deleted');index" comment: "whether the image has been still active or deleted"`
	Width              uint64     `gorm:"type:int unsigned" comment: "width of the image"`
	Height             uint64     `gorm:"type:int unsigned" comment: "hight of the image"`
	Size               uint64     `gorm:"type:int unsigned" comment: "size of the image"`
	Comment            string     `gorm:"type:varchar(1000)" comment: "comment description of the image"`
	CreatedUserID      uint64     `gorm:"type:int unsigned" comment:"ID of Tr posted the review"`
	CreatedURL         string     `gorm:"type:varchar(512)" comment: "source URL of the review creation"`
	UpdatedUserID      int        `gorm:"type:int" comment:"ID of Tr updated the review"`
	UpdatedURL         string     `gorm:"type:varchar(512)" comment: "source URL of the review update"`

	ContentTranslation []ContentTranslation `gorm:"polymorphic:Content;"`
}
