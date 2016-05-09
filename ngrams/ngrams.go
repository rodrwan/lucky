package ngrams

import (
	"regexp"
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/kennygrant/sanitize"
)

func normalize(line string, words []string) string {
	r := regexp.MustCompile("[^\\p{L}]+")
	re := regexp.MustCompile("[\\s+]")
	relen := regexp.MustCompile("[^\\w{2}$]+")
	lower := strings.ToLower(line)
	clean := sanitize.Accents(lower)
	str := r.ReplaceAllString(clean, " ")
	str = relen.ReplaceAllString(str, " ")

	if len(words) > 0 {
		for _, word := range words {
			str = strings.Replace(str, word, "", -1)
		}
	}
	str = strings.TrimSpace(str)
	return re.ReplaceAllString(str, " ")
}

func stopWords(line string) string {
	return stopwords.CleanString(line, "es", true)
}

// Make creates N-grams of a given string
func Make(str string, N int, word []string) (result []string) {
	arr := strings.Fields("$ " + stopWords(normalize(str, word)) + " $")
	words := len(arr)

	for k := 0; k < N; k++ {
		step := (words - k)
		length := k + 1
		tmp := make([]string, length)

		for i := 0; i < step; i++ {
			if len(arr[i:i+length]) == 1 && arr[i : i+length][0] == "$" {
				continue
			}

			copy(tmp, arr[i:i+length])
			result = append(result, strings.Join(tmp, " "))
		}
	}
	return
}
