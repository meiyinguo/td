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

// AccountVerifyEmailRequest represents TL type `account.verifyEmail#ecba39db`.
// Verify an email address for telegram passport¹.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/account.verifyEmail for reference.
type AccountVerifyEmailRequest struct {
	// The email to verify
	Email string
	// The verification code that was received
	Code string
}

// AccountVerifyEmailRequestTypeID is TL type id of AccountVerifyEmailRequest.
const AccountVerifyEmailRequestTypeID = 0xecba39db

func (v *AccountVerifyEmailRequest) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Email == "") {
		return false
	}
	if !(v.Code == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *AccountVerifyEmailRequest) String() string {
	if v == nil {
		return "AccountVerifyEmailRequest(nil)"
	}
	type Alias AccountVerifyEmailRequest
	return fmt.Sprintf("AccountVerifyEmailRequest%+v", Alias(*v))
}

// FillFrom fills AccountVerifyEmailRequest from given interface.
func (v *AccountVerifyEmailRequest) FillFrom(from interface {
	GetEmail() (value string)
	GetCode() (value string)
}) {
	v.Email = from.GetEmail()
	v.Code = from.GetCode()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (v *AccountVerifyEmailRequest) TypeID() uint32 {
	return AccountVerifyEmailRequestTypeID
}

// Encode implements bin.Encoder.
func (v *AccountVerifyEmailRequest) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode account.verifyEmail#ecba39db as nil")
	}
	b.PutID(AccountVerifyEmailRequestTypeID)
	b.PutString(v.Email)
	b.PutString(v.Code)
	return nil
}

// GetEmail returns value of Email field.
func (v *AccountVerifyEmailRequest) GetEmail() (value string) {
	return v.Email
}

// GetCode returns value of Code field.
func (v *AccountVerifyEmailRequest) GetCode() (value string) {
	return v.Code
}

// Decode implements bin.Decoder.
func (v *AccountVerifyEmailRequest) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode account.verifyEmail#ecba39db to nil")
	}
	if err := b.ConsumeID(AccountVerifyEmailRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.verifyEmail#ecba39db: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.verifyEmail#ecba39db: field email: %w", err)
		}
		v.Email = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.verifyEmail#ecba39db: field code: %w", err)
		}
		v.Code = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountVerifyEmailRequest.
var (
	_ bin.Encoder = &AccountVerifyEmailRequest{}
	_ bin.Decoder = &AccountVerifyEmailRequest{}
)

// AccountVerifyEmail invokes method account.verifyEmail#ecba39db returning error if any.
// Verify an email address for telegram passport¹.
//
// Links:
//  1) https://core.telegram.org/passport
//
// Possible errors:
//  400 EMAIL_VERIFY_EXPIRED: The verification email has expired
//
// See https://core.telegram.org/method/account.verifyEmail for reference.
func (c *Client) AccountVerifyEmail(ctx context.Context, request *AccountVerifyEmailRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
