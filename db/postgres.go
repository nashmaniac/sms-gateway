package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sms-gateway/models"
	"strconv"
)

func BuildConnectionObject() models.DBConnectionHolder {
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")
	p, err := strconv.ParseInt(port, 10, 32)
	if err != nil {
		panic("Error in parsing port number")
	}
	return models.DBConnectionHolder{
		Host:     host,
		Password: password,
		Name:     name,
		Username: user,
		Port: int(p),
		Timezone: timezone,
		Sslmode:  sslmode,
	}
}

func GetPostgresConnection() *gorm.DB  {
	dbObj := BuildConnectionObject()
	db, err := gorm.Open(postgres.Open(dbObj.GetDSNString()), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database")
	}
	// run all the automigration in here
	return db
}
