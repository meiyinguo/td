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

// SavedPhoneContact represents TL type `savedPhoneContact#1142bd56`.
// Saved contact
//
// See https://core.telegram.org/constructor/savedPhoneContact for reference.
type SavedPhoneContact struct {
	// Phone number
	Phone string
	// First name
	FirstName string
	// Last name
	LastName string
	// Date added
	Date int
}

// SavedPhoneContactTypeID is TL type id of SavedPhoneContact.
const SavedPhoneContactTypeID = 0x1142bd56

func (s *SavedPhoneContact) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Phone == "") {
		return false
	}
	if !(s.FirstName == "") {
		return false
	}
	if !(s.LastName == "") {
		return false
	}
	if !(s.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SavedPhoneContact) String() string {
	if s == nil {
		return "SavedPhoneContact(nil)"
	}
	type Alias SavedPhoneContact
	return fmt.Sprintf("SavedPhoneContact%+v", Alias(*s))
}

// FillFrom fills SavedPhoneContact from given interface.
func (s *SavedPhoneContact) FillFrom(from interface {
	GetPhone() (value string)
	GetFirstName() (value string)
	GetLastName() (value string)
	GetDate() (value int)
}) {
	s.Phone = from.GetPhone()
	s.FirstName = from.GetFirstName()
	s.LastName = from.GetLastName()
	s.Date = from.GetDate()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SavedPhoneContact) TypeID() uint32 {
	return SavedPhoneContactTypeID
}

// Encode implements bin.Encoder.
func (s *SavedPhoneContact) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode savedPhoneContact#1142bd56 as nil")
	}
	b.PutID(SavedPhoneContactTypeID)
	b.PutString(s.Phone)
	b.PutString(s.FirstName)
	b.PutString(s.LastName)
	b.PutInt(s.Date)
	return nil
}

// GetPhone returns value of Phone field.
func (s *SavedPhoneContact) GetPhone() (value string) {
	return s.Phone
}

// GetFirstName returns value of FirstName field.
func (s *SavedPhoneContact) GetFirstName() (value string) {
	return s.FirstName
}

// GetLastName returns value of LastName field.
func (s *SavedPhoneContact) GetLastName() (value string) {
	return s.LastName
}

// GetDate returns value of Date field.
func (s *SavedPhoneContact) GetDate() (value int) {
	return s.Date
}

// Decode implements bin.Decoder.
func (s *SavedPhoneContact) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode savedPhoneContact#1142bd56 to nil")
	}
	if err := b.ConsumeID(SavedPhoneContactTypeID); err != nil {
		return fmt.Errorf("unable to decode savedPhoneContact#1142bd56: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode savedPhoneContact#1142bd56: field phone: %w", err)
		}
		s.Phone = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode savedPhoneContact#1142bd56: field first_name: %w", err)
		}
		s.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode savedPhoneContact#1142bd56: field last_name: %w", err)
		}
		s.LastName = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode savedPhoneContact#1142bd56: field date: %w", err)
		}
		s.Date = value
	}
	return nil
}

// Ensuring interfaces in compile-time for SavedPhoneContact.
var (
	_ bin.Encoder = &SavedPhoneContact{}
	_ bin.Decoder = &SavedPhoneContact{}
)
