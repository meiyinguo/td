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

// MessagesSendEncryptedServiceRequest represents TL type `messages.sendEncryptedService#32d439a4`.
// Sends a service message to a secret chat.
//
// See https://core.telegram.org/method/messages.sendEncryptedService for reference.
type MessagesSendEncryptedServiceRequest struct {
	// Secret chat ID
	Peer InputEncryptedChat
	// Unique client message ID required to prevent message resending
	RandomID int64
	// TL-serialization of  DecryptedMessage¹ type, encrypted with a key generated during chat initialization
	//
	// Links:
	//  1) https://core.telegram.org/type/DecryptedMessage
	Data []byte
}

// MessagesSendEncryptedServiceRequestTypeID is TL type id of MessagesSendEncryptedServiceRequest.
const MessagesSendEncryptedServiceRequestTypeID = 0x32d439a4

func (s *MessagesSendEncryptedServiceRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer.Zero()) {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}
	if !(s.Data == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendEncryptedServiceRequest) String() string {
	if s == nil {
		return "MessagesSendEncryptedServiceRequest(nil)"
	}
	type Alias MessagesSendEncryptedServiceRequest
	return fmt.Sprintf("MessagesSendEncryptedServiceRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendEncryptedServiceRequest from given interface.
func (s *MessagesSendEncryptedServiceRequest) FillFrom(from interface {
	GetPeer() (value InputEncryptedChat)
	GetRandomID() (value int64)
	GetData() (value []byte)
}) {
	s.Peer = from.GetPeer()
	s.RandomID = from.GetRandomID()
	s.Data = from.GetData()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *MessagesSendEncryptedServiceRequest) TypeID() uint32 {
	return MessagesSendEncryptedServiceRequestTypeID
}

// Encode implements bin.Encoder.
func (s *MessagesSendEncryptedServiceRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendEncryptedService#32d439a4 as nil")
	}
	b.PutID(MessagesSendEncryptedServiceRequestTypeID)
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendEncryptedService#32d439a4: field peer: %w", err)
	}
	b.PutLong(s.RandomID)
	b.PutBytes(s.Data)
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSendEncryptedServiceRequest) GetPeer() (value InputEncryptedChat) {
	return s.Peer
}

// GetRandomID returns value of RandomID field.
func (s *MessagesSendEncryptedServiceRequest) GetRandomID() (value int64) {
	return s.RandomID
}

// GetData returns value of Data field.
func (s *MessagesSendEncryptedServiceRequest) GetData() (value []byte) {
	return s.Data
}

// Decode implements bin.Decoder.
func (s *MessagesSendEncryptedServiceRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendEncryptedService#32d439a4 to nil")
	}
	if err := b.ConsumeID(MessagesSendEncryptedServiceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendEncryptedService#32d439a4: %w", err)
	}
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedService#32d439a4: field peer: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedService#32d439a4: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendEncryptedService#32d439a4: field data: %w", err)
		}
		s.Data = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSendEncryptedServiceRequest.
var (
	_ bin.Encoder = &MessagesSendEncryptedServiceRequest{}
	_ bin.Decoder = &MessagesSendEncryptedServiceRequest{}
)

// MessagesSendEncryptedService invokes method messages.sendEncryptedService#32d439a4 returning error if any.
// Sends a service message to a secret chat.
//
// Possible errors:
//  400 DATA_INVALID: Encrypted data invalid
//  400 ENCRYPTION_DECLINED: The secret chat was declined
//  400 ENCRYPTION_ID_INVALID: The provided secret chat ID is invalid
//  400 MSG_WAIT_FAILED: A waiting call returned an error
//  403 USER_IS_BLOCKED: You were blocked by this user
//
// See https://core.telegram.org/method/messages.sendEncryptedService for reference.
func (c *Client) MessagesSendEncryptedService(ctx context.Context, request *MessagesSendEncryptedServiceRequest) (MessagesSentEncryptedMessageClass, error) {
	var result MessagesSentEncryptedMessageBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.SentEncryptedMessage, nil
}
