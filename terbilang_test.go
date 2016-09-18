package terbilang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTerbilang(t *testing.T) {

	// simple test
	assert.Equal(t, ToTerbilang(109209), "seratus sembilan ribu dua ratus sembilan")
	assert.Equal(t, ToTerbilang(1), "satu")
	assert.Equal(t, ToTerbilang(19), "sembilan belas")
	assert.Equal(t, ToTerbilang(99), "sembilan puluh sembilan")
	assert.Equal(t, ToTerbilang(100), "seratus")
	assert.Equal(t, ToTerbilang(1002), "seribu dua")
	assert.Equal(t, ToTerbilang(1000001), "satu juta satu")
	assert.Equal(t, ToTerbilang(2000055), "dua juta lima puluh lima")
	assert.Equal(t, ToTerbilang(1000000001), "satu milyar satu")
	assert.Equal(t, ToTerbilang(2000000002), "dua milyar dua")
	assert.Equal(t, ToTerbilang(1000000000000), "satu triliun")
	assert.Equal(t, ToTerbilang(2000000000200), "dua triliun dua ratus")

}

func TestTerbilangRp(t *testing.T) {
	assert.Equal(t, ToTerbilangRp(1), "satu rupiah")
}

func TestTerbilangSuff(t *testing.T) {
	assert.Equal(t, ToTerbilangSuffix("buah", 11), "sebelas buah")
}
