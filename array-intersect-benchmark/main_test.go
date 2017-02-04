package main

import (
	"math/rand"
	"testing"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var (
	original         = []string{}
	target           = []string{}
	targetMap        = map[string]bool{}
	targetNoMatch    = []string{}
	targetMapNoMatch = map[string]bool{}
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	numItems := 5000
	for i := 0; i < numItems; i++ {
		original = append(original, randSeq(10))
	}

	i := rand.Intn(numItems)
	if i >= 4500 {
		i = 4499
	}
	stop := i + 500
	for ; i < stop; i++ {
		target = append(target, original[i])
		targetMap[original[i]] = true
		noMatch := randSeq(9)
		targetNoMatch = append(target, noMatch)
		targetMapNoMatch[noMatch] = true
	}

}

func ON(orig []string, tgt map[string]bool) bool {
	for _, val := range orig {
		if tgt[val] {
			return true
		}
	}
	return false
}

func ON2(orig, tgt []string) bool {
	for _, i := range orig {
		for _, x := range tgt {
			if i == x {
				return true
			}
		}
	}
	return false
}

func BenchmarkMapLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ON(original, targetMap)
	}
}

func BenchmarkNestedRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ON2(original, target)
	}
}

func BenchmarkMapLookupNoMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ON(original, targetMapNoMatch)
	}
}

func BenchmarkNestRangeNoMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ON2(original, targetNoMatch)
	}
}
