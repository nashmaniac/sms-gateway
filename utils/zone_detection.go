package utils

import "strings"


func DetectDestinationBasedOnNumber(phoneNumber string) *string {
	var zone string;
	if strings.HasPrefix(phoneNumber, "880") {
		zone = "bd"
		return &zone
	}
	return nil
}
