package models

import (
    "time"
)

type TemplateQuestion struct {
	ID         uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	TemplateID int       `gorm:"type:int unsigned;index"`
	QuestionID int       `gorm:"type:int unsigned;index"`
	CreatedAt  time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:datetime" json:"updated_at"`
}

type TemplateReview struct {
	ID         uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	TemplateID int       `gorm:"type:int unsigned;index"`
	ReviewID   int       `gorm:"type:int unsigned;index"`
	CreatedAt  time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:datetime" json:"updated_at"`
}

type SurveyTemplate struct {
	ID        uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name" comment: "this is not shown to the user but for internal admin purpose"`
        ServiceKey         string    `gorm:"type:enum('activity','ticket','point')" json:"service_key" comment: "key to identify the external service(ES) uses Review Service"`
        ProductID          uint      `gorm:"type:int unsigned" json:"service_target_id" comment: "product id of the ES"`
	LangID    int       `gorm:"type:int unsigned;index" comment: "lang_id of question template"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`

	Questions         []Question         `gorm:"foreignKey:TemplateID"`
	TemplateReviews   []TemplateReview   `gorm:"foreignKey:TemplateID"`
	TemplateQuestions []TemplateQuestion `gorm:"foreignKey:TemplateID"`
}
