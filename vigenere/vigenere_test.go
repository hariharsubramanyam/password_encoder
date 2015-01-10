package vigenere

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccessfulEncode(t *testing.T) {
	actual, err := Encode(1234, "password")
	assert.Nil(t, err)
	assert.Equal(t, "qcvwxquh", actual, "Encoding doesn't work.")
}

func TestInvalidMessage(t *testing.T) {
	_, err := Encode(1234, "PASSWORD")
	assert.NotNil(t, err)
}

func TestLongKey(t *testing.T) {
	actual, err := Encode(1111, "pa")
	assert.Nil(t, err)
	assert.Equal(t, "qb", actual, "Long key encoding doesn't work.")
}

func TestInvalidCharacter(t *testing.T) {
	_, err := Encode(1111, "p!")
	assert.NotNil(t, err)
}
