package models

import (
    "time"
)

type QuestionTemplate struct {
	ID        uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name" comment: "this is not shown to the user but for internal admin purpose"`
	LangID    int       `gorm:"type:int unsigned;index" comment: "lang_id of question template"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`

	QuestionSections []QuestionSection `gorm:"foreignKey:QuestionTemplateID"`
	Questions        []Question        `gorm:"foreignKey:QuestionTemplateID"`
}
