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
	"github.com/davecgh/go-spew/spew"
)

var db = make(map[string]string)

func migrateDatabase() {
	db, err := gorm.Open(mysql.Open("moomin:moomin@tcp(127.0.0.1:3306)/test?parseTime=true"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	_models := []interface{}{
		 &models.Reply{},
		 &models.ReviewImage{},
		 &models.ReviewKeys{},
		 &models.ReviewContent{},
		 &models.ContentTranslation{},
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

        var questions        []models.Question
        var questionSections []models.QuestionSection
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
			questionSections = append(questionSections, questionSection)

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

			// Translations
			translation := getContentTranslationReviewContent(reviewContent.ID, data)
			db.Create(&translation)
		}

		reply := getReply(review.ID, data)
		db.Create(&reply)
		translation := getContentTranslationReply(reply.ID, data)
		db.Create(&translation)


		for _, questionSection := range questionSections {
			answer := getAnswerInt(questionSection.ID, uint(review.ID))
			db.Create(&answer)
		}
	}

	fmt.Println("\n-------------------------------------------------------------------")

	//fmt.Println("\nGet Sections from Templates")
	//executeRawSQLString("select * from question_templates t join question_sections s on s.question_template_id = t.id", db, &[]models.QuestionTemplate{})

	//fmt.Println("\nGet Survey, ReviewID from Answer")
	//executeRawSQLString("select a.value, s.label from answer_ints a join reviews r on r.id = a.review_id join question_sections s on s.id = a.question_section_id", db, &[]models.AnswerInt{})


	if db.Error != nil {
		log.Fatal(db.Error)
	}
}

func executeRawSQLString(sql string, db *gorm.DB, questionsQuery interface{}) {
	result := db.Raw(sql).Find(questionsQuery)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("Raw SQL query:", sql)
	spew.Dump(questionsQuery)
}

func getReviewImage(reviewID uint, data map[string]interface{}) models.ReviewImage {

	return models.ReviewImage{
		ReviewID:         reviewID,
		Filename:         toString(data, "file_name"),
		FilenameBase:     toString(data, "file_name_base"),
		Width:            toUint64(data, "width"),
		Height:           toUint64(data, "height"),
		Size:             toUint64(data, "size"),
		Comment:          toString(data, "comment"),
		CreatedUserID:    toUint64(data, "created_user_id"),
		CreatedURL:       toString(data, "created_url"),
		UpdatedUserID:    toInt(data, "updated_user_id"),
		UpdatedURL:       toString(data, "updated_url"),
		ACConversionFlag: toUint8(data, "acc_conversion_flag"),
	}
}

func getReviewKeys(data map[string]interface{}) models.ReviewKeys {

	return models.ReviewKeys{
		BookingID:     toUint  (data, "booking_id"),
		TrUserBasicID: toUint  (data, "tr_user_basic_id"),
		Hash:          toString(data, "hash"),
		CreatedURL:    toString(data, "created_url"),
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

func getAnswerInt(questionSectionID uint, reviewID uint) models.AnswerInt {
	currentTime := getCurrentTime()
	
	return models.AnswerInt{
		QuestionSectionID:  &questionSectionID,
		QuestionOptionID:   nil,
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
		Status:             toString(data, "status"),
		DisplayUserName:    toString(data, "display_user_name"),
		Title:              toString(data, "title"),
		Advice:             toString(data, "advice"),
		GoWithID:           toUint16(data, "go_with_id"),
		CommentDate:        &currentTime,
		Content:            toString(data, "content_en"),
	 	LangID:             toInt   (data, "lang_id"),
	 	ActivityDate:       &currentTime,
	 	PointCurrency:      toString(data, "point_currency"),
	 	CreatedUserID:      toInt   (data, "created_user_id"),
	 	CreatedURL:         toString(data, "created_url"),
	 	UpdatedUserID:      toInt   (data, "updated_user_id"),
	 	UpdatedURL:         toString(data, "updated_url"),
	 }
}           

func getContentTranslationReviewContent(reviewContentID uint64, data map[string]interface{}) models.ContentTranslation {

	return models.ContentTranslation{
	        TranslatedContent:  "test",
	        ContentType:        "reply",
	        ReviewContentID:    &reviewContentID,
	        ReplyID:            nil,
	        ReviewImageID:      nil,
	        LangID:             0,
	        Translator:         "google",
	        HumanApprovalID:    12345,
	 }
}           

func getContentTranslationReply(replyID uint64, data map[string]interface{}) models.ContentTranslation {

	return models.ContentTranslation{
	        TranslatedContent:  "test",
	        ContentType:        "reply",
	        ReviewContentID:    nil,
	        ReplyID:            &replyID,
	        ReviewImageID:      nil,
	        LangID:             0,
	        Translator:         "google",
	        HumanApprovalID:    12345,
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
