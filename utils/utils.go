package utils

import (
	"errors"

	"github.com/agext/levenshtein"
)

// ClosestMatch will find the closest batch between a string and
// an array of strings. It will give back an warning if two strings
// are two close.
func ClosestMatch(name string, alternatives []string) (string, error) {
	params := levenshtein.NewParams()

	bestName := ""
	currBestScore := 0.0
	nextBestScore := 0.0

	for _, alternative := range alternatives {
		score := levenshtein.Similarity(name, alternative, params)

		if score >= currBestScore {
			bestName = alternative
			nextBestScore = currBestScore
			currBestScore = score
		} else if score >= nextBestScore {
			nextBestScore = score
		}
	}

	if currBestScore > nextBestScore*1.1 {
		return bestName, nil
	}

	return "", errors.New("unable to pick out best match")
}
