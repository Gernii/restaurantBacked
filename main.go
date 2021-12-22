package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"restaurantBacked/common"
	"restaurantBacked/component/appctx"
	"restaurantBacked/component/uploadprovider"
	"restaurantBacked/middleware"
	userstorage "restaurantBacked/modules/user/storage"
)

type Restaurant struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DB_CONN_STR")
	s3BucketName := "g07"                                     // os.Getenv("s3BucketName")
	s3Region := "ap-southeast-1"                              // os.Getenv("s3Region")
	s3APIKey := "AKIA5FKZIMNBV6WLCBZP"                        // os.Getenv("s3APIKey")
	s3SecretKey := "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" // os.Getenv("s3SecretKey")
	s3Domain := "d195cqgdjmjwdm.cloudfront.net"               // os.Getenv("s3Domain")
	secretKey := "Cola"                                       // os.Getenv("SYSTEM_SECRET")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	log.Println(db, err)
	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appCtx := appctx.NewAppContext(db, s3Provider, secretKey)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Static("/static", "./static")
	v1 := r.Group("/v1")
	authStore := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
	mainRoute(v1, appCtx)
	admin := v1.Group("/admin", middleware.RequiredAuth(appCtx, authStore), middleware.RequiredRoles(appCtx, "admin"))
	{
		admin.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, common.SimpleSuccessResponse("ok"))
		})
	}
	r.Run()
}
