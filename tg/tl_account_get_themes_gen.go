// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// AccountGetThemesRequest represents TL type `account.getThemes#285946f8`.
// Get installed themes
//
// See https://core.telegram.org/method/account.getThemes for reference.
type AccountGetThemesRequest struct {
	// Theme format, a string that identifies the theming engines supported by the client
	Format string
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// AccountGetThemesRequestTypeID is TL type id of AccountGetThemesRequest.
const AccountGetThemesRequestTypeID = 0x285946f8

func (g *AccountGetThemesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Format == "") {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetThemesRequest) String() string {
	if g == nil {
		return "AccountGetThemesRequest(nil)"
	}
	type Alias AccountGetThemesRequest
	return fmt.Sprintf("AccountGetThemesRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetThemesRequest from given interface.
func (g *AccountGetThemesRequest) FillFrom(from interface {
	GetFormat() (value string)
	GetHash() (value int)
}) {
	g.Format = from.GetFormat()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetThemesRequest) TypeID() uint32 {
	return AccountGetThemesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetThemesRequest) TypeName() string {
	return "account.getThemes"
}

// TypeInfo returns info about TL type.
func (g *AccountGetThemesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getThemes",
		ID:   AccountGetThemesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Format",
			SchemaName: "format",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetThemesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getThemes#285946f8 as nil")
	}
	b.PutID(AccountGetThemesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetThemesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getThemes#285946f8 as nil")
	}
	b.PutString(g.Format)
	b.PutInt(g.Hash)
	return nil
}

// GetFormat returns value of Format field.
func (g *AccountGetThemesRequest) GetFormat() (value string) {
	return g.Format
}

// GetHash returns value of Hash field.
func (g *AccountGetThemesRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *AccountGetThemesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getThemes#285946f8 to nil")
	}
	if err := b.ConsumeID(AccountGetThemesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getThemes#285946f8: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetThemesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getThemes#285946f8 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.getThemes#285946f8: field format: %w", err)
		}
		g.Format = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.getThemes#285946f8: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetThemesRequest.
var (
	_ bin.Encoder     = &AccountGetThemesRequest{}
	_ bin.Decoder     = &AccountGetThemesRequest{}
	_ bin.BareEncoder = &AccountGetThemesRequest{}
	_ bin.BareDecoder = &AccountGetThemesRequest{}
)

// AccountGetThemes invokes method account.getThemes#285946f8 returning error if any.
// Get installed themes
//
// See https://core.telegram.org/method/account.getThemes for reference.
func (c *Client) AccountGetThemes(ctx context.Context, request *AccountGetThemesRequest) (AccountThemesClass, error) {
	var result AccountThemesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Themes, nil
}
