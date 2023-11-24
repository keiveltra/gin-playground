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

	_models := []interface{}{
		 &models.Reply{},
		 &models.ReviewImage{},
		 &models.ReviewKeys{},
		 &models.ReviewContent{},
		 &models.QuestionTemplate{},
		 &models.Question{},
		 &models.QuestionSection{},
		 &models.QuestionOption{},
		 &models.AnswerInt{},
		 &models.Review{},
	}
	
	if err := db.AutoMigrate(_models...); err != nil {
		log.Fatal(err)
	}

        var questions []models.Question
	for _, data := range getTestModelDataFromYaml("question_template") {	
		fmt.Println("question_template: ", data)
	    	questionTemplate := getQuestionTemplate(data)
	    	db.Create(&questionTemplate)

		question := getQuestion(questionTemplate.ID, data)
	    	db.Create(&question)
		questions = append(questions, question)

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

	fmt.Println("Injecting the dataset ...")

	for _, data := range getTestModelDataFromYaml("review_keys") {	
		fmt.Println("review_keys: ", data)

		reviewKeys := getReviewKeys(data)
		db.Create(&reviewKeys)
	}

	// var questions []models.Question
	// db.Find(&questions)
	// fmt.Println("questions: ", questions)

	for _, data := range getTestModelDataFromYaml("review") {	
		fmt.Println("review: ", data)

		review := getReview(questions[0].ID, data)
		db.Create(&review)

		for _, data := range getTestModelDataFromYaml("review_image") {	
			fmt.Println("review_image: ", data)

			reviewImage := getReviewImage(review.ID, data)
			db.Create(&reviewImage)
		}


		for _, data := range getTestModelDataFromYaml("review_content") {
			reviewContent := getReviewContent(review.ID, data)
			db.Create(&reviewContent)
		}

		reply := getReply(review.ID, data)
		db.Create(&reply)
	}

	//answer := getAnswerInt(questionSection.ID, 
	//      	      questionOption.ID, uint(review.ID))
	//db.Create(&answer)

	if db.Error != nil {
		log.Fatal(db.Error)
	}
}

func getReviewImage(reviewID uint, data map[string]interface{}) models.ReviewImage {
	currentTime := getCurrentTime()

	return models.ReviewImage{
		ReviewID:         reviewID,
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
}

func getReviewKeys(data map[string]interface{}) models.ReviewKeys {
	currentTime := getCurrentTime()

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

func getQuestion(questionTemplateID uint, data map[string]interface{}) models.Question {
	currentTime := getCurrentTime()

	return models.Question{
		QuestionTemplateID: questionTemplateID,
		ServiceKey:         "activity",
		ProductID:          toUint(data, "service_category_id"),
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
	}
}

func getQuestionTemplate(data map[string]interface{}) models.QuestionTemplate {
	currentTime := getCurrentTime()

	return models.QuestionTemplate{
		Name:      toString(data, "name"),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}

func getCurrentTime() time.Time {
	return time.Now()
}

func getReply(reviewID uint, data map[string]interface{}) models.Reply {
	currentTime := getCurrentTime()

	return models.Reply {
		ReviewID:           reviewID,
	 	PtrBasicID:         toInt   (data, "ptr_basic_id"),
		PtrComment:         toString(data, "ptr_comment"),
		PtrStatus:          "pending",
	 	PtrStatusChangeDate: &currentTime,
	 	PtrStatusChangeID:  toInt   (data, "ptr_status_change_id"),
	}
}

func getQuestionSection(data map[string]interface{}, questionTemplateID uint) models.QuestionSection {
	currentTime := getCurrentTime()

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
	currentTime := getCurrentTime()

	return models.QuestionOption{
		QuestionSectionID:  questionSectionID,
		Type:               string(models.OptionTypeCheckbox),
		Label:              toString(data, "label"),
		SortOrder:          toUint  (data, "sort_order"),
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
	}
}

func getAnswerInt(questionSectionID uint, questionOptionID uint, reviewID uint) models.AnswerInt {
	currentTime := getCurrentTime()
	
	return models.AnswerInt{
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

func getReview(questionID uint, data map[string]interface{}) models.Review {
	currentTime := getCurrentTime()

	return models.Review{
		ServiceKey:         models.Activity,
		ProductID:          toUint64(data, "service_category_id"),
		QuestionID:         questionID,
		BookingID:          toUint64(data, "booking_id"),
		UserBasicID:        toUint64(data, "user_basic_id"),
		OrgReviewID:        toUint64(data, "org_review_id"),
		LikeCount:          toUint64(data, "like_count"),
		UseFlag:            toUint8 (data, "use_flag"),
		MappingID:          toInt64 (data, "mapping_id"),
		CdFlag:             toUint8 (data, "cd_flag"),
		PostDate:           &currentTime,
		StatusChangeDate:   &currentTime,
		StatusChangeID:     toInt   (data, "status_change_id"),
	 	MSiteID:            toInt   (data, "m_site_id"),
	 	MOriginID:          toUint64(data, "m_origin_id"),
	 	CreatedUserID:      toInt   (data, "created_user_id"),
	 	CreatedURL:         toString(data, "created_url"),
	 	UpdatedUserID:      toInt   (data, "updated_user_id"),
	 	UpdatedURL:         toString(data, "updated_url"),
	 	ACConversionFlag:   toUint8 (data, "acc_conversion_flag"),
	 }
}           

func getReviewContent(reviewID uint, data map[string]interface{}) models.ReviewContent {
	currentTime := getCurrentTime()

	return models.ReviewContent{
		ReviewID:           reviewID,
		Rate:               toUint8 (data, "rate"),
		VersionID:          toUint  (data, "version_id"),
		LatestContent:      true,
		Status:             toString(data, "status"),
		DisplayUserName:    toString(data, "display_user_name"),
		Title:              toString(data, "title"),
		Advice:             toString(data, "advice"),
		GoWithID:           toUint16(data, "go_with_id"),
		CommentDate:        &currentTime,
		ContentEn:          toString(data, "content_en"),
		ContentJp:          toString(data, "content_jp"),
	 	LangID:             toInt   (data, "lang_id"),
	 	ActivityDate:       &currentTime,
	 	PointCurrency:      toString(data, "point_currency"),
	 	CreatedUserID:      toInt   (data, "created_user_id"),
	 	CreatedURL:         toString(data, "created_url"),
	 	UpdatedUserID:      toInt   (data, "updated_user_id"),
	 	UpdatedURL:         toString(data, "updated_url"),
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
