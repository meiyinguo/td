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

// HelpSetBotUpdatesStatusRequest represents TL type `help.setBotUpdatesStatus#ec22cfcd`.
// Informs the server about the number of pending bot updates if they haven't been processed for a long time; for bots only
//
// See https://core.telegram.org/method/help.setBotUpdatesStatus for reference.
type HelpSetBotUpdatesStatusRequest struct {
	// Number of pending updates
	PendingUpdatesCount int
	// Error message, if present
	Message string
}

// HelpSetBotUpdatesStatusRequestTypeID is TL type id of HelpSetBotUpdatesStatusRequest.
const HelpSetBotUpdatesStatusRequestTypeID = 0xec22cfcd

func (s *HelpSetBotUpdatesStatusRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.PendingUpdatesCount == 0) {
		return false
	}
	if !(s.Message == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *HelpSetBotUpdatesStatusRequest) String() string {
	if s == nil {
		return "HelpSetBotUpdatesStatusRequest(nil)"
	}
	type Alias HelpSetBotUpdatesStatusRequest
	return fmt.Sprintf("HelpSetBotUpdatesStatusRequest%+v", Alias(*s))
}

// FillFrom fills HelpSetBotUpdatesStatusRequest from given interface.
func (s *HelpSetBotUpdatesStatusRequest) FillFrom(from interface {
	GetPendingUpdatesCount() (value int)
	GetMessage() (value string)
}) {
	s.PendingUpdatesCount = from.GetPendingUpdatesCount()
	s.Message = from.GetMessage()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpSetBotUpdatesStatusRequest) TypeID() uint32 {
	return HelpSetBotUpdatesStatusRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpSetBotUpdatesStatusRequest) TypeName() string {
	return "help.setBotUpdatesStatus"
}

// TypeInfo returns info about TL type.
func (s *HelpSetBotUpdatesStatusRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.setBotUpdatesStatus",
		ID:   HelpSetBotUpdatesStatusRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PendingUpdatesCount",
			SchemaName: "pending_updates_count",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *HelpSetBotUpdatesStatusRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode help.setBotUpdatesStatus#ec22cfcd as nil")
	}
	b.PutID(HelpSetBotUpdatesStatusRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *HelpSetBotUpdatesStatusRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode help.setBotUpdatesStatus#ec22cfcd as nil")
	}
	b.PutInt(s.PendingUpdatesCount)
	b.PutString(s.Message)
	return nil
}

// GetPendingUpdatesCount returns value of PendingUpdatesCount field.
func (s *HelpSetBotUpdatesStatusRequest) GetPendingUpdatesCount() (value int) {
	return s.PendingUpdatesCount
}

// GetMessage returns value of Message field.
func (s *HelpSetBotUpdatesStatusRequest) GetMessage() (value string) {
	return s.Message
}

// Decode implements bin.Decoder.
func (s *HelpSetBotUpdatesStatusRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode help.setBotUpdatesStatus#ec22cfcd to nil")
	}
	if err := b.ConsumeID(HelpSetBotUpdatesStatusRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.setBotUpdatesStatus#ec22cfcd: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *HelpSetBotUpdatesStatusRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode help.setBotUpdatesStatus#ec22cfcd to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.setBotUpdatesStatus#ec22cfcd: field pending_updates_count: %w", err)
		}
		s.PendingUpdatesCount = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.setBotUpdatesStatus#ec22cfcd: field message: %w", err)
		}
		s.Message = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpSetBotUpdatesStatusRequest.
var (
	_ bin.Encoder     = &HelpSetBotUpdatesStatusRequest{}
	_ bin.Decoder     = &HelpSetBotUpdatesStatusRequest{}
	_ bin.BareEncoder = &HelpSetBotUpdatesStatusRequest{}
	_ bin.BareDecoder = &HelpSetBotUpdatesStatusRequest{}
)

// HelpSetBotUpdatesStatus invokes method help.setBotUpdatesStatus#ec22cfcd returning error if any.
// Informs the server about the number of pending bot updates if they haven't been processed for a long time; for bots only
//
// See https://core.telegram.org/method/help.setBotUpdatesStatus for reference.
// Can be used by bots.
func (c *Client) HelpSetBotUpdatesStatus(ctx context.Context, request *HelpSetBotUpdatesStatusRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
