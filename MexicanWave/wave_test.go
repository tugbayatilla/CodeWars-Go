package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaveEmpty(t *testing.T) {
	assert.Equal(t, []string{}, wave(""))
}

func TestSingleLetter(t *testing.T) {
	assert.Equal(t, []string{"A"}, wave("a"))
}

func TestTwoCharacters(t *testing.T) {
	assert.Equal(t, []string{"Ab", "aB"}, wave("ab"))
}

func TestThreeCharacters(t *testing.T) {
	assert.Equal(t, []string{"Abc", "aBc", "abC"}, wave("abc"))
}

func TestFiveCharacters(t *testing.T) {
	assert.Equal(t, []string{"Hello", "hEllo", "heLlo", "helLo", "hellO"}, wave("hello"))
}

func TestSixCharactersIncludingSpace(t *testing.T) {
	assert.Equal(t, []string{"Hell o", "hEll o", "heLl o", "helL o", "hell O"}, wave("hell o"))
}

func BenchmarkWave(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wave("lorem ipsum")
	}
}
