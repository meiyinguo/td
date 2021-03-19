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

// Contact represents TL type `contact#f911c994`.
// A contact of the current user that is registered in the system.
//
// See https://core.telegram.org/constructor/contact for reference.
type Contact struct {
	// User identifier
	UserID int
	// Current user is in the user's contact list
	Mutual bool
}

// ContactTypeID is TL type id of Contact.
const ContactTypeID = 0xf911c994

func (c *Contact) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.Mutual == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *Contact) String() string {
	if c == nil {
		return "Contact(nil)"
	}
	type Alias Contact
	return fmt.Sprintf("Contact%+v", Alias(*c))
}

// FillFrom fills Contact from given interface.
func (c *Contact) FillFrom(from interface {
	GetUserID() (value int)
	GetMutual() (value bool)
}) {
	c.UserID = from.GetUserID()
	c.Mutual = from.GetMutual()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Contact) TypeID() uint32 {
	return ContactTypeID
}

// TypeName returns name of type in TL schema.
func (*Contact) TypeName() string {
	return "contact"
}

// TypeInfo returns info about TL type.
func (c *Contact) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contact",
		ID:   ContactTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Mutual",
			SchemaName: "mutual",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *Contact) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contact#f911c994 as nil")
	}
	b.PutID(ContactTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *Contact) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contact#f911c994 as nil")
	}
	b.PutInt(c.UserID)
	b.PutBool(c.Mutual)
	return nil
}

// GetUserID returns value of UserID field.
func (c *Contact) GetUserID() (value int) {
	return c.UserID
}

// GetMutual returns value of Mutual field.
func (c *Contact) GetMutual() (value bool) {
	return c.Mutual
}

// Decode implements bin.Decoder.
func (c *Contact) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contact#f911c994 to nil")
	}
	if err := b.ConsumeID(ContactTypeID); err != nil {
		return fmt.Errorf("unable to decode contact#f911c994: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *Contact) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contact#f911c994 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contact#f911c994: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode contact#f911c994: field mutual: %w", err)
		}
		c.Mutual = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Contact.
var (
	_ bin.Encoder     = &Contact{}
	_ bin.Decoder     = &Contact{}
	_ bin.BareEncoder = &Contact{}
	_ bin.BareDecoder = &Contact{}
)
