// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// MessagesDeleteHistoryRequest represents TL type `messages.deleteHistory#1c015b09`.
// Deletes communication history.
//
// See https://core.telegram.org/method/messages.deleteHistory for reference.
type MessagesDeleteHistoryRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Just clear history for the current user, without actually removing messages for every chat user
	JustClear bool
	// Whether to delete the message history for all chat participants
	Revoke bool
	// User or chat, communication history of which will be deleted
	Peer InputPeerClass
	// Maximum ID of message to delete
	MaxID int
}

// MessagesDeleteHistoryRequestTypeID is TL type id of MessagesDeleteHistoryRequest.
const MessagesDeleteHistoryRequestTypeID = 0x1c015b09

func (d *MessagesDeleteHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.JustClear == false) {
		return false
	}
	if !(d.Revoke == false) {
		return false
	}
	if !(d.Peer == nil) {
		return false
	}
	if !(d.MaxID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDeleteHistoryRequest) String() string {
	if d == nil {
		return "MessagesDeleteHistoryRequest(nil)"
	}
	type Alias MessagesDeleteHistoryRequest
	return fmt.Sprintf("MessagesDeleteHistoryRequest%+v", Alias(*d))
}

// FillFrom fills MessagesDeleteHistoryRequest from given interface.
func (d *MessagesDeleteHistoryRequest) FillFrom(from interface {
	GetJustClear() (value bool)
	GetRevoke() (value bool)
	GetPeer() (value InputPeerClass)
	GetMaxID() (value int)
}) {
	d.Peer = from.GetPeer()
	d.MaxID = from.GetMaxID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *MessagesDeleteHistoryRequest) TypeID() uint32 {
	return MessagesDeleteHistoryRequestTypeID
}

// Encode implements bin.Encoder.
func (d *MessagesDeleteHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteHistory#1c015b09 as nil")
	}
	b.PutID(MessagesDeleteHistoryRequestTypeID)
	if !(d.JustClear == false) {
		d.Flags.Set(0)
	}
	if !(d.Revoke == false) {
		d.Flags.Set(1)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteHistory#1c015b09: field flags: %w", err)
	}
	if d.Peer == nil {
		return fmt.Errorf("unable to encode messages.deleteHistory#1c015b09: field peer is nil")
	}
	if err := d.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteHistory#1c015b09: field peer: %w", err)
	}
	b.PutInt(d.MaxID)
	return nil
}

// SetJustClear sets value of JustClear conditional field.
func (d *MessagesDeleteHistoryRequest) SetJustClear(value bool) {
	if value {
		d.Flags.Set(0)
		d.JustClear = true
	} else {
		d.Flags.Unset(0)
		d.JustClear = false
	}
}

// GetJustClear returns value of JustClear conditional field.
func (d *MessagesDeleteHistoryRequest) GetJustClear() (value bool) {
	return d.Flags.Has(0)
}

// SetRevoke sets value of Revoke conditional field.
func (d *MessagesDeleteHistoryRequest) SetRevoke(value bool) {
	if value {
		d.Flags.Set(1)
		d.Revoke = true
	} else {
		d.Flags.Unset(1)
		d.Revoke = false
	}
}

// GetRevoke returns value of Revoke conditional field.
func (d *MessagesDeleteHistoryRequest) GetRevoke() (value bool) {
	return d.Flags.Has(1)
}

// GetPeer returns value of Peer field.
func (d *MessagesDeleteHistoryRequest) GetPeer() (value InputPeerClass) {
	return d.Peer
}

// GetMaxID returns value of MaxID field.
func (d *MessagesDeleteHistoryRequest) GetMaxID() (value int) {
	return d.MaxID
}

// Decode implements bin.Decoder.
func (d *MessagesDeleteHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteHistory#1c015b09 to nil")
	}
	if err := b.ConsumeID(MessagesDeleteHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deleteHistory#1c015b09: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.deleteHistory#1c015b09: field flags: %w", err)
		}
	}
	d.JustClear = d.Flags.Has(0)
	d.Revoke = d.Flags.Has(1)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteHistory#1c015b09: field peer: %w", err)
		}
		d.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteHistory#1c015b09: field max_id: %w", err)
		}
		d.MaxID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesDeleteHistoryRequest.
var (
	_ bin.Encoder = &MessagesDeleteHistoryRequest{}
	_ bin.Decoder = &MessagesDeleteHistoryRequest{}
)

// MessagesDeleteHistory invokes method messages.deleteHistory#1c015b09 returning error if any.
// Deletes communication history.
//
// Possible errors:
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.deleteHistory for reference.
func (c *Client) MessagesDeleteHistory(ctx context.Context, request *MessagesDeleteHistoryRequest) (*MessagesAffectedHistory, error) {
	var result MessagesAffectedHistory

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
