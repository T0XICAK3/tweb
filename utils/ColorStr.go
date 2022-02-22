package utils

import (
	"fmt"
	"strings"
)

func ColorStrResolve(str, colorStr string) {
	//fontColor,mode,backColor
	//4[10,1,20],9[5,2,14]
	//4[10,1,20]染四个字符，字是10号色，模式1，背景色20
	csParts := strings.Split(colorStr, "],")
	finalChar := string(rune(csParts[len(csParts)-1][len(csParts[len(csParts)-1])-1]))
	if finalChar == "]" {
		csParts[len(csParts)-1] = csParts[len(csParts)-1][:len(csParts[len(csParts)-1])-1]
	}
	outIndex := 0
	endIndex := 0
	for _, temp := range csParts {
		lengthStr, colorPack := Divide(temp, "[")
		length := Atoi(lengthStr)
		if outIndex+length > len(str[outIndex:]) {
			endIndex = outIndex + len(str[outIndex:])
		} else {
			endIndex = outIndex + length
		}
		colors := strings.Split(colorPack, ",")
		d := Atoi(colors[1])
		b := Atoi(colors[2])
		f := Atoi(colors[0])
		//fmt.Print("%c[%d;%d;%dm%s%c[0m",0x1B,d,b,f,str[outIndex:endIndex], 0x1B)
		fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1B, d, b, f, str[outIndex:endIndex], 0x1B)
		outIndex = endIndex
	}
	if endIndex < len(str) {
		fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1B, DEFAULT_MODE, DEFAULT_BCOLOR, DEFAULT_FCOLOR, str[endIndex:], 0x1B)
	}
}
func test() {
	for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
		for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
			for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
				fmt.Print(fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, d, b, f, "N0P3", 0x1B))
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}
