package types

import (
	"fmt"
)

const (
	// ModuleName defines the module name
	ModuleName = "eightball"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_eightball"

	// Version defines the current version the IBC module supports
	Version = "eightball-1"

	// PortID is the default port id that module binds to
	PortID = "eightball"

	DexTransferChannelKeyPrefix = "dexTransferChannel"
	DexConnectionKeyPrefix      = "dexConnection"

	PacketRequestKeyPrefix = "packetRequest"

	TransferSeqKeyPrefix     = "transferSeq"
	TransferRecvSeqKeyPrefix = "transferRecvSeq"

	ICASeqKeyPrefix = "icaSeq"

	SrcOrigin = "src"
	DstOrigin = "dest"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("eightball-port-")
)

func KeyDexTransferChannel(eightballModuleName string) []byte {
	return []byte(fmt.Sprintf("%s/%s", DexTransferChannelKeyPrefix, eightballModuleName))
}

func KeyDexConnection(eightballModuleName string) []byte {
	return []byte(fmt.Sprintf("%s/%s", DexConnectionKeyPrefix, eightballModuleName))
}

func KeyPacketToRequest(origin, portID, channelID string, sequence uint64) []byte {
	return []byte(fmt.Sprintf("%s/%s/%s/%s/%d", PacketRequestKeyPrefix, origin, portID, channelID, sequence))
}

func KeyTransferSeq(transferSeq uint64) []byte {
	return []byte(fmt.Sprintf("%s/%v", TransferSeqKeyPrefix, transferSeq))
}

func KeyICASeq(icaSeq uint64) []byte {
	return []byte(fmt.Sprintf("%s/%v", ICASeqKeyPrefix, icaSeq))
}

func KeyTransferRecvSeq(transferRecvSeq uint64) []byte {
	return []byte(fmt.Sprintf("%s/%v", TransferRecvSeqKeyPrefix, transferRecvSeq))
}
