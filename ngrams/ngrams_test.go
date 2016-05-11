package ngrams

import "testing"

var invalidWords = []string{
	"compra",
	"normal",
}
var testCases = []struct {
	phrases string
	count   int
	want    []string
	equal   bool
}{
	{
		phrases: "adidas parque arauco",
		count:   1,
		want:    []string{"adidas", "parque", "arauco"},
		equal:   true,
	},
	{
		phrases: "adidas parque arauco",
		count:   2,
		want: []string{
			"adidas",
			"parque",
			"arauco",
			"$ adidas",
			"adidas parque",
			"parque arauco",
			"arauco $",
		},
		equal: true,
	},
	{
		phrases: "adidas parque arauco",
		count:   3,
		want: []string{
			"adidas",
			"parque",
			"arauco",
			"$ adidas",
			"adidas parque",
			"parque arauco",
			"arauco $",
			"$ adidas parque",
			"adidas parque arauco",
			"parque arauco $",
		},
		equal: true,
	},
	{
		phrases: "adidas parque arauco",
		count:   3,
		want: []string{
			"adidas",
			"arauco",
			"adidas parque",
			"parque arauco",
			"adidas parque arauco",
		},
		equal: false,
	},
	{
		phrases: "compra castano las condes",
		count:   3,
		want: []string{
			"compra",
			"castano",
			"condes",
			"las",
			"compra castano",
			"castano las",
			"las condes",
			"compra castano las",
			"castano las condes",
		},
		equal: false,
	},
	{
		phrases: "compra castano las condes",
		count:   3,
		want: []string{
			"castano",
			"condes",
			"$ castano",
			"castano condes",
			"condes $",
			"$ castano condes",
			"castano condes $",
		},
		equal: true,
	},
	{
		phrases: "compra ok market sub centro",
		count:   3,
		want: []string{
			"ok",
			"market",
			"sub",
			"centro",
			"$ ok",
			"ok market",
			"market sub",
			"sub centro",
			"centro $",
			"$ ok market",
			"ok market sub",
			"market sub centro",
			"sub centro $",
		},
		equal: true,
	},
}

func TestMake(t *testing.T) {
	for _, c := range testCases {
		got := Make(c.phrases, c.count, invalidWords)
		if testEq(got, c.want) != c.equal {
			t.Errorf("Make(%q, %d) == %q, want %q", c.phrases, c.count, got, c.want)
		}
	}
}

var testCases2 = []struct {
	phrases string
	want    string
	equal   bool
}{
	{
		phrases: "compra normal memits parque arauco",
		want:    "memits parque arauco",
		equal:   true,
	},
}

func TestNormalize(t *testing.T) {
	for _, c := range testCases2 {
		got := normalize(c.phrases, invalidWords)
		if got != c.want {
			t.Errorf("normalize(%q) == %q, want %q", c.phrases, got, c.want)
		}
	}
}

func testEq(lhs, rhs []string) bool {
	if lhs == nil && rhs == nil {
		return true
	}

	if lhs == nil || rhs == nil {
		return false
	}

	if len(lhs) != len(rhs) {
		return false
	}

	for i := range lhs {
		if lhs[i] != rhs[i] {
			return false
		}
	}

	return true
}
