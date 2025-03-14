package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func getData() ([]Problem, []Word, error) {
	// read the result.json file
	result := make([]Problem, 0)
	resultData, err := os.ReadFile("result.json")
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	err = json.Unmarshal(resultData, &result)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	wordsList := make([]Word, 0)
	wordsListData, err := os.ReadFile("wordsList.json")
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	err = json.Unmarshal(wordsListData, &wordsList)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return result, wordsList, nil
}

func upload() {
	app := pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDataDir: "../back/pb_data",
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// get data from json files
		result, wordsList, err := getData()
		if err != nil {
			fmt.Println(err)
			return err
		}

		// upload data to the database
		wordsCollection, err := app.FindCollectionByNameOrId("words")
		if err != nil {
			return err
		}
		problemsCollection, err := app.FindCollectionByNameOrId("problems")
		if err != nil {
			return err
		}

		wordIds := make(map[string]string)
		for i, word := range wordsList {
			record := core.NewRecord(wordsCollection)

			record.Set("word", word.Preprocessed)
			record.Set("raw", word.Raw)

			// validate and persist
			// (use SaveNoValidate to skip fields validation)
			err = app.Save(record)
			if err != nil {
				return err
			}
			wordIds[word.Preprocessed] = record.Id

			if i%1000 == 0 {
				fmt.Println(i)
			}
		}

		for i, problem := range result {
			translatedMatches := make([]string, 0, len(problem.TopMatches))
			for _, match := range problem.TopMatches {
				translatedMatches = append(translatedMatches, wordIds[match])
			}

			record := core.NewRecord(problemsCollection)

			record.Set("letters", problem.Letters)
			record.Set("center_letter", problem.CenterLetter)
			record.Set("top_matches", translatedMatches)
			record.Set("max_points", problem.MaxPoints)

			err = app.Save(record)
			if err != nil {
				return err
			}

			if i%1000 == 0 {
				fmt.Println(i)
			}
		}

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
