package keeper

import (
	"bytes"

	"github.com/mattverse/cardhaus/x/namecard/types"
)

func nameCardStoreKey(address string) []byte {
	return combineKeys(types.KeyNameCard, []byte(address))
}

func combineKeys(keys ...[]byte) []byte {
	return bytes.Join(keys, types.KeyIndexSeparator)
}
