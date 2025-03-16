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

type Problem struct {
	CenterLetter string
	Letters      []string
	TopMatches   []string
	MaxPoints    int
}

type Match struct {
	Word   string
	Points int
}

var translationTable = map[rune]rune{
	'é': 'e', 'à': 'a', 'è': 'e', 'ù': 'u', 'â': 'a', 'ê': 'e', 'î': 'i', 'ô': 'o', 'û': 'u', 'ç': 'c', 'ï': 'i', 'ü': 'u', 'ë': 'e', 'ö': 'o', 'ä': 'a'}
var uniqueWords = make(map[string]struct{})
var wordsList []Word
var wg sync.WaitGroup
var mutex sync.Mutex

func preprocessWord(word string) string {
	translated := ""
	for _, c := range word {
		if _, exists := translationTable[c]; exists {
			translated += string(translationTable[c])
		} else {
			translated += string(c)
		}
	}

	if strings.ContainsAny(word, "-") || len(translated) < 4 {
		return ""
	}
	// check translated string is ascii
	for _, c := range translated {
		if c > 127 {
			fmt.Println("Non ascii character", c)
			fmt.Println("Original word", word)
			return ""
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
	registered_candidates := make(map[int]struct{})
	for _, word := range frenchWords {
		wordStr := strings.ToLower(word)
		uniqueWords[wordStr] = struct{}{}
		preprocessedWord := preprocessWord(wordStr)
		if preprocessedWord != "" {
			mutex.Lock()
			wordsList = append(wordsList, Word{Raw: word, Preprocessed: preprocessedWord})
			i := len(wordsList) - 1
			mutex.Unlock()
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				if len(set(wordsList[i].Preprocessed)) == 7 {
					mutex.Lock()
					mask := calculateMask(wordsList[i].Preprocessed)
					if _, exists := registered_candidates[mask]; !exists {
						registered_candidates[mask] = struct{}{}
						candidates = append(candidates, i)
					}
					mutex.Unlock()
				}
			}(i)
		}
	}
	wg.Wait()

	fmt.Println("Processed all words")

	masks := make([]int, len(wordsList))
	for i, word := range wordsList {
		masks[i] = calculateMask(word.Preprocessed)
	}

	result := make([]Problem, 0)
	min_count := len(wordsList)
	max_count := 0
	for i, index := range candidates {
		currentMask := masks[index]
		count := 0
		matches := make([]Match, 0)
		for j := range wordsList {
			otherMask := masks[j]
			if (otherMask & currentMask) == otherMask {
				count++
				// Always add 5 bonus letter points to the word
				// to calculate max points
				points := getBaseNumberOfPoints(wordsList[j].Preprocessed) + 5
				if (otherMask & currentMask) == currentMask {
					points += 7
				}
				matches = append(matches, Match{wordsList[j].Preprocessed, points})
			}
		}

		if count < min_count {
			min_count = count
			fmt.Println("Word", wordsList[index].Raw)
			fmt.Println("New min count", min_count)
			fmt.Println("Words", matches)
		}

		if count > max_count {
			max_count = count
			fmt.Println("Word", wordsList[index].Raw)
			fmt.Println("New max count", max_count)
		}

		available := make([]string, 0)
		for c := range set(wordsList[index].Preprocessed) {
			count_letter := 0
			for _, match := range matches {
				if strings.Contains(match.Word, string(c)) {
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

		for k := range matches {
			for j := k + 1; j < len(matches); j++ {
				if matches[k].Points < matches[j].Points {
					matches[k], matches[j] = matches[j], matches[k]
				}
			}
		}

		runeLetters := set(wordsList[index].Preprocessed)
		// convert runeLetters to strings
		letters := make([]string, 0)
		for c := range runeLetters {
			letters = append(letters, string(c))
		}

		for _, centerLetter := range available {
			topMatches := make([]string, 0, 50)
			j := 0
			maxPoints := 0
			for len(topMatches) < 50 && j < len(matches) {
				if strings.Contains(matches[j].Word, centerLetter) {
					if len(topMatches) < 12 {
						maxPoints += matches[j].Points
					}
					topMatches = append(topMatches, matches[j].Word)
				}
				j++
			}

			problem := Problem{
				CenterLetter: centerLetter,
				Letters:      letters,
				TopMatches:   topMatches,
				MaxPoints:    maxPoints,
			}
			result = append(result, problem)
		}
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

func getBaseNumberOfPoints(word string) int {
	wordPoints := 0
	if len(word) == 4 {
		wordPoints = 2
	} else if len(word) == 5 {
		wordPoints = 4
	} else if len(word) == 6 {
		wordPoints = 6
	} else if len(word) == 7 {
		wordPoints = 12
	} else {
		wordPoints = 12 + (len(word)-7)*3
	}
	return wordPoints
}

func calculateMask(word string) int {
	uniqueLetters := set(word)
	mask := 0
	for c := range uniqueLetters {
		pos := int(c - 'a')
		if pos < 0 {
			fmt.Println(c)
		}
		mask |= (1 << pos)
	}
	return mask
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
