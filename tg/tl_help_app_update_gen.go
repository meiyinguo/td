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

// HelpAppUpdate represents TL type `help.appUpdate#1da7158f`.
// An update is available for the application.
//
// See https://core.telegram.org/constructor/help.appUpdate for reference.
type HelpAppUpdate struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Unskippable, the new info must be shown to the user (with a popup or something else)
	CanNotSkip bool
	// Update ID
	ID int
	// New version name
	Version string
	// Text description of the update
	Text string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	Entities []MessageEntityClass
	// Application binary
	//
	// Use SetDocument and GetDocument helpers.
	Document DocumentClass
	// Application download URL
	//
	// Use SetURL and GetURL helpers.
	URL string
}

// HelpAppUpdateTypeID is TL type id of HelpAppUpdate.
const HelpAppUpdateTypeID = 0x1da7158f

func (a *HelpAppUpdate) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.CanNotSkip == false) {
		return false
	}
	if !(a.ID == 0) {
		return false
	}
	if !(a.Version == "") {
		return false
	}
	if !(a.Text == "") {
		return false
	}
	if !(a.Entities == nil) {
		return false
	}
	if !(a.Document == nil) {
		return false
	}
	if !(a.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *HelpAppUpdate) String() string {
	if a == nil {
		return "HelpAppUpdate(nil)"
	}
	type Alias HelpAppUpdate
	return fmt.Sprintf("HelpAppUpdate%+v", Alias(*a))
}

// FillFrom fills HelpAppUpdate from given interface.
func (a *HelpAppUpdate) FillFrom(from interface {
	GetCanNotSkip() (value bool)
	GetID() (value int)
	GetVersion() (value string)
	GetText() (value string)
	GetEntities() (value []MessageEntityClass)
	GetDocument() (value DocumentClass, ok bool)
	GetURL() (value string, ok bool)
}) {
	a.CanNotSkip = from.GetCanNotSkip()
	a.ID = from.GetID()
	a.Version = from.GetVersion()
	a.Text = from.GetText()
	a.Entities = from.GetEntities()
	if val, ok := from.GetDocument(); ok {
		a.Document = val
	}

	if val, ok := from.GetURL(); ok {
		a.URL = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpAppUpdate) TypeID() uint32 {
	return HelpAppUpdateTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpAppUpdate) TypeName() string {
	return "help.appUpdate"
}

// TypeInfo returns info about TL type.
func (a *HelpAppUpdate) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.appUpdate",
		ID:   HelpAppUpdateTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CanNotSkip",
			SchemaName: "can_not_skip",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Version",
			SchemaName: "version",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
		},
		{
			Name:       "Document",
			SchemaName: "document",
			Null:       !a.Flags.Has(1),
		},
		{
			Name:       "URL",
			SchemaName: "url",
			Null:       !a.Flags.Has(2),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *HelpAppUpdate) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode help.appUpdate#1da7158f as nil")
	}
	b.PutID(HelpAppUpdateTypeID)
	if !(a.CanNotSkip == false) {
		a.Flags.Set(0)
	}
	if !(a.Document == nil) {
		a.Flags.Set(1)
	}
	if !(a.URL == "") {
		a.Flags.Set(2)
	}
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.appUpdate#1da7158f: field flags: %w", err)
	}
	b.PutInt(a.ID)
	b.PutString(a.Version)
	b.PutString(a.Text)
	b.PutVectorHeader(len(a.Entities))
	for idx, v := range a.Entities {
		if v == nil {
			return fmt.Errorf("unable to encode help.appUpdate#1da7158f: field entities element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.appUpdate#1da7158f: field entities element with index %d: %w", idx, err)
		}
	}
	if a.Flags.Has(1) {
		if a.Document == nil {
			return fmt.Errorf("unable to encode help.appUpdate#1da7158f: field document is nil")
		}
		if err := a.Document.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.appUpdate#1da7158f: field document: %w", err)
		}
	}
	if a.Flags.Has(2) {
		b.PutString(a.URL)
	}
	return nil
}

// SetCanNotSkip sets value of CanNotSkip conditional field.
func (a *HelpAppUpdate) SetCanNotSkip(value bool) {
	if value {
		a.Flags.Set(0)
		a.CanNotSkip = true
	} else {
		a.Flags.Unset(0)
		a.CanNotSkip = false
	}
}

// GetCanNotSkip returns value of CanNotSkip conditional field.
func (a *HelpAppUpdate) GetCanNotSkip() (value bool) {
	return a.Flags.Has(0)
}

// GetID returns value of ID field.
func (a *HelpAppUpdate) GetID() (value int) {
	return a.ID
}

