package unicode

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func RuneCountHelper() error{
	buf := []byte("Hello, 世" +
		"界")
	fmt.Println("bytes =", len(buf))
	fmt.Println("runes =", utf8.RuneCount(buf))
	return nil
}

func RuneCountInString() error{
	str := "   Hello,     世界23   "
	fmt.Println("bytes =", len(str))
	fmt.Println("runes =", utf8.RuneCountInString(str))


	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	wordsNum := utf8.RuneCountInString(str)
	fmt.Println("wordsNum = ", wordsNum)
	return nil
}