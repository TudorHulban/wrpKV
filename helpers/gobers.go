package helpers

import (
	"bytes"
	"encoding/gob"
)

// The type of the passed value should be recognizable.
func Encoder(toEncode interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)

	if errEncode := gob.NewEncoder(buf).Encode(toEncode); errEncode != nil {
		return nil, errEncode
	}

	return buf.Bytes(),
		nil
}

// anyDecoder decodes byte slice into the correct type which should be a pointer type.
// The type should be recognizable.
func Decoder(toDecode []byte, decodeInTo interface{}) error {
	return gob.NewDecoder(bytes.NewReader(toDecode)).Decode(decodeInTo)
}
