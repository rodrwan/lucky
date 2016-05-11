package model

import "math"

// Sample asdfs
type Sample struct {
	Ngram    string
	Freq     float64
	Classes  map[uint]float64
	Prob     float64
	Tfidf    map[uint]float64
	Maximum  float64
	Minimum  float64
	Weighted bool
}

func (s *Sample) add() {
	s.Freq++
}

// this method return max category key looking for best tfidf value.
func (s *Sample) maxKey() (maxKey uint) {
	var maximum float64
	for key, value := range s.Classes {
		if value > maximum {
			maximum = value
			maxKey = key
		}
	}
	return
}

// General probability of 's' according to it general frequency
// and which category seen this ngram.
func (s *Sample) toProb() {
	for _, value := range s.Classes {
		s.Prob = value / s.Freq
		if s.Prob > s.Maximum {
			s.Maximum = s.Prob
		}
		if s.Prob < s.Minimum {
			s.Minimum = s.Prob
		}
	}
	return
}

// Tfidf value of sample s for each class detected
func (s *Sample) toTfIdf(catsLen float64) {
	if !s.Weighted {
		sampLen := float64(len(s.Classes))
		div := catsLen / sampLen
		idf := math.Log10(1.0 + div)

		for key, tf := range s.Classes {
			tf = 1 + math.Log(tf)
			s.Tfidf[key] = tf * idf // * s.Prob
			if s.Tfidf[key] > s.Maximum {
				s.Maximum = s.Tfidf[key]
			}
			if s.Tfidf[key] < s.Minimum {
				s.Minimum = s.Tfidf[key]
			}
		}
	}
}
