package controllers

import (
	"github.com/gin-gonic/gin"
	"veltra.com/gin_playground/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
		"strings"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is index!!!!",
	})
}

//
// In some project/framework usually BE has Costraint class,
// i.e. AnswerConstraint.class to handle the validation on the backend.
// In another case (i.e. rails) it is defined in model class.
// You can choose the practice of your preference, yet it is generally
// recommended that this validation is defined per class.
//
// TODO: add function arguments to make it sufficient
func validate(product_id string, service_key string) ([]string, int) {
	var errors []string
        httpStatus := 200

	if product_id == "" {
		errors = append(errors, "product_id is missing.")
                httpStatus = 400 // for example. put whatever number you need
        }

	if service_key == "" {
		errors = append(errors, "service_key is missing.")
                httpStatus = 400 // for example. put whatever number you need
        }

        return errors, httpStatus
}

func GetProduct(c *gin.Context) {
	product_id         := c.Params.ByName("product_id")
	service_key        := c.Query("service_key")
	lang_id            := c.Query("lang_id")
	review_status      := c.Query("review_status")
	reply_status       := c.Query("reply_status")
	activity_date_from := c.Query("activity_date_from")
	activity_date_to   := c.Query("activity_date_to")
	posted_date_from   := c.Query("posted_date_from")
	posted_date_to     := c.Query("posted_date_to")
	plan               := c.Query("plan")
	has_image          := c.Query("has_image")
	rating_range       := c.Query("rating_range")
	attended_as        := c.Query("attended_as")
	participated_month := c.Query("participated_month")
	survey_id          := c.Query("survey_id")
	survey_score       := c.Query("survey_score")
	sort_by            := c.Query("sort_by")
	limit              := c.Query("limit")
	page               := c.Query("page")

	db, err := getDatabase()
	if err != nil {
			// raise error
	}

	errors, httpStatus := validate(product_id, service_key)

	review_list := getReviewList(
			db,
			product_id,
			service_key,
			lang_id, 
			review_status,
			reply_status,
			activity_date_from,
			activity_date_to, 
			posted_date_from, 
			posted_date_to,
			plan,
			has_image,
			rating_range,
			attended_as,
			participated_month,
			survey_id,
			survey_score,
			sort_by,
			limit,
			page,
	)

	var total_page_count     = 1 // TODO: implement
	var total_items_count    = 1 // TODO: implement
	var total_review_count   = 1 // TODO: implement
	var average_review_score = 3 // TODO: implement

	c.JSON(httpStatus, gin.H{
		"page":                 page,
		"limit":                limit,
		"total_page_count":     total_page_count,
		"total_items_count":    total_items_count,
		"total_review_count":   total_review_count,
		"average_review_score": average_review_score,
		"review_list":          review_list,
                "errors":               errors,
	})
}

// TODO: move it to some other class
func getDatabase() (*gorm.DB, error) {
	dsn := "moomin:moomin@tcp(127.0.0.1:3306)/test?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

func addCondition(
	qb        *strings.Builder,
	args      *[]interface{}, 
	condition string,
	value     string,
) {
	if value != "" {
		if qb.Len() > 0 {
			qb.WriteString(" AND ")
		}
		qb.WriteString(condition + " = ?")
		*args = append(*args, value)
	}
}

// TODO: move it to helper class or sth.
func getReviewList(
	db                 *gorm.DB,
	product_id         string,
	service_key        string,
	lang_id            string, 
	review_status      string,
	reply_status       string,
	activity_date_from string,
	activity_date_to   string, 
	posted_date_from   string, 
	posted_date_to     string,
	plan               string,
	has_image          string,
	rating_range       string,
	attended_as        string,
	participated_month string,
	survey_id          string,
	survey_score       string,
	sort_by            string,
	limit              string,
	page               string,
) []models.Review {
	var reviews []models.Review
	var qb strings.Builder
	var args []interface{}

	addCondition(&qb, &args, "product_id",    product_id)
	addCondition(&qb, &args, "service_key",   service_key)
	addCondition(&qb, &args, "lang_id",       lang_id)
	addCondition(&qb, &args, "review_status", review_status)
	addCondition(&qb, &args, "reply_status", reply_status)
	addCondition(&qb, &args, "activity_date_from", activity_date_from)
	addCondition(&qb, &args, "activity_date_to", activity_date_to)
	addCondition(&qb, &args, "posted_date_from", posted_date_from)
	addCondition(&qb, &args, "posted_date_to", posted_date_to)
	addCondition(&qb, &args, "plan", plan)
	addCondition(&qb, &args, "has_image", has_image)
	addCondition(&qb, &args, "rating_range", rating_range)
	addCondition(&qb, &args, "attended_as", attended_as)
	addCondition(&qb, &args, "participated_month", participated_month)
	addCondition(&qb, &args, "survey_id", survey_id)
	addCondition(&qb, &args, "survey_score", survey_score)
	addCondition(&qb, &args, "sort_by", sort_by)
	addCondition(&qb, &args, "limit", limit)
	addCondition(&qb, &args, "page", page)

	query := qb.String()

	result := db.Where(query, args...).Find(&reviews)
	if result.Error != nil {
	    // handle error
	}
	return reviews
}
