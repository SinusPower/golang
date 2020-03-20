package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

type dictionary struct {
	word string
	freq int
}

// Top10 returns 10 the most frequent words in the input string.
func Top10(in string) []string {
	var res []string
	words := strings.Fields(in)
	dict := make(map[string]int)
	for _, word := range words {
		dict[word]++
	}

	var sorted []dictionary
	for word, freq := range dict {
		sorted = append(sorted, dictionary{word, freq})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].freq > sorted[j].freq
	})

	if len(sorted) >= 10 {
		sorted = sorted[:10]
	}
	for i := range sorted {
		res = append(res, sorted[i].word)
	}
	return res
}
