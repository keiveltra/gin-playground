package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
	"veltra.com/gin_playground/models"
)

var db = make(map[string]string)

func migrateDatabase() {
	db, err := gorm.Open(mysql.Open("moomin:moomin@tcp(127.0.0.1:3306)/test"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	models := []interface{}{
	    &models.Review{},
	    &models.ReviewImage{},
	    &models.ReviewKeys{},
	    &models.QuestionTemplate{},
	    &models.Question{},
	    &models.QuestionSection{},
	    &models.QuestionOption{},
	    &models.Answer{},
	}
	
	if err := db.AutoMigrate(models...); err != nil {
	    log.Fatal(err)
	}

	fmt.Println("Injecting the dataset ...")

        for _, data := range getTestModelDataFromYaml() {	
		fmt.Println("...%s", data)

		review := getReview(data)
		db.Create(&review)
	}

	reviewImage := getReviewImage()
	db.Create(&reviewImage)

	//reviewKeys := getReviewKeys()
	//db.Create(&reviewKeys)

	//questionTemplate := getQuestionTemplate()
	//db.Create(&questionTemplate)

	//questionSection := getQuestionSection(questionTemplate.ID)
	//db.Create(&questionSection)

	//questionOption := getQuestionOption(questionSection.ID)
	//db.Create(&questionOption)

	//answer := getAnswer(questionSection.ID, 
	//      	      questionOption.ID, uint(review.ID))
	//db.Create(&answer)

	if db.Error != nil {
		log.Fatal(db.Error)
	}
}

func getReviewImage() models.ReviewImage {
	return models.ReviewImage{}
}

func getReviewKeys() models.ReviewKeys {
    	currentTime := time.Now()
	return models.ReviewKeys{
		BookingID: 1,
		TrUserBasicID: 100,
		Hash: "abc123abc123",
		Created: currentTime,
		Updated: currentTime,
	}
}

func getQuestionTemplate() models.QuestionTemplate {
	return models.QuestionTemplate{
		Name: "Type 1",
	}
}

func getQuestionSection(questionTemplateID uint) models.QuestionSection {
	currentTime := time.Now()

	fmt.Println("questionTemplateID is:", questionTemplateID)
	return models.QuestionSection{
		QuestionTemplateID: questionTemplateID,
		Type:               models.SectionTypeNormal,
		Label:              "Type 1",
		SortOrder:          1,
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
	}
}

func getQuestionOption(questionSectionID uint) models.QuestionOption {
	currentTime := time.Now()

	fmt.Println("questionSectionID is:", questionSectionID)
	return models.QuestionOption{
		QuestionSectionID:  questionSectionID,
		Type:               string(models.OptionTypeCheckbox),
		Label:              "How is the staff?",
		SortOrder:          1,
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
	}
}

func getAnswer(questionSectionID uint, questionOptionID uint, reviewID uint) models.Answer {
	currentTime := time.Now()
	
	return models.Answer{
		QuestionSectionID:  questionSectionID,
		QuestionOptionID:   questionOptionID,
		ReviewID:           reviewID,
		Value:              3,
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
	}
}

func getTestModelDataFromYaml() []map[string]interface{} {
	viper.SetConfigFile("test/fixtures/test.yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		os.Exit(1)
	}

	var result []map[string]interface{}
	reviews := viper.Get("data").([]interface{})
	for _, review := range reviews {
		reviewMap         := review.(map[string]interface{})
		result             = append(result, reviewMap)
	}

	return result
}

func toUint64(data map[string]interface{}, key string) uint64 {
	value, _ := data[key].(uint64)	
	return value
}

func toUint8(data map[string]interface{}, key string) uint8 {
	value, _ := data[key].(uint8)	
	return value
}

func toUint16(data map[string]interface{}, key string) uint16 {
	value, _ := data[key].(uint16)	
	return value
}

func toString(data map[string]interface{}, key string) string {
	value, _ := data[key].(string)	
	return value
}

func getReview(data map[string]interface{}) models.Review {
	currentTime := time.Now()
	
	orgReviewID, _       := data["org_review_id"].(uint64)
	ptrComment, _        := data["ptr_comment"].(string)
	likeCount, _         := data["like_count"].(uint64)
	status, _            := data["status"].(string)
	ptrStatus, _         := data["ptr_status"].(string)
	useFlag, _           := data["use_flag"].(uint8)
	mappingID, _         := data["mapping_id"].(int64)
	cdFlag, _            := data["cd_flag"].(uint8)
	statusChangeID, _    := data["status_change_id"].(int)
	ptrStatusChangeID, _ := data["ptr_status_change_id"].(int)
	mSiteID, _           := data["m_site_id"].(int)
	langID, _            := data["lang_id"].(int)
	mOriginID, _         := data["m_origin_id"].(uint64)
	ptrBasicID, _        := data["ptr_basic_id"].(int)
	pointCurrency, _     := data["point_currency"].(string)

	return models.Review{
		ServiceKey:         models.Activity,
		ServiceCategoryID:  toUint64(data, "service_category_id"),
		BookingID:          toUint64(data, "booking_id"),
		UserBasicID:        toUint64(data, "user_basic_id"),
		Rate:               toUint8(data, "rate"),
		DisplayUserName:    toString(data, "display_user_name"),
		Advice:             toString(data, "advice"),
		GoWithID:           toUint16(data, "go_with_id"),
		FirstReviewID:      toUint64(data, "first_review_id"),
		OrgReviewID:        orgReviewID,
		PtrComment:         ptrComment,
		LikeCount:          likeCount,
		Status:             status,
		PtrStatus:          ptrStatus,
		UseFlag:            useFlag,
		MappingID:          mappingID,
		CdFlag:             cdFlag,
		PostDate:           &currentTime,
		CommentDate:        &currentTime,
		StatusChangeDate:   &currentTime,
		StatusChangeID:     statusChangeID,
	    	PtrStatusChangeDate: &currentTime,
	    	PtrStatusChangeID:  ptrStatusChangeID,
	    	MSiteID:            mSiteID,
	    	LangID:             langID,
	    	MOriginID:          mOriginID,
	    	ActivityDate:       &currentTime,
	    	PtrBasicID:         ptrBasicID,
	    	PointCurrency:      pointCurrency,
	    	Created:            &currentTime,
	    	CreatedUserID:      30,
	    	CreatedURL:         "http://veltra.com/ac",
	    	Updated:            &currentTime,
	    	UpdatedUserID:      30,
	    	UpdatedURL:         "http://veltra.com/ac",
	    	ACConversionFlag:   0,
	    }
}           
