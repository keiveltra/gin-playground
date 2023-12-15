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
		 &models.ReplyContent{},
		 &models.ReviewImage{},
		 &models.ReviewContent{},
		 &models.ContentTranslation{},
		 &models.SurveyTemplate{},
		 &models.Question{},
		 &models.QuestionOption{},
		 &models.Answer{},
		 &models.Review{},
		 &models.Vote{},
		 &models.LikeCountStat{},
		 &models.QuestionAverageStat{},
	}

	if err := db.AutoMigrate(_models...); err != nil {
		log.Fatal(err)
	}

        var questions []models.Question
	for _, data := range getTestModelDataFromYaml("question_template") {
		fmt.Println("question_template: ", data)
	    	questionTemplate := getSurveyTemplate(data)
	    	db.Create(&questionTemplate)

		for _, data := range getTestModelDataFromYaml("question_section") {
			fmt.Println("question_section: ", data)
		    	question := getQuestion(data, questionTemplate.ID)
			db.Create(&question)
			questions = append(questions, question)

			if(question.Type == "multi_choice") {
				questionOption := getQuestionOption(question.ID, data)
				db.Create(&questionOption)


			}
		}
	}

	fmt.Println("Injecting the dataset ...")

	for _, data := range getTestModelDataFromYaml("review") {
		fmt.Println("review: ", data)

		review := getReview(data)
		db.Create(&review)

		likeCount := getLikeCountStat(review.ID, data)
                db.Create(&likeCount)

                for _, data := range getTestModelDataFromYaml("vote") {
			fmt.Println("vote: ", data)

			vote := getVote(review.ID, data)
			db.Create(&vote)
                }

		var images []models.ReviewImage
		for _, data := range getTestModelDataFromYaml("review_image") {
			fmt.Println("review_image: ", data)

			reviewImage := getReviewImage(review.ID, data)
			db.Create(&reviewImage)
                        images = append(images, reviewImage)

			translation := getContentTranslation(reviewImage.ID, "image", data)
			db.Create(&translation)
		}

		for _, data := range getTestModelDataFromYaml("review_content") {
			reviewContent := getReviewContent(review.ID, data)
			db.Create(&reviewContent)

			// Translations
			translation := getContentTranslation(reviewContent.ID, "review", data)
			db.Create(&translation)
		}

		reply := getReply(review.ID, data)
		db.Create(&reply)

		replyContent := getReplyContent(reply.ID, data)
		db.Create(&replyContent)

		translation := getContentTranslation(reply.ID, "reply", data)
		db.Create(&translation)


		for _, question := range questions {
			answer := getAnswer(question.ID, uint(review.ID))
			db.Create(&answer)
			stat := getQuestionAverageStat(question.ID, data)
			db.Create(&stat)
		}
	}

	fmt.Println("\n-------------------------------------------------------------------")

	//fmt.Println("\nGet Sections from Templates")
	//executeRawSQLString("select * from question_templates t join question_sections s on s.question_template_id = t.id", db, &[]models.SurveyTemplate{})

	//fmt.Println("\nGet Survey, ReviewID from Answer")
	//executeRawSQLString("select a.value, s.label from answer_ints a join reviews r on r.id = a.review_id join question_sections s on s.id = a.question_section_id", db, &[]models.Answer{})


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

func getQuestionAverageStat(questionID uint, data map[string]interface{}) models.QuestionAverageStat {
	return models.QuestionAverageStat{
		ServiceKey: "ac",
		ProductID:  1123,
		QuestionID: questionID,
		Average:    100,
	}
}

func getLikeCountStat(reviewID uint, data map[string]interface{}) models.LikeCountStat {
	return models.LikeCountStat{
		ServiceKey: "ac",
		ProductID:  1123,
		CategoryID: 12,
		ReviewID:   reviewID,
		LikeCount:  1000,
	}
}

func getVote(reviewID uint, data map[string]interface{}) models.Vote {

	return models.Vote{
		ReviewID:         reviewID,
		TrUserBasicID:    toUint64(data, "tr_user_basic_id"),
	}
}

func getReviewImage(reviewID uint, data map[string]interface{}) models.ReviewImage {

	return models.ReviewImage{
		ReviewID:         reviewID,
		Filename:         toString(data, "file_name"),
		FilenameBase:     toString(data, "file_name_base"),
		Status:           toString(data, "status"),
		Width:            toUint64(data, "width"),
		Height:           toUint64(data, "height"),
		Size:             toUint64(data, "size"),
		Comment:          toString(data, "comment"),
		CreatedUserID:    toUint64(data, "created_user_id"),
		CreatedURL:       toString(data, "created_url"),
		UpdatedUserID:    toInt(data, "updated_user_id"),
		UpdatedURL:       toString(data, "updated_url"),
	}
}

