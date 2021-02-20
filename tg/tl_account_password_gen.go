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

// AccountPassword represents TL type `account.password#ad2641f8`.
// Configuration for two-factor authorization
//
// See https://core.telegram.org/constructor/account.password for reference.
type AccountPassword struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the user has a recovery method configured
	HasRecovery bool
	// Whether telegram passport¹ is enabled
	//
	// Links:
	//  1) https://core.telegram.org/passport
	HasSecureValues bool
	// Whether the user has a password
	HasPassword bool
	// The KDF algorithm for SRP two-factor authentication¹ of the current password
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	//
	// Use SetCurrentAlgo and GetCurrentAlgo helpers.
	CurrentAlgo PasswordKdfAlgoClass
	// Srp B param for SRP authorization¹
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	//
	// Use SetSRPB and GetSRPB helpers.
	SRPB []byte
	// Srp ID param for SRP authorization¹
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	//
	// Use SetSRPID and GetSRPID helpers.
	SRPID int64
	// Text hint for the password
	//
	// Use SetHint and GetHint helpers.
	Hint string
	// A password recovery email¹ with the specified pattern² is still awaiting verification
	//
	// Links:
	//  1) https://core.telegram.org/api/srp#email-verification
	//  2) https://core.telegram.org/api/pattern
	//
	// Use SetEmailUnconfirmedPattern and GetEmailUnconfirmedPattern helpers.
	EmailUnconfirmedPattern string
	// The KDF algorithm for SRP two-factor authentication¹ to use when creating new passwords
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	NewAlgo PasswordKdfAlgoClass
	// The KDF algorithm for telegram passport¹
	//
	// Links:
	//  1) https://core.telegram.org/passport
	NewSecureAlgo SecurePasswordKdfAlgoClass
	// Secure random string
	SecureRandom []byte
}

// AccountPasswordTypeID is TL type id of AccountPassword.
const AccountPasswordTypeID = 0xad2641f8

func (p *AccountPassword) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.HasRecovery == false) {
		return false
	}
	if !(p.HasSecureValues == false) {
		return false
	}
	if !(p.HasPassword == false) {
		return false
	}
	if !(p.CurrentAlgo == nil) {
		return false
	}
	if !(p.SRPB == nil) {
		return false
	}
	if !(p.SRPID == 0) {
		return false
	}
	if !(p.Hint == "") {
		return false
	}
	if !(p.EmailUnconfirmedPattern == "") {
		return false
	}
	if !(p.NewAlgo == nil) {
		return false
	}
	if !(p.NewSecureAlgo == nil) {
		return false
	}
	if !(p.SecureRandom == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *AccountPassword) String() string {
	if p == nil {
		return "AccountPassword(nil)"
	}
	type Alias AccountPassword
	return fmt.Sprintf("AccountPassword%+v", Alias(*p))
}

// FillFrom fills AccountPassword from given interface.
func (p *AccountPassword) FillFrom(from interface {
	GetHasRecovery() (value bool)
	GetHasSecureValues() (value bool)
	GetHasPassword() (value bool)
	GetCurrentAlgo() (value PasswordKdfAlgoClass, ok bool)
	GetSRPB() (value []byte, ok bool)
	GetSRPID() (value int64, ok bool)
	GetHint() (value string, ok bool)
	GetEmailUnconfirmedPattern() (value string, ok bool)
	GetNewAlgo() (value PasswordKdfAlgoClass)
	GetNewSecureAlgo() (value SecurePasswordKdfAlgoClass)
	GetSecureRandom() (value []byte)
}) {
	if val, ok := from.GetCurrentAlgo(); ok {
		p.CurrentAlgo = val
	}
	if val, ok := from.GetSRPB(); ok {
		p.SRPB = val
	}
	if val, ok := from.GetSRPID(); ok {
		p.SRPID = val
	}
	if val, ok := from.GetHint(); ok {
		p.Hint = val
	}
	if val, ok := from.GetEmailUnconfirmedPattern(); ok {
		p.EmailUnconfirmedPattern = val
	}
	p.NewAlgo = from.GetNewAlgo()
	p.NewSecureAlgo = from.GetNewSecureAlgo()
	p.SecureRandom = from.GetSecureRandom()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *AccountPassword) TypeID() uint32 {
	return AccountPasswordTypeID
}

