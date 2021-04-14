package scripts

import "sms-gateway/db"

func RunTestScripts() {
	db.GetPostgresConnection()
	//PopulateMessageTemplateDB()
	//TestSMSFormatting()
	//PopulateSenderToDB()
	//PopulateBusinessEntityToDB()
	TestSMSSending()
	//TestSMSSendingCarrier()
}
