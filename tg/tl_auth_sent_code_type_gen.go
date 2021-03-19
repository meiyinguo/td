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

// AuthSentCodeTypeApp represents TL type `auth.sentCodeTypeApp#3dbb5986`.
// The code was sent through the telegram app
//
// See https://core.telegram.org/constructor/auth.sentCodeTypeApp for reference.
type AuthSentCodeTypeApp struct {
	// Length of the code in bytes
	Length int
}

// AuthSentCodeTypeAppTypeID is TL type id of AuthSentCodeTypeApp.
const AuthSentCodeTypeAppTypeID = 0x3dbb5986

func (s *AuthSentCodeTypeApp) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Length == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AuthSentCodeTypeApp) String() string {
	if s == nil {
		return "AuthSentCodeTypeApp(nil)"
	}
	type Alias AuthSentCodeTypeApp
	return fmt.Sprintf("AuthSentCodeTypeApp%+v", Alias(*s))
}

// FillFrom fills AuthSentCodeTypeApp from given interface.
func (s *AuthSentCodeTypeApp) FillFrom(from interface {
	GetLength() (value int)
}) {
	s.Length = from.GetLength()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthSentCodeTypeApp) TypeID() uint32 {
	return AuthSentCodeTypeAppTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthSentCodeTypeApp) TypeName() string {
	return "auth.sentCodeTypeApp"
}

// TypeInfo returns info about TL type.
func (s *AuthSentCodeTypeApp) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.sentCodeTypeApp",
		ID:   AuthSentCodeTypeAppTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Length",
			SchemaName: "length",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AuthSentCodeTypeApp) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeApp#3dbb5986 as nil")
	}
	b.PutID(AuthSentCodeTypeAppTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AuthSentCodeTypeApp) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeApp#3dbb5986 as nil")
	}
	b.PutInt(s.Length)
	return nil
}

// GetLength returns value of Length field.
func (s *AuthSentCodeTypeApp) GetLength() (value int) {
	return s.Length
}

// Decode implements bin.Decoder.
func (s *AuthSentCodeTypeApp) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeApp#3dbb5986 to nil")
	}
	if err := b.ConsumeID(AuthSentCodeTypeAppTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sentCodeTypeApp#3dbb5986: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AuthSentCodeTypeApp) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeApp#3dbb5986 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCodeTypeApp#3dbb5986: field length: %w", err)
		}
		s.Length = value
	}
	return nil
}

// construct implements constructor of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeApp) construct() AuthSentCodeTypeClass { return &s }

// Ensuring interfaces in compile-time for AuthSentCodeTypeApp.
var (
	_ bin.Encoder     = &AuthSentCodeTypeApp{}
	_ bin.Decoder     = &AuthSentCodeTypeApp{}
	_ bin.BareEncoder = &AuthSentCodeTypeApp{}
	_ bin.BareDecoder = &AuthSentCodeTypeApp{}

	_ AuthSentCodeTypeClass = &AuthSentCodeTypeApp{}
)

// AuthSentCodeTypeSMS represents TL type `auth.sentCodeTypeSms#c000bba2`.
// The code was sent via SMS
//
// See https://core.telegram.org/constructor/auth.sentCodeTypeSms for reference.
type AuthSentCodeTypeSMS struct {
	// Length of the code in bytes
	Length int
}

// AuthSentCodeTypeSMSTypeID is TL type id of AuthSentCodeTypeSMS.
const AuthSentCodeTypeSMSTypeID = 0xc000bba2

func (s *AuthSentCodeTypeSMS) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Length == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AuthSentCodeTypeSMS) String() string {
	if s == nil {
		return "AuthSentCodeTypeSMS(nil)"
	}
	type Alias AuthSentCodeTypeSMS
	return fmt.Sprintf("AuthSentCodeTypeSMS%+v", Alias(*s))
}

// FillFrom fills AuthSentCodeTypeSMS from given interface.
func (s *AuthSentCodeTypeSMS) FillFrom(from interface {
	GetLength() (value int)
}) {
	s.Length = from.GetLength()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthSentCodeTypeSMS) TypeID() uint32 {
	return AuthSentCodeTypeSMSTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthSentCodeTypeSMS) TypeName() string {
	return "auth.sentCodeTypeSms"
}

