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

// GroupCallDiscarded represents TL type `groupCallDiscarded#7780bcb4`.
//
// See https://core.telegram.org/constructor/groupCallDiscarded for reference.
type GroupCallDiscarded struct {
	// ID field of GroupCallDiscarded.
	ID int64
	// AccessHash field of GroupCallDiscarded.
	AccessHash int64
	// Duration field of GroupCallDiscarded.
	Duration int
}

// GroupCallDiscardedTypeID is TL type id of GroupCallDiscarded.
const GroupCallDiscardedTypeID = 0x7780bcb4

func (g *GroupCallDiscarded) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ID == 0) {
		return false
	}
	if !(g.AccessHash == 0) {
		return false
	}
	if !(g.Duration == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GroupCallDiscarded) String() string {
	if g == nil {
		return "GroupCallDiscarded(nil)"
	}
	type Alias GroupCallDiscarded
	return fmt.Sprintf("GroupCallDiscarded%+v", Alias(*g))
}

// FillFrom fills GroupCallDiscarded from given interface.
func (g *GroupCallDiscarded) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetDuration() (value int)
}) {
	g.ID = from.GetID()
	g.AccessHash = from.GetAccessHash()
	g.Duration = from.GetDuration()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *GroupCallDiscarded) TypeID() uint32 {
	return GroupCallDiscardedTypeID
}

// Encode implements bin.Encoder.
func (g *GroupCallDiscarded) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallDiscarded#7780bcb4 as nil")
	}
	b.PutID(GroupCallDiscardedTypeID)
	b.PutLong(g.ID)
	b.PutLong(g.AccessHash)
	b.PutInt(g.Duration)
	return nil
}

// GetID returns value of ID field.
func (g *GroupCallDiscarded) GetID() (value int64) {
	return g.ID
}

// GetAccessHash returns value of AccessHash field.
func (g *GroupCallDiscarded) GetAccessHash() (value int64) {
	return g.AccessHash
}

// GetDuration returns value of Duration field.
func (g *GroupCallDiscarded) GetDuration() (value int) {
	return g.Duration
}

// Decode implements bin.Decoder.
func (g *GroupCallDiscarded) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallDiscarded#7780bcb4 to nil")
	}
	if err := b.ConsumeID(GroupCallDiscardedTypeID); err != nil {
		return fmt.Errorf("unable to decode groupCallDiscarded#7780bcb4: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallDiscarded#7780bcb4: field id: %w", err)
		}
		g.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallDiscarded#7780bcb4: field access_hash: %w", err)
		}
		g.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallDiscarded#7780bcb4: field duration: %w", err)
		}
		g.Duration = value
	}
	return nil
}

// construct implements constructor of GroupCallClass.
func (g GroupCallDiscarded) construct() GroupCallClass { return &g }

// Ensuring interfaces in compile-time for GroupCallDiscarded.
var (
	_ bin.Encoder = &GroupCallDiscarded{}
	_ bin.Decoder = &GroupCallDiscarded{}

	_ GroupCallClass = &GroupCallDiscarded{}
)

// GroupCall represents TL type `groupCall#55903081`.
//
// See https://core.telegram.org/constructor/groupCall for reference.
type GroupCall struct {
	// Flags field of GroupCall.
	Flags bin.Fields
	// JoinMuted field of GroupCall.
	JoinMuted bool
	// CanChangeJoinMuted field of GroupCall.
	CanChangeJoinMuted bool
	// ID field of GroupCall.
	ID int64
	// AccessHash field of GroupCall.
	AccessHash int64
	// ParticipantsCount field of GroupCall.
	ParticipantsCount int
	// Params field of GroupCall.
	//
	// Use SetParams and GetParams helpers.
	Params DataJSON
	// Version field of GroupCall.
	Version int
}

// GroupCallTypeID is TL type id of GroupCall.
const GroupCallTypeID = 0x55903081

func (g *GroupCall) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.JoinMuted == false) {
		return false
	}
	if !(g.CanChangeJoinMuted == false) {
		return false
	}
	if !(g.ID == 0) {
		return false
	}
	if !(g.AccessHash == 0) {
		return false
	}
	if !(g.ParticipantsCount == 0) {
		return false
	}
	if !(g.Params.Zero()) {
		return false
	}
	if !(g.Version == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GroupCall) String() string {
	if g == nil {
		return "GroupCall(nil)"
	}
	type Alias GroupCall
	return fmt.Sprintf("GroupCall%+v", Alias(*g))
}

// FillFrom fills GroupCall from given interface.
func (g *GroupCall) FillFrom(from interface {
	GetJoinMuted() (value bool)
	GetCanChangeJoinMuted() (value bool)
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetParticipantsCount() (value int)
	GetParams() (value DataJSON, ok bool)
	GetVersion() (value int)
}) {
	g.ID = from.GetID()
	g.AccessHash = from.GetAccessHash()
	g.ParticipantsCount = from.GetParticipantsCount()
	if val, ok := from.GetParams(); ok {
		g.Params = val
	}
	g.Version = from.GetVersion()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *GroupCall) TypeID() uint32 {
	return GroupCallTypeID
}

// Encode implements bin.Encoder.
func (g *GroupCall) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCall#55903081 as nil")
	}
	b.PutID(GroupCallTypeID)
	if !(g.JoinMuted == false) {
		g.Flags.Set(1)
	}
	if !(g.CanChangeJoinMuted == false) {
		g.Flags.Set(2)
	}
	if !(g.Params.Zero()) {
		g.Flags.Set(0)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode groupCall#55903081: field flags: %w", err)
	}
	b.PutLong(g.ID)
	b.PutLong(g.AccessHash)
	b.PutInt(g.ParticipantsCount)
	if g.Flags.Has(0) {
		if err := g.Params.Encode(b); err != nil {
			return fmt.Errorf("unable to encode groupCall#55903081: field params: %w", err)
		}
	}
	b.PutInt(g.Version)
	return nil
}

