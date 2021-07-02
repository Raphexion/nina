package utils

import "github.com/agext/levenshtein"

func ClosestMatch(name string, alternatives []string) string {
	var best_match = ""
	var best_score = 0.0

	for _, alternative := range alternatives {
		params := levenshtein.NewParams()

		score := levenshtein.Similarity(name, alternative, params)
		if score > best_score {
			best_match = alternative
			best_score = score
		}
	}

	return best_match
}
