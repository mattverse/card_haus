package types

const (
	// ModuleName defines the module name
	ModuleName = "namecard"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_namecard"
)

var (
	KeyNameCard = []byte{0x01}
	// KeyIndexSeparator defines separator between keys when combine, it should be one that is not used in denom expression
	KeyIndexSeparator = []byte{0xFF}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
