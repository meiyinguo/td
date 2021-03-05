// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is
var _ = sort.Ints
var _ = tdp.Format

// ReplyKeyboardHide represents TL type `replyKeyboardHide#a03e5b85`.
// Hide sent bot keyboard
//
// See https://core.telegram.org/constructor/replyKeyboardHide for reference.
type ReplyKeyboardHide struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Use this flag if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet
	Selective bool
}

// ReplyKeyboardHideTypeID is TL type id of ReplyKeyboardHide.
const ReplyKeyboardHideTypeID = 0xa03e5b85

func (r *ReplyKeyboardHide) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Selective == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyKeyboardHide) String() string {
	if r == nil {
		return "ReplyKeyboardHide(nil)"
	}
	type Alias ReplyKeyboardHide
	return fmt.Sprintf("ReplyKeyboardHide%+v", Alias(*r))
}

// FillFrom fills ReplyKeyboardHide from given interface.
func (r *ReplyKeyboardHide) FillFrom(from interface {
	GetSelective() (value bool)
}) {
	r.Selective = from.GetSelective()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplyKeyboardHide) TypeID() uint32 {
	return ReplyKeyboardHideTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplyKeyboardHide) TypeName() string {
	return "replyKeyboardHide"
}

// TypeInfo returns info about TL type.
func (r *ReplyKeyboardHide) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replyKeyboardHide",
		ID:   ReplyKeyboardHideTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Selective",
			SchemaName: "selective",
			Null:       !r.Flags.Has(2),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReplyKeyboardHide) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyKeyboardHide#a03e5b85 as nil")
	}
	b.PutID(ReplyKeyboardHideTypeID)
	if !(r.Selective == false) {
		r.Flags.Set(2)
	}
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode replyKeyboardHide#a03e5b85: field flags: %w", err)
	}
	return nil
}

// SetSelective sets value of Selective conditional field.
func (r *ReplyKeyboardHide) SetSelective(value bool) {
	if value {
		r.Flags.Set(2)
		r.Selective = true
	} else {
		r.Flags.Unset(2)
		r.Selective = false
	}
}

// GetSelective returns value of Selective conditional field.
func (r *ReplyKeyboardHide) GetSelective() (value bool) {
	return r.Flags.Has(2)
}

// Decode implements bin.Decoder.
func (r *ReplyKeyboardHide) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyKeyboardHide#a03e5b85 to nil")
	}
	if err := b.ConsumeID(ReplyKeyboardHideTypeID); err != nil {
		return fmt.Errorf("unable to decode replyKeyboardHide#a03e5b85: %w", err)
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode replyKeyboardHide#a03e5b85: field flags: %w", err)
		}
	}
	r.Selective = r.Flags.Has(2)
	return nil
}

// construct implements constructor of ReplyMarkupClass.
func (r ReplyKeyboardHide) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyKeyboardHide.
var (
	_ bin.Encoder = &ReplyKeyboardHide{}
	_ bin.Decoder = &ReplyKeyboardHide{}

	_ ReplyMarkupClass = &ReplyKeyboardHide{}
)

// ReplyKeyboardForceReply represents TL type `replyKeyboardForceReply#f4108aa0`.
// Force the user to send a reply
//
// See https://core.telegram.org/constructor/replyKeyboardForceReply for reference.
type ReplyKeyboardForceReply struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again.
	SingleUse bool
	// Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message. Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard.
	Selective bool
}

// ReplyKeyboardForceReplyTypeID is TL type id of ReplyKeyboardForceReply.
const ReplyKeyboardForceReplyTypeID = 0xf4108aa0

func (r *ReplyKeyboardForceReply) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.SingleUse == false) {
		return false
	}
	if !(r.Selective == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyKeyboardForceReply) String() string {
	if r == nil {
		return "ReplyKeyboardForceReply(nil)"
	}
	type Alias ReplyKeyboardForceReply
	return fmt.Sprintf("ReplyKeyboardForceReply%+v", Alias(*r))
}

