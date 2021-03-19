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

// HelpPassportConfigNotModified represents TL type `help.passportConfigNotModified#bfb9f457`.
// Password configuration not modified
//
// See https://core.telegram.org/constructor/help.passportConfigNotModified for reference.
type HelpPassportConfigNotModified struct {
}

// HelpPassportConfigNotModifiedTypeID is TL type id of HelpPassportConfigNotModified.
const HelpPassportConfigNotModifiedTypeID = 0xbfb9f457

func (p *HelpPassportConfigNotModified) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *HelpPassportConfigNotModified) String() string {
	if p == nil {
		return "HelpPassportConfigNotModified(nil)"
	}
	type Alias HelpPassportConfigNotModified
	return fmt.Sprintf("HelpPassportConfigNotModified%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpPassportConfigNotModified) TypeID() uint32 {
	return HelpPassportConfigNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpPassportConfigNotModified) TypeName() string {
	return "help.passportConfigNotModified"
}

// TypeInfo returns info about TL type.
func (p *HelpPassportConfigNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.passportConfigNotModified",
		ID:   HelpPassportConfigNotModifiedTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (p *HelpPassportConfigNotModified) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.passportConfigNotModified#bfb9f457 as nil")
	}
	b.PutID(HelpPassportConfigNotModifiedTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *HelpPassportConfigNotModified) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.passportConfigNotModified#bfb9f457 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *HelpPassportConfigNotModified) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.passportConfigNotModified#bfb9f457 to nil")
	}
	if err := b.ConsumeID(HelpPassportConfigNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode help.passportConfigNotModified#bfb9f457: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *HelpPassportConfigNotModified) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.passportConfigNotModified#bfb9f457 to nil")
	}
	return nil
}

// construct implements constructor of HelpPassportConfigClass.
func (p HelpPassportConfigNotModified) construct() HelpPassportConfigClass { return &p }

// Ensuring interfaces in compile-time for HelpPassportConfigNotModified.
var (
	_ bin.Encoder     = &HelpPassportConfigNotModified{}
	_ bin.Decoder     = &HelpPassportConfigNotModified{}
	_ bin.BareEncoder = &HelpPassportConfigNotModified{}
	_ bin.BareDecoder = &HelpPassportConfigNotModified{}

	_ HelpPassportConfigClass = &HelpPassportConfigNotModified{}
)

// HelpPassportConfig represents TL type `help.passportConfig#a098d6af`.
// Telegram passport¹ configuration
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/constructor/help.passportConfig for reference.
type HelpPassportConfig struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// Localization
	CountriesLangs DataJSON
}

// HelpPassportConfigTypeID is TL type id of HelpPassportConfig.
const HelpPassportConfigTypeID = 0xa098d6af

func (p *HelpPassportConfig) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Hash == 0) {
		return false
	}
	if !(p.CountriesLangs.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *HelpPassportConfig) String() string {
	if p == nil {
		return "HelpPassportConfig(nil)"
	}
	type Alias HelpPassportConfig
	return fmt.Sprintf("HelpPassportConfig%+v", Alias(*p))
}

// FillFrom fills HelpPassportConfig from given interface.
func (p *HelpPassportConfig) FillFrom(from interface {
	GetHash() (value int)
	GetCountriesLangs() (value DataJSON)
}) {
	p.Hash = from.GetHash()
	p.CountriesLangs = from.GetCountriesLangs()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpPassportConfig) TypeID() uint32 {
	return HelpPassportConfigTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpPassportConfig) TypeName() string {
	return "help.passportConfig"
}

// TypeInfo returns info about TL type.
func (p *HelpPassportConfig) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.passportConfig",
		ID:   HelpPassportConfigTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "CountriesLangs",
			SchemaName: "countries_langs",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *HelpPassportConfig) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.passportConfig#a098d6af as nil")
	}
	b.PutID(HelpPassportConfigTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *HelpPassportConfig) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.passportConfig#a098d6af as nil")
	}
	b.PutInt(p.Hash)
	if err := p.CountriesLangs.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.passportConfig#a098d6af: field countries_langs: %w", err)
	}
	return nil
}

// GetHash returns value of Hash field.
func (p *HelpPassportConfig) GetHash() (value int) {
	return p.Hash
}

// GetCountriesLangs returns value of CountriesLangs field.
func (p *HelpPassportConfig) GetCountriesLangs() (value DataJSON) {
	return p.CountriesLangs
}

// Decode implements bin.Decoder.
func (p *HelpPassportConfig) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.passportConfig#a098d6af to nil")
	}
	if err := b.ConsumeID(HelpPassportConfigTypeID); err != nil {
		return fmt.Errorf("unable to decode help.passportConfig#a098d6af: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *HelpPassportConfig) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.passportConfig#a098d6af to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.passportConfig#a098d6af: field hash: %w", err)
		}
		p.Hash = value
	}
	{
		if err := p.CountriesLangs.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.passportConfig#a098d6af: field countries_langs: %w", err)
		}
	}
	return nil
}

// construct implements constructor of HelpPassportConfigClass.
func (p HelpPassportConfig) construct() HelpPassportConfigClass { return &p }

// Ensuring interfaces in compile-time for HelpPassportConfig.
var (
	_ bin.Encoder     = &HelpPassportConfig{}
	_ bin.Decoder     = &HelpPassportConfig{}
	_ bin.BareEncoder = &HelpPassportConfig{}
	_ bin.BareDecoder = &HelpPassportConfig{}

	_ HelpPassportConfigClass = &HelpPassportConfig{}
)

