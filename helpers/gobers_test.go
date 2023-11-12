package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	resNil, errEncodeNil := Encoder(nil)
	assert.Error(t, errEncodeNil)
	require.Zero(t, resNil)

	resEnc, errEncode := Encoder(v)
	assert.NoError(t, errEncode)
	assert.NotNil(t, resEnc)

	var p tstruct

	assert.NoError(t,
		Decoder(resEnc, &p),
	)
	assert.Equal(t, v, p)
}
