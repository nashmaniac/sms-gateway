package scripts

import (
	"fmt"
	"log"
	"sms-gateway/repository"
	"sms-gateway/services"
)

func PopulateMessageTemplateDB() {
	messages := [][]string {
		{"ভাল থাকবেন %v", "OTP"},
		{"আপনার অ্যাাপটি সচল করতে এই কোডটি প্রবেশ করান %v", "OTP"},
		{"আপনার গোপন কোড %v", "OTP"},
		{"%v হচ্ছে আপনার লগইন কোড", "OTP"},
		{"সালাম। আপনার কোড %v। ধন্যবাদ। ", "OTP"},
	}
	repo := repository.NewSmsRepository()
	smsService := services.NewSmsService(repo)
	for _, message := range messages {
		savedTemplate := smsService.CreateMessageTemplate(message[0], message[1])
		log.Println(savedTemplate.Id)
	}
}

func PopulateSenderToDB() {
	senders := []string {
		"8801886267494",
		"8801886267495",
		"8801719267494",
		"8801720267494",
	}
	repo := repository.NewSmsRepository()
	smsService := services.NewSmsService(repo)

	for _, sender := range senders {
		savedmodel := smsService.CreateSender(sender)
		log.Println(savedmodel)
	}
}

func PopulateBusinessEntityToDB() {
	names := []string {"BAT", "Unilever", "Trust Bank"}

	repo := repository.NewSmsRepository()
	smsService := services.NewSmsService(repo)

	for i, name := range names {
		model := smsService.CreateBusinessEntity(name)
		log.Println(fmt.Sprintf("%v - %v - %v", i, model.Name, model.Id))
	}
}
