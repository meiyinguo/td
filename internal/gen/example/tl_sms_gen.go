// Code generated by gotdgen, DO NOT EDIT.

package td

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

// SMS represents TL type `sms#ed8bebfe`.
//
// See https://localhost:80/doc/constructor/sms for reference.
type SMS struct {
	// Text field of SMS.
	Text string
}

// SMSTypeID is TL type id of SMS.
const SMSTypeID = 0xed8bebfe

func (s *SMS) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SMS) String() string {
	if s == nil {
		return "SMS(nil)"
	}
	type Alias SMS
	return fmt.Sprintf("SMS%+v", Alias(*s))
}

// FillFrom fills SMS from given interface.
func (s *SMS) FillFrom(from interface {
	GetText() (value string)
}) {
	s.Text = from.GetText()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SMS) TypeID() uint32 {
	return SMSTypeID
}

// Encode implements bin.Encoder.
func (s *SMS) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sms#ed8bebfe as nil")
	}
	b.PutID(SMSTypeID)
	b.PutString(s.Text)
	return nil
}

// GetText returns value of Text field.
func (s *SMS) GetText() (value string) {
	return s.Text
}

// Decode implements bin.Decoder.
func (s *SMS) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sms#ed8bebfe to nil")
	}
	if err := b.ConsumeID(SMSTypeID); err != nil {
		return fmt.Errorf("unable to decode sms#ed8bebfe: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sms#ed8bebfe: field text: %w", err)
		}
		s.Text = value
	}
	return nil
}

// Ensuring interfaces in compile-time for SMS.
var (
	_ bin.Encoder = &SMS{}
	_ bin.Decoder = &SMS{}
)