// TypeInfo returns info about TL type.
func (s *AuthSentCodeTypeSMS) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.sentCodeTypeSms",
		ID:   AuthSentCodeTypeSMSTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Length",
			SchemaName: "length",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AuthSentCodeTypeSMS) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeSms#c000bba2 as nil")
	}
	b.PutID(AuthSentCodeTypeSMSTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AuthSentCodeTypeSMS) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeSms#c000bba2 as nil")
	}
	b.PutInt(s.Length)
	return nil
}

// GetLength returns value of Length field.
func (s *AuthSentCodeTypeSMS) GetLength() (value int) {
	return s.Length
}

// Decode implements bin.Decoder.
func (s *AuthSentCodeTypeSMS) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeSms#c000bba2 to nil")
	}
	if err := b.ConsumeID(AuthSentCodeTypeSMSTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sentCodeTypeSms#c000bba2: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AuthSentCodeTypeSMS) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeSms#c000bba2 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCodeTypeSms#c000bba2: field length: %w", err)
		}
		s.Length = value
	}
	return nil
}

// construct implements constructor of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeSMS) construct() AuthSentCodeTypeClass { return &s }

// Ensuring interfaces in compile-time for AuthSentCodeTypeSMS.
var (
	_ bin.Encoder     = &AuthSentCodeTypeSMS{}
	_ bin.Decoder     = &AuthSentCodeTypeSMS{}
	_ bin.BareEncoder = &AuthSentCodeTypeSMS{}
	_ bin.BareDecoder = &AuthSentCodeTypeSMS{}

	_ AuthSentCodeTypeClass = &AuthSentCodeTypeSMS{}
)

// AuthSentCodeTypeCall represents TL type `auth.sentCodeTypeCall#5353e5a7`.
// The code will be sent via a phone call: a synthesized voice will tell the user which verification code to input.
//
// See https://core.telegram.org/constructor/auth.sentCodeTypeCall for reference.
type AuthSentCodeTypeCall struct {
	// Length of the verification code
	Length int
}

// AuthSentCodeTypeCallTypeID is TL type id of AuthSentCodeTypeCall.
const AuthSentCodeTypeCallTypeID = 0x5353e5a7

func (s *AuthSentCodeTypeCall) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Length == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AuthSentCodeTypeCall) String() string {
	if s == nil {
		return "AuthSentCodeTypeCall(nil)"
	}
	type Alias AuthSentCodeTypeCall
	return fmt.Sprintf("AuthSentCodeTypeCall%+v", Alias(*s))
}

// FillFrom fills AuthSentCodeTypeCall from given interface.
func (s *AuthSentCodeTypeCall) FillFrom(from interface {
	GetLength() (value int)
}) {
	s.Length = from.GetLength()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthSentCodeTypeCall) TypeID() uint32 {
	return AuthSentCodeTypeCallTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthSentCodeTypeCall) TypeName() string {
	return "auth.sentCodeTypeCall"
}

// TypeInfo returns info about TL type.
func (s *AuthSentCodeTypeCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.sentCodeTypeCall",
		ID:   AuthSentCodeTypeCallTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Length",
			SchemaName: "length",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AuthSentCodeTypeCall) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeCall#5353e5a7 as nil")
	}
	b.PutID(AuthSentCodeTypeCallTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AuthSentCodeTypeCall) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeCall#5353e5a7 as nil")
	}
	b.PutInt(s.Length)
	return nil
}

// GetLength returns value of Length field.
func (s *AuthSentCodeTypeCall) GetLength() (value int) {
	return s.Length
}

// Decode implements bin.Decoder.
func (s *AuthSentCodeTypeCall) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeCall#5353e5a7 to nil")
	}
	if err := b.ConsumeID(AuthSentCodeTypeCallTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sentCodeTypeCall#5353e5a7: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AuthSentCodeTypeCall) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeCall#5353e5a7 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCodeTypeCall#5353e5a7: field length: %w", err)
		}
		s.Length = value
	}
	return nil
}

// construct implements constructor of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeCall) construct() AuthSentCodeTypeClass { return &s }

// Ensuring interfaces in compile-time for AuthSentCodeTypeCall.
var (
	_ bin.Encoder     = &AuthSentCodeTypeCall{}
	_ bin.Decoder     = &AuthSentCodeTypeCall{}
	_ bin.BareEncoder = &AuthSentCodeTypeCall{}
	_ bin.BareDecoder = &AuthSentCodeTypeCall{}

	_ AuthSentCodeTypeClass = &AuthSentCodeTypeCall{}
)

