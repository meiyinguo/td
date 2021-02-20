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

// SecureValueHash represents TL type `secureValueHash#ed1ecdb0`.
// Secure value hash
//
// See https://core.telegram.org/constructor/secureValueHash for reference.
type SecureValueHash struct {
	// Secure value type
	Type SecureValueTypeClass
	// Hash
	Hash []byte
}

// SecureValueHashTypeID is TL type id of SecureValueHash.
const SecureValueHashTypeID = 0xed1ecdb0

func (s *SecureValueHash) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Type == nil) {
		return false
	}
	if !(s.Hash == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueHash) String() string {
	if s == nil {
		return "SecureValueHash(nil)"
	}
	type Alias SecureValueHash
	return fmt.Sprintf("SecureValueHash%+v", Alias(*s))
}

// FillFrom fills SecureValueHash from given interface.
func (s *SecureValueHash) FillFrom(from interface {
	GetType() (value SecureValueTypeClass)
	GetHash() (value []byte)
}) {
	s.Type = from.GetType()
	s.Hash = from.GetHash()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueHash) TypeID() uint32 {
	return SecureValueHashTypeID
}

// Encode implements bin.Encoder.
func (s *SecureValueHash) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueHash#ed1ecdb0 as nil")
	}
	b.PutID(SecureValueHashTypeID)
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureValueHash#ed1ecdb0: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValueHash#ed1ecdb0: field type: %w", err)
	}
	b.PutBytes(s.Hash)
	return nil
}

// GetType returns value of Type field.
func (s *SecureValueHash) GetType() (value SecureValueTypeClass) {
	return s.Type
}

// GetHash returns value of Hash field.
func (s *SecureValueHash) GetHash() (value []byte) {
	return s.Hash
}

// Decode implements bin.Decoder.
func (s *SecureValueHash) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueHash#ed1ecdb0 to nil")
	}
	if err := b.ConsumeID(SecureValueHashTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueHash#ed1ecdb0: %w", err)
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValueHash#ed1ecdb0: field type: %w", err)
		}
		s.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureValueHash#ed1ecdb0: field hash: %w", err)
		}
		s.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for SecureValueHash.
var (
	_ bin.Encoder = &SecureValueHash{}
	_ bin.Decoder = &SecureValueHash{}
)
