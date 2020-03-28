package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(sInput string) (string, error) {
	// Place your code here
	var sResult string
	rInput := []rune(sInput)
	sResult = ""
	// Если получили пустой слайс - вернем пустоту - нечего распаковывать
	if len(rInput) == 0 {
		return sResult, nil
	}
	bParseAsLetter := false
	for key, val := range rInput {
		// Сохраним порядковое значение
		iNextKey := key + 1
		iPrevKey := key - 1

		if string(val) == "\\" {
			// Установим флаг что началось экранирование
			bParseAsLetter = true
		}
		if key > 0 && unicode.IsDigit(val) {
			iMulti, err := strconv.Atoi(string(val))
			if err != nil {
				log.Fatal(err)
			}
			if unicode.IsDigit(rInput[iPrevKey]) && !bParseAsLetter {
				// 2 цифри подряд
				return "", ErrInvalidString
			} else if bParseAsLetter {
				if iNextKey < len(rInput) && !unicode.IsDigit(rInput[iNextKey]) {
					sResult += string(val)
				} else if iNextKey == len(rInput) {
					sResult += string(val)
				}
				// Не будем дублировать экранизирующий символ
				//sResult += string(rInput[key])
				// Сбросим флаг экранирования
				bParseAsLetter = false
				continue
			} else {
				sResult += strings.Repeat(string(rInput[iPrevKey]), iMulti)
			}
		} else if unicode.IsLetter(val) {
			if iNextKey < len(rInput) && !unicode.IsDigit(rInput[iNextKey]) {
				sResult += string(val)
			} else if iNextKey == len(rInput) {
				sResult += string(val)
			}
		}
	}
	if len(sResult) > 0 {
		return sResult, nil
	} else {
		return "", ErrInvalidString
	}
}
