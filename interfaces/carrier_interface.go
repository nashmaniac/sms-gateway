package interfaces

type CarrierInterface interface {
	Send(from string, to string, content string)
}