// Encode implements bin.Encoder.
func (p *AccountPassword) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode account.password#ad2641f8 as nil")
	}
	b.PutID(AccountPasswordTypeID)
	if !(p.HasRecovery == false) {
		p.Flags.Set(0)
	}
	if !(p.HasSecureValues == false) {
		p.Flags.Set(1)
	}
	if !(p.HasPassword == false) {
		p.Flags.Set(2)
	}
	if !(p.CurrentAlgo == nil) {
		p.Flags.Set(2)
	}
	if !(p.SRPB == nil) {
		p.Flags.Set(2)
	}
	if !(p.SRPID == 0) {
		p.Flags.Set(2)
	}
	if !(p.Hint == "") {
		p.Flags.Set(3)
	}
	if !(p.EmailUnconfirmedPattern == "") {
		p.Flags.Set(4)
	}
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.password#ad2641f8: field flags: %w", err)
	}
	if p.Flags.Has(2) {
		if p.CurrentAlgo == nil {
			return fmt.Errorf("unable to encode account.password#ad2641f8: field current_algo is nil")
		}
		if err := p.CurrentAlgo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.password#ad2641f8: field current_algo: %w", err)
		}
	}
	if p.Flags.Has(2) {
		b.PutBytes(p.SRPB)
	}
	if p.Flags.Has(2) {
		b.PutLong(p.SRPID)
	}
	if p.Flags.Has(3) {
		b.PutString(p.Hint)
	}
	if p.Flags.Has(4) {
		b.PutString(p.EmailUnconfirmedPattern)
	}
	if p.NewAlgo == nil {
		return fmt.Errorf("unable to encode account.password#ad2641f8: field new_algo is nil")
	}
	if err := p.NewAlgo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.password#ad2641f8: field new_algo: %w", err)
	}
	if p.NewSecureAlgo == nil {
		return fmt.Errorf("unable to encode account.password#ad2641f8: field new_secure_algo is nil")
	}
	if err := p.NewSecureAlgo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.password#ad2641f8: field new_secure_algo: %w", err)
	}
	b.PutBytes(p.SecureRandom)
	return nil
}

// SetHasRecovery sets value of HasRecovery conditional field.
func (p *AccountPassword) SetHasRecovery(value bool) {
	if value {
		p.Flags.Set(0)
		p.HasRecovery = true
	} else {
		p.Flags.Unset(0)
		p.HasRecovery = false
	}
}

// GetHasRecovery returns value of HasRecovery conditional field.
func (p *AccountPassword) GetHasRecovery() (value bool) {
	return p.Flags.Has(0)
}

// SetHasSecureValues sets value of HasSecureValues conditional field.
func (p *AccountPassword) SetHasSecureValues(value bool) {
	if value {
		p.Flags.Set(1)
		p.HasSecureValues = true
	} else {
		p.Flags.Unset(1)
		p.HasSecureValues = false
	}
}

// GetHasSecureValues returns value of HasSecureValues conditional field.
func (p *AccountPassword) GetHasSecureValues() (value bool) {
	return p.Flags.Has(1)
}

// SetHasPassword sets value of HasPassword conditional field.
func (p *AccountPassword) SetHasPassword(value bool) {
	if value {
		p.Flags.Set(2)
		p.HasPassword = true
	} else {
		p.Flags.Unset(2)
		p.HasPassword = false
	}
}

// GetHasPassword returns value of HasPassword conditional field.
func (p *AccountPassword) GetHasPassword() (value bool) {
	return p.Flags.Has(2)
}

// SetCurrentAlgo sets value of CurrentAlgo conditional field.
func (p *AccountPassword) SetCurrentAlgo(value PasswordKdfAlgoClass) {
	p.Flags.Set(2)
	p.CurrentAlgo = value
}

// GetCurrentAlgo returns value of CurrentAlgo conditional field and
// boolean which is true if field was set.
func (p *AccountPassword) GetCurrentAlgo() (value PasswordKdfAlgoClass, ok bool) {
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.CurrentAlgo, true
}

