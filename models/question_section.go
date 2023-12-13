package models

import (
	"time"
)

type SectionType string
const (
    SectionTypeNormal  SectionType = "normal"
    SectionTypeWeather SectionType = "weather"
    SectionTypeMultiChoice SectionType = "multi_choice"
)

type QuestionSection struct {
	ID                 uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	QuestionTemplateID uint     `gorm:"type:int unsigned" json:"question_template_id"`
        Type               SectionType `gorm:"type:enum('normal','weather','multi_choice')" json:"type" comment: "form of each question item. if it is usual 5-digit score evaluation, it is normal"`
        Name               string    `gorm:"type:varchar(100)" json:"label" comment: "summary label of question item (shorter version). This is chiefly used for the statistics view"`
        Label              string    `gorm:"type:varchar(100)" json:"label" comment: "detailed label"`
        //
        // As discussed with PdM on 2023-11-28, they concluded this show field is not needed sor far.
        //
	// Show               bool      `gorm:"column:show"`  // Whether show/display on the UI or not.
        Required           bool      `gorm:"column:required" comment: "If this is False, Tr must input this question item."`
        SortOrder          uint      `gorm:"type:int unsigned" json:"sort_order" comment: "On which index (order) this item should be positioned."`
	CreatedAt          time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"type:datetime" json:"updated_at"`

	Answer Answer `gorm:"foreignKey:QuestionSectionID"`
	QuestionOptions []QuestionOption `gorm:"foreignKey:QuestionSectionID"`
 	QuestionSectionAverageStat QuestionSectionAverageStat
}
