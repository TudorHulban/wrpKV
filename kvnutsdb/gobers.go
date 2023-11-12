package kvnuts

import (
	"bytes"
	"encoding/gob"
)

// Encode encodes passed value into byte slice.
// The type of the passed value should be recognizable.
func Encode[T any](toEncode T) ([]byte, error) {
	var buf bytes.Buffer

	if errEncode := gob.NewEncoder(&buf).Encode(toEncode); errEncode != nil {
		return []byte{}, errEncode
	}

	return buf.Bytes(), nil
}

// Decode decodes byte slice into the correct type which should be a pointer type.
// The type should be recognizable.
func Decode[T any](toDecode []byte, decodeInTo T) error {
	return gob.NewDecoder(bytes.NewReader(toDecode)).Decode(decodeInTo)
}
