package model

import "math"

// Sample asdfs
type Sample struct {
	Ngram    string
	Freq     float64
	Classes  map[uint]float64
	Probs    map[uint]float64
	Maximum  float64
	Minimum  float64
	Weighted bool
}

func (s *Sample) add() {
	s.Freq++
}

func (s *Sample) maxKey() (maxKey uint) {
	var maximum float64
	for key, value := range s.Probs {
		if value > maximum {
			maximum = value
			maxKey = key
		}
	}
	return
}

func (s *Sample) toTfIdf(catsLen float64) {
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
func (s *Sample) scale() (prob float64) {
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
