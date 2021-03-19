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

// InvokeWithLayerRequest represents TL type `invokeWithLayer#da9b0d0d`.
// Invoke the specified query using the specified API layer¹
//
// Links:
//  1) https://core.telegram.org/api/invoking#layers
//
// See https://core.telegram.org/constructor/invokeWithLayer for reference.
type InvokeWithLayerRequest struct {
	// The layer to use
	Layer int
	// The query
	Query bin.Object
}

// InvokeWithLayerRequestTypeID is TL type id of InvokeWithLayerRequest.
const InvokeWithLayerRequestTypeID = 0xda9b0d0d

func (i *InvokeWithLayerRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Layer == 0) {
		return false
	}
	if !(i.Query == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InvokeWithLayerRequest) String() string {
	if i == nil {
		return "InvokeWithLayerRequest(nil)"
	}
	type Alias InvokeWithLayerRequest
	return fmt.Sprintf("InvokeWithLayerRequest%+v", Alias(*i))
}

// FillFrom fills InvokeWithLayerRequest from given interface.
func (i *InvokeWithLayerRequest) FillFrom(from interface {
	GetLayer() (value int)
	GetQuery() (value bin.Object)
}) {
	i.Layer = from.GetLayer()
	i.Query = from.GetQuery()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InvokeWithLayerRequest) TypeID() uint32 {
	return InvokeWithLayerRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*InvokeWithLayerRequest) TypeName() string {
	return "invokeWithLayer"
}

// TypeInfo returns info about TL type.
func (i *InvokeWithLayerRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "invokeWithLayer",
		ID:   InvokeWithLayerRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Layer",
			SchemaName: "layer",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InvokeWithLayerRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeWithLayer#da9b0d0d as nil")
	}
	b.PutID(InvokeWithLayerRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InvokeWithLayerRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeWithLayer#da9b0d0d as nil")
	}
	b.PutInt(i.Layer)
	if err := i.Query.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invokeWithLayer#da9b0d0d: field query: %w", err)
	}
	return nil
}

// GetLayer returns value of Layer field.
func (i *InvokeWithLayerRequest) GetLayer() (value int) {
	return i.Layer
}

// GetQuery returns value of Query field.
func (i *InvokeWithLayerRequest) GetQuery() (value bin.Object) {
	return i.Query
}

// Decode implements bin.Decoder.
func (i *InvokeWithLayerRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeWithLayer#da9b0d0d to nil")
	}
	if err := b.ConsumeID(InvokeWithLayerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode invokeWithLayer#da9b0d0d: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InvokeWithLayerRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeWithLayer#da9b0d0d to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode invokeWithLayer#da9b0d0d: field layer: %w", err)
		}
		i.Layer = value
	}
	{
		if err := i.Query.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invokeWithLayer#da9b0d0d: field query: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for InvokeWithLayerRequest.
var (
	_ bin.Encoder     = &InvokeWithLayerRequest{}
	_ bin.Decoder     = &InvokeWithLayerRequest{}
	_ bin.BareEncoder = &InvokeWithLayerRequest{}
	_ bin.BareDecoder = &InvokeWithLayerRequest{}
)
