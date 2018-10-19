package main

type pair struct {
	first  string
	second string
}

// Variations returns all 1 edit variations for the given string
func Variations(s string) []string {
	variations := []string{}
	pairs := pairs(s)
	variations = append(variations, deletions(pairs)...)
	variations = append(variations, transposes(pairs)...)
	variations = append(variations, replaces(pairs)...)
	return variations
}

func pairs(s string) []pair {
	ps := []pair{}
	for i := 0; i < len(s)+1; i++ {
		pair := pair{first: s[:i], second: s[i:]}
		ps = append(ps, pair)
	}
	return ps
}

func deletions(pairs []pair) []string {
	deletions := []string{}
	for _, p := range pairs {
		if len(p.second) == 0 {
			continue
		}
		del := p.first + p.second[1:]
		deletions = append(deletions, del)
	}
	return deletions
}

func transposes(pairs []pair) []string {
	ts := []string{}
	for _, p := range pairs {
		if len(p.second) <= 1 {
			continue
		}
		t := p.first + string(p.second[1]) + string(p.second[0]) + p.second[2:]
		ts = append(ts, t)
	}
	return ts
}

func replaces(pairs []pair) []string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	rs := []string{}
	for _, p := range pairs {
		for _, l := range letters {
			i := p.first + string(l) + p.second
			rs = append(rs, i)
			if len(p.second) > 0 {
				r := p.first + string(l) + p.second[1:]
				rs = append(rs, r)
			}
		}
	}
	return rs
}
