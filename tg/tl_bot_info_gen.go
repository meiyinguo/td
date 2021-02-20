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

// BotInfo represents TL type `botInfo#98e81d3a`.
// Info about bots (available bot commands, etc)
//
// See https://core.telegram.org/constructor/botInfo for reference.
type BotInfo struct {
	// ID of the bot
	UserID int
	// Description of the bot
	Description string
	// Bot commands that can be used in the chat
	Commands []BotCommand
}

// BotInfoTypeID is TL type id of BotInfo.
const BotInfoTypeID = 0x98e81d3a

func (b *BotInfo) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.UserID == 0) {
		return false
	}
	if !(b.Description == "") {
		return false
	}
	if !(b.Commands == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotInfo) String() string {
	if b == nil {
		return "BotInfo(nil)"
	}
	type Alias BotInfo
	return fmt.Sprintf("BotInfo%+v", Alias(*b))
}

// FillFrom fills BotInfo from given interface.
func (b *BotInfo) FillFrom(from interface {
	GetUserID() (value int)
	GetDescription() (value string)
	GetCommands() (value []BotCommand)
}) {
	b.UserID = from.GetUserID()
	b.Description = from.GetDescription()
	b.Commands = from.GetCommands()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (b *BotInfo) TypeID() uint32 {
	return BotInfoTypeID
}

// Encode implements bin.Encoder.
func (b *BotInfo) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInfo#98e81d3a as nil")
	}
	buf.PutID(BotInfoTypeID)
	buf.PutInt(b.UserID)
	buf.PutString(b.Description)
	buf.PutVectorHeader(len(b.Commands))
	for idx, v := range b.Commands {
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInfo#98e81d3a: field commands element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetUserID returns value of UserID field.
func (b *BotInfo) GetUserID() (value int) {
	return b.UserID
}

// GetDescription returns value of Description field.
func (b *BotInfo) GetDescription() (value string) {
	return b.Description
}

// GetCommands returns value of Commands field.
func (b *BotInfo) GetCommands() (value []BotCommand) {
	return b.Commands
}

// Decode implements bin.Decoder.
func (b *BotInfo) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInfo#98e81d3a to nil")
	}
	if err := buf.ConsumeID(BotInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode botInfo#98e81d3a: %w", err)
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#98e81d3a: field user_id: %w", err)
		}
		b.UserID = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#98e81d3a: field description: %w", err)
		}
		b.Description = value
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode botInfo#98e81d3a: field commands: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BotCommand
			if err := value.Decode(buf); err != nil {
				return fmt.Errorf("unable to decode botInfo#98e81d3a: field commands: %w", err)
			}
			b.Commands = append(b.Commands, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for BotInfo.
var (
	_ bin.Encoder = &BotInfo{}
	_ bin.Decoder = &BotInfo{}
)
