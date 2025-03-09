package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Word struct {
	Raw          string
	Preprocessed string
}

type Candidate struct {
	Word             string
	AvailableLetters []string
}

var translationTable = map[rune]rune{
	'é': 'e', 'à': 'a', 'è': 'e', 'ù': 'u', 'â': 'a', 'ê': 'e', 'î': 'i', 'ô': 'o', 'û': 'u', 'ç': 'c'}
var uniqueWords = make(map[string]struct{})
var wordsList []Word
var wg sync.WaitGroup
var mutex sync.Mutex

func preprocessWord(word string) string {
	translated := ""
	if strings.ContainsAny(word, "-") || len(word) < 4 {
		return ""
	}
	for _, c := range word {
		if _, exists := translationTable[c]; exists {
			translated += string(translationTable[c])
		} else {
			translated += string(c)
		}
	}
	return strings.ToLower(translated)
}

func parseList() {
	frenchWords, err := readCSV("./liste_francais_full.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range frenchWords[:5] {
		fmt.Println(row)
	}

	candidates := make([]int, 0)
	for _, word := range frenchWords {
		wordStr := strings.ToLower(word)
		uniqueWords[wordStr] = struct{}{}
		preprocessedWord := preprocessWord(wordStr)
		if preprocessedWord != "" {
			mutex.Lock()
			wordsList = append(wordsList, Word{Raw: word, Preprocessed: preprocessWord(word)})
			i := len(wordsList) - 1
			mutex.Unlock()
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				if len(set(wordsList[i].Preprocessed)) == 7 {
					candidates = append(candidates, i)
				}
			}(i)
		}
	}
	wg.Wait()

	fmt.Println("Processed all words")

	masks := make([]int, len(wordsList))
	candidatesSet := make(map[int]struct{})
	for i, word := range wordsList {
		if _, exists := candidatesSet[i]; !exists {
			uniqueLetters := set(word.Preprocessed)
			mask := 0
			for c := range uniqueLetters {
				pos := int(c - 'a')
				if pos < 0 {
					fmt.Println(c)
				}
				mask |= (1 << pos)
			}
			masks[i] = mask
		}
	}

	result := make([]Candidate, 0)
	min_count := len(wordsList)
	for i, index := range candidates {
		currentMask := masks[index]
		count := 0
		matches := make([]string, 0)
		for j := range wordsList {
			if index == j {
				continue
			}
			otherMask := masks[j]
			if (otherMask & currentMask) == otherMask {
				count++
				matches = append(matches, wordsList[j].Raw)
			}
		}

		if count < min_count {
			min_count = count
			fmt.Println("Word", wordsList[index].Raw)
			fmt.Println("New min count", min_count)
			fmt.Println("Words", matches)
		}

		available := make([]string, 0)
		for c := range set(wordsList[index].Preprocessed) {
			count_letter := 0
			for _, word := range matches {
				if strings.Contains(word, string(c)) {
					count_letter++
				}
			}
			if count_letter >= 12 {
				available = append(available, string(c))
			}
		}

		if len(available) < 1 {
			fmt.Println("Skipping word:", wordsList[index].Raw, " count:", count, " matches:", matches)
			continue
		}

		result = append(result, Candidate{
			Word:             wordsList[index].Preprocessed,
			AvailableLetters: available,
		})
		if i%1000 == 0 {
			fmt.Println(i)
		}
	}

	// Dump as json
	jsonResult, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonWordsList, err := json.Marshal(wordsList)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile("result.json", jsonResult, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile("wordsList.json", jsonWordsList, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func readCSV(file string) ([]string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	rows := make([]string, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		rows = append(rows, line)
	}

	return rows, nil
}

func set(s string) map[rune]struct{} {
	m := make(map[rune]struct{})
	for _, c := range s {
		m[c] = struct{}{}
	}
	return m
}
