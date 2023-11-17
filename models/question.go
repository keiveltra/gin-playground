package models

import (
    "time"
)

type Question struct {
	ID                 uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	QuestionTemplateID uint      `gorm:"type:int unsigned" json:"question_template_id"`
	ServiceKey         string    `gorm:"type:enum('activity','ticket','point')" json:"service_key"`
	ServiceCategoryID  uint      `gorm:"type:int unsigned" json:"service_target_id"`
	CreatedAt          time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"type:datetime" json:"updated_at"`
}