// AuthSentCodeTypeFlashCall represents TL type `auth.sentCodeTypeFlashCall#ab03c6d9`.
// The code will be sent via a flash phone call, that will be closed immediately. The phone code will then be the phone number itself, just make sure that the phone number matches the specified pattern.
//
// See https://core.telegram.org/constructor/auth.sentCodeTypeFlashCall for reference.
type AuthSentCodeTypeFlashCall struct {
	// pattern¹ to match
	//
	// Links:
	//  1) https://core.telegram.org/api/pattern
	Pattern string
}

// AuthSentCodeTypeFlashCallTypeID is TL type id of AuthSentCodeTypeFlashCall.
const AuthSentCodeTypeFlashCallTypeID = 0xab03c6d9

func (s *AuthSentCodeTypeFlashCall) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Pattern == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AuthSentCodeTypeFlashCall) String() string {
	if s == nil {
		return "AuthSentCodeTypeFlashCall(nil)"
	}
	type Alias AuthSentCodeTypeFlashCall
	return fmt.Sprintf("AuthSentCodeTypeFlashCall%+v", Alias(*s))
}

// FillFrom fills AuthSentCodeTypeFlashCall from given interface.
func (s *AuthSentCodeTypeFlashCall) FillFrom(from interface {
	GetPattern() (value string)
}) {
	s.Pattern = from.GetPattern()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthSentCodeTypeFlashCall) TypeID() uint32 {
	return AuthSentCodeTypeFlashCallTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthSentCodeTypeFlashCall) TypeName() string {
	return "auth.sentCodeTypeFlashCall"
}

// TypeInfo returns info about TL type.
func (s *AuthSentCodeTypeFlashCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.sentCodeTypeFlashCall",
		ID:   AuthSentCodeTypeFlashCallTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Pattern",
			SchemaName: "pattern",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AuthSentCodeTypeFlashCall) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeFlashCall#ab03c6d9 as nil")
	}
	b.PutID(AuthSentCodeTypeFlashCallTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AuthSentCodeTypeFlashCall) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeFlashCall#ab03c6d9 as nil")
	}
	b.PutString(s.Pattern)
	return nil
}

// GetPattern returns value of Pattern field.
func (s *AuthSentCodeTypeFlashCall) GetPattern() (value string) {
	return s.Pattern
}

// Decode implements bin.Decoder.
func (s *AuthSentCodeTypeFlashCall) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeFlashCall#ab03c6d9 to nil")
	}
	if err := b.ConsumeID(AuthSentCodeTypeFlashCallTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sentCodeTypeFlashCall#ab03c6d9: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AuthSentCodeTypeFlashCall) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeFlashCall#ab03c6d9 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCodeTypeFlashCall#ab03c6d9: field pattern: %w", err)
		}
		s.Pattern = value
	}
	return nil
}

// construct implements constructor of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeFlashCall) construct() AuthSentCodeTypeClass { return &s }

// Ensuring interfaces in compile-time for AuthSentCodeTypeFlashCall.
var (
	_ bin.Encoder     = &AuthSentCodeTypeFlashCall{}
	_ bin.Decoder     = &AuthSentCodeTypeFlashCall{}
	_ bin.BareEncoder = &AuthSentCodeTypeFlashCall{}
	_ bin.BareDecoder = &AuthSentCodeTypeFlashCall{}

	_ AuthSentCodeTypeClass = &AuthSentCodeTypeFlashCall{}
)

// AuthSentCodeTypeClass represents auth.SentCodeType generic type.
//
// See https://core.telegram.org/type/auth.SentCodeType for reference.
//
// Example:
//  g, err := tg.DecodeAuthSentCodeType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.AuthSentCodeTypeApp: // auth.sentCodeTypeApp#3dbb5986
//  case *tg.AuthSentCodeTypeSMS: // auth.sentCodeTypeSms#c000bba2
//  case *tg.AuthSentCodeTypeCall: // auth.sentCodeTypeCall#5353e5a7
//  case *tg.AuthSentCodeTypeFlashCall: // auth.sentCodeTypeFlashCall#ab03c6d9
//  default: panic(v)
//  }
type AuthSentCodeTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AuthSentCodeTypeClass

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

