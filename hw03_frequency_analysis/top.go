package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

func Top10(text string) []string {
	re := regexp.MustCompile(`\s+`)
	words := re.Split(text, -1)
	freq := make(map[string]int)
	for _, word := range words {
		if word != "" {
			freq[word]++
		}
	}

	wordFreqs := make([]string, 0)
	for word := range freq {
		wordFreqs = append(wordFreqs, word)
	}

	sort.SliceStable(wordFreqs, func(i, j int) bool {
		if freq[wordFreqs[i]] == freq[wordFreqs[j]] {
			return wordFreqs[i] < wordFreqs[j]
		}
		return freq[wordFreqs[i]] > freq[wordFreqs[j]]
	})

	if len(wordFreqs) > 10 {
		return wordFreqs[:10]
	}
	return wordFreqs
}
