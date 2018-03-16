package main

import (
	"testing"
)

func BenchmarkPrefix1_1(b *testing.B) {
	benchFixture(prefixSearch, 1, 1, b)
}

func BenchmarkPrefix1_1P(b *testing.B) {
	benchFixture(prefixSearchP, 1, 1, b)
}

func BenchmarkPrefix2_2(b *testing.B) {
	benchFixture(prefixSearch, 1, 1, b)
}

func BenchmarkPrefix2_2P(b *testing.B) {
	benchFixture(prefixSearchP, 1, 1, b)
}

func TestPrefix1_1(t *testing.T) {
	testFixture(prefixSearch, 1, 1, t)
}

func TestPrefix1_1P(t *testing.T) {
	testFixture(prefixSearchP, 1, 1, t)
}

func TestPrefix2_2(t *testing.T) {
	testFixture(prefixSearch, 2, 2, t)
}

func TestPrefix2_2P(t *testing.T) {
	testFixture(prefixSearchP, 2, 2, t)
}