// GetVersion returns value of Version field.
func (a *HelpAppUpdate) GetVersion() (value string) {
	return a.Version
}

// GetText returns value of Text field.
func (a *HelpAppUpdate) GetText() (value string) {
	return a.Text
}

// GetEntities returns value of Entities field.
func (a *HelpAppUpdate) GetEntities() (value []MessageEntityClass) {
	return a.Entities
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (a *HelpAppUpdate) MapEntities() (value MessageEntityClassArray) {
	return MessageEntityClassArray(a.Entities)
}

// SetDocument sets value of Document conditional field.
func (a *HelpAppUpdate) SetDocument(value DocumentClass) {
	a.Flags.Set(1)
	a.Document = value
}

// GetDocument returns value of Document conditional field and
// boolean which is true if field was set.
func (a *HelpAppUpdate) GetDocument() (value DocumentClass, ok bool) {
	if !a.Flags.Has(1) {
		return value, false
	}
	return a.Document, true
}

// SetURL sets value of URL conditional field.
func (a *HelpAppUpdate) SetURL(value string) {
	a.Flags.Set(2)
	a.URL = value
}

// GetURL returns value of URL conditional field and
// boolean which is true if field was set.
func (a *HelpAppUpdate) GetURL() (value string, ok bool) {
	if !a.Flags.Has(2) {
		return value, false
	}
	return a.URL, true
}

// Decode implements bin.Decoder.
func (a *HelpAppUpdate) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode help.appUpdate#1da7158f to nil")
	}
	if err := b.ConsumeID(HelpAppUpdateTypeID); err != nil {
		return fmt.Errorf("unable to decode help.appUpdate#1da7158f: %w", err)
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field flags: %w", err)
		}
	}
	a.CanNotSkip = a.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field id: %w", err)
		}
		a.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field version: %w", err)
		}
		a.Version = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field text: %w", err)
		}
		a.Text = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field entities: %w", err)
			}
			a.Entities = append(a.Entities, value)
		}
	}
	if a.Flags.Has(1) {
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field document: %w", err)
		}
		a.Document = value
	}
	if a.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.appUpdate#1da7158f: field url: %w", err)
		}
		a.URL = value
	}
	return nil
}

// construct implements constructor of HelpAppUpdateClass.
func (a HelpAppUpdate) construct() HelpAppUpdateClass { return &a }

// Ensuring interfaces in compile-time for HelpAppUpdate.
var (
	_ bin.Encoder = &HelpAppUpdate{}
	_ bin.Decoder = &HelpAppUpdate{}

	_ HelpAppUpdateClass = &HelpAppUpdate{}
)

// HelpNoAppUpdate represents TL type `help.noAppUpdate#c45a6536`.
// No updates are available for the application.
//
// See https://core.telegram.org/constructor/help.noAppUpdate for reference.
type HelpNoAppUpdate struct {
}

// HelpNoAppUpdateTypeID is TL type id of HelpNoAppUpdate.
const HelpNoAppUpdateTypeID = 0xc45a6536

func (n *HelpNoAppUpdate) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *HelpNoAppUpdate) String() string {
	if n == nil {
		return "HelpNoAppUpdate(nil)"
	}
	type Alias HelpNoAppUpdate
	return fmt.Sprintf("HelpNoAppUpdate%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpNoAppUpdate) TypeID() uint32 {
	return HelpNoAppUpdateTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpNoAppUpdate) TypeName() string {
	return "help.noAppUpdate"
}

// TypeInfo returns info about TL type.
func (n *HelpNoAppUpdate) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.noAppUpdate",
		ID:   HelpNoAppUpdateTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (n *HelpNoAppUpdate) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode help.noAppUpdate#c45a6536 as nil")
	}
	b.PutID(HelpNoAppUpdateTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (n *HelpNoAppUpdate) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode help.noAppUpdate#c45a6536 to nil")
	}
	if err := b.ConsumeID(HelpNoAppUpdateTypeID); err != nil {
		return fmt.Errorf("unable to decode help.noAppUpdate#c45a6536: %w", err)
	}
	return nil
}

// construct implements constructor of HelpAppUpdateClass.
func (n HelpNoAppUpdate) construct() HelpAppUpdateClass { return &n }

// Ensuring interfaces in compile-time for HelpNoAppUpdate.
var (
	_ bin.Encoder = &HelpNoAppUpdate{}
	_ bin.Decoder = &HelpNoAppUpdate{}

	_ HelpAppUpdateClass = &HelpNoAppUpdate{}
)

