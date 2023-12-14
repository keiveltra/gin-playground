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
	QuestionID uint      `gorm:"type:int unsigned" json:"question_section_id"`
        Type              string    `gorm:"type:enum('checkbox')" json:"type" comment: "Type of option. (i.e. checkbox)"`
        Label             string    `gorm:"type:varchar(100)" json:"label" comment: "Label of each option"`
        SortOrder         uint      `gorm:"type:int unsigned" json:"sort_order" comment: "sort order of each option"`
	CreatedAt         time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"type:datetime" json:"updated_at"`

	Answer Answer `gorm:"foreignKey:QuestionOptionID"`
}
