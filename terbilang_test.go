package terbilang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTerbilang(t *testing.T) {

	// simple test
	assert.Equal(t, ToTerbilang(109209), "seratus sembilan ribu dua ratus sembilan")
	assert.Equal(t, ToTerbilang(1), "satu")
	assert.Equal(t, ToTerbilang(100), "seratus")
	assert.Equal(t, ToTerbilang(1000001), "satu juta satu")
}

func TestTerbilangRp(t *testing.T) {
	assert.Equal(t, ToTerbilangRp(1), "satu rupiah")
}

func TestTerbilangSuff(t *testing.T) {
	assert.Equal(t, ToTerbilangSuffix("buah", 11), "sebelas buah")
}
