package main

import (
	"time"
)

type QuestionSection struct {
	ID                uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	QuestionTemplateID uint     `gorm:"type:int unsigned" json:"question_template_id"`
	Type              string    `gorm:"type:enum('normal','weather','multi_choice')" json:"type"`
	Label             string    `gorm:"type:varchar(100)" json:"label"`
	SortOrder         uint      `gorm:"type:int unsigned" json:"sort_order"`
	CreatedAt         time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"type:datetime" json:"updated_at"`

	QuestionOptions []QuestionOption `gorm:"foreignKey:QuestionSectionID"`
}
