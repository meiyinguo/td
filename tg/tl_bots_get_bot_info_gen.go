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
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// BotsGetBotInfoRequest represents TL type `bots.getBotInfo#75ec12e6`.
//
// See https://core.telegram.org/method/bots.getBotInfo for reference.
type BotsGetBotInfoRequest struct {
	// LangCode field of BotsGetBotInfoRequest.
	LangCode string
}

// BotsGetBotInfoRequestTypeID is TL type id of BotsGetBotInfoRequest.
const BotsGetBotInfoRequestTypeID = 0x75ec12e6

// Ensuring interfaces in compile-time for BotsGetBotInfoRequest.
var (
	_ bin.Encoder     = &BotsGetBotInfoRequest{}
	_ bin.Decoder     = &BotsGetBotInfoRequest{}
	_ bin.BareEncoder = &BotsGetBotInfoRequest{}
	_ bin.BareDecoder = &BotsGetBotInfoRequest{}
)

func (g *BotsGetBotInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *BotsGetBotInfoRequest) String() string {
	if g == nil {
		return "BotsGetBotInfoRequest(nil)"
	}
	type Alias BotsGetBotInfoRequest
	return fmt.Sprintf("BotsGetBotInfoRequest%+v", Alias(*g))
}

// FillFrom fills BotsGetBotInfoRequest from given interface.
func (g *BotsGetBotInfoRequest) FillFrom(from interface {
	GetLangCode() (value string)
}) {
	g.LangCode = from.GetLangCode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsGetBotInfoRequest) TypeID() uint32 {
	return BotsGetBotInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsGetBotInfoRequest) TypeName() string {
	return "bots.getBotInfo"
}

// TypeInfo returns info about TL type.
func (g *BotsGetBotInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.getBotInfo",
		ID:   BotsGetBotInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *BotsGetBotInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode bots.getBotInfo#75ec12e6 as nil")
	}
	b.PutID(BotsGetBotInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *BotsGetBotInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode bots.getBotInfo#75ec12e6 as nil")
	}
	b.PutString(g.LangCode)
	return nil
}

// Decode implements bin.Decoder.
func (g *BotsGetBotInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode bots.getBotInfo#75ec12e6 to nil")
	}
	if err := b.ConsumeID(BotsGetBotInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.getBotInfo#75ec12e6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *BotsGetBotInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode bots.getBotInfo#75ec12e6 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.getBotInfo#75ec12e6: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	return nil
}

// GetLangCode returns value of LangCode field.
func (g *BotsGetBotInfoRequest) GetLangCode() (value string) {
	if g == nil {
		return
	}
	return g.LangCode
}

// BotsGetBotInfo invokes method bots.getBotInfo#75ec12e6 returning error if any.
//
// See https://core.telegram.org/method/bots.getBotInfo for reference.
func (c *Client) BotsGetBotInfo(ctx context.Context, langcode string) ([]string, error) {
	var result StringVector

	request := &BotsGetBotInfoRequest{
		LangCode: langcode,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []string(result.Elems), nil
}