// DecodeAuthSentCodeType implements binary de-serialization for AuthSentCodeTypeClass.
func DecodeAuthSentCodeType(buf *bin.Buffer) (AuthSentCodeTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthSentCodeTypeAppTypeID:
		// Decoding auth.sentCodeTypeApp#3dbb5986.
		v := AuthSentCodeTypeApp{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthSentCodeTypeSMSTypeID:
		// Decoding auth.sentCodeTypeSms#c000bba2.
		v := AuthSentCodeTypeSMS{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthSentCodeTypeCallTypeID:
		// Decoding auth.sentCodeTypeCall#5353e5a7.
		v := AuthSentCodeTypeCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthSentCodeTypeFlashCallTypeID:
		// Decoding auth.sentCodeTypeFlashCall#ab03c6d9.
		v := AuthSentCodeTypeFlashCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// AuthSentCodeType boxes the AuthSentCodeTypeClass providing a helper.
type AuthSentCodeTypeBox struct {
	SentCodeType AuthSentCodeTypeClass
}

// Decode implements bin.Decoder for AuthSentCodeTypeBox.
func (b *AuthSentCodeTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthSentCodeTypeBox to nil")
	}
	v, err := DecodeAuthSentCodeType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SentCodeType = v
	return nil
}

// Encode implements bin.Encode for AuthSentCodeTypeBox.
func (b *AuthSentCodeTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SentCodeType == nil {
		return fmt.Errorf("unable to encode AuthSentCodeTypeClass as nil")
	}
	return b.SentCodeType.Encode(buf)
}

// AuthSentCodeTypeClassArray is adapter for slice of AuthSentCodeTypeClass.
type AuthSentCodeTypeClassArray []AuthSentCodeTypeClass

// Sort sorts slice of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeClassArray) Sort(less func(a, b AuthSentCodeTypeClass) bool) AuthSentCodeTypeClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeClassArray) SortStable(less func(a, b AuthSentCodeTypeClass) bool) AuthSentCodeTypeClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeClassArray) Retain(keep func(x AuthSentCodeTypeClass) bool) AuthSentCodeTypeClassArray {
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
func (s AuthSentCodeTypeClassArray) First() (v AuthSentCodeTypeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthSentCodeTypeClassArray) Last() (v AuthSentCodeTypeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeClassArray) PopFirst() (v AuthSentCodeTypeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthSentCodeTypeClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeClassArray) Pop() (v AuthSentCodeTypeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsAuthSentCodeTypeApp returns copy with only AuthSentCodeTypeApp constructors.
func (s AuthSentCodeTypeClassArray) AsAuthSentCodeTypeApp() (to AuthSentCodeTypeAppArray) {
	for _, elem := range s {
		value, ok := elem.(*AuthSentCodeTypeApp)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsAuthSentCodeTypeSMS returns copy with only AuthSentCodeTypeSMS constructors.
func (s AuthSentCodeTypeClassArray) AsAuthSentCodeTypeSMS() (to AuthSentCodeTypeSMSArray) {
	for _, elem := range s {
		value, ok := elem.(*AuthSentCodeTypeSMS)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsAuthSentCodeTypeCall returns copy with only AuthSentCodeTypeCall constructors.
func (s AuthSentCodeTypeClassArray) AsAuthSentCodeTypeCall() (to AuthSentCodeTypeCallArray) {
	for _, elem := range s {
		value, ok := elem.(*AuthSentCodeTypeCall)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsAuthSentCodeTypeFlashCall returns copy with only AuthSentCodeTypeFlashCall constructors.
func (s AuthSentCodeTypeClassArray) AsAuthSentCodeTypeFlashCall() (to AuthSentCodeTypeFlashCallArray) {
	for _, elem := range s {
		value, ok := elem.(*AuthSentCodeTypeFlashCall)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AuthSentCodeTypeAppArray is adapter for slice of AuthSentCodeTypeApp.
type AuthSentCodeTypeAppArray []AuthSentCodeTypeApp

// Sort sorts slice of AuthSentCodeTypeApp.
func (s AuthSentCodeTypeAppArray) Sort(less func(a, b AuthSentCodeTypeApp) bool) AuthSentCodeTypeAppArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthSentCodeTypeApp.
func (s AuthSentCodeTypeAppArray) SortStable(less func(a, b AuthSentCodeTypeApp) bool) AuthSentCodeTypeAppArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthSentCodeTypeApp.
func (s AuthSentCodeTypeAppArray) Retain(keep func(x AuthSentCodeTypeApp) bool) AuthSentCodeTypeAppArray {
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
func (s AuthSentCodeTypeAppArray) First() (v AuthSentCodeTypeApp, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthSentCodeTypeAppArray) Last() (v AuthSentCodeTypeApp, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeAppArray) PopFirst() (v AuthSentCodeTypeApp, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthSentCodeTypeApp
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeAppArray) Pop() (v AuthSentCodeTypeApp, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AuthSentCodeTypeSMSArray is adapter for slice of AuthSentCodeTypeSMS.
type AuthSentCodeTypeSMSArray []AuthSentCodeTypeSMS

// Sort sorts slice of AuthSentCodeTypeSMS.
func (s AuthSentCodeTypeSMSArray) Sort(less func(a, b AuthSentCodeTypeSMS) bool) AuthSentCodeTypeSMSArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthSentCodeTypeSMS.
func (s AuthSentCodeTypeSMSArray) SortStable(less func(a, b AuthSentCodeTypeSMS) bool) AuthSentCodeTypeSMSArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthSentCodeTypeSMS.
func (s AuthSentCodeTypeSMSArray) Retain(keep func(x AuthSentCodeTypeSMS) bool) AuthSentCodeTypeSMSArray {
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
func (s AuthSentCodeTypeSMSArray) First() (v AuthSentCodeTypeSMS, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthSentCodeTypeSMSArray) Last() (v AuthSentCodeTypeSMS, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeSMSArray) PopFirst() (v AuthSentCodeTypeSMS, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthSentCodeTypeSMS
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeSMSArray) Pop() (v AuthSentCodeTypeSMS, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AuthSentCodeTypeCallArray is adapter for slice of AuthSentCodeTypeCall.
type AuthSentCodeTypeCallArray []AuthSentCodeTypeCall

// Sort sorts slice of AuthSentCodeTypeCall.
func (s AuthSentCodeTypeCallArray) Sort(less func(a, b AuthSentCodeTypeCall) bool) AuthSentCodeTypeCallArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthSentCodeTypeCall.
func (s AuthSentCodeTypeCallArray) SortStable(less func(a, b AuthSentCodeTypeCall) bool) AuthSentCodeTypeCallArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthSentCodeTypeCall.
func (s AuthSentCodeTypeCallArray) Retain(keep func(x AuthSentCodeTypeCall) bool) AuthSentCodeTypeCallArray {
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
func (s AuthSentCodeTypeCallArray) First() (v AuthSentCodeTypeCall, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthSentCodeTypeCallArray) Last() (v AuthSentCodeTypeCall, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeCallArray) PopFirst() (v AuthSentCodeTypeCall, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthSentCodeTypeCall
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeCallArray) Pop() (v AuthSentCodeTypeCall, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AuthSentCodeTypeFlashCallArray is adapter for slice of AuthSentCodeTypeFlashCall.
type AuthSentCodeTypeFlashCallArray []AuthSentCodeTypeFlashCall

// Sort sorts slice of AuthSentCodeTypeFlashCall.
func (s AuthSentCodeTypeFlashCallArray) Sort(less func(a, b AuthSentCodeTypeFlashCall) bool) AuthSentCodeTypeFlashCallArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthSentCodeTypeFlashCall.
func (s AuthSentCodeTypeFlashCallArray) SortStable(less func(a, b AuthSentCodeTypeFlashCall) bool) AuthSentCodeTypeFlashCallArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthSentCodeTypeFlashCall.
func (s AuthSentCodeTypeFlashCallArray) Retain(keep func(x AuthSentCodeTypeFlashCall) bool) AuthSentCodeTypeFlashCallArray {
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
func (s AuthSentCodeTypeFlashCallArray) First() (v AuthSentCodeTypeFlashCall, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthSentCodeTypeFlashCallArray) Last() (v AuthSentCodeTypeFlashCall, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeFlashCallArray) PopFirst() (v AuthSentCodeTypeFlashCall, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthSentCodeTypeFlashCall
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeFlashCallArray) Pop() (v AuthSentCodeTypeFlashCall, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
