package controllers

import (
	"github.com/gin-gonic/gin"
	"veltra.com/gin_playground/models"
)

func Index(c *gin.Context) {
    	c.JSON(200, gin.H{
		"message": "this is index!!!!",
	})
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

        review_list := getReviewList(
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

    	c.JSON(200, gin.H{
		"page":  page,
                "limit": limit,
                "total_page_count":     total_page_count,
                "total_items_count":    total_items_count,
                "total_review_count":   total_review_count,
                "average_review_score": average_review_score,
                "review_list":          review_list,
	})
}

// TODO: move it to helper class or sth.
func getReviewList(
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
        review_list := []models.Review{}
	//
        // Query whatever you want.
	//
	return review_list
}
