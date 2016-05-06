package model

import (
	"fmt"
	"testing"
)

var testVotes = map[uint]uint{
	72: 2,
	89: 11,
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
	vote, equals := maxVote(testVotes)
	expectedVote := uint(89)
	expectedEquels := false

	if vote != expectedVote {
		t.Errorf("Vote should be %d, got %d", expectedVote, vote)
	}
	if equals != expectedEquels {
		t.Errorf("Equals should be %v, got %v", expectedEquels, equals)
	}
	fmt.Printf("MaxVote: %d\n\n", vote)
}

func TestMaxFreq(t *testing.T) {
	expected := uint(72)
	actual := maxFreq(testFreq)
	if actual != expected {
		t.Errorf("Freq should be %d, got %d", expected, actual)
	}
	fmt.Printf("MaxFreq: %d\n\n", actual)
}

func TestBestOption(t *testing.T) {
	expected := uint(89)
	actual := bestOption(testVotes, testFreq)
	if actual != expected {
		t.Errorf("Freq should be %d, got %d", expected, actual)
	}
	fmt.Printf("BestOption: %d\n\n", actual)
}

var testVotesEquals = map[uint]uint{
	25: 2,
	87: 2,
}

var testFreq2 = map[float64]uint{
	7:    25,
	1413: 87,
	6:    25,
	990:  87,
}

func TestMaxVotesEquals(t *testing.T) {
	vote, equals := maxVote(testVotesEquals)
	expectedVote := uint(25)
	expectedEquels := true

	if vote != expectedVote {
		t.Errorf("Vote should be %d, got %d", expectedVote, vote)
	}
	if equals != expectedEquels {
		t.Errorf("Equals should be %v, got %v", expectedEquels, equals)
	}
	fmt.Printf("MaxVote: %d\n\n", vote)
}

func TestMaxFreqEquals(t *testing.T) {
	expected := uint(87)
	actual := maxFreq(testFreq2)
	if actual != expected {
		t.Errorf("Freq should be %d, got %d", expected, actual)
	}
	fmt.Printf("MaxFreq: %d\n\n", actual)
}

func TestBestOptionEquals(t *testing.T) {
	expected := uint(87)
	actual := bestOption(testVotesEquals, testFreq2)
	if actual != expected {
		t.Errorf("Freq should be %d, got %d", expected, actual)
	}
	fmt.Printf("BestOption: %d\n\n", actual)
}
