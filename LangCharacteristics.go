package main

import "sort"

type LangCharacteristics struct {
	runes []rune
	words []string
	pairs []string
}

func (lang LangCharacteristics) identify(wc TokenCount, token string) int {
	var score int
	sort.Strings(lang.pairs)
	sort.Strings(lang.words)

	for key, val := range wc {
		if token == "words" {
			if stringInSlice(key, lang.words) {
				score += val
			}
		}
		if token == "pairs" {
			if stringInSlice(key, lang.pairs) {
				score += val
			}
		}
		if token == "runes" {
			for _, r := range lang.runes {
				if string(r) == key {
					score += val
				}
			}
		}
	}

	return score
}
