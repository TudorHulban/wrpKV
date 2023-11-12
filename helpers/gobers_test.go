package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tstruct struct {
	F1 int
	F2 []byte
}

func TestGobers(t *testing.T) {
	v := tstruct{
		F1: 1,
		F2: []byte("a"),
	}

	resEnc, errEncode := Encode(v)
	assert.NoError(t, errEncode)
	assert.NotNil(t, resEnc)

	var p tstruct

	assert.NoError(t,
		Decode(resEnc, &p),
	)
	assert.Equal(t, v, p)
}
