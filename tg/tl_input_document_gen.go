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

// InputDocumentEmpty represents TL type `inputDocumentEmpty#72f0eaae`.
// Empty constructor.
//
// See https://core.telegram.org/constructor/inputDocumentEmpty for reference.
type InputDocumentEmpty struct {
}

// InputDocumentEmptyTypeID is TL type id of InputDocumentEmpty.
const InputDocumentEmptyTypeID = 0x72f0eaae

func (i *InputDocumentEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputDocumentEmpty) String() string {
	if i == nil {
		return "InputDocumentEmpty(nil)"
	}
	type Alias InputDocumentEmpty
	return fmt.Sprintf("InputDocumentEmpty%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputDocumentEmpty) TypeID() uint32 {
	return InputDocumentEmptyTypeID
}

// Encode implements bin.Encoder.
func (i *InputDocumentEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputDocumentEmpty#72f0eaae as nil")
	}
	b.PutID(InputDocumentEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputDocumentEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputDocumentEmpty#72f0eaae to nil")
	}
	if err := b.ConsumeID(InputDocumentEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputDocumentEmpty#72f0eaae: %w", err)
	}
	return nil
}

// construct implements constructor of InputDocumentClass.
func (i InputDocumentEmpty) construct() InputDocumentClass { return &i }

// Ensuring interfaces in compile-time for InputDocumentEmpty.
var (
	_ bin.Encoder = &InputDocumentEmpty{}
	_ bin.Decoder = &InputDocumentEmpty{}

	_ InputDocumentClass = &InputDocumentEmpty{}
)

// InputDocument represents TL type `inputDocument#1abfb575`.
// Defines a video for subsequent interaction.
//
// See https://core.telegram.org/constructor/inputDocument for reference.
type InputDocument struct {
	// Document ID
	ID int64
	// access_hash parameter from the document¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/constructor/document
	AccessHash int64
	// File reference¹
	//
	// Links:
	//  1) https://core.telegram.org/api/file_reference
	FileReference []byte
}

// InputDocumentTypeID is TL type id of InputDocument.
const InputDocumentTypeID = 0x1abfb575

func (i *InputDocument) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}
	if !(i.FileReference == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputDocument) String() string {
	if i == nil {
		return "InputDocument(nil)"
	}
	type Alias InputDocument
	return fmt.Sprintf("InputDocument%+v", Alias(*i))
}

// FillFrom fills InputDocument from given interface.
func (i *InputDocument) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetFileReference() (value []byte)
}) {
	i.ID = from.GetID()
	i.AccessHash = from.GetAccessHash()
	i.FileReference = from.GetFileReference()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputDocument) TypeID() uint32 {
	return InputDocumentTypeID
}

// Encode implements bin.Encoder.
func (i *InputDocument) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputDocument#1abfb575 as nil")
	}
	b.PutID(InputDocumentTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	b.PutBytes(i.FileReference)
	return nil
}

// GetID returns value of ID field.
func (i *InputDocument) GetID() (value int64) {
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputDocument) GetAccessHash() (value int64) {
	return i.AccessHash
}

// GetFileReference returns value of FileReference field.
func (i *InputDocument) GetFileReference() (value []byte) {
	return i.FileReference
}

// Decode implements bin.Decoder.
func (i *InputDocument) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputDocument#1abfb575 to nil")
	}
	if err := b.ConsumeID(InputDocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode inputDocument#1abfb575: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputDocument#1abfb575: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputDocument#1abfb575: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputDocument#1abfb575: field file_reference: %w", err)
		}
		i.FileReference = value
	}
	return nil
}

// construct implements constructor of InputDocumentClass.
func (i InputDocument) construct() InputDocumentClass { return &i }

// Ensuring interfaces in compile-time for InputDocument.
var (
	_ bin.Encoder = &InputDocument{}
	_ bin.Decoder = &InputDocument{}

	_ InputDocumentClass = &InputDocument{}
)

// InputDocumentClass represents InputDocument generic type.
//
// See https://core.telegram.org/type/InputDocument for reference.
//
// Example:
//  g, err := tg.DecodeInputDocument(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputDocumentEmpty: // inputDocumentEmpty#72f0eaae
//  case *tg.InputDocument: // inputDocument#1abfb575
//  default: panic(v)
//  }
type InputDocumentClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputDocumentClass

	// AsNotEmpty tries to map InputDocumentClass to InputDocument.
	AsNotEmpty() (*InputDocument, bool)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// AsNotEmpty tries to map InputDocumentClass to InputDocument.
func (i *InputDocumentEmpty) AsNotEmpty() (*InputDocument, bool) {
	return nil, false
}

// AsNotEmpty tries to map InputDocumentClass to InputDocument.
func (i *InputDocument) AsNotEmpty() (*InputDocument, bool) {
	return i, true
}

// DecodeInputDocument implements binary de-serialization for InputDocumentClass.
func DecodeInputDocument(buf *bin.Buffer) (InputDocumentClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputDocumentEmptyTypeID:
		// Decoding inputDocumentEmpty#72f0eaae.
		v := InputDocumentEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputDocumentClass: %w", err)
		}
		return &v, nil
	case InputDocumentTypeID:
		// Decoding inputDocument#1abfb575.
		v := InputDocument{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputDocumentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputDocumentClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputDocument boxes the InputDocumentClass providing a helper.
type InputDocumentBox struct {
	InputDocument InputDocumentClass
}

// Decode implements bin.Decoder for InputDocumentBox.
func (b *InputDocumentBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputDocumentBox to nil")
	}
	v, err := DecodeInputDocument(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputDocument = v
	return nil
}

// Encode implements bin.Encode for InputDocumentBox.
func (b *InputDocumentBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputDocument == nil {
		return fmt.Errorf("unable to encode InputDocumentClass as nil")
	}
	return b.InputDocument.Encode(buf)
}

// InputDocumentClassSlice is adapter for slice of InputDocumentClass.
type InputDocumentClassSlice []InputDocumentClass

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s InputDocumentClassSlice) AppendOnlyNotEmpty(to []*InputDocument) []*InputDocument {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s InputDocumentClassSlice) AsNotEmpty() (to []*InputDocument) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s InputDocumentClassSlice) FirstAsNotEmpty() (v *InputDocument, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s InputDocumentClassSlice) LastAsNotEmpty() (v *InputDocument, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// First returns first element of slice (if exists).
func (s InputDocumentClassSlice) First() (v InputDocumentClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputDocumentClassSlice) Last() (v InputDocumentClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputDocumentClassSlice) PopFirst() (v InputDocumentClass, ok bool) {
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
func (s *InputDocumentClassSlice) Pop() (v InputDocumentClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
