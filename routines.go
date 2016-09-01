package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func identifyLanguage(wc, pc, rc TokenCount) {
	var max = func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	var score = map[string]int{}
	score["eng_w"] = engLanguage.identify(wc, "words")
	score["esp_w"] = espLanguage.identify(wc, "words")
	score["ger_w"] = gerLanguage.identify(wc, "words")

	score["eng_p"] = engLanguage.identify(pc, "pairs")
	score["esp_p"] = espLanguage.identify(pc, "pairs")
	score["ger_p"] = gerLanguage.identify(pc, "pairs")

	score["eng_r"] = engLanguage.identify(rc, "runes")
	score["esp_r"] = espLanguage.identify(rc, "runes")
	score["ger_r"] = gerLanguage.identify(rc, "runes")

	eng := score["eng_w"]*score["eng_w"] + score["eng_p"]*2 + score["eng_r"]
	ger := score["ger_w"]*score["ger_w"] + score["ger_p"]*2 + score["ger_r"]
	esp := score["esp_w"]*score["esp_w"] + score["esp_p"]*2 + score["esp_r"]

	hr("lang detection")
	fmt.Println("eng: ", eng, score["eng_w"], score["eng_p"], score["eng_r"])
	fmt.Println("ger: ", ger, score["ger_w"], score["ger_p"], score["ger_r"])
	fmt.Println("esp: ", esp, score["esp_w"], score["esp_p"], score["esp_r"])

	if eng >= ger && eng >= esp {
		fmt.Println("gap: ", (eng - max(ger, esp)), "=> eng")
	} else if ger >= eng && ger >= esp {
		fmt.Println("gap: ", (ger - max(eng, esp)), "=> ger")
	} else if esp >= ger && esp >= eng {
		fmt.Println("gap: ", (esp - max(eng, ger)), "=> esp")
	} else {
		fmt.Println("=> null")
	}
}

func initialize(file string) ([]string, []string) {
	// Initialize
	lines, _ := readLines(file)
	stopwords := loadStopwords()
	sort.Strings(stopwords)
	fmt.Println("Stopwords: ", len(stopwords))

	// preprocessing
	res := strings.Join(lines, "")
	loRes := strings.ToLower(res)
	loTrRes := stripchars(loRes, ".,$“\"'->+<:’?_!çœ[](){}‘”;\t\r\n")
	tokens := strings.Split(loTrRes, " ")
	fmt.Println("Tokens: ", len(tokens))

	tokensClean := normalize(tokens, stopwords)
	fmt.Println("Tokens w/o Stopwords: ", len(tokensClean))

	return tokensClean, stopwords
}

func normalize(input, stopwords []string) []string {
	var rm func(pos int, input, stopwords []string) []string
	rm = func(pos int, input, stopwords []string) []string {
		if pos >= len(input) {
			return input
		}
		if _, err := strconv.Atoi(input[pos]); err == nil || stringInSlice(input[pos], stopwords) || strings.TrimSpace(input[pos]) == "" {
			input = append(input[:pos], input[pos+1:]...)
			return rm(pos, input, stopwords)
		} else {
			return rm(pos+1, input, stopwords)
		}
	}

	return rm(0, input, stopwords)
}

func loadStopwords() []string {
	var res []string
	swFiles := []string{
		"dutch",
		"finnish",
		"german",
		"italian",
		"norwegian",
		"russian",
		"swedish",
		"danish",
		"english",
		"french",
		"hungarian",
		"kazakh",
		"portuguese",
		"spanish",
		"turkish",
	}

	for _, lang := range swFiles {
		lines, _ := readLines("./stopwords/" + lang)
		res = append(res, lines...)
	}

	return res
}
