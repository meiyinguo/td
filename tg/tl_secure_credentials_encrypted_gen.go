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

// SecureCredentialsEncrypted represents TL type `secureCredentialsEncrypted#33f0ea47`.
// Encrypted credentials required to decrypt telegram passport¹ data.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/constructor/secureCredentialsEncrypted for reference.
type SecureCredentialsEncrypted struct {
	// Encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication, as described in decrypting data »¹
	//
	// Links:
	//  1) https://core.telegram.org/passport#decrypting-data
	Data []byte
	// Data hash for data authentication as described in decrypting data »¹
	//
	// Links:
	//  1) https://core.telegram.org/passport#decrypting-data
	Hash []byte
	// Secret, encrypted with the bot's public RSA key, required for data decryption as described in decrypting data »¹
	//
	// Links:
	//  1) https://core.telegram.org/passport#decrypting-data
	Secret []byte
}

// SecureCredentialsEncryptedTypeID is TL type id of SecureCredentialsEncrypted.
const SecureCredentialsEncryptedTypeID = 0x33f0ea47

func (s *SecureCredentialsEncrypted) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Data == nil) {
		return false
	}
	if !(s.Hash == nil) {
		return false
	}
	if !(s.Secret == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureCredentialsEncrypted) String() string {
	if s == nil {
		return "SecureCredentialsEncrypted(nil)"
	}
	type Alias SecureCredentialsEncrypted
	return fmt.Sprintf("SecureCredentialsEncrypted%+v", Alias(*s))
}

// FillFrom fills SecureCredentialsEncrypted from given interface.
func (s *SecureCredentialsEncrypted) FillFrom(from interface {
	GetData() (value []byte)
	GetHash() (value []byte)
	GetSecret() (value []byte)
}) {
	s.Data = from.GetData()
	s.Hash = from.GetHash()
	s.Secret = from.GetSecret()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SecureCredentialsEncrypted) TypeID() uint32 {
	return SecureCredentialsEncryptedTypeID
}

// TypeName returns name of type in TL schema.
func (*SecureCredentialsEncrypted) TypeName() string {
	return "secureCredentialsEncrypted"
}

// TypeInfo returns info about TL type.
func (s *SecureCredentialsEncrypted) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "secureCredentialsEncrypted",
		ID:   SecureCredentialsEncryptedTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Data",
			SchemaName: "data",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Secret",
			SchemaName: "secret",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SecureCredentialsEncrypted) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureCredentialsEncrypted#33f0ea47 as nil")
	}
	b.PutID(SecureCredentialsEncryptedTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SecureCredentialsEncrypted) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureCredentialsEncrypted#33f0ea47 as nil")
	}
	b.PutBytes(s.Data)
	b.PutBytes(s.Hash)
	b.PutBytes(s.Secret)
	return nil
}

// GetData returns value of Data field.
func (s *SecureCredentialsEncrypted) GetData() (value []byte) {
	return s.Data
}

// GetHash returns value of Hash field.
func (s *SecureCredentialsEncrypted) GetHash() (value []byte) {
	return s.Hash
}

// GetSecret returns value of Secret field.
func (s *SecureCredentialsEncrypted) GetSecret() (value []byte) {
	return s.Secret
}

// Decode implements bin.Decoder.
func (s *SecureCredentialsEncrypted) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureCredentialsEncrypted#33f0ea47 to nil")
	}
	if err := b.ConsumeID(SecureCredentialsEncryptedTypeID); err != nil {
		return fmt.Errorf("unable to decode secureCredentialsEncrypted#33f0ea47: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SecureCredentialsEncrypted) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureCredentialsEncrypted#33f0ea47 to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureCredentialsEncrypted#33f0ea47: field data: %w", err)
		}
		s.Data = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureCredentialsEncrypted#33f0ea47: field hash: %w", err)
		}
		s.Hash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureCredentialsEncrypted#33f0ea47: field secret: %w", err)
		}
		s.Secret = value
	}
	return nil
}

// Ensuring interfaces in compile-time for SecureCredentialsEncrypted.
var (
	_ bin.Encoder     = &SecureCredentialsEncrypted{}
	_ bin.Decoder     = &SecureCredentialsEncrypted{}
	_ bin.BareEncoder = &SecureCredentialsEncrypted{}
	_ bin.BareDecoder = &SecureCredentialsEncrypted{}
)
