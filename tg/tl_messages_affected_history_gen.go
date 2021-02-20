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

// MessagesAffectedHistory represents TL type `messages.affectedHistory#b45c69d1`.
// Affected part of communication history with the user or in a chat.
//
// See https://core.telegram.org/constructor/messages.affectedHistory for reference.
type MessagesAffectedHistory struct {
	// Number of events occured in a text box
	Pts int
	// Number of affected events
	PtsCount int
	// If a parameter contains positive value, it is necessary to repeat the method call using the given value; during the proceeding of all the history the value itself shall gradually decrease
	Offset int
}

// MessagesAffectedHistoryTypeID is TL type id of MessagesAffectedHistory.
const MessagesAffectedHistoryTypeID = 0xb45c69d1

func (a *MessagesAffectedHistory) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Pts == 0) {
		return false
	}
	if !(a.PtsCount == 0) {
		return false
	}
	if !(a.Offset == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *MessagesAffectedHistory) String() string {
	if a == nil {
		return "MessagesAffectedHistory(nil)"
	}
	type Alias MessagesAffectedHistory
	return fmt.Sprintf("MessagesAffectedHistory%+v", Alias(*a))
}

// FillFrom fills MessagesAffectedHistory from given interface.
func (a *MessagesAffectedHistory) FillFrom(from interface {
	GetPts() (value int)
	GetPtsCount() (value int)
	GetOffset() (value int)
}) {
	a.Pts = from.GetPts()
	a.PtsCount = from.GetPtsCount()
	a.Offset = from.GetOffset()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (a *MessagesAffectedHistory) TypeID() uint32 {
	return MessagesAffectedHistoryTypeID
}

// Encode implements bin.Encoder.
func (a *MessagesAffectedHistory) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.affectedHistory#b45c69d1 as nil")
	}
	b.PutID(MessagesAffectedHistoryTypeID)
	b.PutInt(a.Pts)
	b.PutInt(a.PtsCount)
	b.PutInt(a.Offset)
	return nil
}

// GetPts returns value of Pts field.
func (a *MessagesAffectedHistory) GetPts() (value int) {
	return a.Pts
}

// GetPtsCount returns value of PtsCount field.
func (a *MessagesAffectedHistory) GetPtsCount() (value int) {
	return a.PtsCount
}

// GetOffset returns value of Offset field.
func (a *MessagesAffectedHistory) GetOffset() (value int) {
	return a.Offset
}

// Decode implements bin.Decoder.
func (a *MessagesAffectedHistory) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.affectedHistory#b45c69d1 to nil")
	}
	if err := b.ConsumeID(MessagesAffectedHistoryTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.affectedHistory#b45c69d1: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedHistory#b45c69d1: field pts: %w", err)
		}
		a.Pts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedHistory#b45c69d1: field pts_count: %w", err)
		}
		a.PtsCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedHistory#b45c69d1: field offset: %w", err)
		}
		a.Offset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesAffectedHistory.
var (
	_ bin.Encoder = &MessagesAffectedHistory{}
	_ bin.Decoder = &MessagesAffectedHistory{}
)