// FillFrom fills ReplyKeyboardForceReply from given interface.
func (r *ReplyKeyboardForceReply) FillFrom(from interface {
	GetSingleUse() (value bool)
	GetSelective() (value bool)
}) {
	r.SingleUse = from.GetSingleUse()
	r.Selective = from.GetSelective()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplyKeyboardForceReply) TypeID() uint32 {
	return ReplyKeyboardForceReplyTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplyKeyboardForceReply) TypeName() string {
	return "replyKeyboardForceReply"
}

// TypeInfo returns info about TL type.
func (r *ReplyKeyboardForceReply) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replyKeyboardForceReply",
		ID:   ReplyKeyboardForceReplyTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SingleUse",
			SchemaName: "single_use",
			Null:       !r.Flags.Has(1),
		},
		{
			Name:       "Selective",
			SchemaName: "selective",
			Null:       !r.Flags.Has(2),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReplyKeyboardForceReply) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyKeyboardForceReply#f4108aa0 as nil")
	}
	b.PutID(ReplyKeyboardForceReplyTypeID)
	if !(r.SingleUse == false) {
		r.Flags.Set(1)
	}
	if !(r.Selective == false) {
		r.Flags.Set(2)
	}
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode replyKeyboardForceReply#f4108aa0: field flags: %w", err)
	}
	return nil
}

// SetSingleUse sets value of SingleUse conditional field.
func (r *ReplyKeyboardForceReply) SetSingleUse(value bool) {
	if value {
		r.Flags.Set(1)
		r.SingleUse = true
	} else {
		r.Flags.Unset(1)
		r.SingleUse = false
	}
}

// GetSingleUse returns value of SingleUse conditional field.
func (r *ReplyKeyboardForceReply) GetSingleUse() (value bool) {
	return r.Flags.Has(1)
}

// SetSelective sets value of Selective conditional field.
func (r *ReplyKeyboardForceReply) SetSelective(value bool) {
	if value {
		r.Flags.Set(2)
		r.Selective = true
	} else {
		r.Flags.Unset(2)
		r.Selective = false
	}
}

// GetSelective returns value of Selective conditional field.
func (r *ReplyKeyboardForceReply) GetSelective() (value bool) {
	return r.Flags.Has(2)
}

// Decode implements bin.Decoder.
func (r *ReplyKeyboardForceReply) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyKeyboardForceReply#f4108aa0 to nil")
	}
	if err := b.ConsumeID(ReplyKeyboardForceReplyTypeID); err != nil {
		return fmt.Errorf("unable to decode replyKeyboardForceReply#f4108aa0: %w", err)
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode replyKeyboardForceReply#f4108aa0: field flags: %w", err)
		}
	}
	r.SingleUse = r.Flags.Has(1)
	r.Selective = r.Flags.Has(2)
	return nil
}

// construct implements constructor of ReplyMarkupClass.
func (r ReplyKeyboardForceReply) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyKeyboardForceReply.
var (
	_ bin.Encoder = &ReplyKeyboardForceReply{}
	_ bin.Decoder = &ReplyKeyboardForceReply{}

	_ ReplyMarkupClass = &ReplyKeyboardForceReply{}
)

// ReplyKeyboardMarkup represents TL type `replyKeyboardMarkup#3502758c`.
// Bot keyboard
//
// See https://core.telegram.org/constructor/replyKeyboardMarkup for reference.
type ReplyKeyboardMarkup struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). If not set, the custom keyboard is always of the same height as the app's standard keyboard.
	Resize bool
	// Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again.
	SingleUse bool
	// Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard.
	Selective bool
	// Button row
	Rows []KeyboardButtonRow
}

// ReplyKeyboardMarkupTypeID is TL type id of ReplyKeyboardMarkup.
const ReplyKeyboardMarkupTypeID = 0x3502758c

func (r *ReplyKeyboardMarkup) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Resize == false) {
		return false
	}
	if !(r.SingleUse == false) {
		return false
	}
	if !(r.Selective == false) {
		return false
	}
	if !(r.Rows == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyKeyboardMarkup) String() string {
	if r == nil {
		return "ReplyKeyboardMarkup(nil)"
	}
	type Alias ReplyKeyboardMarkup
	return fmt.Sprintf("ReplyKeyboardMarkup%+v", Alias(*r))
}

