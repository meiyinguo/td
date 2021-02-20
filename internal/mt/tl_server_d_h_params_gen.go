// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// ServerDHParamsFail represents TL type `server_DH_params_fail#79cb045d`.
type ServerDHParamsFail struct {
	// Nonce field of ServerDHParamsFail.
	Nonce bin.Int128
	// ServerNonce field of ServerDHParamsFail.
	ServerNonce bin.Int128
	// NewNonceHash field of ServerDHParamsFail.
	NewNonceHash bin.Int128
}

// ServerDHParamsFailTypeID is TL type id of ServerDHParamsFail.
const ServerDHParamsFailTypeID = 0x79cb045d

func (s *ServerDHParamsFail) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Nonce == bin.Int128{}) {
		return false
	}
	if !(s.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(s.NewNonceHash == bin.Int128{}) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ServerDHParamsFail) String() string {
	if s == nil {
		return "ServerDHParamsFail(nil)"
	}
	type Alias ServerDHParamsFail
	return fmt.Sprintf("ServerDHParamsFail%+v", Alias(*s))
}

// FillFrom fills ServerDHParamsFail from given interface.
func (s *ServerDHParamsFail) FillFrom(from interface {
	GetNonce() (value bin.Int128)
	GetServerNonce() (value bin.Int128)
	GetNewNonceHash() (value bin.Int128)
}) {
	s.Nonce = from.GetNonce()
	s.ServerNonce = from.GetServerNonce()
	s.NewNonceHash = from.GetNewNonceHash()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *ServerDHParamsFail) TypeID() uint32 {
	return ServerDHParamsFailTypeID
}

// Encode implements bin.Encoder.
func (s *ServerDHParamsFail) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode server_DH_params_fail#79cb045d as nil")
	}
	b.PutID(ServerDHParamsFailTypeID)
	b.PutInt128(s.Nonce)
	b.PutInt128(s.ServerNonce)
	b.PutInt128(s.NewNonceHash)
	return nil
}

// GetNonce returns value of Nonce field.
func (s *ServerDHParamsFail) GetNonce() (value bin.Int128) {
	return s.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (s *ServerDHParamsFail) GetServerNonce() (value bin.Int128) {
	return s.ServerNonce
}

// GetNewNonceHash returns value of NewNonceHash field.
func (s *ServerDHParamsFail) GetNewNonceHash() (value bin.Int128) {
	return s.NewNonceHash
}

// Decode implements bin.Decoder.
func (s *ServerDHParamsFail) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode server_DH_params_fail#79cb045d to nil")
	}
	if err := b.ConsumeID(ServerDHParamsFailTypeID); err != nil {
		return fmt.Errorf("unable to decode server_DH_params_fail#79cb045d: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_fail#79cb045d: field nonce: %w", err)
		}
		s.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_fail#79cb045d: field server_nonce: %w", err)
		}
		s.ServerNonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_fail#79cb045d: field new_nonce_hash: %w", err)
		}
		s.NewNonceHash = value
	}
	return nil
}

// construct implements constructor of ServerDHParamsClass.
func (s ServerDHParamsFail) construct() ServerDHParamsClass { return &s }

// Ensuring interfaces in compile-time for ServerDHParamsFail.
var (
	_ bin.Encoder = &ServerDHParamsFail{}
	_ bin.Decoder = &ServerDHParamsFail{}

	_ ServerDHParamsClass = &ServerDHParamsFail{}
)

// ServerDHParamsOk represents TL type `server_DH_params_ok#d0e8075c`.
type ServerDHParamsOk struct {
	// Nonce field of ServerDHParamsOk.
	Nonce bin.Int128
	// ServerNonce field of ServerDHParamsOk.
	ServerNonce bin.Int128
	// EncryptedAnswer field of ServerDHParamsOk.
	EncryptedAnswer []byte
}

// ServerDHParamsOkTypeID is TL type id of ServerDHParamsOk.
const ServerDHParamsOkTypeID = 0xd0e8075c

