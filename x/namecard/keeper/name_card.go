package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/mattverse/cardhaus/x/namecard/types"
)

func (k Keeper) GetNameCardByAddress(ctx sdk.Context, address string) (*types.NameCard, error) {
	nameCard := types.NameCard{}
	store := ctx.KVStore(k.storeKey)
	nameCardKey := nameCardStoreKey(address)
	if !store.Has(nameCardKey) {
		return nil, fmt.Errorf("name card with address %s does not exist", address)
	}
	bz := store.Get(nameCardKey)
	if err := proto.Unmarshal(bz, &nameCard); err != nil {
		return nil, err
	}
	return &nameCard, nil
}

func (k Keeper) SetNameCard(ctx sdk.Context, nameCard *types.NameCard) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := proto.Marshal(nameCard)
	if err != nil {
		return err
	}
	store.Set(nameCardStoreKey(nameCard.Address), bz)
	return nil
}
