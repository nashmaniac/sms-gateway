package utils

import (
	"github.com/lab-smart/sms-gateway/carrier_config"
)

type MessageDispatcher struct {
	MessageId string
	From      string
	To        string
	Content   string
}

func GetMessageDispatcher(messageId string, from string, to string, content string) *MessageDispatcher {
	return &MessageDispatcher{
		MessageId: messageId,
		From:      from,
		To:        to,
		Content:   content,
	}
}

func (d *MessageDispatcher) Send() *carrier_config.CarrierResponse {
	// TODO select the lowest cost carrier
	// by default send by ada reach
	carrier := carrier_config.GetAdaInstance()
	return carrier.Send(d.From, d.To, d.Content)
}
