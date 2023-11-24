package models

import (
    "time"
)

type Answer struct {
	ID                uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	QuestionSectionID *uint     `gorm:"type:int unsigned" json:"question_section_id"`
	QuestionOptionID  *uint     `gorm:"type:int unsigned" json:"question_field_id"`
	ReviewID          uint      `gorm:"type:int unsigned" json:"review_id"`

	// TODO: Discussion
	NumberValue       *uint      `gorm:"type:int unsigned" json:"value"`
	BooleanValue      *bool      `gorm:"column:latest_content"`
	TextValue         *string    `gorm:"type:varchar(100)" json:"label"`

	CreatedAt         time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
