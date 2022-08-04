package utils

import (
	"bufio"
	"fmt"
	"os"
	path2 "path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var DefaultColor = map[string]int{
	"leftColor":  0,
	"rightColor": 0}

func PatchStringMap(oldMap, patch map[string]string) map[string]string {
	for key, value := range patch {
		oldMap[key] = value
	}
	return oldMap
}
func Atoi(number string) int {
	num, err := strconv.Atoi(number)
	if err != nil {
		errStr := fmt.Sprintln(err)
		ConsoleOut("src.Common.Atoi", errStr[:len(errStr)-1], Red, White)
	}
	return num
}
func Itoa(number int) string {
	num := strconv.Itoa(number)
	return num
}
func InsertStr(targetStr string, index int, str string) string {
	return targetStr[:index] + str + targetStr[index:]
}

func GetKeys(m map[string]string) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func IsIn(value string, values []string) bool {
	for _, v := range values {
		if value == v {
			return true
		}
	}
	return false
}

func Esc(targetStr string, specialStr string) string {
	//置空将匹配转译正则字符
	if specialStr == "" {
		specialStr = "$()*+.[]?^\\{}|"
	}
	set := make(map[rune]struct{})
	for _, char := range specialStr {
		set[char] = struct{}{}
	}
	index := 0
	replaceCount := 0
	for _, char := range targetStr {
		if _, ok := set[char]; ok {
			targetStr = InsertStr(targetStr, index+replaceCount, "\\")
			replaceCount += 1
			index += 1
		}
	}
	//DebugOut('ESC',retStr)
	return targetStr
}
func Match(left, right, text string, whole bool) []string {
	re := regexp.MustCompile(Esc(left, "") + ".+?" + Esc(right, ""))
	match := re.FindAllString(text, -1)
	if !whole {
		for index, result := range match {
			match[index] = result[len(left) : len(result)-len(right)]
		}
	}
	return match
}
func stringCreate(length int, chr byte) string {
	var str = " "
	for index := 0; index < length; index += 1 {
		str += string(chr)
	}
	return str
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func RawOut(word interface{}) {
	fmt.Print(word)
}
func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		ConsoleOut("src.Common.Input", "Input error!", Red, Red)
		return ""
	} else {
		return scanner.Text()
	}
}
func Divide(str, split string) (string, string) {
	result := strings.Split(str, split)
	length := len(result)
	if length == 2 {
		return result[0], result[1]
	} else if length > 2 {
		return result[0], strings.Join(result[1:], split)
	} else {
		return result[0], ""
	}
}
func ConsoleOut(identity, word string, leftColor, rightColor Color) {
	leftColor += 30
	rightColor += 30
	//SPACE  TAB
	if leftColor == TAB+30 {
		identity = "[" + stringCreate(len(identity), ' ') + "]:"
	} else if leftColor == SPACE+30 {
		identity = ""
	} else {
		identity = "[" + identity + "]:"
	}
	//--------------------------------
	identityColorStr := Itoa(len(identity)) + "[" + Itoa(int(leftColor)) + "," + Itoa(int(DEFAULT_MODE)) + "," + Itoa(int(DEFAULT_BCOLOR+40)) + "]"
	wordColorStr := Itoa(len(word)) + "[" + Itoa(int(rightColor)) + "," + Itoa(int(DEFAULT_MODE)) + "," + Itoa(int(DEFAULT_BCOLOR+40)) + "]"
	//colorStr:="4[10,1,20],9[5,2,14]"
	colorStr := identityColorStr + "," + wordColorStr
	ColorStrResolve(identity+word, colorStr)
	//print(colorStr)
	fmt.Print("\n")
}
func FillEcho(word string, color Color, sp Mode) {
	colorStr := Itoa(len(word)) + "[" + Itoa(int(color)+30) + ",0," + Itoa(int(sp)) + "]"
	ColorStrResolve(word, colorStr)
}

func RangeExtend(str string) []string {
	commonStr, right := Divide(str, "[")
	if right == "" {
		return []string{str}
	}
	left, suffix := Divide(right, "]")
	left, right = Divide(left, "-")
	start, err := strconv.Atoi(left)
	if err != nil {
		return []string{str}
	}
	end, err := strconv.Atoi(right)
	if err != nil {
		return []string{str}
	}

	if end < start {
		start, end = end, start
	}

	var result []string
	for i := start; i <= end; i++ {
		result = append(result, commonStr+strconv.Itoa(i)+suffix)
	}
	return result
}

func Time2Exit(second int, color Color) {
	for i := second; i > 0; i-- {
		Echo("Exit in "+Itoa(i)+"s", color)
		time.Sleep(time.Duration(1) * time.Second)
	}
	os.Exit(0)
}

func GetNameFromPath(path string, suffix string) string {
	_, fileName := path2.Split(strings.ReplaceAll(path, "\\", "/"))
	fileName, _ = Divide(fileName, suffix)
	return fileName
}
func DebugOut(identity, word string, leftColor, rightColor Color, Debug_Level int) {
	if DEBUG && DEBUG_LEVEL >= Debug_Level {
		ConsoleOut(identity, word, leftColor, rightColor)
	}
}
func Echo(word string, color Color) {
	ConsoleOut("", word, SPACE, color)
}

func init() {
	/*
		if err := termbox.Init(); err != nil {
			Echo("TERM_BOX_INIT_ERROR", Red)
			os.Exit(0)
		}

	*/
}
