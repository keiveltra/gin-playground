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

	for _, data := range getTestModelDataFromYaml("review") {	
		fmt.Println("review: ", data)

		review := getReview(data)
		db.Create(&review)
	}

	for _, data := range getTestModelDataFromYaml("review_image") {	
		fmt.Println("review_image: ", data)

		reviewImage := getReviewImage(data)
		db.Create(&reviewImage)
	}

	for _, data := range getTestModelDataFromYaml("review_image") {	
		fmt.Println("review_keys: ", data)

		reviewKeys := getReviewKeys(data)
		db.Create(&reviewKeys)
	}

	for _, data := range getTestModelDataFromYaml("question_template") {	
		fmt.Println("question_template: ", data)
	    	questionTemplate := getQuestionTemplate(data)
	    	db.Create(&questionTemplate)

		for _, data := range getTestModelDataFromYaml("question_section") {	
			fmt.Println("question_section: ", data)
		    	questionSection := getQuestionSection(data, questionTemplate.ID)
			db.Create(&questionSection)

			if(questionSection.Type == "multi_choice") {
				questionOption := getQuestionOption(questionSection.ID, data)
				db.Create(&questionOption)


			}
		}
	}

	//answer := getAnswer(questionSection.ID, 
	//      	      questionOption.ID, uint(review.ID))
	//db.Create(&answer)

	if db.Error != nil {
		log.Fatal(db.Error)
	}
}

func getReviewImage(data map[string]interface{}) models.ReviewImage {
		currentTime := time.Now()
	return models.ReviewImage{
		Filename:         toString(data, "file_name"),
		FilenameBase:     toString(data, "file_name_base"),
		Width:            toUint64(data, "width"),
		Height:           toUint64(data, "height"),
		Size:             toUint64(data, "size"),
		Comment:          toString(data, "comment"),
		Created:          &currentTime,
		CreatedUserID:    toUint64(data, "created_user_id"),
		CreatedURL:       toString(data, "created_url"),
		Updated:          &currentTime,
		UpdatedUserID:    toInt(data, "updated_user_id"),
		UpdatedURL:       toString(data, "updated_url"),
		ACConversionFlag: toUint8(data, "acc_conversion_flag"),
	}
	
	return models.ReviewImage{}
}

func getReviewKeys(data map[string]interface{}) models.ReviewKeys {
		currentTime := time.Now()

	return models.ReviewKeys{
		BookingID:     toUint  (data, "booking_id"),
		TrUserBasicID: toUint  (data, "tr_user_basic_id"),
		Hash:          toString(data, "hash"),
		Created:       currentTime,
		CreatedURL:    toString(data, "created_url"),
		Updated:       currentTime,
		UpdatedURL:    toString(data, "updated_url"),
	}
}

func getQuestionTemplate(data map[string]interface{}) models.QuestionTemplate {
	return models.QuestionTemplate{
		Name: toString(data, "name"),
	}
}

func getQuestionSection(data map[string]interface{}, questionTemplateID uint) models.QuestionSection {
	currentTime := time.Now()

	return models.QuestionSection{
		QuestionTemplateID: questionTemplateID,
		Type:               toEnum(data, "type"),
		Label:              toString(data, "label"),
		SortOrder:          toUint  (data, "sort_order"),
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
	}
}

func getQuestionOption(questionSectionID uint, data map[string]interface{}) models.QuestionOption {
	currentTime := time.Now()

	return models.QuestionOption{
		QuestionSectionID:  questionSectionID,
		Type:               string(models.OptionTypeCheckbox),
		Label:              toString(data, "label"),
		SortOrder:          toUint  (data, "sort_order"),
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

func getTestModelDataFromYaml(target string) []map[string]interface{} {
	viper.SetConfigFile(fmt.Sprintf("test/fixtures/%s.yaml", target))

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

func getReview(data map[string]interface{}) models.Review {
	currentTime := time.Now()

	return models.Review{
		ServiceKey:         models.Activity,
		ServiceCategoryID:  toUint64(data, "service_category_id"),
		BookingID:          toUint64(data, "booking_id"),
		UserBasicID:        toUint64(data, "user_basic_id"),
		Rate:               toUint8 (data, "rate"),
		DisplayUserName:    toString(data, "display_user_name"),
		Advice:             toString(data, "advice"),
		GoWithID:           toUint16(data, "go_with_id"),
		FirstReviewID:      toUint64(data, "first_review_id"),
		OrgReviewID:        toUint64(data, "org_review_id"),
		PtrComment:         toString(data, "ptr_comment"),
		LikeCount:          toUint64(data, "like_count"),
		Status:             toString(data, "status"),
		PtrStatus:          toString(data, "ptr_status"),
		UseFlag:            toUint8 (data, "use_flag"),
		MappingID:          toInt64 (data, "mapping_id"),
		CdFlag:             toUint8 (data, "cd_flag"),
		PostDate:           &currentTime,
		CommentDate:        &currentTime,
		StatusChangeDate:   &currentTime,
		StatusChangeID:     toInt   (data, "status_change_id"),
	 	PtrStatusChangeDate: &currentTime,
	 	PtrStatusChangeID:  toInt   (data, "ptr_status_change_id"),
	 	MSiteID:            toInt   (data, "m_site_id"),
	 	LangID:             toInt   (data, "lang_id"),
	 	MOriginID:          toUint64(data, "m_origin_id"),
	 	ActivityDate:       &currentTime,
	 	PtrBasicID:         toInt   (data, "ptr_basic_id"),
	 	PointCurrency:      toString(data, "point_currency"),
	 	Created:            &currentTime,
	 	CreatedUserID:      toInt   (data, "created_user_id"),
	 	CreatedURL:         toString(data, "created_url"),
	 	Updated:            &currentTime,
	 	UpdatedUserID:      toInt   (data, "updated_user_id"),
	 	UpdatedURL:         toString(data, "updated_url"),
	 	ACConversionFlag:   toUint8 (data, "acc_conversion_flag"),
	 }
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

func toUint(data map[string]interface{}, key string) uint {
	value, _ := data[key].(uint)	
	return value
}

func toString(data map[string]interface{}, key string) string {
	value, _ := data[key].(string)	
	return value
}

func toInt(data map[string]interface{}, key string) int {
	value, _ := data[key].(int)	
	return value
}

func toInt64(data map[string]interface{}, key string) int64 {
	value, _ := data[key].(int64)	
	return value
}

func toEnum(data map[string]interface{}, key string) models.SectionType {
	value, ok := data[key].(string)
	if !ok {
		return models.SectionTypeNormal
	}

	switch value {
	case "normal":
		return models.SectionTypeNormal
	case "weather":
		return models.SectionTypeWeather
	case "multi_choice":
		return models.SectionTypeMultiChoice
	default:
		return models.SectionTypeNormal
	}
}
