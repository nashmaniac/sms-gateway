package app

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lab-smart/sms-gateway/db"
	"github.com/lab-smart/sms-gateway/routes"
	"gorm.io/gorm"
)

var dbObj *gorm.DB

func checkEnvironmentVariables() {
	envVar := []string{"DB_NAME", "DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT",
		"DB_SSLMODE", "DB_TIMEZONE", "ADAREACH_URL", "ADAREACH_USERNAME", "ADAREACH_PASSWORD"}
	for _, key := range envVar {
		val, isPresent := os.LookupEnv(key)
		if !isPresent {
			panic(fmt.Sprintf("Environement variable %v is not present", key))
		}
		log.Println(fmt.Sprintf("%v - %v", key, val))
	}
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r = routes.ConfigureSmSRouter(r, db)
	r = routes.ConfigureBusinessRouter(r, db)
	return r
}

func StartApp() {
	checkEnvironmentVariables()
	dbObj = db.GetPostgresConnection()
	dbInstance, _ := dbObj.DB()
	defer dbInstance.Close()
	r := setupRouter(dbObj)
	r.Use(cors.Default())
	r.Run()
}
