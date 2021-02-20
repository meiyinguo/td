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

// AccountSendVerifyPhoneCodeRequest represents TL type `account.sendVerifyPhoneCode#a5a356f9`.
// Send the verification phone code for telegram passport¹.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/account.sendVerifyPhoneCode for reference.
type AccountSendVerifyPhoneCodeRequest struct {
	// The phone number to verify
	PhoneNumber string
	// Phone code settings
	Settings CodeSettings
}

// AccountSendVerifyPhoneCodeRequestTypeID is TL type id of AccountSendVerifyPhoneCodeRequest.
const AccountSendVerifyPhoneCodeRequestTypeID = 0xa5a356f9

func (s *AccountSendVerifyPhoneCodeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.PhoneNumber == "") {
		return false
	}
	if !(s.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSendVerifyPhoneCodeRequest) String() string {
	if s == nil {
		return "AccountSendVerifyPhoneCodeRequest(nil)"
	}
	type Alias AccountSendVerifyPhoneCodeRequest
	return fmt.Sprintf("AccountSendVerifyPhoneCodeRequest%+v", Alias(*s))
}

// FillFrom fills AccountSendVerifyPhoneCodeRequest from given interface.
func (s *AccountSendVerifyPhoneCodeRequest) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetSettings() (value CodeSettings)
}) {
	s.PhoneNumber = from.GetPhoneNumber()
	s.Settings = from.GetSettings()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *AccountSendVerifyPhoneCodeRequest) TypeID() uint32 {
	return AccountSendVerifyPhoneCodeRequestTypeID
}

// Encode implements bin.Encoder.
func (s *AccountSendVerifyPhoneCodeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sendVerifyPhoneCode#a5a356f9 as nil")
	}
	b.PutID(AccountSendVerifyPhoneCodeRequestTypeID)
	b.PutString(s.PhoneNumber)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.sendVerifyPhoneCode#a5a356f9: field settings: %w", err)
	}
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (s *AccountSendVerifyPhoneCodeRequest) GetPhoneNumber() (value string) {
	return s.PhoneNumber
}

// GetSettings returns value of Settings field.
func (s *AccountSendVerifyPhoneCodeRequest) GetSettings() (value CodeSettings) {
	return s.Settings
}

// Decode implements bin.Decoder.
func (s *AccountSendVerifyPhoneCodeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sendVerifyPhoneCode#a5a356f9 to nil")
	}
	if err := b.ConsumeID(AccountSendVerifyPhoneCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.sendVerifyPhoneCode#a5a356f9: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.sendVerifyPhoneCode#a5a356f9: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.sendVerifyPhoneCode#a5a356f9: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSendVerifyPhoneCodeRequest.
var (
	_ bin.Encoder = &AccountSendVerifyPhoneCodeRequest{}
	_ bin.Decoder = &AccountSendVerifyPhoneCodeRequest{}
)

// AccountSendVerifyPhoneCode invokes method account.sendVerifyPhoneCode#a5a356f9 returning error if any.
// Send the verification phone code for telegram passport¹.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/account.sendVerifyPhoneCode for reference.
func (c *Client) AccountSendVerifyPhoneCode(ctx context.Context, request *AccountSendVerifyPhoneCodeRequest) (*AuthSentCode, error) {
	var result AuthSentCode

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
