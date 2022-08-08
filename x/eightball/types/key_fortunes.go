package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FortunesKeyPrefix is the prefix to retrieve all Fortunes
	FortunesKeyPrefix = "Fortunes/value/"
)

// FortunesKey returns the store key to retrieve a Fortunes from the index fields
func FortunesKey(
	owner string,
) []byte {
	var key []byte

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)
	key = append(key, []byte("/")...)

	return key
}
