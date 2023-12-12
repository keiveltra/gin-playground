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
        TranslatedContent   string  `gorm:"type:varchar(1000)" comment: "translated content here."`
        ContentType         string  `gorm:"type:enum('reply', 'review', 'image')" json:"type" comment: "translation of which content (reply/review/image)"`
        ContentID           uint64  `gorm:"type:int unsigned;index;foreignKey:Content" json:"content_id" comment: "since this is child table of content of [content_type], this column clarify which this instance belongs to."`
        LangID              int     `gorm:"type:int unsigned;index" comment: "The language ID of the translation object"`
        Translator          string  `gorm:"type:enum('google', 'deepl')" json:"type" comment: "Google, DeepL, etc... Translation Service identifier"`
        HumanApprovalID     uint    `gorm:"type:int unsigned" json:"human_approval_id" comment: "when translation approved, the ID of the user who approved the translation."`
}
