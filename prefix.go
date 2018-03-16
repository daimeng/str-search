package main

import (
	"strings"
	"sync"
)

// MatchResults map to indices
type MatchResults map[string][]int

func prefixSearchOne(word string, text string) (matches []int) {
	for i := range text {
		if strings.HasPrefix(text[i:], word) {
			matches = append(matches, i)
		}
	}
	return matches
}

func prefixSearch(matches MatchResults, words []string, text string) {
	for _, word := range words {
		res := prefixSearchOne(word, text)
		if res != nil {
			matches[word] = res
		}
	}
}

var sem = make(chan struct{}, 16)

func prefixSearchP(matches MatchResults, words []string, text string) {
	wg := sync.WaitGroup{}
	for _, word := range words {
		wg.Add(1)
		select {
		case sem <- struct{}{}:
			go func(w string) {
				res := prefixSearchOne(w, text)
				if res != nil {
					matches[w] = res
				}
				<-sem
				wg.Done()
			}(word)
		default:
			res := prefixSearchOne(word, text)
			if res != nil {
				matches[word] = res
			}
			wg.Done()
		}
		wg.Wait()
	}
}
