package carrier

import (
	"github.com/lab-smart/sms-gateway/carrier_config"
)

type CarrierInterface interface {
	Send(from string, to string, content string) *carrier_config.CarrierResponse
}
