package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FortuneKeyPrefix is the prefix to retrieve Fortunes
	FortuneKeyPrefix    = "Fortune/value/"
	AllFortuneKeyString = "AllFortunes"
)

// FortuneKey returns the store key to retrieve a Fortune from the index fields
func FortuneKey(
	owner string,
) []byte {
	key := []byte("ownedFortune/")

	ownerBytes := []byte(owner)
	key = append(key, ownerBytes...)

	return key
}

func AllFortuneKey() []byte {
	return []byte(AllFortuneKeyString)
}