func getSurveyTemplate(data map[string]interface{}) models.SurveyTemplate {
	currentTime := getCurrentTime()

	return models.SurveyTemplate{
		Name:      toString(data, "name"),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}

func getCurrentTime() time.Time {
	return time.Now()
}

func getReply(reviewID uint, data map[string]interface{}) models.Reply {

	return models.Reply {
		ReviewID:           reviewID,
	 	PtrBasicID:         toInt   (data, "ptr_basic_id"),
	}
}

func getReplyContent(replyID uint64, data map[string]interface{}) models.ReplyContent {
	currentTime := getCurrentTime()

	return models.ReplyContent {
		ReplyID:            replyID,
		PtrComment:         toString(data, "ptr_comment"),
		PtrStatus:          "pending",
	 	PtrStatusChangeDate: &currentTime,
	 	PtrStatusChangeID:  toInt   (data, "ptr_status_change_id"),
	}
}

func getQuestion(data map[string]interface{}, questionTemplateID uint) models.Question {
	currentTime := getCurrentTime()

	return models.Question{
		SurveyTemplateID: questionTemplateID,
		ServiceKey:         "activity",
		ProductID:          toUint(data, "service_category_id"),
		Type:               toEnum  (data, "type"),
		Label:              toString(data, "label"),
		SortOrder:          toUint  (data, "sort_order"),
		// Show:               toBool  (data, "show"), // As discussion with PdM, this field is not needed
		Required:           toBool  (data, "required"),
		CreatedAt:          currentTime,
		UpdatedAt:          currentTime,
	}
}

func getQuestionOption(questionID uint, data map[string]interface{}) models.QuestionOption {
	currentTime := getCurrentTime()

	return models.QuestionOption{
		QuestionID: questionID,
		Type:       string(models.OptionTypeCheckbox),
		Label:      toString(data, "label"),
		SortOrder:  toUint  (data, "sort_order"),
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}
}

func getAnswer(questionID uint, reviewID uint) models.Answer {
	currentTime := getCurrentTime()
	numberValue := uint(3)

	return models.Answer{
		QuestionID:         &questionID,
		QuestionOptionID:   nil,
		ReviewID:           reviewID,
		NumberValue:        &numberValue,
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
	currentTime := getCurrentTime()

	return models.Review{
		ServiceKey:         models.Activity,
		ProductID:          toUint64(data, "service_category_id"),
		UserBasicID:        toUint64(data, "user_basic_id"),
		VoteCount:          toUint64(data, "vote_count"),
		Hash:               "aab12394893bbffe",
		PostDate:           &currentTime,
		StatusChangeDate:   &currentTime,
		StatusChangeID:     toInt   (data, "status_change_id"),
	 	LangID:             toInt   (data, "lang_id"),
	 	MSiteID:            toInt   (data, "m_site_id"),
	 	MOriginID:          toUint64(data, "m_origin_id"),
		AttendedAsID:       toUint16(data, "go_with_id"),
	 	ActivityDate:       &currentTime,
	 	CreatedUserID:      toInt   (data, "created_user_id"),
	 	CreatedURL:         toString(data, "created_url"),
	 	UpdatedUserID:      toInt   (data, "updated_user_id"),
	 	UpdatedURL:         toString(data, "updated_url"),
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
		CommentDate:        &currentTime,
		Content:            toString(data, "content_en"),
	 	CreatedUserID:      toInt   (data, "created_user_id"),
	 	CreatedURL:         toString(data, "created_url"),
	 	UpdatedUserID:      toInt   (data, "updated_user_id"),
	 	UpdatedURL:         toString(data, "updated_url"),
	 }
}

func getContentTranslation(contentID uint64, contentType string, data map[string]interface{}) models.ContentTranslation {

	return models.ContentTranslation{
	        TranslatedContent:  "test",
	        ContentType:        contentType,
	        ContentID:          contentID,
	        LangID:             0,
	        Translator:         "google",
	        HumanApprovalID:    12345,
	 }
}

func toBool(data map[string]interface{}, key string) bool {
	value, _ := data[key].(bool)
	return value
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
		return models.SectionTypeRating
	}

	switch value {
	case "normal":
		return models.SectionTypeRating
	case "weather":
		return models.SectionTypeText
	case "multi_choice":
		return models.SectionTypeMultipleAnswers
	default:
		return models.SectionTypeRating
	}
}
