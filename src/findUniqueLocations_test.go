package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var result int

func BenchmarkFindUniqueLocations(b *testing.B) {
	var output int

	for n := 0; n < b.N; n++ {
		output = FindUniqueLocations(strings.Repeat("N", 1<<20) + strings.Repeat("E", 1<<20) + strings.Repeat("S", 1<<20) + strings.Repeat("W", 1<<20))
	}

	result = output
}

func TestFindUniqueLocations(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    string
		expected int
	}{
		{"NESW", 4},
		{"N", 1},
		{"S", 1},
		{"E", 1},
		{"W", 1},
		{"SSEEENEE", 8},
		{"SSNESSN", 5},
		{"EEESS", 5},
		{"ENENNSS", 5},
		{"NESNESN", 5},
		{"ESSSSS", 6},
		{"ESENENNES", 9},
		{"S", 1},
		{"SNNNNE", 6},
		{"SSNEESE", 6},
		{"NENEESNSEEENSSNSS", 12},
		{"EESEEENNSES", 10},
		{"ESNESSSESNSNES", 10},
		{"NNSESEESSSES", 11},
		{"ESEEENNESSSNSENESE", 16},
		{"NSSSEESNES", 9},
		{"SNEEEEESSSS", 11},
		{"SEESSSENNEENEESNS", 15},
		{"NNESESESNSENESEN", 14},
		{"NSSNNNNNESSNNEESSSE", 15},
		{"SSSN", 3},
		{"", 0},
		{"E", 1},
		{"", 0},
		{"ESNSS", 3},
		{"EENNSNNEEENNESEESNE", 16},
		{"NNSSNNSNENE", 6},
		{"SSSNNSSSENSENESENES", 14},
		{"ESENNNNENEEN", 12},
		{"ESESNNENNN", 9},
		{"SSEEENEEESSN", 11},
		{"SSENSNSSNSNSNNENEEN", 10},
		{"EENSESSEENNNNNSEENN", 17},
		{"NSSNSNNSENENNENNNSS", 12},
		{"SEENNSNESNESNENSEN", 13},
		{"EN",2},
{"SSSWWE",5},
{"SEENENWE",7},
{"WN",2},
{"EW",2},
{"NWNENWWENWESNEENS",12},
{"SNNSNEEENWN",9},
{"SESEWWSEENNSWNESN",9},
{"NWSENNWNWNNNES",13},
{"EEEWNSSWEWWWWSESWWE",14},
{"SNSWWSEWSWW",9},
{"EEWNSSSWESWSWWEN",12},
{"ESESWSSESEEWENWESW",13},
{"WSSWSEENNESE",12},
{"EWNWEWNSWWESWWSNSN",11},
{"EEEEEWWWWWWWWWWWWWWWWWWWWSSSSSEEEEEWWWWWSSSSSWWWWWSSSSSSSSSSNNNNNEEEEENNNNNSSSSSWWWWWWWWWWNNNNNNNNNNWWWWWEEEEEWWWWWEEEEENNNNNWWWWWNNNNNEEEEENNNNNSSSSSSSSSSNNNNNEEEEEWWWWWWWWWWNNNNNSSSSSNNNNNSSSSSEEEEESSSSS",119},
	}

	for _, test := range tests {
		assert.Equal(FindUniqueLocations(test.input), test.expected)
	}
}
