package model

import (
	"fmt"
	"testing"
)

var testVotes = map[uint]*Vote{
	72: {
		Count: 72,
		Score: 2,
	},
	89: {
		Count: 89,
		Score: 11,
	},
}

var testFreq = map[float64]uint{
	9477: 72,
	1510: 89,
	1001: 89,
	897:  89,
	1164: 89,
	293:  89,
	755:  89,
	3738: 89,
	1700: 89,
	7388: 72,
	382:  89,
}

func TestMaxVotes(t *testing.T) {
	vote, _ := maxVote(testVotes)
	expectedVote := uint(89)

	if vote != expectedVote {
		t.Errorf("Vote should be %d, got %d", expectedVote, vote)
	}
	fmt.Printf("MaxVote: %d\n\n", vote)
}

func TestBestOption(t *testing.T) {
	expected := uint(89)
	actual, _ := maxVote(testVotes)
	if actual != expected {
		t.Errorf("Freq should be %d, got %d", expected, actual)
	}
	fmt.Printf("BestOption: %d\n\n", actual)
}

var testVotesEquals = map[uint]*Vote{
	25: {
		Count: 4,
		Score: 63,
	},
	87: {
		Count: 4,
		Score: 100,
	},
}

var testFreq2 = map[float64]uint{
	7:    25,
	1413: 87,
	6:    25,
	990:  87,
}

func TestMaxVotesEquals(t *testing.T) {
	vote, _ := maxVote(testVotesEquals)
	expectedVote := uint(87)

	if vote != expectedVote {
		t.Errorf("Vote should be %d, got %d", expectedVote, vote)
	}
	fmt.Printf("MaxVote: %d\n\n", vote)
}

func TestBestOptionEquals(t *testing.T) {
	expected := uint(87)
	actual, _ := maxVote(testVotesEquals)
	if actual != expected {
		t.Errorf("Freq should be %d, got %d", expected, actual)
	}
	fmt.Printf("BestOption: %d\n\n", actual)
}
