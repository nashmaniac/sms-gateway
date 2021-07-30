package scripts

func RunTestScripts() {
	// db.GetPostgresConnection()
	//TestSMSFormatting()
	PopulateBusinessEntityToDB()
	PopulateSenderToDB()
	PopulateMessageTemplateDB()
	// TestSMSSending()
	//TestSMSSendingCarrier()
}
