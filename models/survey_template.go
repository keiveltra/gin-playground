package models

import (
    "time"
)

type SurveyTemplate struct {
	ID        uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name" comment: "this is not shown to the user but for internal admin purpose"`
        ServiceKey         string    `gorm:"type:enum('activity','ticket','point')" json:"service_key" comment: "key to identify the external service(ES) uses Review Service"`
        ProductID          uint      `gorm:"type:int unsigned" json:"service_target_id" comment: "product id of the ES"`
	LangID    int       `gorm:"type:int unsigned;index" comment: "lang_id of question template"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`

	Questions []Question `gorm:"foreignKey:SurveyTemplateID"`
}
