package types

const (
	// ModuleName defines the module name
	ModuleName = "igniteexample"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_igniteexample"
)

var (
	ParamsKey = []byte("p_igniteexample")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
