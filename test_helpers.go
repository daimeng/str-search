package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func load(tnum int, wnum int, fatal func(args ...interface{})) (text string, words []string, results MatchResults) {
	rfile := filepath.Join("fixtures", fmt.Sprintf("results%d-%d.txt", tnum, wnum))
	tfile := filepath.Join("fixtures", fmt.Sprintf("text%d.txt", tnum))
	wfile := filepath.Join("fixtures", fmt.Sprintf("words%d.txt", wnum))

	// parse text file
	b, err := ioutil.ReadFile(tfile)
	if err != nil {
		fatal(err)
	}
	text = strings.ToLower(string(b))

	// parse words file
	b, err = ioutil.ReadFile(wfile)
	if err != nil {
		fatal(err)
	}
	words = []string{}
	for _, r := range strings.Split(string(b), "\n") {
		if r != "" {
			words = append(words, r)
		}
	}

	// parse results file
	b, err = ioutil.ReadFile(rfile)
	if err != nil {
		fatal(err)
	}
	results = MatchResults{}
	for i, r := range strings.Split(string(b), "\n") {
		if r == "" {
			continue
		}

		sa := strings.Split(r, " ")
		ia := make([]int, len(sa))
		for ii, s := range sa {
			v, _ := strconv.Atoi(s)
			ia[ii] = v
		}
		results[words[i]] = ia
	}

	return
}

func testFixture(search func(MatchResults, []string, string), tnum int, wnum int, t *testing.T) {
	matches := MatchResults{}
	text, words, results := load(tnum, wnum, t.Fatal)
	search(matches, words, text)

	if !reflect.DeepEqual(matches, results) {
		t.Errorf("%v\n!=\n%v", matches, results)
	}
}

func benchFixture(search func(MatchResults, []string, string), tnum int, wnum int, b *testing.B) {
	b.StopTimer()
	text, words, _ := load(tnum, wnum, b.Fatal)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		matches := MatchResults{}
		search(matches, words, text)
	}
}
