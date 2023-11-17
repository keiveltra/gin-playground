package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

type Review struct {
	gorm.Model
	ID                  uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	ServiceKey          string     `gorm:"type:enum('ac','ticket');index"`
	ServiceCategoryID   uint64     `gorm:"type:int unsigned"`
	BookingID           uint64     `gorm:"type:int unsigned"`
	UserBasicID         uint64     `gorm:"type:int unsigned;index"`
	Rate                uint8      `gorm:"type:tinyint unsigned;index:idx_rate;default:5"`
	DisplayUserName     string     `gorm:"type:varchar(64)"`
	Title               string     `gorm:"type:varchar(256)"`
	Review              string     `gorm:"type:varchar(4000)"`
	Advice              string     `gorm:"type:varchar(4000)"`
	GoWithID            uint16     `gorm:"type:smallint unsigned"`
	FirstReviewID       uint64     `gorm:"type:int unsigned;index"`
	OrgReviewID         uint64     `gorm:"type:int unsigned"`
	PtrComment          string     `gorm:"type:varchar(1000)"`
	LikeCount           uint64     `gorm:"type:int unsigned"`
	Status              string     `gorm:"type:enum('new','pending','published','declined','deleted');index"`
	PtrStatus           string     `gorm:"type:enum('pending','published','declined')"`
	UseFlag             uint8      `gorm:"type:tinyint unsigned;index"`
	MappingID           int64      `gorm:"type:int"`
	CdFlag              uint8      `gorm:"type:tinyint unsigned;default:0"`
	PostDate            *time.Time `gorm:"type:datetime"`
	CommentDate         *time.Time `gorm:"type:datetime"`
	StatusChangeDate    *time.Time `gorm:"type:datetime"`
	StatusChangeID      int        `gorm:"type:int"`
	PtrStatusChangeDate *time.Time `gorm:"type:datetime"`
	PtrStatusChangeID   int        `gorm:"type:int"`
	MSiteID             int        `gorm:"type:int"`
	LangID              int        `gorm:"type:int unsigned;index"`
	MOriginID           uint64     `gorm:"type:int unsigned"`
	ActivityDate        *time.Time `gorm:"type:date"`
	PtrBasicID          int        `gorm:"type:int unsigned;index"`
	PointCurrency       string     `gorm:"type:varchar(10)"`
	Created             *time.Time `gorm:"type:datetime"`
	CreatedUserID       int        `gorm:"type:int"`
	CreatedURL          string     `gorm:"type:varchar(512)"`
	Updated             *time.Time `gorm:"type:datetime"`
	UpdatedUserID       int        `gorm:"type:int"`
	UpdatedURL          string     `gorm:"type:varchar(512)"`
	ACConversionFlag    uint8      `gorm:"type:tinyint unsigned;index;default:0"`

	Answers []Answer `gorm:"foreignKey:ReviewID"`
}

type ReviewImage struct {
	gorm.Model
	ID                 uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
	Filename           string     `gorm:"type:varchar(128)"`
	FilenameBase       string     `gorm:"type:varchar(128)"`
	Width              uint64     `gorm:"type:int unsigned"`
	Height             uint64     `gorm:"type:int unsigned"`
	Size               uint64     `gorm:"type:int unsigned"`
	Comment            string     `gorm:"type:varchar(1000)"`
	Created            *time.Time `gorm:"type:datetime"`
	CreatedUserID      uint64     `gorm:"type:int unsigned"`
	CreatedURL         string     `gorm:"type:varchar(512)"`
	Updated            *time.Time `gorm:"type:datetime"`
	UpdatedUserID      int        `gorm:"type:int"`
	UpdatedURL         string     `gorm:"type:varchar(512)"`
	ACConversionFlag   uint8      `gorm:"type:tinyint unsigned;index;default:0"`
}

type ReviewKeys struct {
	ID                 uint      `gorm:"column:id;primaryKey"`
	BookingID          uint      `gorm:"column:booking_id;index"`
	TrUserBasicID      uint      `gorm:"column:tr_user_basic_id"`
	Hash               string    `gorm:"column:hash;unique"`
	Created            time.Time `gorm:"column:created"`
	CreatedUserID      uint      `gorm:"column:created_user_id"`
	CreatedURL         string    `gorm:"column:created_url"`
	Updated            time.Time `gorm:"column:updated"`
	UpdatedUserID      uint      `gorm:"column:updated_user_id"`
	UpdatedURL         string    `gorm:"column:updated_url"`
}

type Question struct {
	ID                 uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	QuestionTemplateID uint      `gorm:"type:int unsigned" json:"question_template_id"`
	ServiceKey         string    `gorm:"type:enum('activity','ticket','point')" json:"service_key"`
	ServiceCategoryID  uint      `gorm:"type:int unsigned" json:"service_target_id"`
	CreatedAt          time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"type:datetime" json:"updated_at"`
	
	QuestionSections []QuestionSection `gorm:"foreignKey:QuestionTemplateID"`
}

type QuestionTemplate struct {
	ID        uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`

	Questions []Question `gorm:"foreignKey:QuestionTemplateID"`
}

type QuestionSection struct {
	ID                uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	QuestionTemplateID uint     `gorm:"type:int unsigned" json:"question_template_id"`
	Type              string    `gorm:"type:enum('normal','weather','multi_choice')" json:"type"`
	Label             string    `gorm:"type:varchar(100)" json:"label"`
	SortOrder         uint      `gorm:"type:int unsigned" json:"sort_order"`
	CreatedAt         time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"type:datetime" json:"updated_at"`

	QuestionOptions []QuestionOption `gorm:"foreignKey:QuestionSectionID"`
}
 
type QuestionOption struct {
	ID                uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	QuestionSectionID uint      `gorm:"type:int unsigned" json:"question_section_id"`
	Type              uint8     `gorm:"type:tinyint unsigned" json:"type"`
	Label             string    `gorm:"type:varchar(100)" json:"label"`
	SortOrder         uint      `gorm:"type:int unsigned" json:"sort_order"`
	CreatedAt         time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"type:datetime" json:"updated_at"`

	Answers []Answer `gorm:"foreignKey:QuestionOptionID"`
}

type Answer struct {
	ID                uint      `gorm:"type:int unsigned;primary_key;auto_increment" json:"id"`
	QuestionSectionID uint      `gorm:"type:int unsigned" json:"question_section_id"`
	QuestionOptionID  uint      `gorm:"type:int unsigned" json:"question_field_id"`
	ReviewID          uint      `gorm:"type:int unsigned" json:"review_id"`
	Value             uint      `gorm:"type:int unsigned" json:"value"`
	CreatedAt         time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"updated_at"`

	Review Review `gorm:"foreignKey:ReviewID"`
	QuestionOption QuestionOption `gorm:"foreignKey:QuestionOptionID"`
}

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
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	http://localhost:8080/admin \
	-H 'authorization: Basic Zm9vOmJhcg==' \
	-H 'content-type: application/json' \
	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
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
