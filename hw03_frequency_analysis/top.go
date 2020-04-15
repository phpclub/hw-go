package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

const TopLimit = 10

func Top10(sInput string) []string {
	aResult := make([]string, 0)
	if len(sInput) == 0 {
		return aResult
	}
	// Очистим от спецсимволов и двойных пробелов
	oReqexpSpecial := regexp.MustCompile(`\t\n\r`)
	oReqexpSpaces := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	sFilteredInput := oReqexpSpaces.ReplaceAllString(oReqexpSpecial.ReplaceAllString(sInput, ""), " ")

	// Заполним map уникальными словами
	mWords := make(map[string]int)
	for _, sWord := range strings.Split(sFilteredInput, " ") {
		_, ok := mWords[sWord]
		// Если значение уже задано - инкрементируем счетчик
		if ok {
			mWords[sWord]++
		} else {
			mWords[sWord] = 1
		}
	}

	// Сделал дополнительный тест если передали короткий текст для анализа, нет смысла продолжать
	if len(mWords) < TopLimit {
		for sWord := range mWords {
			aResult = append(aResult, sWord)
		}
		return aResult
	}

	// Мы не можем сортировать внутри map - переложим полученный результат в структуру и применим sort.Slice
	type tWords struct {
		Name  string
		Count int
	}
	stWords := make([]tWords, 0)
	for kWord, Count := range mWords {
		stWords = append(stWords, tWords{kWord, Count})
	}
	sort.Slice(stWords, func(i, j int) bool { return stWords[i].Count > stWords[j].Count })

	// Вернем первые 10 повторяющихся слов
	for _, sWord := range stWords[:TopLimit] {
		aResult = append(aResult, sWord.Name)
	}
	return aResult
}