// HelpAppUpdateClass represents help.AppUpdate generic type.
//
// See https://core.telegram.org/type/help.AppUpdate for reference.
//
// Example:
//  g, err := tg.DecodeHelpAppUpdate(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.HelpAppUpdate: // help.appUpdate#1da7158f
//  case *tg.HelpNoAppUpdate: // help.noAppUpdate#c45a6536
//  default: panic(v)
//  }
type HelpAppUpdateClass interface {
	bin.Encoder
	bin.Decoder
	construct() HelpAppUpdateClass

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

// DecodeHelpAppUpdate implements binary de-serialization for HelpAppUpdateClass.
func DecodeHelpAppUpdate(buf *bin.Buffer) (HelpAppUpdateClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpAppUpdateTypeID:
		// Decoding help.appUpdate#1da7158f.
		v := HelpAppUpdate{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpAppUpdateClass: %w", err)
		}
		return &v, nil
	case HelpNoAppUpdateTypeID:
		// Decoding help.noAppUpdate#c45a6536.
		v := HelpNoAppUpdate{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpAppUpdateClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpAppUpdateClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpAppUpdate boxes the HelpAppUpdateClass providing a helper.
type HelpAppUpdateBox struct {
	AppUpdate HelpAppUpdateClass
}

// Decode implements bin.Decoder for HelpAppUpdateBox.
func (b *HelpAppUpdateBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpAppUpdateBox to nil")
	}
	v, err := DecodeHelpAppUpdate(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.AppUpdate = v
	return nil
}

// Encode implements bin.Encode for HelpAppUpdateBox.
func (b *HelpAppUpdateBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.AppUpdate == nil {
		return fmt.Errorf("unable to encode HelpAppUpdateClass as nil")
	}
	return b.AppUpdate.Encode(buf)
}

// HelpAppUpdateClassArray is adapter for slice of HelpAppUpdateClass.
type HelpAppUpdateClassArray []HelpAppUpdateClass

// Sort sorts slice of HelpAppUpdateClass.
func (s HelpAppUpdateClassArray) Sort(less func(a, b HelpAppUpdateClass) bool) HelpAppUpdateClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpAppUpdateClass.
func (s HelpAppUpdateClassArray) SortStable(less func(a, b HelpAppUpdateClass) bool) HelpAppUpdateClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpAppUpdateClass.
func (s HelpAppUpdateClassArray) Retain(keep func(x HelpAppUpdateClass) bool) HelpAppUpdateClassArray {
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
func (s HelpAppUpdateClassArray) First() (v HelpAppUpdateClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpAppUpdateClassArray) Last() (v HelpAppUpdateClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpAppUpdateClassArray) PopFirst() (v HelpAppUpdateClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpAppUpdateClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpAppUpdateClassArray) Pop() (v HelpAppUpdateClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsHelpAppUpdate returns copy with only HelpAppUpdate constructors.
func (s HelpAppUpdateClassArray) AsHelpAppUpdate() (to HelpAppUpdateArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpAppUpdate)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// HelpAppUpdateArray is adapter for slice of HelpAppUpdate.
type HelpAppUpdateArray []HelpAppUpdate

// Sort sorts slice of HelpAppUpdate.
func (s HelpAppUpdateArray) Sort(less func(a, b HelpAppUpdate) bool) HelpAppUpdateArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpAppUpdate.
func (s HelpAppUpdateArray) SortStable(less func(a, b HelpAppUpdate) bool) HelpAppUpdateArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpAppUpdate.
func (s HelpAppUpdateArray) Retain(keep func(x HelpAppUpdate) bool) HelpAppUpdateArray {
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
func (s HelpAppUpdateArray) First() (v HelpAppUpdate, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpAppUpdateArray) Last() (v HelpAppUpdate, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpAppUpdateArray) PopFirst() (v HelpAppUpdate, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpAppUpdate
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpAppUpdateArray) Pop() (v HelpAppUpdate, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of HelpAppUpdate by ID.
func (s HelpAppUpdateArray) SortByID() HelpAppUpdateArray {
	return s.Sort(func(a, b HelpAppUpdate) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of HelpAppUpdate by ID.
func (s HelpAppUpdateArray) SortStableByID() HelpAppUpdateArray {
	return s.SortStable(func(a, b HelpAppUpdate) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s HelpAppUpdateArray) FillMap(to map[int]HelpAppUpdate) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s HelpAppUpdateArray) ToMap() map[int]HelpAppUpdate {
	r := make(map[int]HelpAppUpdate, len(s))
	s.FillMap(r)
	return r
}
