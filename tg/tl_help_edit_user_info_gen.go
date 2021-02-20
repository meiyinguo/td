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

// HelpEditUserInfoRequest represents TL type `help.editUserInfo#66b91b70`.
// Internal use
//
// See https://core.telegram.org/method/help.editUserInfo for reference.
type HelpEditUserInfoRequest struct {
	// User
	UserID InputUserClass
	// Message
	Message string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	Entities []MessageEntityClass
}

// HelpEditUserInfoRequestTypeID is TL type id of HelpEditUserInfoRequest.
const HelpEditUserInfoRequestTypeID = 0x66b91b70

func (e *HelpEditUserInfoRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.UserID == nil) {
		return false
	}
	if !(e.Message == "") {
		return false
	}
	if !(e.Entities == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *HelpEditUserInfoRequest) String() string {
	if e == nil {
		return "HelpEditUserInfoRequest(nil)"
	}
	type Alias HelpEditUserInfoRequest
	return fmt.Sprintf("HelpEditUserInfoRequest%+v", Alias(*e))
}

// FillFrom fills HelpEditUserInfoRequest from given interface.
func (e *HelpEditUserInfoRequest) FillFrom(from interface {
	GetUserID() (value InputUserClass)
	GetMessage() (value string)
	GetEntities() (value []MessageEntityClass)
}) {
	e.UserID = from.GetUserID()
	e.Message = from.GetMessage()
	e.Entities = from.GetEntities()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (e *HelpEditUserInfoRequest) TypeID() uint32 {
	return HelpEditUserInfoRequestTypeID
}

// Encode implements bin.Encoder.
func (e *HelpEditUserInfoRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode help.editUserInfo#66b91b70 as nil")
	}
	b.PutID(HelpEditUserInfoRequestTypeID)
	if e.UserID == nil {
		return fmt.Errorf("unable to encode help.editUserInfo#66b91b70: field user_id is nil")
	}
	if err := e.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.editUserInfo#66b91b70: field user_id: %w", err)
	}
	b.PutString(e.Message)
	b.PutVectorHeader(len(e.Entities))
	for idx, v := range e.Entities {
		if v == nil {
			return fmt.Errorf("unable to encode help.editUserInfo#66b91b70: field entities element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.editUserInfo#66b91b70: field entities element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetUserID returns value of UserID field.
func (e *HelpEditUserInfoRequest) GetUserID() (value InputUserClass) {
	return e.UserID
}

// GetMessage returns value of Message field.
func (e *HelpEditUserInfoRequest) GetMessage() (value string) {
	return e.Message
}

// GetEntities returns value of Entities field.
func (e *HelpEditUserInfoRequest) GetEntities() (value []MessageEntityClass) {
	return e.Entities
}

// MapEntities returns field Entities wrapped in MessageEntityClassSlice helper.
func (e *HelpEditUserInfoRequest) MapEntities() (value MessageEntityClassSlice) {
	return MessageEntityClassSlice(e.Entities)
}

// Decode implements bin.Decoder.
func (e *HelpEditUserInfoRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode help.editUserInfo#66b91b70 to nil")
	}
	if err := b.ConsumeID(HelpEditUserInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.editUserInfo#66b91b70: %w", err)
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode help.editUserInfo#66b91b70: field user_id: %w", err)
		}
		e.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.editUserInfo#66b91b70: field message: %w", err)
		}
		e.Message = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.editUserInfo#66b91b70: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.editUserInfo#66b91b70: field entities: %w", err)
			}
			e.Entities = append(e.Entities, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpEditUserInfoRequest.
var (
	_ bin.Encoder = &HelpEditUserInfoRequest{}
	_ bin.Decoder = &HelpEditUserInfoRequest{}
)

// HelpEditUserInfo invokes method help.editUserInfo#66b91b70 returning error if any.
// Internal use
//
// See https://core.telegram.org/method/help.editUserInfo for reference.
func (c *Client) HelpEditUserInfo(ctx context.Context, request *HelpEditUserInfoRequest) (HelpUserInfoClass, error) {
	var result HelpUserInfoBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.UserInfo, nil
}
