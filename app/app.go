package app

import (
	"fmt"
	"log"
	"os"
	"sms-gateway/scripts"
)

func checkEnvironmentVariables() {
	envVar := []string {"DB_NAME", "DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_SSLMODE", "DB_TIMEZONE"} 
	for _, key := range envVar {
		val, isPresent := os.LookupEnv(key)
		if !isPresent {
			panic(fmt.Sprintf("Environement variable %v is not present", key))
		}
		log.Println(fmt.Sprintf("%v - %v", key, val))
	}
}

func StartApp() {
	checkEnvironmentVariables()
	scripts.RunTestScripts()
}
