package models

import (
	"gorm.io/gorm"
)

// @CS DB this is defined as `ac_h_packages`
type Plan struct {
	gorm.Model
	ID                  uint       `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ReviewID            uint       `gorm:"type:int unsigned" json:"review_id"`
	PlanID              uint64     `gorm:"type:int unsigned"`
	Name                string     `gorm:"type:varchar(512)"`
}
