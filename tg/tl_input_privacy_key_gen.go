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

// InputPrivacyKeyStatusTimestamp represents TL type `inputPrivacyKeyStatusTimestamp#4f96cb18`.
// Whether we can see the exact last online timestamp of the user
//
// See https://core.telegram.org/constructor/inputPrivacyKeyStatusTimestamp for reference.
type InputPrivacyKeyStatusTimestamp struct {
}

// InputPrivacyKeyStatusTimestampTypeID is TL type id of InputPrivacyKeyStatusTimestamp.
const InputPrivacyKeyStatusTimestampTypeID = 0x4f96cb18

func (i *InputPrivacyKeyStatusTimestamp) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyStatusTimestamp) String() string {
	if i == nil {
		return "InputPrivacyKeyStatusTimestamp(nil)"
	}
	type Alias InputPrivacyKeyStatusTimestamp
	return fmt.Sprintf("InputPrivacyKeyStatusTimestamp%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyKeyStatusTimestamp) TypeID() uint32 {
	return InputPrivacyKeyStatusTimestampTypeID
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyStatusTimestamp) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyStatusTimestamp#4f96cb18 as nil")
	}
	b.PutID(InputPrivacyKeyStatusTimestampTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyStatusTimestamp) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyStatusTimestamp#4f96cb18 to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyStatusTimestampTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyStatusTimestamp#4f96cb18: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyStatusTimestamp) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyStatusTimestamp.
var (
	_ bin.Encoder = &InputPrivacyKeyStatusTimestamp{}
	_ bin.Decoder = &InputPrivacyKeyStatusTimestamp{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyStatusTimestamp{}
)

// InputPrivacyKeyChatInvite represents TL type `inputPrivacyKeyChatInvite#bdfb0426`.
// Whether the user can be invited to chats
//
// See https://core.telegram.org/constructor/inputPrivacyKeyChatInvite for reference.
type InputPrivacyKeyChatInvite struct {
}

// InputPrivacyKeyChatInviteTypeID is TL type id of InputPrivacyKeyChatInvite.
const InputPrivacyKeyChatInviteTypeID = 0xbdfb0426

func (i *InputPrivacyKeyChatInvite) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyChatInvite) String() string {
	if i == nil {
		return "InputPrivacyKeyChatInvite(nil)"
	}
	type Alias InputPrivacyKeyChatInvite
	return fmt.Sprintf("InputPrivacyKeyChatInvite%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyKeyChatInvite) TypeID() uint32 {
	return InputPrivacyKeyChatInviteTypeID
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyChatInvite) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyChatInvite#bdfb0426 as nil")
	}
	b.PutID(InputPrivacyKeyChatInviteTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyChatInvite) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyChatInvite#bdfb0426 to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyChatInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyChatInvite#bdfb0426: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyChatInvite) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyChatInvite.
var (
	_ bin.Encoder = &InputPrivacyKeyChatInvite{}
	_ bin.Decoder = &InputPrivacyKeyChatInvite{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyChatInvite{}
)

// InputPrivacyKeyPhoneCall represents TL type `inputPrivacyKeyPhoneCall#fabadc5f`.
// Whether the user will accept phone calls
//
// See https://core.telegram.org/constructor/inputPrivacyKeyPhoneCall for reference.
type InputPrivacyKeyPhoneCall struct {
}

// InputPrivacyKeyPhoneCallTypeID is TL type id of InputPrivacyKeyPhoneCall.
const InputPrivacyKeyPhoneCallTypeID = 0xfabadc5f

func (i *InputPrivacyKeyPhoneCall) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyPhoneCall) String() string {
	if i == nil {
		return "InputPrivacyKeyPhoneCall(nil)"
	}
	type Alias InputPrivacyKeyPhoneCall
	return fmt.Sprintf("InputPrivacyKeyPhoneCall%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyKeyPhoneCall) TypeID() uint32 {
	return InputPrivacyKeyPhoneCallTypeID
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyPhoneCall) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyPhoneCall#fabadc5f as nil")
	}
	b.PutID(InputPrivacyKeyPhoneCallTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyPhoneCall) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyPhoneCall#fabadc5f to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyPhoneCallTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyPhoneCall#fabadc5f: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyPhoneCall) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyPhoneCall.
