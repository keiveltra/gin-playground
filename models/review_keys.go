package models

type ReviewKeys struct {
	ID                 uint      `gorm:"column:id;primaryKey"`
	BookingID          uint      `gorm:"column:booking_id;index"`
	TrUserBasicID      uint      `gorm:"column:tr_user_basic_id"`
	Hash               string    `gorm:"column:hash;unique"` // TODO: ask KL whether this is still used/not
	CreatedUserID      uint      `gorm:"column:created_user_id"`
	CreatedURL         string    `gorm:"column:created_url"`
	UpdatedUserID      uint      `gorm:"column:updated_user_id"`
	UpdatedURL         string    `gorm:"column:updated_url"`
}
