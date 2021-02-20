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

// SecureValueVector is a box for Vector<SecureValue>
type SecureValueVector struct {
	// Elements of Vector<SecureValue>
	Elems []SecureValue
}

// SecureValueVectorTypeID is TL type id of SecureValueVector.
const SecureValueVectorTypeID = bin.TypeVector

func (vec *SecureValueVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *SecureValueVector) String() string {
	if vec == nil {
		return "SecureValueVector(nil)"
	}
	type Alias SecureValueVector
	return fmt.Sprintf("SecureValueVector%+v", Alias(*vec))
}

// FillFrom fills SecureValueVector from given interface.
func (vec *SecureValueVector) FillFrom(from interface {
	GetElems() (value []SecureValue)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (vec *SecureValueVector) TypeID() uint32 {
	return SecureValueVectorTypeID
}

// Encode implements bin.Encoder.
func (vec *SecureValueVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<SecureValue> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<SecureValue>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *SecureValueVector) GetElems() (value []SecureValue) {
	return vec.Elems
}

// Decode implements bin.Decoder.
func (vec *SecureValueVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<SecureValue> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<SecureValue>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SecureValue
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<SecureValue>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for SecureValueVector.
var (
	_ bin.Encoder = &SecureValueVector{}
	_ bin.Decoder = &SecureValueVector{}
)
