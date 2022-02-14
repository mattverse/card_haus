package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateNamecard = "create_namecard"

var _ sdk.Msg = &MsgCreateNamecard{}

func NewMsgCreateNamecard(creator string, address string, name string, imageUrl string, selfIntro string) *MsgCreateNamecard {
	return &MsgCreateNamecard{
		Address:   address,
		Name:      name,
		ImageUrl:  imageUrl,
		SelfIntro: selfIntro,
	}
}

func (msg *MsgCreateNamecard) Route() string {
	return RouterKey
}

func (msg *MsgCreateNamecard) Type() string {
	return TypeMsgCreateNamecard
}

func (msg *MsgCreateNamecard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNamecard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNamecard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
