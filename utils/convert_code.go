package utils

type CodeConverter struct {
	Code string
	Source string
	Destination string
}

func (c CodeConverter) ConvertMessage() *string {
	var convertedString = ""
	if c.Source == "en" && c.Destination == "bd" {
		convertedString = ConvertEngToBD(c.Code)
	}
	var stringPointer *string = nil
	if convertedString != "" {
		stringPointer = &convertedString
	}
	return stringPointer
}
