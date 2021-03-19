// Code generated by gotdgen, DO NOT EDIT.

package td

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

// SendRequest represents TL type `send#f74488a`.
//
// See https://localhost:80/doc/method/send for reference.
type SendRequest struct {
	// Msg field of SendRequest.
	Msg SMS
}

// SendRequestTypeID is TL type id of SendRequest.
const SendRequestTypeID = 0xf74488a

func (s *SendRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Msg.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendRequest) String() string {
	if s == nil {
		return "SendRequest(nil)"
	}
	type Alias SendRequest
	return fmt.Sprintf("SendRequest%+v", Alias(*s))
}

// FillFrom fills SendRequest from given interface.
func (s *SendRequest) FillFrom(from interface {
	GetMsg() (value SMS)
}) {
	s.Msg = from.GetMsg()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendRequest) TypeID() uint32 {
	return SendRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendRequest) TypeName() string {
	return "send"
}

// TypeInfo returns info about TL type.
func (s *SendRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "send",
		ID:   SendRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Msg",
			SchemaName: "msg",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode send#f74488a as nil")
	}
	b.PutID(SendRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode send#f74488a as nil")
	}
	if err := s.Msg.Encode(b); err != nil {
		return fmt.Errorf("unable to encode send#f74488a: field msg: %w", err)
	}
	return nil
}

// GetMsg returns value of Msg field.
func (s *SendRequest) GetMsg() (value SMS) {
	return s.Msg
}

// Decode implements bin.Decoder.
func (s *SendRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode send#f74488a to nil")
	}
	if err := b.ConsumeID(SendRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode send#f74488a: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode send#f74488a to nil")
	}
	{
		if err := s.Msg.Decode(b); err != nil {
			return fmt.Errorf("unable to decode send#f74488a: field msg: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for SendRequest.
var (
	_ bin.Encoder     = &SendRequest{}
	_ bin.Decoder     = &SendRequest{}
	_ bin.BareEncoder = &SendRequest{}
	_ bin.BareDecoder = &SendRequest{}
)

// Send invokes method send#f74488a returning error if any.
//
// See https://localhost:80/doc/method/send for reference.
func (c *Client) Send(ctx context.Context, msg SMS) (*SMS, error) {
	var result SMS

	request := &SendRequest{
		Msg: msg,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
