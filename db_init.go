package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db = make(map[string]string)

func migrateDatabase() {
	db, err := gorm.Open(mysql.Open("moomin:moomin@tcp(127.0.0.1:3306)/test"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	models := []interface{}{
	    &Review{},
	    &ReviewImage{},
	    &ReviewKeys{},
	    &QuestionTemplate{},
	    &Question{},
	    &QuestionSection{},
	    &QuestionOption{},
	    &Answer{},
	}
	
	if err := db.AutoMigrate(models...); err != nil {
	    log.Fatal(err)
	}

	fmt.Println("Injecting the dataset ...")
	review           := getReview()
	db.Create(&review)

	reviewImage      := getReviewImage()
	db.Create(&reviewImage)

	reviewKeys       := getReviewKeys()
	db.Create(&reviewKeys)

	questionTemplate := getQuestionTemplate()
	db.Create(&questionTemplate)

	questionSection  := getQuestionSection(questionTemplate.ID)
	db.Create(&questionSection)

	questionOption   := getQuestionOption(questionSection.ID)
	db.Create(&questionOption)

	answer           := getAnswer(questionSection.ID, 
				      questionOption.ID, uint(review.ID))
	db.Create(&answer)

	if db.Error != nil {
		log.Fatal(db.Error)
	}
}

func getReviewImage() ReviewImage {
	return ReviewImage{}
}

func getReviewKeys() ReviewKeys {
    	currentTime := time.Now()
	return ReviewKeys{
		BookingID: 1,
		TrUserBasicID: 100,
		Hash: "abc123abc123",
		Created: currentTime,
		Updated: currentTime,
	}
}

func getQuestionTemplate() QuestionTemplate {
	return QuestionTemplate{
		Name: "Type 1",
	}
}

func getQuestionSection(questionTemplateID uint) QuestionSection {
	currentTime := time.Now()

	fmt.Println("questionTemplateID is:", questionTemplateID)
	return QuestionSection{
		QuestionTemplateID: questionTemplateID,
		Type:               SectionTypeNormal,
		Label:              "Type 1",
		SortOrder:          1,
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
	}
}

func getQuestionOption(questionSectionID uint) QuestionOption {
	currentTime := time.Now()

	fmt.Println("questionSectionID is:", questionSectionID)
	return QuestionOption{
		QuestionSectionID:  questionSectionID,
		Type:               string(OptionTypeCheckbox),
		Label:              "How is the staff?",
		SortOrder:          1,
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
	}
}

func getAnswer(questionSectionID uint, questionOptionID uint, reviewID uint) Answer {
	currentTime := time.Now()
	
	return Answer{
		QuestionSectionID:  questionSectionID,
		QuestionOptionID:   questionOptionID,
		ReviewID:           reviewID,
		Value:              3,
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
	}
}

func getReview() Review {
	currentTime := time.Now()
	
	return Review{
		ServiceKey:         Activity,
		ServiceCategoryID:  10,
		BookingID:          12345,
		UserBasicID:        1,
		Rate:               4,
		DisplayUserName:    "David Solomon a.k.a D-sol",
		Title:              "I LOVE THIS PLACE!!!!!",
		Review:             "This is one of the best stadium in the world I can dance, and forget about my daily job",
		Advice:             "Well, do not wear any suits. you are here to enjoy.",
		GoWithID:           777,
		FirstReviewID:      0,
		OrgReviewID:        10,
		PtrComment:         "Thanks. Our team is very impressed about your feedback.",
		LikeCount:          200000,
		Status:             "new",
		PtrStatus:          "pending",
		UseFlag:            10,
		MappingID:          10,
		CdFlag:             10,
		PostDate:           &currentTime,
		CommentDate:        &currentTime,
		StatusChangeDate:   &currentTime,
		StatusChangeID:     10,
	    	PtrStatusChangeDate: &currentTime,
	    	PtrStatusChangeID:  20,
	    	MSiteID:            10,
	    	LangID:             2,
	    	MOriginID:          10,
	    	ActivityDate:       &currentTime,
	    	PtrBasicID:         23456,
	    	PointCurrency:      "JPY",
	    	Created:            &currentTime,
	    	CreatedUserID:      30,
	    	CreatedURL:         "http://veltra.com/ac",
	    	Updated:            &currentTime,
	    	UpdatedUserID:      30,
	    	UpdatedURL:         "http://veltra.com/ac",
	    	ACConversionFlag:   0,
	    }
}           
