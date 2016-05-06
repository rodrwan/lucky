package ngrams

import (
	"regexp"
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/kennygrant/sanitize"
)

func normalize(line string) string {
	r := regexp.MustCompile("[^\\p{L}]+")
	re := regexp.MustCompile("[\\s+]")
	lower := strings.ToLower(line)
	clean := sanitize.Accents(lower)
	str := r.ReplaceAllString(clean, " ")
	str = strings.Replace(str, "compra", "", -1)
	str = strings.Replace(str, "pago", "", -1)
	str = strings.Replace(str, "normal", "", -1)
	str = strings.TrimSpace(str)
	return re.ReplaceAllString(str, " ")
}

func stopWords(line string) string {
	return stopwords.CleanString(line, "es", true)
}

// Make creates N-grams of a given string
func Make(str string, N int) (result []string) {
	arr := strings.Fields("$ " + stopWords(normalize(str)) + " $")
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
