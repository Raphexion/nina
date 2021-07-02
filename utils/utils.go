package utils

import (
	"errors"

	"github.com/agext/levenshtein"
)

func ClosestMatch(name string, alternatives []string) (string, error) {
	params := levenshtein.NewParams()

	best_name := ""
	curr_best_score := 0.0
	next_best_score := 0.0

	for _, alternative := range alternatives {
		score := levenshtein.Similarity(name, alternative, params)

		if score >= curr_best_score {
			best_name = alternative
			next_best_score = curr_best_score
			curr_best_score = score
		} else if score >= next_best_score {
			next_best_score = score
		}
	}

	if curr_best_score > next_best_score*1.5 {
		return best_name, nil
	}

	return "", errors.New("unable to pick out best match")
}
