package myrandom

import "github.com/cowabungal/solutions/wordz"

func Digit() string {
	wordz.Words = []string{"1", "2", "3", "4", "5"}
	return wordz.Random()
}
