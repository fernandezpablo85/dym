package main

import (
	"sort"
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

// Corrections returns the possible corrections ordered by frequency or nil if none
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
