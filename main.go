package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type Review struct {
    gorm.Model
    ID                 uint64     `gorm:"type:int unsigned;primaryKey;autoIncrement"`
    AcActivityID       uint64     `gorm:"type:int unsigned;index"`
    BookingID          uint64     `gorm:"type:int unsigned"`
    UserBasicID        uint64     `gorm:"type:int unsigned;index"`
    Rate               uint8      `gorm:"type:tinyint unsigned;index:idx_rate;default:5"`
    DisplayUserName    string     `gorm:"type:varchar(64)"`
    Title              string     `gorm:"type:varchar(256)"`
    Review             string     `gorm:"type:varchar(4000)"`
    GoWithID           uint16     `gorm:"type:smallint unsigned"`
    FirstReviewID      uint64     `gorm:"type:int unsigned;index"`
    OrgReviewID        uint64     `gorm:"type:int unsigned"`
    PtrComment         string     `gorm:"type:varchar(1000)"`
    LikeCount          uint64     `gorm:"type:int unsigned"`
    Status             string     `gorm:"type:enum('new','pending','published','declined','deleted');index"`
    PtrStatus          string     `gorm:"type:enum('pending','published','declined')"`
    UseFlag            uint8      `gorm:"type:tinyint unsigned;index"`
    MappingID          int64      `gorm:"type:int"`
    CdFlag             uint8      `gorm:"type:tinyint unsigned;default:0"`
    PostDate           *time.Time `gorm:"type:datetime"`
    CommentDate        *time.Time `gorm:"type:datetime"`
    StatusChangeDate   *time.Time `gorm:"type:datetime"`
    StatusChangeID     int        `gorm:"type:int"`
    PtrStatusChangeDate *time.Time `gorm:"type:datetime"`
    PtrStatusChangeID   int        `gorm:"type:int"`
    MSiteID            int        `gorm:"type:int"`
    LangID             int        `gorm:"type:int unsigned;index"`
    MOriginID          uint64     `gorm:"type:int unsigned"`
    ActivityDate       *time.Time `gorm:"type:date"`
    PtrBasicID         int        `gorm:"type:int unsigned;index"`
    PointCurrency      string     `gorm:"type:varchar(10)"`
    Created            *time.Time `gorm:"type:datetime"`
    CreatedUserID      int        `gorm:"type:int"`
    CreatedURL         string     `gorm:"type:varchar(512)"`
    Updated            *time.Time `gorm:"type:datetime"`
    UpdatedUserID      int        `gorm:"type:int"`
    UpdatedURL         string     `gorm:"type:varchar(512)"`
    ACConversionFlag   uint8      `gorm:"type:tinyint unsigned;index;default:0"`
}

var db = make(map[string]string)

func migrateDatabase() {
    db, err := gorm.Open(mysql.Open("username:password@tcp(127.0.0.1:3306)/your_database"), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    db.AutoMigrate(&Review{})
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
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
