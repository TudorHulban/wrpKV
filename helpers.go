package badgerwrap

import (
	"bytes"
	"encoding/gob"
)

// anyEncoder encodes passed value into byte slice. The type of the passed value should be recognizable.
func anyEncoder(toEncode interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	if errEncode := gob.NewEncoder(buffer).Encode(toEncode); errEncode != nil {
		return []byte{}, errEncode
	}
	return buffer.Bytes(), nil
}

// anyDecoder decodes byte slice into the correct type which should be a pointer type. The type should be recognizable.
func anyDecoder(toDecode []byte, decodeInTo interface{}) error {
	return gob.NewDecoder(bytes.NewReader(toDecode)).Decode(decodeInTo)
}
