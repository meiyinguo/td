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

// MessagesEditChatAdminRequest represents TL type `messages.editChatAdmin#a9e69f2e`.
// Make a user admin in a legacy group¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/messages.editChatAdmin for reference.
type MessagesEditChatAdminRequest struct {
	// The ID of the group
	ChatID int
	// The user to make admin
	UserID InputUserClass
	// Whether to make him admin
	IsAdmin bool
}

// MessagesEditChatAdminRequestTypeID is TL type id of MessagesEditChatAdminRequest.
const MessagesEditChatAdminRequestTypeID = 0xa9e69f2e

func (e *MessagesEditChatAdminRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.UserID == nil) {
		return false
	}
	if !(e.IsAdmin == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesEditChatAdminRequest) String() string {
	if e == nil {
		return "MessagesEditChatAdminRequest(nil)"
	}
	type Alias MessagesEditChatAdminRequest
	return fmt.Sprintf("MessagesEditChatAdminRequest%+v", Alias(*e))
}

// FillFrom fills MessagesEditChatAdminRequest from given interface.
func (e *MessagesEditChatAdminRequest) FillFrom(from interface {
	GetChatID() (value int)
	GetUserID() (value InputUserClass)
	GetIsAdmin() (value bool)
}) {
	e.ChatID = from.GetChatID()
	e.UserID = from.GetUserID()
	e.IsAdmin = from.GetIsAdmin()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesEditChatAdminRequest) TypeID() uint32 {
	return MessagesEditChatAdminRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesEditChatAdminRequest) TypeName() string {
	return "messages.editChatAdmin"
}

// TypeInfo returns info about TL type.
func (e *MessagesEditChatAdminRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.editChatAdmin",
		ID:   MessagesEditChatAdminRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "IsAdmin",
			SchemaName: "is_admin",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *MessagesEditChatAdminRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatAdmin#a9e69f2e as nil")
	}
	b.PutID(MessagesEditChatAdminRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *MessagesEditChatAdminRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatAdmin#a9e69f2e as nil")
	}
	b.PutInt(e.ChatID)
	if e.UserID == nil {
		return fmt.Errorf("unable to encode messages.editChatAdmin#a9e69f2e: field user_id is nil")
	}
	if err := e.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editChatAdmin#a9e69f2e: field user_id: %w", err)
	}
	b.PutBool(e.IsAdmin)
	return nil
}

// GetChatID returns value of ChatID field.
func (e *MessagesEditChatAdminRequest) GetChatID() (value int) {
	return e.ChatID
}

// GetUserID returns value of UserID field.
func (e *MessagesEditChatAdminRequest) GetUserID() (value InputUserClass) {
	return e.UserID
}

// GetIsAdmin returns value of IsAdmin field.
func (e *MessagesEditChatAdminRequest) GetIsAdmin() (value bool) {
	return e.IsAdmin
}

// Decode implements bin.Decoder.
func (e *MessagesEditChatAdminRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatAdmin#a9e69f2e to nil")
	}
	if err := b.ConsumeID(MessagesEditChatAdminRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.editChatAdmin#a9e69f2e: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *MessagesEditChatAdminRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatAdmin#a9e69f2e to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatAdmin#a9e69f2e: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatAdmin#a9e69f2e: field user_id: %w", err)
		}
		e.UserID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatAdmin#a9e69f2e: field is_admin: %w", err)
		}
		e.IsAdmin = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesEditChatAdminRequest.
var (
	_ bin.Encoder     = &MessagesEditChatAdminRequest{}
	_ bin.Decoder     = &MessagesEditChatAdminRequest{}
	_ bin.BareEncoder = &MessagesEditChatAdminRequest{}
	_ bin.BareDecoder = &MessagesEditChatAdminRequest{}
)

// MessagesEditChatAdmin invokes method messages.editChatAdmin#a9e69f2e returning error if any.
// Make a user admin in a legacy group¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 USER_ID_INVALID: The provided user ID is invalid
//  400 USER_NOT_PARTICIPANT: You're not a member of this supergroup/channel
//
// See https://core.telegram.org/method/messages.editChatAdmin for reference.
func (c *Client) MessagesEditChatAdmin(ctx context.Context, request *MessagesEditChatAdminRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