// FillFrom fills ReplyKeyboardMarkup from given interface.
func (r *ReplyKeyboardMarkup) FillFrom(from interface {
	GetResize() (value bool)
	GetSingleUse() (value bool)
	GetSelective() (value bool)
	GetRows() (value []KeyboardButtonRow)
}) {
	r.Resize = from.GetResize()
	r.SingleUse = from.GetSingleUse()
	r.Selective = from.GetSelective()
	r.Rows = from.GetRows()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplyKeyboardMarkup) TypeID() uint32 {
	return ReplyKeyboardMarkupTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplyKeyboardMarkup) TypeName() string {
	return "replyKeyboardMarkup"
}

// TypeInfo returns info about TL type.
func (r *ReplyKeyboardMarkup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replyKeyboardMarkup",
		ID:   ReplyKeyboardMarkupTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Resize",
			SchemaName: "resize",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "SingleUse",
			SchemaName: "single_use",
			Null:       !r.Flags.Has(1),
		},
		{
			Name:       "Selective",
			SchemaName: "selective",
			Null:       !r.Flags.Has(2),
		},
		{
			Name:       "Rows",
			SchemaName: "rows",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReplyKeyboardMarkup) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyKeyboardMarkup#3502758c as nil")
	}
	b.PutID(ReplyKeyboardMarkupTypeID)
	if !(r.Resize == false) {
		r.Flags.Set(0)
	}
	if !(r.SingleUse == false) {
		r.Flags.Set(1)
	}
	if !(r.Selective == false) {
		r.Flags.Set(2)
	}
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode replyKeyboardMarkup#3502758c: field flags: %w", err)
	}
	b.PutVectorHeader(len(r.Rows))
	for idx, v := range r.Rows {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode replyKeyboardMarkup#3502758c: field rows element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetResize sets value of Resize conditional field.
func (r *ReplyKeyboardMarkup) SetResize(value bool) {
	if value {
		r.Flags.Set(0)
		r.Resize = true
	} else {
		r.Flags.Unset(0)
		r.Resize = false
	}
}

// GetResize returns value of Resize conditional field.
func (r *ReplyKeyboardMarkup) GetResize() (value bool) {
	return r.Flags.Has(0)
}

// SetSingleUse sets value of SingleUse conditional field.
func (r *ReplyKeyboardMarkup) SetSingleUse(value bool) {
	if value {
		r.Flags.Set(1)
		r.SingleUse = true
	} else {
		r.Flags.Unset(1)
		r.SingleUse = false
	}
}

// GetSingleUse returns value of SingleUse conditional field.
func (r *ReplyKeyboardMarkup) GetSingleUse() (value bool) {
	return r.Flags.Has(1)
}

// SetSelective sets value of Selective conditional field.
func (r *ReplyKeyboardMarkup) SetSelective(value bool) {
	if value {
		r.Flags.Set(2)
		r.Selective = true
	} else {
		r.Flags.Unset(2)
		r.Selective = false
	}
}

// GetSelective returns value of Selective conditional field.
func (r *ReplyKeyboardMarkup) GetSelective() (value bool) {
	return r.Flags.Has(2)
}

// GetRows returns value of Rows field.
func (r *ReplyKeyboardMarkup) GetRows() (value []KeyboardButtonRow) {
	return r.Rows
}

// Decode implements bin.Decoder.
func (r *ReplyKeyboardMarkup) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyKeyboardMarkup#3502758c to nil")
	}
	if err := b.ConsumeID(ReplyKeyboardMarkupTypeID); err != nil {
		return fmt.Errorf("unable to decode replyKeyboardMarkup#3502758c: %w", err)
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode replyKeyboardMarkup#3502758c: field flags: %w", err)
		}
	}
	r.Resize = r.Flags.Has(0)
	r.SingleUse = r.Flags.Has(1)
	r.Selective = r.Flags.Has(2)
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode replyKeyboardMarkup#3502758c: field rows: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value KeyboardButtonRow
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode replyKeyboardMarkup#3502758c: field rows: %w", err)
			}
			r.Rows = append(r.Rows, value)
		}
	}
	return nil
}