func (s *ServerDHParamsOk) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Nonce == bin.Int128{}) {
		return false
	}
	if !(s.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(s.EncryptedAnswer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ServerDHParamsOk) String() string {
	if s == nil {
		return "ServerDHParamsOk(nil)"
	}
	type Alias ServerDHParamsOk
	return fmt.Sprintf("ServerDHParamsOk%+v", Alias(*s))
}

// FillFrom fills ServerDHParamsOk from given interface.
func (s *ServerDHParamsOk) FillFrom(from interface {
	GetNonce() (value bin.Int128)
	GetServerNonce() (value bin.Int128)
	GetEncryptedAnswer() (value []byte)
}) {
	s.Nonce = from.GetNonce()
	s.ServerNonce = from.GetServerNonce()
	s.EncryptedAnswer = from.GetEncryptedAnswer()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *ServerDHParamsOk) TypeID() uint32 {
	return ServerDHParamsOkTypeID
}

// Encode implements bin.Encoder.
func (s *ServerDHParamsOk) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode server_DH_params_ok#d0e8075c as nil")
	}
	b.PutID(ServerDHParamsOkTypeID)
	b.PutInt128(s.Nonce)
	b.PutInt128(s.ServerNonce)
	b.PutBytes(s.EncryptedAnswer)
	return nil
}

// GetNonce returns value of Nonce field.
func (s *ServerDHParamsOk) GetNonce() (value bin.Int128) {
	return s.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (s *ServerDHParamsOk) GetServerNonce() (value bin.Int128) {
	return s.ServerNonce
}

// GetEncryptedAnswer returns value of EncryptedAnswer field.
func (s *ServerDHParamsOk) GetEncryptedAnswer() (value []byte) {
	return s.EncryptedAnswer
}

// Decode implements bin.Decoder.
func (s *ServerDHParamsOk) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode server_DH_params_ok#d0e8075c to nil")
	}
	if err := b.ConsumeID(ServerDHParamsOkTypeID); err != nil {
		return fmt.Errorf("unable to decode server_DH_params_ok#d0e8075c: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_ok#d0e8075c: field nonce: %w", err)
		}
		s.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_ok#d0e8075c: field server_nonce: %w", err)
		}
		s.ServerNonce = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_ok#d0e8075c: field encrypted_answer: %w", err)
		}
		s.EncryptedAnswer = value
	}
	return nil
}

// construct implements constructor of ServerDHParamsClass.
func (s ServerDHParamsOk) construct() ServerDHParamsClass { return &s }

// Ensuring interfaces in compile-time for ServerDHParamsOk.
var (
	_ bin.Encoder = &ServerDHParamsOk{}
	_ bin.Decoder = &ServerDHParamsOk{}

	_ ServerDHParamsClass = &ServerDHParamsOk{}
)

// ServerDHParamsClass represents Server_DH_Params generic type.
//
// Example:
//  g, err := mt.DecodeServerDHParams(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *mt.ServerDHParamsFail: // server_DH_params_fail#79cb045d
//  case *mt.ServerDHParamsOk: // server_DH_params_ok#d0e8075c
//  default: panic(v)
//  }
type ServerDHParamsClass interface {
	bin.Encoder
	bin.Decoder
	construct() ServerDHParamsClass

	// Nonce field of ServerDHParamsFail.
	GetNonce() (value bin.Int128)
	// ServerNonce field of ServerDHParamsFail.
	GetServerNonce() (value bin.Int128)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeServerDHParams implements binary de-serialization for ServerDHParamsClass.
func DecodeServerDHParams(buf *bin.Buffer) (ServerDHParamsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ServerDHParamsFailTypeID:
		// Decoding server_DH_params_fail#79cb045d.
		v := ServerDHParamsFail{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ServerDHParamsClass: %w", err)
		}
		return &v, nil
	case ServerDHParamsOkTypeID:
		// Decoding server_DH_params_ok#d0e8075c.
		v := ServerDHParamsOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ServerDHParamsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ServerDHParamsClass: %w", bin.NewUnexpectedID(id))
	}
}

// ServerDHParams boxes the ServerDHParamsClass providing a helper.
type ServerDHParamsBox struct {
	Server_DH_Params ServerDHParamsClass
}

// Decode implements bin.Decoder for ServerDHParamsBox.
func (b *ServerDHParamsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ServerDHParamsBox to nil")
	}
	v, err := DecodeServerDHParams(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Server_DH_Params = v
	return nil
}

// Encode implements bin.Encode for ServerDHParamsBox.
func (b *ServerDHParamsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Server_DH_Params == nil {
		return fmt.Errorf("unable to encode ServerDHParamsClass as nil")
	}
	return b.Server_DH_Params.Encode(buf)
}

// ServerDHParamsClassSlice is adapter for slice of ServerDHParamsClass.
type ServerDHParamsClassSlice []ServerDHParamsClass

// First returns first element of slice (if exists).
func (s ServerDHParamsClassSlice) First() (v ServerDHParamsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ServerDHParamsClassSlice) Last() (v ServerDHParamsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ServerDHParamsClassSlice) PopFirst() (v ServerDHParamsClass, ok bool) {
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
func (s *ServerDHParamsClassSlice) Pop() (v ServerDHParamsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