var (
	_ bin.Encoder = &InputPrivacyKeyPhoneCall{}
	_ bin.Decoder = &InputPrivacyKeyPhoneCall{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyPhoneCall{}
)

// InputPrivacyKeyPhoneP2P represents TL type `inputPrivacyKeyPhoneP2P#db9e70d2`.
// Whether the user allows P2P communication during VoIP calls
//
// See https://core.telegram.org/constructor/inputPrivacyKeyPhoneP2P for reference.
type InputPrivacyKeyPhoneP2P struct {
}

// InputPrivacyKeyPhoneP2PTypeID is TL type id of InputPrivacyKeyPhoneP2P.
const InputPrivacyKeyPhoneP2PTypeID = 0xdb9e70d2

func (i *InputPrivacyKeyPhoneP2P) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyPhoneP2P) String() string {
	if i == nil {
		return "InputPrivacyKeyPhoneP2P(nil)"
	}
	type Alias InputPrivacyKeyPhoneP2P
	return fmt.Sprintf("InputPrivacyKeyPhoneP2P%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyKeyPhoneP2P) TypeID() uint32 {
	return InputPrivacyKeyPhoneP2PTypeID
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyPhoneP2P) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyPhoneP2P#db9e70d2 as nil")
	}
	b.PutID(InputPrivacyKeyPhoneP2PTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyPhoneP2P) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyPhoneP2P#db9e70d2 to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyPhoneP2PTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyPhoneP2P#db9e70d2: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyPhoneP2P) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyPhoneP2P.
var (
	_ bin.Encoder = &InputPrivacyKeyPhoneP2P{}
	_ bin.Decoder = &InputPrivacyKeyPhoneP2P{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyPhoneP2P{}
)

// InputPrivacyKeyForwards represents TL type `inputPrivacyKeyForwards#a4dd4c08`.
// Whether messages forwarded from this user will be anonymous¹
//
// Links:
//  1) https://telegram.org/blog/unsend-privacy-emoji#anonymous-forwarding
//
// See https://core.telegram.org/constructor/inputPrivacyKeyForwards for reference.
type InputPrivacyKeyForwards struct {
}

// InputPrivacyKeyForwardsTypeID is TL type id of InputPrivacyKeyForwards.
const InputPrivacyKeyForwardsTypeID = 0xa4dd4c08

func (i *InputPrivacyKeyForwards) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyForwards) String() string {
	if i == nil {
		return "InputPrivacyKeyForwards(nil)"
	}
	type Alias InputPrivacyKeyForwards
	return fmt.Sprintf("InputPrivacyKeyForwards%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyKeyForwards) TypeID() uint32 {
	return InputPrivacyKeyForwardsTypeID
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyForwards) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyForwards#a4dd4c08 as nil")
	}
	b.PutID(InputPrivacyKeyForwardsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyForwards) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyForwards#a4dd4c08 to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyForwardsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyForwards#a4dd4c08: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyForwards) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyForwards.
