package models

//
// This is temporal table definition easier to understand what is needed outside of
// main tables, for the stat use.
// The following needs to be discussed:
//
// 1. Redis or View, or RDB table (as-is)
// 2. Is the following table  definition is sufficient?
// 3. Is the following column definition is sufficient?
//

type LikeCountStat struct {
	ServiceKey          ServiceKey `gorm:"type:enum('ac','ticket');index"`
	ProductID           uint64     `gorm:"type:int unsigned"`
	CategoryID          uint64     `gorm:"type:int unsigned"`
	ReviewID            uint64     `gorm:"type:int unsigned"`
	LikeCount           uint64     `gorm:"type:int unsigned"`
}

type QuestionSectionAverageStat struct {
	ServiceKey          ServiceKey `gorm:"type:enum('ac','ticket');index"`
	ProductID           uint64     `gorm:"type:int unsigned"`
	QuestionSectionID   uint64     `gorm:"type:int unsigned"`
	Average             uint64     `gorm:"type:int unsigned"` // Average Score, per question_section
}
