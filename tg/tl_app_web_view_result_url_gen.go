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
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// AppWebViewResultURL represents TL type `appWebViewResultUrl#3c1b4f0d`.
//
// See https://core.telegram.org/constructor/appWebViewResultUrl for reference.
type AppWebViewResultURL struct {
	// URL field of AppWebViewResultURL.
	URL string
}

// AppWebViewResultURLTypeID is TL type id of AppWebViewResultURL.
const AppWebViewResultURLTypeID = 0x3c1b4f0d

// Ensuring interfaces in compile-time for AppWebViewResultURL.
var (
	_ bin.Encoder     = &AppWebViewResultURL{}
	_ bin.Decoder     = &AppWebViewResultURL{}
	_ bin.BareEncoder = &AppWebViewResultURL{}
	_ bin.BareDecoder = &AppWebViewResultURL{}
)

func (a *AppWebViewResultURL) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AppWebViewResultURL) String() string {
	if a == nil {
		return "AppWebViewResultURL(nil)"
	}
	type Alias AppWebViewResultURL
	return fmt.Sprintf("AppWebViewResultURL%+v", Alias(*a))
}

// FillFrom fills AppWebViewResultURL from given interface.
func (a *AppWebViewResultURL) FillFrom(from interface {
	GetURL() (value string)
}) {
	a.URL = from.GetURL()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AppWebViewResultURL) TypeID() uint32 {
	return AppWebViewResultURLTypeID
}

// TypeName returns name of type in TL schema.
func (*AppWebViewResultURL) TypeName() string {
	return "appWebViewResultUrl"
}

// TypeInfo returns info about TL type.
func (a *AppWebViewResultURL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "appWebViewResultUrl",
		ID:   AppWebViewResultURLTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AppWebViewResultURL) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode appWebViewResultUrl#3c1b4f0d as nil")
	}
	b.PutID(AppWebViewResultURLTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AppWebViewResultURL) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode appWebViewResultUrl#3c1b4f0d as nil")
	}
	b.PutString(a.URL)
	return nil
}

// Decode implements bin.Decoder.
func (a *AppWebViewResultURL) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode appWebViewResultUrl#3c1b4f0d to nil")
	}
	if err := b.ConsumeID(AppWebViewResultURLTypeID); err != nil {
		return fmt.Errorf("unable to decode appWebViewResultUrl#3c1b4f0d: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AppWebViewResultURL) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode appWebViewResultUrl#3c1b4f0d to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode appWebViewResultUrl#3c1b4f0d: field url: %w", err)
		}
		a.URL = value
	}
	return nil
}

// GetURL returns value of URL field.
func (a *AppWebViewResultURL) GetURL() (value string) {
	if a == nil {
		return
	}
	return a.URL
}
