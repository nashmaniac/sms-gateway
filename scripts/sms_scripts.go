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
