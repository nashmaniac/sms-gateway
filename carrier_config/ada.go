package carrier_config

import (
	"log"
	"os"
)

type AdaReach struct {
	From string
	To string
	Message string
	URL string
	Method string
	Username string
	Password string
}

func GetAdaInstance() *AdaReach{
	return &AdaReach{}
}


func (ada AdaReach) BuildQueryParams() map[string]string {
	m := make(map[string]string)
	m["Username"] = ada.Username
	m["Message"] = ada.Message
	m["Password"] = ada.Password
	m["From"] = ada.From
	m["To"] = ada.To
	return m
}

func (ada *AdaReach) Send(from string, to string, content string) {
	log.Println(ada.From, ada.To, ada.Message)
	ada.InitCarrier(from, to, content)
	log.Println(ada.From, ada.To, ada.Message)
	queryParams := ada.BuildQueryParams()
	log.Println(queryParams)
}

func (ada *AdaReach) InitCarrier(from string, to string, content string) *AdaReach {
	url, url_exists := os.LookupEnv("ADAREACH_URL")
	if !url_exists {
		url = "https://api.mobireach.com.bd/SendTextMessage"
	}

	username, username_exists := os.LookupEnv("ADAREACH_USERNAME")
	if !username_exists {
		username = "shagor"
	}

	password, password_exists := os.LookupEnv("ADAREACH_URL")
	if !password_exists {
		password = "Sh@g0R21AdmiN"
	}
	ada.Username = username
	ada.Password = password
	ada.URL = url
	ada.Method = "GET"
	ada.From = from
	ada.To = to
	ada.Message = content
	return ada
}


