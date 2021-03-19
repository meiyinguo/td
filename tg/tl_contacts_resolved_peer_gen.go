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

// ContactsResolvedPeer represents TL type `contacts.resolvedPeer#7f077ad9`.
// Resolved peer
//
// See https://core.telegram.org/constructor/contacts.resolvedPeer for reference.
type ContactsResolvedPeer struct {
	// The peer
	Peer PeerClass
	// Chats
	Chats []ChatClass
	// Users
	Users []UserClass
}

// ContactsResolvedPeerTypeID is TL type id of ContactsResolvedPeer.
const ContactsResolvedPeerTypeID = 0x7f077ad9

func (r *ContactsResolvedPeer) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.Chats == nil) {
		return false
	}
	if !(r.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ContactsResolvedPeer) String() string {
	if r == nil {
		return "ContactsResolvedPeer(nil)"
	}
	type Alias ContactsResolvedPeer
	return fmt.Sprintf("ContactsResolvedPeer%+v", Alias(*r))
}

// FillFrom fills ContactsResolvedPeer from given interface.
func (r *ContactsResolvedPeer) FillFrom(from interface {
	GetPeer() (value PeerClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	r.Peer = from.GetPeer()
	r.Chats = from.GetChats()
	r.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsResolvedPeer) TypeID() uint32 {
	return ContactsResolvedPeerTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsResolvedPeer) TypeName() string {
	return "contacts.resolvedPeer"
}

// TypeInfo returns info about TL type.
func (r *ContactsResolvedPeer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.resolvedPeer",
		ID:   ContactsResolvedPeerTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ContactsResolvedPeer) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode contacts.resolvedPeer#7f077ad9 as nil")
	}
	b.PutID(ContactsResolvedPeerTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ContactsResolvedPeer) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode contacts.resolvedPeer#7f077ad9 as nil")
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode contacts.resolvedPeer#7f077ad9: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.resolvedPeer#7f077ad9: field peer: %w", err)
	}
	b.PutVectorHeader(len(r.Chats))
	for idx, v := range r.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.resolvedPeer#7f077ad9: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.resolvedPeer#7f077ad9: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(r.Users))
	for idx, v := range r.Users {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.resolvedPeer#7f077ad9: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.resolvedPeer#7f077ad9: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *ContactsResolvedPeer) GetPeer() (value PeerClass) {
	return r.Peer
}

// GetChats returns value of Chats field.
func (r *ContactsResolvedPeer) GetChats() (value []ChatClass) {
	return r.Chats
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (r *ContactsResolvedPeer) MapChats() (value ChatClassArray) {
	return ChatClassArray(r.Chats)
}

// GetUsers returns value of Users field.
func (r *ContactsResolvedPeer) GetUsers() (value []UserClass) {
	return r.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (r *ContactsResolvedPeer) MapUsers() (value UserClassArray) {
	return UserClassArray(r.Users)
}

// Decode implements bin.Decoder.
func (r *ContactsResolvedPeer) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode contacts.resolvedPeer#7f077ad9 to nil")
	}
	if err := b.ConsumeID(ContactsResolvedPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.resolvedPeer#7f077ad9: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ContactsResolvedPeer) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode contacts.resolvedPeer#7f077ad9 to nil")
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.resolvedPeer#7f077ad9: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.resolvedPeer#7f077ad9: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.resolvedPeer#7f077ad9: field chats: %w", err)
			}
			r.Chats = append(r.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.resolvedPeer#7f077ad9: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.resolvedPeer#7f077ad9: field users: %w", err)
			}
			r.Users = append(r.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsResolvedPeer.
var (
	_ bin.Encoder     = &ContactsResolvedPeer{}
	_ bin.Decoder     = &ContactsResolvedPeer{}
	_ bin.BareEncoder = &ContactsResolvedPeer{}
	_ bin.BareDecoder = &ContactsResolvedPeer{}
)