// HelpPassportConfigClass represents help.PassportConfig generic type.
//
// See https://core.telegram.org/type/help.PassportConfig for reference.
//
// Example:
//  g, err := tg.DecodeHelpPassportConfig(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.HelpPassportConfigNotModified: // help.passportConfigNotModified#bfb9f457
//  case *tg.HelpPassportConfig: // help.passportConfig#a098d6af
//  default: panic(v)
//  }
type HelpPassportConfigClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() HelpPassportConfigClass

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

	// AsModified tries to map HelpPassportConfigClass to HelpPassportConfig.
	AsModified() (*HelpPassportConfig, bool)
}

// AsModified tries to map HelpPassportConfigNotModified to HelpPassportConfig.
func (p *HelpPassportConfigNotModified) AsModified() (*HelpPassportConfig, bool) {
	return nil, false
}

// AsModified tries to map HelpPassportConfig to HelpPassportConfig.
func (p *HelpPassportConfig) AsModified() (*HelpPassportConfig, bool) {
	return p, true
}

// DecodeHelpPassportConfig implements binary de-serialization for HelpPassportConfigClass.
func DecodeHelpPassportConfig(buf *bin.Buffer) (HelpPassportConfigClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpPassportConfigNotModifiedTypeID:
		// Decoding help.passportConfigNotModified#bfb9f457.
		v := HelpPassportConfigNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpPassportConfigClass: %w", err)
		}
		return &v, nil
	case HelpPassportConfigTypeID:
		// Decoding help.passportConfig#a098d6af.
		v := HelpPassportConfig{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpPassportConfigClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpPassportConfigClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpPassportConfig boxes the HelpPassportConfigClass providing a helper.
type HelpPassportConfigBox struct {
	PassportConfig HelpPassportConfigClass
}

// Decode implements bin.Decoder for HelpPassportConfigBox.
func (b *HelpPassportConfigBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpPassportConfigBox to nil")
	}
	v, err := DecodeHelpPassportConfig(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PassportConfig = v
	return nil
}

// Encode implements bin.Encode for HelpPassportConfigBox.
func (b *HelpPassportConfigBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PassportConfig == nil {
		return fmt.Errorf("unable to encode HelpPassportConfigClass as nil")
	}
	return b.PassportConfig.Encode(buf)
}

// HelpPassportConfigClassArray is adapter for slice of HelpPassportConfigClass.
type HelpPassportConfigClassArray []HelpPassportConfigClass

// Sort sorts slice of HelpPassportConfigClass.
func (s HelpPassportConfigClassArray) Sort(less func(a, b HelpPassportConfigClass) bool) HelpPassportConfigClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpPassportConfigClass.
func (s HelpPassportConfigClassArray) SortStable(less func(a, b HelpPassportConfigClass) bool) HelpPassportConfigClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpPassportConfigClass.
func (s HelpPassportConfigClassArray) Retain(keep func(x HelpPassportConfigClass) bool) HelpPassportConfigClassArray {
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
func (s HelpPassportConfigClassArray) First() (v HelpPassportConfigClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpPassportConfigClassArray) Last() (v HelpPassportConfigClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpPassportConfigClassArray) PopFirst() (v HelpPassportConfigClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpPassportConfigClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpPassportConfigClassArray) Pop() (v HelpPassportConfigClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsHelpPassportConfig returns copy with only HelpPassportConfig constructors.
func (s HelpPassportConfigClassArray) AsHelpPassportConfig() (to HelpPassportConfigArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpPassportConfig)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s HelpPassportConfigClassArray) AppendOnlyModified(to []*HelpPassportConfig) []*HelpPassportConfig {
	for _, elem := range s {
		value, ok := elem.AsModified()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsModified returns copy with only Modified constructors.
func (s HelpPassportConfigClassArray) AsModified() (to []*HelpPassportConfig) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s HelpPassportConfigClassArray) FirstAsModified() (v *HelpPassportConfig, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s HelpPassportConfigClassArray) LastAsModified() (v *HelpPassportConfig, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *HelpPassportConfigClassArray) PopFirstAsModified() (v *HelpPassportConfig, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *HelpPassportConfigClassArray) PopAsModified() (v *HelpPassportConfig, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// HelpPassportConfigArray is adapter for slice of HelpPassportConfig.
type HelpPassportConfigArray []HelpPassportConfig

// Sort sorts slice of HelpPassportConfig.
func (s HelpPassportConfigArray) Sort(less func(a, b HelpPassportConfig) bool) HelpPassportConfigArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpPassportConfig.
func (s HelpPassportConfigArray) SortStable(less func(a, b HelpPassportConfig) bool) HelpPassportConfigArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpPassportConfig.
func (s HelpPassportConfigArray) Retain(keep func(x HelpPassportConfig) bool) HelpPassportConfigArray {
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
func (s HelpPassportConfigArray) First() (v HelpPassportConfig, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpPassportConfigArray) Last() (v HelpPassportConfig, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpPassportConfigArray) PopFirst() (v HelpPassportConfig, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpPassportConfig
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpPassportConfigArray) Pop() (v HelpPassportConfig, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
