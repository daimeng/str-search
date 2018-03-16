package main

// Trie type
type Trie struct {
	next  map[rune]*Trie
	fail  *Trie
	entry string
}

// NewTrie creates new Trie
func NewTrie() *Trie {
	return &Trie{
		next: map[rune]*Trie{},
	}
}

func (t *Trie) insert(key []rune, entry string) {
	if len(key) == 0 {
		t.entry = entry
		return
	}

	r, rest := key[0], key[1:]
	if _, ok := t.next[r]; !ok {
		t.next[r] = NewTrie()
	}
	t.next[r].insert(rest, entry)
}

func (t *Trie) match(text string) MatchResults {
	matches := MatchResults{}

	for i := range text {
		p := t
		for _, c := range text[i:] {
			if v, ok := p.next[c]; ok {
				p = v
				if p.entry != "" {
					matches[p.entry] = append(matches[p.entry], i)
				}
			} else {
				break
			}
		}
	}

	return matches
}
