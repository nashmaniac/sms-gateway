package utils

import (
	"strings"
)

func ConvertEngToBD(code string) string {
	/*
	Convert the code from english to bd
	 */
	// declare the map
	//characterMap := map[string]string {
	//	"1": "১",
	//	"2": "২",
	//	"3": "৩",
	//	"4": "৪",
	//	"5": "৫",
	//	"6": "৬",
	//	"7": "৭",
	//	"8": "৮",
	//	"9": "৯",
	//	"0": "০",
	//}
	characterMap := map[int32]string {
		48: "০",
		49: "১",
		50: "২",
		51: "৩",
		52: "৪",
		53: "৫",
		54: "৬",
		55: "৭",
		56: "৮",
		57: "৯",
	}

	builder := strings.Builder{}
	for _, c := range code {
		//log.Println(fmt.Sprintf("%v - %v - %v", index, c, characterMap[c]))
		builder.WriteString(characterMap[c])
	}

	//log.Println(fmt.Sprintf("%v - %v", code, builder.String()))
	return builder.String()
}
