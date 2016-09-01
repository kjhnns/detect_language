package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func hr(title string) {
	title = strings.ToUpper(title)
	fmt.Printf("\n################## %20s ##################\n", title)
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func stringInSlice(term string, list []string) bool {
	i := sort.SearchStrings(list, term)
	if i >= len(list) {
		return false
	} else {
		return (list[i] == term)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
