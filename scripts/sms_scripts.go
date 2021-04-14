package scripts

import (
	"log"
	"math/rand"
	"sms-gateway/repository"
	"sms-gateway/services"
	"sms-gateway/utils"
	"strconv"
)

func TestSMSFormatting() {
	repo := repository.NewSmsRepository()
	service := services.NewSmsService(repo)
	messageTemplate := service.FindLeastUsedMessageTemplate()
	n := rand.Int()
	englishString := strconv.FormatInt(int64(n), 10)
	converter := utils.CodeConverter{
		Code:        englishString,
		Source:      "en",
		Destination: "bd",
	}
	convertedMessage := converter.ConvertMessage()
	messageToSend := messageTemplate.OutputFormattedMessage(*convertedMessage)
	log.Println(messageToSend)
}

func TestSMSSending() {
	apiKey := "XoEFfRsWxPLDnJObCsNV"
	repo := repository.NewSmsRepository()
	service := services.NewSmsService(repo)
	n := rand.Int()
	englishString := strconv.FormatInt(int64(n), 10)
	to := "8801886267494"
	log.Println(apiKey, to, englishString)
	log.Println(service)
	model, err := service.SendTextMessage(apiKey, englishString, to, "en", "bd")
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(model.Id)

}

func TestSMSSendingCarrier() {
	dispatcher := utils.GetMessageDispatcher("1618420310962851000", "adareach", "01886267494", "Hello World")
	dispatcher.Send()
}