var (
	_ bin.Encoder = &InputPrivacyKeyForwards{}
	_ bin.Decoder = &InputPrivacyKeyForwards{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyForwards{}
)

// InputPrivacyKeyProfilePhoto represents TL type `inputPrivacyKeyProfilePhoto#5719bacc`.
// Whether people will be able to see the user's profile picture
//
// See https://core.telegram.org/constructor/inputPrivacyKeyProfilePhoto for reference.
type InputPrivacyKeyProfilePhoto struct {
}

// InputPrivacyKeyProfilePhotoTypeID is TL type id of InputPrivacyKeyProfilePhoto.
const InputPrivacyKeyProfilePhotoTypeID = 0x5719bacc

func (i *InputPrivacyKeyProfilePhoto) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyProfilePhoto) String() string {
	if i == nil {
		return "InputPrivacyKeyProfilePhoto(nil)"
	}
	type Alias InputPrivacyKeyProfilePhoto
	return fmt.Sprintf("InputPrivacyKeyProfilePhoto%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyKeyProfilePhoto) TypeID() uint32 {
	return InputPrivacyKeyProfilePhotoTypeID
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyProfilePhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyProfilePhoto#5719bacc as nil")
	}
	b.PutID(InputPrivacyKeyProfilePhotoTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyProfilePhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyProfilePhoto#5719bacc to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyProfilePhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyProfilePhoto#5719bacc: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyProfilePhoto) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyProfilePhoto.
var (
	_ bin.Encoder = &InputPrivacyKeyProfilePhoto{}
	_ bin.Decoder = &InputPrivacyKeyProfilePhoto{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyProfilePhoto{}
)

// InputPrivacyKeyPhoneNumber represents TL type `inputPrivacyKeyPhoneNumber#352dafa`.
// Whether people will be able to see the user's phone number
//
// See https://core.telegram.org/constructor/inputPrivacyKeyPhoneNumber for reference.
type InputPrivacyKeyPhoneNumber struct {
}

// InputPrivacyKeyPhoneNumberTypeID is TL type id of InputPrivacyKeyPhoneNumber.
const InputPrivacyKeyPhoneNumberTypeID = 0x352dafa

func (i *InputPrivacyKeyPhoneNumber) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyPhoneNumber) String() string {
	if i == nil {
		return "InputPrivacyKeyPhoneNumber(nil)"
	}
	type Alias InputPrivacyKeyPhoneNumber
	return fmt.Sprintf("InputPrivacyKeyPhoneNumber%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyKeyPhoneNumber) TypeID() uint32 {
	return InputPrivacyKeyPhoneNumberTypeID
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyPhoneNumber) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyPhoneNumber#352dafa as nil")
	}
	b.PutID(InputPrivacyKeyPhoneNumberTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyPhoneNumber) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyPhoneNumber#352dafa to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyPhoneNumberTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyPhoneNumber#352dafa: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyPhoneNumber) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyPhoneNumber.
var (
	_ bin.Encoder = &InputPrivacyKeyPhoneNumber{}
	_ bin.Decoder = &InputPrivacyKeyPhoneNumber{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyPhoneNumber{}
)

// InputPrivacyKeyAddedByPhone represents TL type `inputPrivacyKeyAddedByPhone#d1219bdd`.
// Whether people can add you to their contact list by your phone number
//
// See https://core.telegram.org/constructor/inputPrivacyKeyAddedByPhone for reference.
type InputPrivacyKeyAddedByPhone struct {
}

// InputPrivacyKeyAddedByPhoneTypeID is TL type id of InputPrivacyKeyAddedByPhone.
const InputPrivacyKeyAddedByPhoneTypeID = 0xd1219bdd

func (i *InputPrivacyKeyAddedByPhone) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyAddedByPhone) String() string {
	if i == nil {
		return "InputPrivacyKeyAddedByPhone(nil)"
	}
	type Alias InputPrivacyKeyAddedByPhone
	return fmt.Sprintf("InputPrivacyKeyAddedByPhone%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPrivacyKeyAddedByPhone) TypeID() uint32 {
	return InputPrivacyKeyAddedByPhoneTypeID
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyAddedByPhone) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyAddedByPhone#d1219bdd as nil")
	}
	b.PutID(InputPrivacyKeyAddedByPhoneTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyAddedByPhone) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyAddedByPhone#d1219bdd to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyAddedByPhoneTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyAddedByPhone#d1219bdd: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyAddedByPhone) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyAddedByPhone.
var (
	_ bin.Encoder = &InputPrivacyKeyAddedByPhone{}
	_ bin.Decoder = &InputPrivacyKeyAddedByPhone{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyAddedByPhone{}
)

// InputPrivacyKeyClass represents InputPrivacyKey generic type.
//
// See https://core.telegram.org/type/InputPrivacyKey for reference.
//
// Example:
//  g, err := tg.DecodeInputPrivacyKey(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputPrivacyKeyStatusTimestamp: // inputPrivacyKeyStatusTimestamp#4f96cb18
//  case *tg.InputPrivacyKeyChatInvite: // inputPrivacyKeyChatInvite#bdfb0426
//  case *tg.InputPrivacyKeyPhoneCall: // inputPrivacyKeyPhoneCall#fabadc5f
//  case *tg.InputPrivacyKeyPhoneP2P: // inputPrivacyKeyPhoneP2P#db9e70d2
//  case *tg.InputPrivacyKeyForwards: // inputPrivacyKeyForwards#a4dd4c08
//  case *tg.InputPrivacyKeyProfilePhoto: // inputPrivacyKeyProfilePhoto#5719bacc
//  case *tg.InputPrivacyKeyPhoneNumber: // inputPrivacyKeyPhoneNumber#352dafa
//  case *tg.InputPrivacyKeyAddedByPhone: // inputPrivacyKeyAddedByPhone#d1219bdd
//  default: panic(v)
//  }
type InputPrivacyKeyClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputPrivacyKeyClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputPrivacyKey implements binary de-serialization for InputPrivacyKeyClass.
func DecodeInputPrivacyKey(buf *bin.Buffer) (InputPrivacyKeyClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPrivacyKeyStatusTimestampTypeID:
		// Decoding inputPrivacyKeyStatusTimestamp#4f96cb18.
		v := InputPrivacyKeyStatusTimestamp{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyChatInviteTypeID:
		// Decoding inputPrivacyKeyChatInvite#bdfb0426.
		v := InputPrivacyKeyChatInvite{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyPhoneCallTypeID:
		// Decoding inputPrivacyKeyPhoneCall#fabadc5f.
		v := InputPrivacyKeyPhoneCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyPhoneP2PTypeID:
		// Decoding inputPrivacyKeyPhoneP2P#db9e70d2.
		v := InputPrivacyKeyPhoneP2P{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyForwardsTypeID:
		// Decoding inputPrivacyKeyForwards#a4dd4c08.
		v := InputPrivacyKeyForwards{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyProfilePhotoTypeID:
		// Decoding inputPrivacyKeyProfilePhoto#5719bacc.
		v := InputPrivacyKeyProfilePhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyPhoneNumberTypeID:
		// Decoding inputPrivacyKeyPhoneNumber#352dafa.
		v := InputPrivacyKeyPhoneNumber{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyAddedByPhoneTypeID:
		// Decoding inputPrivacyKeyAddedByPhone#d1219bdd.
		v := InputPrivacyKeyAddedByPhone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputPrivacyKey boxes the InputPrivacyKeyClass providing a helper.
type InputPrivacyKeyBox struct {
	InputPrivacyKey InputPrivacyKeyClass
}

// Decode implements bin.Decoder for InputPrivacyKeyBox.
func (b *InputPrivacyKeyBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPrivacyKeyBox to nil")
	}
	v, err := DecodeInputPrivacyKey(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPrivacyKey = v
	return nil
}

// Encode implements bin.Encode for InputPrivacyKeyBox.
func (b *InputPrivacyKeyBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPrivacyKey == nil {
		return fmt.Errorf("unable to encode InputPrivacyKeyClass as nil")
	}
	return b.InputPrivacyKey.Encode(buf)
}

// InputPrivacyKeyClassSlice is adapter for slice of InputPrivacyKeyClass.
type InputPrivacyKeyClassSlice []InputPrivacyKeyClass

// First returns first element of slice (if exists).
func (s InputPrivacyKeyClassSlice) First() (v InputPrivacyKeyClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPrivacyKeyClassSlice) Last() (v InputPrivacyKeyClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPrivacyKeyClassSlice) PopFirst() (v InputPrivacyKeyClass, ok bool) {
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
func (s *InputPrivacyKeyClassSlice) Pop() (v InputPrivacyKeyClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
