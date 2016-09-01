package main

import "os"

func main() {
	hr("initialize")
	tokens, _ := initialize(os.Args[1])

	hr("word count")
	wordCount := wordCount(tokens)
	wordCount.print(true)
	wordCount.plot(true, "word")
	wc := wordCount.appearances()
	wc.plot(true, "word")

	hr("rune count")
	runeCount := runeCount(tokens)
	runeCount.print(true)
	runeCount.plot(true, "rune")
	rc := runeCount.appearances()
	rc.plot(true, "rune")

	hr("pair count")
	pairCount := pairCount(tokens)
	pairCount.print(true)
	pairCount.plot(true, "pair")
	pc := pairCount.appearances()
	pc.plot(true, "pair")

	identifyLanguage(wordCount, pairCount, runeCount)
}
