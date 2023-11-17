package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
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
	reviewImage      := getReviewImage()
	reviewKeys       := getReviewKeys()
	questionTemplate := getQuestionTemplate()
	questionSection  := QuestionSection{}
	questionOption   := getQuestionOption(questionSection.ID)
	answer           := getAnswer(questionSection.ID, 
				      questionOption.ID, review.ID)
	db.Create(&review)
	db.Create(&questionSection)
	db.Create(&questionOption)
	db.Create(&answer)
	if db.Error != nil {
		log.Fatal(db.Error)
	}
}

func getReviewImage() ReviewImage {

}

func getReviewKeys() ReviewKeys {

}


func main() {
	args := os.Args
	if len(args) > 1 && (args[1] == "m" || args[1] == "migrate") {
		migrateDatabase()
		return
	}

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func getQuestionOption(questionSectionID uint) QuestionOption {
    currentTime := time.Now()

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
        ServiceKey:         string(Activity),
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