// SetJoinMuted sets value of JoinMuted conditional field.
func (g *GroupCall) SetJoinMuted(value bool) {
	if value {
		g.Flags.Set(1)
		g.JoinMuted = true
	} else {
		g.Flags.Unset(1)
		g.JoinMuted = false
	}
}

// GetJoinMuted returns value of JoinMuted conditional field.
func (g *GroupCall) GetJoinMuted() (value bool) {
	return g.Flags.Has(1)
}

// SetCanChangeJoinMuted sets value of CanChangeJoinMuted conditional field.
func (g *GroupCall) SetCanChangeJoinMuted(value bool) {
	if value {
		g.Flags.Set(2)
		g.CanChangeJoinMuted = true
	} else {
		g.Flags.Unset(2)
		g.CanChangeJoinMuted = false
	}
}

// GetCanChangeJoinMuted returns value of CanChangeJoinMuted conditional field.
func (g *GroupCall) GetCanChangeJoinMuted() (value bool) {
	return g.Flags.Has(2)
}

// GetID returns value of ID field.
func (g *GroupCall) GetID() (value int64) {
	return g.ID
}

// GetAccessHash returns value of AccessHash field.
func (g *GroupCall) GetAccessHash() (value int64) {
	return g.AccessHash
}

// GetParticipantsCount returns value of ParticipantsCount field.
func (g *GroupCall) GetParticipantsCount() (value int) {
	return g.ParticipantsCount
}

// SetParams sets value of Params conditional field.
func (g *GroupCall) SetParams(value DataJSON) {
	g.Flags.Set(0)
	g.Params = value
}

// GetParams returns value of Params conditional field and
// boolean which is true if field was set.
func (g *GroupCall) GetParams() (value DataJSON, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Params, true
}

// GetVersion returns value of Version field.
func (g *GroupCall) GetVersion() (value int) {
	return g.Version
}

// Decode implements bin.Decoder.
func (g *GroupCall) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCall#55903081 to nil")
	}
	if err := b.ConsumeID(GroupCallTypeID); err != nil {
		return fmt.Errorf("unable to decode groupCall#55903081: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode groupCall#55903081: field flags: %w", err)
		}
	}
	g.JoinMuted = g.Flags.Has(1)
	g.CanChangeJoinMuted = g.Flags.Has(2)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#55903081: field id: %w", err)
		}
		g.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#55903081: field access_hash: %w", err)
		}
		g.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#55903081: field participants_count: %w", err)
		}
		g.ParticipantsCount = value
	}
	if g.Flags.Has(0) {
		if err := g.Params.Decode(b); err != nil {
			return fmt.Errorf("unable to decode groupCall#55903081: field params: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#55903081: field version: %w", err)
		}
		g.Version = value
	}
	return nil
}

// construct implements constructor of GroupCallClass.
func (g GroupCall) construct() GroupCallClass { return &g }

// Ensuring interfaces in compile-time for GroupCall.
var (
	_ bin.Encoder = &GroupCall{}
	_ bin.Decoder = &GroupCall{}

	_ GroupCallClass = &GroupCall{}
)

// GroupCallClass represents GroupCall generic type.
//
// See https://core.telegram.org/type/GroupCall for reference.
//
// Example:
//  g, err := tg.DecodeGroupCall(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.GroupCallDiscarded: // groupCallDiscarded#7780bcb4
//  case *tg.GroupCall: // groupCall#55903081
//  default: panic(v)
//  }
type GroupCallClass interface {
	bin.Encoder
	bin.Decoder
	construct() GroupCallClass

	// ID field of GroupCallDiscarded.
	GetID() (value int64)
	// AccessHash field of GroupCallDiscarded.
	GetAccessHash() (value int64)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeGroupCall implements binary de-serialization for GroupCallClass.
func DecodeGroupCall(buf *bin.Buffer) (GroupCallClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case GroupCallDiscardedTypeID:
		// Decoding groupCallDiscarded#7780bcb4.
		v := GroupCallDiscarded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode GroupCallClass: %w", err)
		}
		return &v, nil
	case GroupCallTypeID:
		// Decoding groupCall#55903081.
		v := GroupCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode GroupCallClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode GroupCallClass: %w", bin.NewUnexpectedID(id))
	}
}

// GroupCall boxes the GroupCallClass providing a helper.
type GroupCallBox struct {
	GroupCall GroupCallClass
}

// Decode implements bin.Decoder for GroupCallBox.
func (b *GroupCallBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode GroupCallBox to nil")
	}
	v, err := DecodeGroupCall(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.GroupCall = v
	return nil
}

// Encode implements bin.Encode for GroupCallBox.
func (b *GroupCallBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.GroupCall == nil {
		return fmt.Errorf("unable to encode GroupCallClass as nil")
	}
	return b.GroupCall.Encode(buf)
}

// GroupCallClassSlice is adapter for slice of GroupCallClass.
type GroupCallClassSlice []GroupCallClass

// First returns first element of slice (if exists).
func (s GroupCallClassSlice) First() (v GroupCallClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s GroupCallClassSlice) Last() (v GroupCallClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *GroupCallClassSlice) PopFirst() (v GroupCallClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	a[len(a)-1] = nil
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *GroupCallClassSlice) Pop() (v GroupCallClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
