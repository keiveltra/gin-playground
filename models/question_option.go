package models

import (
    "time"
)

type OptionType string
const (
    OptionTypeCheckbox OptionType = "Checkbox"
)

type QuestionOption struct {
	ID                uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	QuestionSectionID uint      `gorm:"type:int unsigned" json:"question_section_id"`
	Type              string    `gorm:"type:enum('checkbox')" json:"type"`
	Label             string    `gorm:"type:varchar(100)" json:"label"`
	SortOrder         uint      `gorm:"type:int unsigned" json:"sort_order"`
	CreatedAt         time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"type:datetime" json:"updated_at"`

	Answer Answer `gorm:"foreignKey:QuestionOptionID"`
}