// construct implements constructor of ReplyMarkupClass.
func (r ReplyKeyboardMarkup) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyKeyboardMarkup.
var (
	_ bin.Encoder = &ReplyKeyboardMarkup{}
	_ bin.Decoder = &ReplyKeyboardMarkup{}

	_ ReplyMarkupClass = &ReplyKeyboardMarkup{}
)

// ReplyInlineMarkup represents TL type `replyInlineMarkup#48a30254`.
// Bot or inline keyboard
//
// See https://core.telegram.org/constructor/replyInlineMarkup for reference.
type ReplyInlineMarkup struct {
	// Bot or inline keyboard rows
	Rows []KeyboardButtonRow
}

// ReplyInlineMarkupTypeID is TL type id of ReplyInlineMarkup.
const ReplyInlineMarkupTypeID = 0x48a30254

func (r *ReplyInlineMarkup) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Rows == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyInlineMarkup) String() string {
	if r == nil {
		return "ReplyInlineMarkup(nil)"
	}
	type Alias ReplyInlineMarkup
	return fmt.Sprintf("ReplyInlineMarkup%+v", Alias(*r))
}

// FillFrom fills ReplyInlineMarkup from given interface.
func (r *ReplyInlineMarkup) FillFrom(from interface {
	GetRows() (value []KeyboardButtonRow)
}) {
	r.Rows = from.GetRows()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplyInlineMarkup) TypeID() uint32 {
	return ReplyInlineMarkupTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplyInlineMarkup) TypeName() string {
	return "replyInlineMarkup"
}

// TypeInfo returns info about TL type.
func (r *ReplyInlineMarkup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replyInlineMarkup",
		ID:   ReplyInlineMarkupTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Rows",
			SchemaName: "rows",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReplyInlineMarkup) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyInlineMarkup#48a30254 as nil")
	}
	b.PutID(ReplyInlineMarkupTypeID)
	b.PutVectorHeader(len(r.Rows))
	for idx, v := range r.Rows {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode replyInlineMarkup#48a30254: field rows element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetRows returns value of Rows field.
func (r *ReplyInlineMarkup) GetRows() (value []KeyboardButtonRow) {
	return r.Rows
}

// Decode implements bin.Decoder.
func (r *ReplyInlineMarkup) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyInlineMarkup#48a30254 to nil")
	}
	if err := b.ConsumeID(ReplyInlineMarkupTypeID); err != nil {
		return fmt.Errorf("unable to decode replyInlineMarkup#48a30254: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode replyInlineMarkup#48a30254: field rows: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value KeyboardButtonRow
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode replyInlineMarkup#48a30254: field rows: %w", err)
			}
			r.Rows = append(r.Rows, value)
		}
	}
	return nil
}

// construct implements constructor of ReplyMarkupClass.
func (r ReplyInlineMarkup) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyInlineMarkup.
var (
	_ bin.Encoder = &ReplyInlineMarkup{}
	_ bin.Decoder = &ReplyInlineMarkup{}

	_ ReplyMarkupClass = &ReplyInlineMarkup{}
)

