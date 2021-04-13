package scripts

import (
	"fmt"
	"math/rand"
	"sms-gateway/utils"
	"strconv"
)

func TestEngToDBCodeConversion() {
	for i := 0 ; i<100; i++ {
		n := rand.Int()
		englishString := strconv.FormatInt(int64(n), 10)
		//log.Println(fmt.Sprintf("%v %v", n, englishString))
		converter := utils.CodeConverter{
			Code:        englishString,
			Source:      "en",
			Destination: "bd",
		}
		message := converter.ConvertMessage()
		fmt.Println(fmt.Sprintf("%v %v", englishString, *message))
		//log.Println(*message)
	}

}
