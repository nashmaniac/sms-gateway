package carrier_config

import (
	"encoding/xml"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
	"sms-gateway/carrier_wise_response"
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

func (ada *AdaReach) Send(from string, to string, content string) *CarrierResponse {
	log.Println(ada.From, ada.To, ada.Message)
	ada.InitCarrier(from, to, content)
	log.Println(ada.From, ada.To, ada.Message)
	queryParams := ada.BuildQueryParams()

	client := resty.New()

	resp, err := client.R().EnableTrace().SetQueryParams(queryParams).Get(ada.URL)
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	if resp.StatusCode() == 200 {
		var serviceClassArray carrier_wise_response.ArrayofServiceClass
		xml.Unmarshal(resp.Body(), &serviceClassArray)
		if len(serviceClassArray.ServiceClass) > 0 {
			return &CarrierResponse{ResponseId: serviceClassArray.ServiceClass[0].MessageId, IsSuccess: true}
		} else {
			return &CarrierResponse{
				IsSuccess: true,
			}
		}
	} else {
		return &CarrierResponse{
			IsSuccess:  false,
			ErrorText:  "Error in Ada Api",
		}
	}
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
	ada.From = "adareach"
	ada.To = to
	ada.Message = content
	return ada
}