// ReplyMarkupClass represents ReplyMarkup generic type.
//
// See https://core.telegram.org/type/ReplyMarkup for reference.
//
// Example:
//  g, err := tg.DecodeReplyMarkup(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.ReplyKeyboardHide: // replyKeyboardHide#a03e5b85
//  case *tg.ReplyKeyboardForceReply: // replyKeyboardForceReply#f4108aa0
//  case *tg.ReplyKeyboardMarkup: // replyKeyboardMarkup#3502758c
//  case *tg.ReplyInlineMarkup: // replyInlineMarkup#48a30254
//  default: panic(v)
//  }
type ReplyMarkupClass interface {
	bin.Encoder
	bin.Decoder
	construct() ReplyMarkupClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeReplyMarkup implements binary de-serialization for ReplyMarkupClass.
func DecodeReplyMarkup(buf *bin.Buffer) (ReplyMarkupClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ReplyKeyboardHideTypeID:
		// Decoding replyKeyboardHide#a03e5b85.
		v := ReplyKeyboardHide{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case ReplyKeyboardForceReplyTypeID:
		// Decoding replyKeyboardForceReply#f4108aa0.
		v := ReplyKeyboardForceReply{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case ReplyKeyboardMarkupTypeID:
		// Decoding replyKeyboardMarkup#3502758c.
		v := ReplyKeyboardMarkup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case ReplyInlineMarkupTypeID:
		// Decoding replyInlineMarkup#48a30254.
		v := ReplyInlineMarkup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", bin.NewUnexpectedID(id))
	}
}

// ReplyMarkup boxes the ReplyMarkupClass providing a helper.
type ReplyMarkupBox struct {
	ReplyMarkup ReplyMarkupClass
}

// Decode implements bin.Decoder for ReplyMarkupBox.
func (b *ReplyMarkupBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReplyMarkupBox to nil")
	}
	v, err := DecodeReplyMarkup(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ReplyMarkup = v
	return nil
}

// Encode implements bin.Encode for ReplyMarkupBox.
func (b *ReplyMarkupBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode ReplyMarkupClass as nil")
	}
	return b.ReplyMarkup.Encode(buf)
}

// ReplyMarkupClassArray is adapter for slice of ReplyMarkupClass.
type ReplyMarkupClassArray []ReplyMarkupClass

// Sort sorts slice of ReplyMarkupClass.
func (s ReplyMarkupClassArray) Sort(less func(a, b ReplyMarkupClass) bool) ReplyMarkupClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ReplyMarkupClass.
func (s ReplyMarkupClassArray) SortStable(less func(a, b ReplyMarkupClass) bool) ReplyMarkupClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ReplyMarkupClass.
func (s ReplyMarkupClassArray) Retain(keep func(x ReplyMarkupClass) bool) ReplyMarkupClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s ReplyMarkupClassArray) First() (v ReplyMarkupClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReplyMarkupClassArray) Last() (v ReplyMarkupClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReplyMarkupClassArray) PopFirst() (v ReplyMarkupClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ReplyMarkupClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ReplyMarkupClassArray) Pop() (v ReplyMarkupClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsReplyKeyboardHide returns copy with only ReplyKeyboardHide constructors.
func (s ReplyMarkupClassArray) AsReplyKeyboardHide() (to ReplyKeyboardHideArray) {
	for _, elem := range s {
		value, ok := elem.(*ReplyKeyboardHide)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsReplyKeyboardForceReply returns copy with only ReplyKeyboardForceReply constructors.
func (s ReplyMarkupClassArray) AsReplyKeyboardForceReply() (to ReplyKeyboardForceReplyArray) {
	for _, elem := range s {
		value, ok := elem.(*ReplyKeyboardForceReply)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsReplyKeyboardMarkup returns copy with only ReplyKeyboardMarkup constructors.
func (s ReplyMarkupClassArray) AsReplyKeyboardMarkup() (to ReplyKeyboardMarkupArray) {
	for _, elem := range s {
		value, ok := elem.(*ReplyKeyboardMarkup)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsReplyInlineMarkup returns copy with only ReplyInlineMarkup constructors.
func (s ReplyMarkupClassArray) AsReplyInlineMarkup() (to ReplyInlineMarkupArray) {
	for _, elem := range s {
		value, ok := elem.(*ReplyInlineMarkup)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// ReplyKeyboardHideArray is adapter for slice of ReplyKeyboardHide.
type ReplyKeyboardHideArray []ReplyKeyboardHide

// Sort sorts slice of ReplyKeyboardHide.
func (s ReplyKeyboardHideArray) Sort(less func(a, b ReplyKeyboardHide) bool) ReplyKeyboardHideArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ReplyKeyboardHide.
func (s ReplyKeyboardHideArray) SortStable(less func(a, b ReplyKeyboardHide) bool) ReplyKeyboardHideArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ReplyKeyboardHide.
func (s ReplyKeyboardHideArray) Retain(keep func(x ReplyKeyboardHide) bool) ReplyKeyboardHideArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s ReplyKeyboardHideArray) First() (v ReplyKeyboardHide, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReplyKeyboardHideArray) Last() (v ReplyKeyboardHide, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReplyKeyboardHideArray) PopFirst() (v ReplyKeyboardHide, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ReplyKeyboardHide
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ReplyKeyboardHideArray) Pop() (v ReplyKeyboardHide, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ReplyKeyboardForceReplyArray is adapter for slice of ReplyKeyboardForceReply.
type ReplyKeyboardForceReplyArray []ReplyKeyboardForceReply

// Sort sorts slice of ReplyKeyboardForceReply.
func (s ReplyKeyboardForceReplyArray) Sort(less func(a, b ReplyKeyboardForceReply) bool) ReplyKeyboardForceReplyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ReplyKeyboardForceReply.
func (s ReplyKeyboardForceReplyArray) SortStable(less func(a, b ReplyKeyboardForceReply) bool) ReplyKeyboardForceReplyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ReplyKeyboardForceReply.
func (s ReplyKeyboardForceReplyArray) Retain(keep func(x ReplyKeyboardForceReply) bool) ReplyKeyboardForceReplyArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s ReplyKeyboardForceReplyArray) First() (v ReplyKeyboardForceReply, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReplyKeyboardForceReplyArray) Last() (v ReplyKeyboardForceReply, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReplyKeyboardForceReplyArray) PopFirst() (v ReplyKeyboardForceReply, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ReplyKeyboardForceReply
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ReplyKeyboardForceReplyArray) Pop() (v ReplyKeyboardForceReply, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ReplyKeyboardMarkupArray is adapter for slice of ReplyKeyboardMarkup.
type ReplyKeyboardMarkupArray []ReplyKeyboardMarkup

// Sort sorts slice of ReplyKeyboardMarkup.
func (s ReplyKeyboardMarkupArray) Sort(less func(a, b ReplyKeyboardMarkup) bool) ReplyKeyboardMarkupArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ReplyKeyboardMarkup.
func (s ReplyKeyboardMarkupArray) SortStable(less func(a, b ReplyKeyboardMarkup) bool) ReplyKeyboardMarkupArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ReplyKeyboardMarkup.
func (s ReplyKeyboardMarkupArray) Retain(keep func(x ReplyKeyboardMarkup) bool) ReplyKeyboardMarkupArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s ReplyKeyboardMarkupArray) First() (v ReplyKeyboardMarkup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReplyKeyboardMarkupArray) Last() (v ReplyKeyboardMarkup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReplyKeyboardMarkupArray) PopFirst() (v ReplyKeyboardMarkup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ReplyKeyboardMarkup
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ReplyKeyboardMarkupArray) Pop() (v ReplyKeyboardMarkup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ReplyInlineMarkupArray is adapter for slice of ReplyInlineMarkup.
type ReplyInlineMarkupArray []ReplyInlineMarkup

// Sort sorts slice of ReplyInlineMarkup.
func (s ReplyInlineMarkupArray) Sort(less func(a, b ReplyInlineMarkup) bool) ReplyInlineMarkupArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ReplyInlineMarkup.
func (s ReplyInlineMarkupArray) SortStable(less func(a, b ReplyInlineMarkup) bool) ReplyInlineMarkupArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ReplyInlineMarkup.
func (s ReplyInlineMarkupArray) Retain(keep func(x ReplyInlineMarkup) bool) ReplyInlineMarkupArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s ReplyInlineMarkupArray) First() (v ReplyInlineMarkup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReplyInlineMarkupArray) Last() (v ReplyInlineMarkup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReplyInlineMarkupArray) PopFirst() (v ReplyInlineMarkup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ReplyInlineMarkup
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ReplyInlineMarkupArray) Pop() (v ReplyInlineMarkup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