// SetSRPB sets value of SRPB conditional field.
func (p *AccountPassword) SetSRPB(value []byte) {
	p.Flags.Set(2)
	p.SRPB = value
}

// GetSRPB returns value of SRPB conditional field and
// boolean which is true if field was set.
func (p *AccountPassword) GetSRPB() (value []byte, ok bool) {
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.SRPB, true
}

// SetSRPID sets value of SRPID conditional field.
func (p *AccountPassword) SetSRPID(value int64) {
	p.Flags.Set(2)
	p.SRPID = value
}

// GetSRPID returns value of SRPID conditional field and
// boolean which is true if field was set.
func (p *AccountPassword) GetSRPID() (value int64, ok bool) {
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.SRPID, true
}

// SetHint sets value of Hint conditional field.
func (p *AccountPassword) SetHint(value string) {
	p.Flags.Set(3)
	p.Hint = value
}

// GetHint returns value of Hint conditional field and
// boolean which is true if field was set.
func (p *AccountPassword) GetHint() (value string, ok bool) {
	if !p.Flags.Has(3) {
		return value, false
	}
	return p.Hint, true
}

// SetEmailUnconfirmedPattern sets value of EmailUnconfirmedPattern conditional field.
func (p *AccountPassword) SetEmailUnconfirmedPattern(value string) {
	p.Flags.Set(4)
	p.EmailUnconfirmedPattern = value
}

// GetEmailUnconfirmedPattern returns value of EmailUnconfirmedPattern conditional field and
// boolean which is true if field was set.
func (p *AccountPassword) GetEmailUnconfirmedPattern() (value string, ok bool) {
	if !p.Flags.Has(4) {
		return value, false
	}
	return p.EmailUnconfirmedPattern, true
}

// GetNewAlgo returns value of NewAlgo field.
func (p *AccountPassword) GetNewAlgo() (value PasswordKdfAlgoClass) {
	return p.NewAlgo
}

// GetNewSecureAlgo returns value of NewSecureAlgo field.
func (p *AccountPassword) GetNewSecureAlgo() (value SecurePasswordKdfAlgoClass) {
	return p.NewSecureAlgo
}

// GetSecureRandom returns value of SecureRandom field.
func (p *AccountPassword) GetSecureRandom() (value []byte) {
	return p.SecureRandom
}

// Decode implements bin.Decoder.
func (p *AccountPassword) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode account.password#ad2641f8 to nil")
	}
	if err := b.ConsumeID(AccountPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode account.password#ad2641f8: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.password#ad2641f8: field flags: %w", err)
		}
	}
	p.HasRecovery = p.Flags.Has(0)
	p.HasSecureValues = p.Flags.Has(1)
	p.HasPassword = p.Flags.Has(2)
	if p.Flags.Has(2) {
		value, err := DecodePasswordKdfAlgo(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.password#ad2641f8: field current_algo: %w", err)
		}
		p.CurrentAlgo = value
	}
	if p.Flags.Has(2) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode account.password#ad2641f8: field srp_B: %w", err)
		}
		p.SRPB = value
	}
	if p.Flags.Has(2) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.password#ad2641f8: field srp_id: %w", err)
		}
		p.SRPID = value
	}
	if p.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.password#ad2641f8: field hint: %w", err)
		}
		p.Hint = value
	}
	if p.Flags.Has(4) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.password#ad2641f8: field email_unconfirmed_pattern: %w", err)
		}
		p.EmailUnconfirmedPattern = value
	}
	{
		value, err := DecodePasswordKdfAlgo(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.password#ad2641f8: field new_algo: %w", err)
		}
		p.NewAlgo = value
	}
	{
		value, err := DecodeSecurePasswordKdfAlgo(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.password#ad2641f8: field new_secure_algo: %w", err)
		}
		p.NewSecureAlgo = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode account.password#ad2641f8: field secure_random: %w", err)
		}
		p.SecureRandom = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountPassword.
var (
	_ bin.Encoder = &AccountPassword{}
	_ bin.Decoder = &AccountPassword{}
)
