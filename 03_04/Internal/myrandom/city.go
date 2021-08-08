package myrandom

import (
	"github.com/cowabungal/solutions/pkg/wordz"
)

func City() string {
	wordz.Words = []string{"Los Angeles", "Moscow", "Ufa", "Toronto", "Berlin"}
	return wordz.Random()
}
