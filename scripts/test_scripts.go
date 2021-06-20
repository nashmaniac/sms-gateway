package scripts

import "github.com/lab-smart/sms-gateway/db"

func RunTestScripts() {
	db.GetPostgresConnection()
	//PopulateMessageTemplateDB()
	//TestSMSFormatting()
	//PopulateSenderToDB()
	//PopulateBusinessEntityToDB()
	TestSMSSending()
	//TestSMSSendingCarrier()
}
