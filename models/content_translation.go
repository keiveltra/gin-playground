package models

type Translator string
const (
	TranslatorGoogle Translator = "google"
	TranslatorDeepl  Translator = "deepl"
)

type ContentType string
const (
	ContentTypeReply  ContentType = "reply"
	ContentTypeReview ContentType = "review"
	ContentTypeImage  ContentType = "image"
)

type ContentTranslation struct {
	ID                  uint    `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	TranslatedContent   string  `gorm:"type:varchar(1000)"`
	ContentType         string  `gorm:"type:enum('reply', 'review', 'image')" json:"type"`
	// Issue: Reuse ContentID among multiple tables caused error.
	ContentID           uint64  `gorm:"type:int unsigned;index"`
	LangID              int     `gorm:"type:int unsigned;index"`
	Translator          string  `gorm:"type:enum('google', 'deepl')" json:"type"`
	HumanApprovalID     uint    `gorm:"type:int unsigned" json:"content_id"`
}
