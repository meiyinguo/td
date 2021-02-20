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

// AccountAuthorizations represents TL type `account.authorizations#1250abde`.
// Logged-in sessions
//
// See https://core.telegram.org/constructor/account.authorizations for reference.
type AccountAuthorizations struct {
	// Logged-in sessions
	Authorizations []Authorization
}

// AccountAuthorizationsTypeID is TL type id of AccountAuthorizations.
const AccountAuthorizationsTypeID = 0x1250abde

func (a *AccountAuthorizations) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Authorizations == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AccountAuthorizations) String() string {
	if a == nil {
		return "AccountAuthorizations(nil)"
	}
	type Alias AccountAuthorizations
	return fmt.Sprintf("AccountAuthorizations%+v", Alias(*a))
}

// FillFrom fills AccountAuthorizations from given interface.
func (a *AccountAuthorizations) FillFrom(from interface {
	GetAuthorizations() (value []Authorization)
}) {
	a.Authorizations = from.GetAuthorizations()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (a *AccountAuthorizations) TypeID() uint32 {
	return AccountAuthorizationsTypeID
}

// Encode implements bin.Encoder.
func (a *AccountAuthorizations) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode account.authorizations#1250abde as nil")
	}
	b.PutID(AccountAuthorizationsTypeID)
	b.PutVectorHeader(len(a.Authorizations))
	for idx, v := range a.Authorizations {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.authorizations#1250abde: field authorizations element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetAuthorizations returns value of Authorizations field.
func (a *AccountAuthorizations) GetAuthorizations() (value []Authorization) {
	return a.Authorizations
}

// Decode implements bin.Decoder.
func (a *AccountAuthorizations) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode account.authorizations#1250abde to nil")
	}
	if err := b.ConsumeID(AccountAuthorizationsTypeID); err != nil {
		return fmt.Errorf("unable to decode account.authorizations#1250abde: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.authorizations#1250abde: field authorizations: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Authorization
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode account.authorizations#1250abde: field authorizations: %w", err)
			}
			a.Authorizations = append(a.Authorizations, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountAuthorizations.
var (
	_ bin.Encoder = &AccountAuthorizations{}
	_ bin.Decoder = &AccountAuthorizations{}
)
