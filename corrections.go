package main

import (
	"sort"
	"strings"
)

type alternative struct {
	word string
	freq int
}

type byFreq []alternative

func (l byFreq) Len() int {
	return len(l)
}

func (l byFreq) Less(i, j int) bool {
	return l[i].freq > l[j].freq
}

func (l byFreq) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Corrections returns the possible corrections ordered by frequency or nil if none.
func Corrections(s string) []string {
	vars := Variations(s)
	sCount := Dict[s]
	alts := []alternative{}
	for _, v := range vars {
		if Dict[v] > sCount {
			alts = append(alts, alternative{word: v, freq: Dict[v]})
		}
	}
	sort.Sort(byFreq(alts))
	ws := []string{}
	for _, a := range alts {
		ws = append(ws, a.word)
	}
	return ws
}

// Suggestion returns the best possible suggestion.
func Suggestion(s string) string {
	if isFrequent(s) {
		return s
	}
	tks := TokenizeLowerCase(s)
	suggestion := make([]string, len(tks))
	for i, t := range tks {
		if len(t) < 3 || isFrequent(t) {
			suggestion[i] = t
		} else {
			cs := Corrections(t)
			if len(cs) > 0 {
				suggestion[i] = cs[0]
			} else {
				suggestion[i] = t
			}
		}
	}
	return strings.Join(suggestion, " ")
}

func isFrequent(token string) bool {
	return Dict[token] >= 3
}
