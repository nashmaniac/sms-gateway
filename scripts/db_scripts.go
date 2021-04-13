package scripts

import (
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
