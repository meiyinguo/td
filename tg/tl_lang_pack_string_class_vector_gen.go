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

// LangPackStringClassVector is a box for Vector<LangPackString>
type LangPackStringClassVector struct {
	// Elements of Vector<LangPackString>
	Elems []LangPackStringClass
}

// LangPackStringClassVectorTypeID is TL type id of LangPackStringClassVector.
const LangPackStringClassVectorTypeID = bin.TypeVector

func (vec *LangPackStringClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *LangPackStringClassVector) String() string {
	if vec == nil {
		return "LangPackStringClassVector(nil)"
	}
	type Alias LangPackStringClassVector
	return fmt.Sprintf("LangPackStringClassVector%+v", Alias(*vec))
}

// FillFrom fills LangPackStringClassVector from given interface.
func (vec *LangPackStringClassVector) FillFrom(from interface {
	GetElems() (value []LangPackStringClass)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (vec *LangPackStringClassVector) TypeID() uint32 {
	return LangPackStringClassVectorTypeID
}

// Encode implements bin.Encoder.
func (vec *LangPackStringClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<LangPackString> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<LangPackString>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<LangPackString>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *LangPackStringClassVector) GetElems() (value []LangPackStringClass) {
	return vec.Elems
}

// MapElems returns field Elems wrapped in LangPackStringClassSlice helper.
func (vec *LangPackStringClassVector) MapElems() (value LangPackStringClassSlice) {
	return LangPackStringClassSlice(vec.Elems)
}

// Decode implements bin.Decoder.
func (vec *LangPackStringClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<LangPackString> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<LangPackString>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeLangPackString(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<LangPackString>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for LangPackStringClassVector.
var (
	_ bin.Encoder = &LangPackStringClassVector{}
	_ bin.Decoder = &LangPackStringClassVector{}
)
