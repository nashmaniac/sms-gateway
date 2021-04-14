package carrier_wise_response

import "encoding/xml"

type ArrayofServiceClass struct {
	XMLName xml.Name `xml:"ArrayOfServiceClass"`
	ServiceClass []serviceClass `xml:"ServiceClass"`
}

type serviceClass struct {
	XMLName xml.Name `xml:"ServiceClass"`
	MessageId string
	Status int
	StatusText string
	ErrorCode int
	ErrorText string
	SMSCount int
	CurrentCredit float32
}
