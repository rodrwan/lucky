package model

import (
	"strings"

	"github.com/rodrwan/lucky/ngrams"
)

// BestCategory ...
type BestCategory struct {
	ID    uint    // category id
	Name  string  // category name
	Score float64 // category prob
}

// Vote ...
type Vote struct {
	Count uint
	Score float64
}

// BestGram ...
type BestGram struct {
	Gram string
	Prob float64
}

// EPSILON ...
var EPSILON = 0.00000001

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

func maxVote(votes map[uint]*Vote, threshold float64) (uint, float64) {
	if len(votes) == 0 {
		return 1, 0.0
	}

	var maxKey uint
	var maxCount uint
	var equals bool
	var maximum float64

	last := uint(10000)
	for key, vote := range votes {
		if last == maxCount {
			equals = true
		} else {
			equals = false
		}
		last = maxCount

		if maxCount < vote.Count {
			maxCount = vote.Count
			maxKey = key
		}
	}

	maximum = votes[maxKey].Score

	if equals {
		maximum := 0.0
		last := 10000.0
		equals = false

		for key, vote := range votes {
			if floatEquals(last, vote.Score) {
				equals = true
			} else {
				equals = false
			}
			last = vote.Score

			if vote.Score > maximum {
				maximum = vote.Score
				maxKey = key
			}
		}
	}

	if equals {
		return 1, 0.0
	}

	return maxKey, maximum
}

func getBestGram(grams map[uint]*BestGram) (string, uint, float64) {
	var maxKey uint
	var maxGram string
	var maxProb float64

	for key, value := range grams {
		if value.Prob > maxProb {
			maxKey = key
			maxGram = value.Gram
			maxProb = value.Prob
		}
	}
	return maxGram, maxKey, maxProb
}

func createBestGram(grams []string, m map[string]*Sample) map[uint]*BestGram {
	bestGrams := make(map[uint]*BestGram)
	for _, value := range grams {
		sample := m[value]
		if sample != nil {
			key := sample.maxKey()
			if _, ok := bestGrams[key]; ok {
				if bestGrams[key].Prob < sample.Prob {
					bestGrams[key].Prob = sample.Prob
					bestGrams[key].Gram = value
				}
			} else {
				bestGrams[key] = &BestGram{
					Prob: sample.Prob,
					Gram: value,
				}
			}

		}
	}
	return bestGrams
}

// Predict function to get best category
func Predict(m map[string]*Sample, test string, cats map[uint]string, words []string, threshold float64) *BestCategory {
	votes := make(map[uint]*Vote)

	unis := ngrams.Make(test, 1, words)
	unigrams := createBestGram(unis, m)
	maxGram, maxKey, maxProb := getBestGram(unigrams)
	if maxKey != 0 {
		if _, ok := votes[maxKey]; ok {
			votes[maxKey].Count++
			votes[maxKey].Score += maxProb
		} else {
			votes[maxKey] = &Vote{
				Count: 1.0,
				Score: maxProb,
			}
		}
	}

	bis := ngrams.Make(test, 2, words)
	var useBis []string
	for _, value := range bis {
		if strings.Index(value, maxGram) > -1 {
			useBis = append(useBis, value)
		}
	}
	bigrams := createBestGram(useBis, m)
	maxGram, maxKey, maxProb = getBestGram(bigrams)
	if maxKey != 0 {
		if _, ok := votes[maxKey]; ok {
			votes[maxKey].Count++
			votes[maxKey].Score += maxProb
		} else {
			votes[maxKey] = &Vote{
				Count: 1.0,
				Score: maxProb,
			}
		}
	}

	tris := ngrams.Make(test, 3, words)
	var useTris []string
	for _, value := range tris {
		if strings.Index(value, maxGram) > -1 {
			useTris = append(useTris, value)
		}
	}
	trigrams := createBestGram(useTris, m)
	_, maxKey, maxProb = getBestGram(trigrams)
	if maxKey != 0 {
		if _, ok := votes[maxKey]; ok {
			votes[maxKey].Count++
			votes[maxKey].Score += maxProb
		} else {
			votes[maxKey] = &Vote{
				Count: 1.0,
				Score: maxProb,
			}
		}
	}

	bestID, bestProb := maxVote(votes, threshold)

	return &BestCategory{
		ID:    bestID,
		Name:  cats[bestID],
		Score: bestProb,
	}
}
