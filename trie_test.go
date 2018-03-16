package main

import (
	"reflect"
	"testing"
)

func testFixtureTrie(tnum int, wnum int, t *testing.T) {
	text, words, results := load(tnum, wnum, t.Fatal)

	trie := NewTrie()
	for _, word := range words {
		trie.insert([]rune(word), word)
	}

	matches := trie.match(text)

	if !reflect.DeepEqual(matches, results) {
		t.Errorf("%v\n!=\n%v", matches, results)
	}
}

func benchFixtureTrie(tnum int, wnum int, b *testing.B) {
	b.StopTimer()
	text, words, _ := load(tnum, wnum, b.Fatal)
	b.StartTimer()

	trie := NewTrie()
	for _, word := range words {
		trie.insert([]rune(word), word)
	}

	for i := 0; i < b.N; i++ {
		trie.match(text)
	}
}

func TestTrie1_1(t *testing.T) {
	testFixtureTrie(1, 1, t)
}

func TestTrie2_2(t *testing.T) {
	testFixtureTrie(2, 2, t)
}

func BenchmarkTrie1_1(b *testing.B) {
	benchFixtureTrie(1, 1, b)
}

func BenchmarkTrie2_2(b *testing.B) {
	benchFixtureTrie(2, 2, b)
}
