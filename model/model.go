package model

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/rodrwan/categorizer/ngrams"
)

// Samples asdfs
type Samples struct {
	Ngram    string
	Freq     float64
	Classes  map[uint]float64
	Probs    map[uint]float64
	Maximum  float64
	Minimum  float64
	Weighted bool
}

func (s *Samples) add() {
	s.Freq++
}

func (s *Samples) maxKey() (maxKey uint) {
	var maximum float64
	for key, value := range s.Probs {
		if value > maximum {
			maximum = value
			maxKey = key
		}
	}
	return
}

func (s *Samples) toTfIdf(catsLen float64) {
	if !s.Weighted {
		sampLen := float64(len(s.Classes))
		div := catsLen / sampLen
		idf := math.Log10(1.0 + div)

		for key, tf := range s.Classes {
			s.Probs[key] = tf * idf
			if s.Probs[key] > s.Maximum {
				s.Maximum = s.Probs[key]
			}
			if s.Probs[key] < s.Minimum {
				s.Minimum = s.Probs[key]
			}
		}
	}
}

// feature scaling 0 .. 1
func (s *Samples) scale() (prob float64) {
	if !s.Weighted {
		maximum := s.Maximum
		minimum := s.Minimum
		divisor := maximum - minimum

		for key, value := range s.Probs {
			prob = 1.0
			if divisor > 0 {
				prob = (value - minimum) / divisor
			}
			s.Probs[key] = prob
		}
		s.Weighted = true
	}

	return
}

// Fit create a map of ngrams from file
func Fit(path string) (map[string]*Samples, map[uint]float64) {
	modelPath := "model.bin"
	catsPath := "categories.bin"
	if Exists(modelPath) && Exists(catsPath) {
		m, c := Load(modelPath, catsPath)
		return m, c
	}

	m := make(map[string]*Samples)
	cats := make(map[uint]float64)
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		desc := scanner.Text()
		splitedStr := strings.Split(desc, "#")
		category, description := splitedStr[0], splitedStr[1]
		total := ngrams.Make(description, 3)
		i, err := strconv.Atoi(category)
		for _, value := range total {
			cats[uint(i)]++
			if err != nil {
				log.Fatalln(err)
				return nil, nil
			}

			if _, ok := m[value]; ok {
				m[value].Classes[uint(i)]++
				m[value].add()
			} else {
				// init
				m[value] = &Samples{
					Ngram:    value,
					Freq:     1.0,
					Classes:  make(map[uint]float64),
					Probs:    make(map[uint]float64),
					Maximum:  0.0,
					Minimum:  100000000.0,
					Weighted: false,
				}
				m[value].Classes[uint(i)]++
			}
		}
	}

	catsLen := float64(len(cats))
	for _, sample := range m {
		sample.toTfIdf(catsLen)
		sample.scale()
	}

	SaveModel(modelPath, m)
	SaveCats(catsPath, cats)
	return m, cats
}

func bestOption(votes map[uint]uint, freq map[float64]uint) uint {
	var maxFreqKey uint = 1
	var maximum uint
	var maxKey uint
	maxFreq := 0.0
	equals := true

	for key, value := range freq {
		if key > maxFreq {
			maxFreq = key
			maxFreqKey = value
		}
	}

	for key, count := range votes {
		if count > maximum {
			maximum = count
			maxKey = key
		}
		if count != maximum {
			equals = false
		}
	}

	if equals {
		return maxFreqKey
	}
	return maxKey
}

// BestCategory ...
type BestCategory struct {
	ID    uint    // category id
	Name  string  // category name
	Score float64 // category prob
}

// Best ...
type Best struct {
	Prob  float64
	Key   uint
	Ngram string
}

func (b *Best) resetValues() {
	b.Prob = 0.0
	b.Key = 0
	b.Ngram = ""
}

func (b *Best) setValues(prob float64, key uint, ngram string) {
	b.Prob = prob
	b.Key = key
	b.Ngram = ngram
}

// Predict function to get best category
func Predict(m map[string]*Samples, test string, cats map[uint]string, c map[uint]float64) *BestCategory {
	fmt.Printf("test: %s\n", test)
	total := ngrams.Make(test, 3)
	freq := make(map[float64]uint)
	votes := make(map[uint]uint)

	bestOpt := &Best{
		Prob:  0.0,
		Key:   0,
		Ngram: "",
	}

	for _, value := range total {
		sample := m[value]
		if sample != nil {
			maxKey := sample.maxKey()
			prob := sample.Classes[maxKey]
			bestOpt.setValues(prob, maxKey, value)
			freq[sample.Freq] = maxKey
		}
		if bestOpt.Prob > 0.0 {
			votes[bestOpt.Key]++
		}
		bestOpt.resetValues()
	}

	best := bestOption(votes, freq)

	return &BestCategory{
		ID:    best,
		Name:  cats[best],
		Score: bestOpt.Prob,
	}
}
