// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// DocumentAttributeImageSize represents TL type `documentAttributeImageSize#6c37c15c`.
// Defines the width and height of an image uploaded as document
//
// See https://core.telegram.org/constructor/documentAttributeImageSize for reference.
type DocumentAttributeImageSize struct {
	// Width of image
	W int
	// Height of image
	H int
}

// DocumentAttributeImageSizeTypeID is TL type id of DocumentAttributeImageSize.
const DocumentAttributeImageSizeTypeID = 0x6c37c15c

// String implements fmt.Stringer.
func (d *DocumentAttributeImageSize) String() string {
	if d == nil {
		return "DocumentAttributeImageSize(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DocumentAttributeImageSize")
	sb.WriteString("{\n")
	sb.WriteString("\tW: ")
	sb.WriteString(fmt.Sprint(d.W))
	sb.WriteString(",\n")
	sb.WriteString("\tH: ")
	sb.WriteString(fmt.Sprint(d.H))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeImageSize) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeImageSize#6c37c15c as nil")
	}
	b.PutID(DocumentAttributeImageSizeTypeID)
	b.PutInt(d.W)
	b.PutInt(d.H)
	return nil
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeImageSize) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeImageSize#6c37c15c to nil")
	}
	if err := b.ConsumeID(DocumentAttributeImageSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeImageSize#6c37c15c: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeImageSize#6c37c15c: field w: %w", err)
		}
		d.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeImageSize#6c37c15c: field h: %w", err)
		}
		d.H = value
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeImageSize) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeImageSize.
var (
	_ bin.Encoder = &DocumentAttributeImageSize{}
	_ bin.Decoder = &DocumentAttributeImageSize{}

	_ DocumentAttributeClass = &DocumentAttributeImageSize{}
)

// DocumentAttributeAnimated represents TL type `documentAttributeAnimated#11b58939`.
// Defines an animated GIF
//
// See https://core.telegram.org/constructor/documentAttributeAnimated for reference.
type DocumentAttributeAnimated struct {
}

// DocumentAttributeAnimatedTypeID is TL type id of DocumentAttributeAnimated.
const DocumentAttributeAnimatedTypeID = 0x11b58939

// String implements fmt.Stringer.
func (d *DocumentAttributeAnimated) String() string {
	if d == nil {
		return "DocumentAttributeAnimated(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DocumentAttributeAnimated")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeAnimated) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeAnimated#11b58939 as nil")
	}
	b.PutID(DocumentAttributeAnimatedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeAnimated) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeAnimated#11b58939 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeAnimatedTypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeAnimated#11b58939: %w", err)
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeAnimated) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeAnimated.
var (
	_ bin.Encoder = &DocumentAttributeAnimated{}
	_ bin.Decoder = &DocumentAttributeAnimated{}

	_ DocumentAttributeClass = &DocumentAttributeAnimated{}
)

// DocumentAttributeSticker represents TL type `documentAttributeSticker#6319d612`.
// Defines a sticker
//
// See https://core.telegram.org/constructor/documentAttributeSticker for reference.
type DocumentAttributeSticker struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is a mask sticker
	Mask bool
	// Alternative emoji representation of sticker
	Alt string
	// Associated stickerset
	Stickerset InputStickerSetClass
	// Mask coordinates (if this is a mask sticker, attached to a photo)
	//
	// Use SetMaskCoords and GetMaskCoords helpers.
	MaskCoords MaskCoords
}

// DocumentAttributeStickerTypeID is TL type id of DocumentAttributeSticker.
const DocumentAttributeStickerTypeID = 0x6319d612

// String implements fmt.Stringer.
func (d *DocumentAttributeSticker) String() string {
	if d == nil {
		return "DocumentAttributeSticker(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DocumentAttributeSticker")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(d.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tAlt: ")
	sb.WriteString(fmt.Sprint(d.Alt))
	sb.WriteString(",\n")
	sb.WriteString("\tStickerset: ")
	sb.WriteString(d.Stickerset.String())
	sb.WriteString(",\n")
	if d.Flags.Has(0) {
		sb.WriteString("\tMaskCoords: ")
		sb.WriteString(d.MaskCoords.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeSticker) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeSticker#6319d612 as nil")
	}
	b.PutID(DocumentAttributeStickerTypeID)
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode documentAttributeSticker#6319d612: field flags: %w", err)
	}
	b.PutString(d.Alt)
	if d.Stickerset == nil {
		return fmt.Errorf("unable to encode documentAttributeSticker#6319d612: field stickerset is nil")
	}
	if err := d.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode documentAttributeSticker#6319d612: field stickerset: %w", err)
	}
	if d.Flags.Has(0) {
		if err := d.MaskCoords.Encode(b); err != nil {
			return fmt.Errorf("unable to encode documentAttributeSticker#6319d612: field mask_coords: %w", err)
		}
	}
	return nil
}

// SetMask sets value of Mask conditional field.
func (d *DocumentAttributeSticker) SetMask(value bool) {
	if value {
		d.Flags.Set(1)
		d.Mask = true
	} else {
		d.Flags.Unset(1)
		d.Mask = false
	}
}

// SetMaskCoords sets value of MaskCoords conditional field.
func (d *DocumentAttributeSticker) SetMaskCoords(value MaskCoords) {
	d.Flags.Set(0)
	d.MaskCoords = value
}

// GetMaskCoords returns value of MaskCoords conditional field and
// boolean which is true if field was set.
func (d *DocumentAttributeSticker) GetMaskCoords() (value MaskCoords, ok bool) {
	if !d.Flags.Has(0) {
		return value, false
	}
	return d.MaskCoords, true
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeSticker) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeSticker#6319d612 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeStickerTypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeSticker#6319d612: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode documentAttributeSticker#6319d612: field flags: %w", err)
		}
	}
	d.Mask = d.Flags.Has(1)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeSticker#6319d612: field alt: %w", err)
		}
		d.Alt = value
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeSticker#6319d612: field stickerset: %w", err)
		}
		d.Stickerset = value
	}
	if d.Flags.Has(0) {
		if err := d.MaskCoords.Decode(b); err != nil {
			return fmt.Errorf("unable to decode documentAttributeSticker#6319d612: field mask_coords: %w", err)
		}
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeSticker) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeSticker.
var (
	_ bin.Encoder = &DocumentAttributeSticker{}
	_ bin.Decoder = &DocumentAttributeSticker{}

	_ DocumentAttributeClass = &DocumentAttributeSticker{}
)

// DocumentAttributeVideo represents TL type `documentAttributeVideo#ef02ce6`.
// Defines a video
//
// See https://core.telegram.org/constructor/documentAttributeVideo for reference.
type DocumentAttributeVideo struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is a round video
	RoundMessage bool
	// Whether the video supports streaming
	SupportsStreaming bool
	// Duration in seconds
	Duration int
	// Video width
	W int
	// Video height
	H int
}

// DocumentAttributeVideoTypeID is TL type id of DocumentAttributeVideo.
const DocumentAttributeVideoTypeID = 0xef02ce6

// String implements fmt.Stringer.
func (d *DocumentAttributeVideo) String() string {
	if d == nil {
		return "DocumentAttributeVideo(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DocumentAttributeVideo")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(d.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tDuration: ")
	sb.WriteString(fmt.Sprint(d.Duration))
	sb.WriteString(",\n")
	sb.WriteString("\tW: ")
	sb.WriteString(fmt.Sprint(d.W))
	sb.WriteString(",\n")
	sb.WriteString("\tH: ")
	sb.WriteString(fmt.Sprint(d.H))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeVideo) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeVideo#ef02ce6 as nil")
	}
	b.PutID(DocumentAttributeVideoTypeID)
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode documentAttributeVideo#ef02ce6: field flags: %w", err)
	}
	b.PutInt(d.Duration)
	b.PutInt(d.W)
	b.PutInt(d.H)
	return nil
}

// SetRoundMessage sets value of RoundMessage conditional field.
func (d *DocumentAttributeVideo) SetRoundMessage(value bool) {
	if value {
		d.Flags.Set(0)
		d.RoundMessage = true
	} else {
		d.Flags.Unset(0)
		d.RoundMessage = false
	}
}

// SetSupportsStreaming sets value of SupportsStreaming conditional field.
func (d *DocumentAttributeVideo) SetSupportsStreaming(value bool) {
	if value {
		d.Flags.Set(1)
		d.SupportsStreaming = true
	} else {
		d.Flags.Unset(1)
		d.SupportsStreaming = false
	}
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeVideo) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeVideo#ef02ce6 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeVideoTypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeVideo#ef02ce6: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode documentAttributeVideo#ef02ce6: field flags: %w", err)
		}
	}
	d.RoundMessage = d.Flags.Has(0)
	d.SupportsStreaming = d.Flags.Has(1)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeVideo#ef02ce6: field duration: %w", err)
		}
		d.Duration = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeVideo#ef02ce6: field w: %w", err)
		}
		d.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeVideo#ef02ce6: field h: %w", err)
		}
		d.H = value
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeVideo) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeVideo.
var (
	_ bin.Encoder = &DocumentAttributeVideo{}
	_ bin.Decoder = &DocumentAttributeVideo{}

	_ DocumentAttributeClass = &DocumentAttributeVideo{}
)

// DocumentAttributeAudio represents TL type `documentAttributeAudio#9852f9c6`.
// Represents an audio file
//
// See https://core.telegram.org/constructor/documentAttributeAudio for reference.
type DocumentAttributeAudio struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is a voice message
	Voice bool
	// Duration in seconds
	Duration int
	// Name of song
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// Performer
	//
	// Use SetPerformer and GetPerformer helpers.
	Performer string
	// Waveform
	//
	// Use SetWaveform and GetWaveform helpers.
	Waveform []byte
}

// DocumentAttributeAudioTypeID is TL type id of DocumentAttributeAudio.
const DocumentAttributeAudioTypeID = 0x9852f9c6

// String implements fmt.Stringer.
func (d *DocumentAttributeAudio) String() string {
	if d == nil {
		return "DocumentAttributeAudio(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DocumentAttributeAudio")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(d.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tDuration: ")
	sb.WriteString(fmt.Sprint(d.Duration))
	sb.WriteString(",\n")
	if d.Flags.Has(0) {
		sb.WriteString("\tTitle: ")
		sb.WriteString(fmt.Sprint(d.Title))
		sb.WriteString(",\n")
	}
	if d.Flags.Has(1) {
		sb.WriteString("\tPerformer: ")
		sb.WriteString(fmt.Sprint(d.Performer))
		sb.WriteString(",\n")
	}
	if d.Flags.Has(2) {
		sb.WriteString("\tWaveform: ")
		sb.WriteString(fmt.Sprint(d.Waveform))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeAudio) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeAudio#9852f9c6 as nil")
	}
	b.PutID(DocumentAttributeAudioTypeID)
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode documentAttributeAudio#9852f9c6: field flags: %w", err)
	}
	b.PutInt(d.Duration)
	if d.Flags.Has(0) {
		b.PutString(d.Title)
	}
	if d.Flags.Has(1) {
		b.PutString(d.Performer)
	}
	if d.Flags.Has(2) {
		b.PutBytes(d.Waveform)
	}
	return nil
}

// SetVoice sets value of Voice conditional field.
func (d *DocumentAttributeAudio) SetVoice(value bool) {
	if value {
		d.Flags.Set(10)
		d.Voice = true
	} else {
		d.Flags.Unset(10)
		d.Voice = false
	}
}

// SetTitle sets value of Title conditional field.
func (d *DocumentAttributeAudio) SetTitle(value string) {
	d.Flags.Set(0)
	d.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (d *DocumentAttributeAudio) GetTitle() (value string, ok bool) {
	if !d.Flags.Has(0) {
		return value, false
	}
	return d.Title, true
}

// SetPerformer sets value of Performer conditional field.
func (d *DocumentAttributeAudio) SetPerformer(value string) {
	d.Flags.Set(1)
	d.Performer = value
}

// GetPerformer returns value of Performer conditional field and
// boolean which is true if field was set.
func (d *DocumentAttributeAudio) GetPerformer() (value string, ok bool) {
	if !d.Flags.Has(1) {
		return value, false
	}
	return d.Performer, true
}

// SetWaveform sets value of Waveform conditional field.
func (d *DocumentAttributeAudio) SetWaveform(value []byte) {
	d.Flags.Set(2)
	d.Waveform = value
}

// GetWaveform returns value of Waveform conditional field and
// boolean which is true if field was set.
func (d *DocumentAttributeAudio) GetWaveform() (value []byte, ok bool) {
	if !d.Flags.Has(2) {
		return value, false
	}
	return d.Waveform, true
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeAudio) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeAudio#9852f9c6 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeAudioTypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeAudio#9852f9c6: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio#9852f9c6: field flags: %w", err)
		}
	}
	d.Voice = d.Flags.Has(10)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio#9852f9c6: field duration: %w", err)
		}
		d.Duration = value
	}
	if d.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio#9852f9c6: field title: %w", err)
		}
		d.Title = value
	}
	if d.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio#9852f9c6: field performer: %w", err)
		}
		d.Performer = value
	}
	if d.Flags.Has(2) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeAudio#9852f9c6: field waveform: %w", err)
		}
		d.Waveform = value
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeAudio) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeAudio.
var (
	_ bin.Encoder = &DocumentAttributeAudio{}
	_ bin.Decoder = &DocumentAttributeAudio{}

	_ DocumentAttributeClass = &DocumentAttributeAudio{}
)

// DocumentAttributeFilename represents TL type `documentAttributeFilename#15590068`.
// A simple document with a file name
//
// See https://core.telegram.org/constructor/documentAttributeFilename for reference.
type DocumentAttributeFilename struct {
	// The file name
	FileName string
}

// DocumentAttributeFilenameTypeID is TL type id of DocumentAttributeFilename.
const DocumentAttributeFilenameTypeID = 0x15590068

// String implements fmt.Stringer.
func (d *DocumentAttributeFilename) String() string {
	if d == nil {
		return "DocumentAttributeFilename(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DocumentAttributeFilename")
	sb.WriteString("{\n")
	sb.WriteString("\tFileName: ")
	sb.WriteString(fmt.Sprint(d.FileName))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeFilename) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeFilename#15590068 as nil")
	}
	b.PutID(DocumentAttributeFilenameTypeID)
	b.PutString(d.FileName)
	return nil
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeFilename) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeFilename#15590068 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeFilenameTypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeFilename#15590068: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode documentAttributeFilename#15590068: field file_name: %w", err)
		}
		d.FileName = value
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeFilename) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeFilename.
var (
	_ bin.Encoder = &DocumentAttributeFilename{}
	_ bin.Decoder = &DocumentAttributeFilename{}

	_ DocumentAttributeClass = &DocumentAttributeFilename{}
)

// DocumentAttributeHasStickers represents TL type `documentAttributeHasStickers#9801d2f7`.
// Whether the current document has stickers attached
//
// See https://core.telegram.org/constructor/documentAttributeHasStickers for reference.
type DocumentAttributeHasStickers struct {
}

// DocumentAttributeHasStickersTypeID is TL type id of DocumentAttributeHasStickers.
const DocumentAttributeHasStickersTypeID = 0x9801d2f7

// String implements fmt.Stringer.
func (d *DocumentAttributeHasStickers) String() string {
	if d == nil {
		return "DocumentAttributeHasStickers(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DocumentAttributeHasStickers")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DocumentAttributeHasStickers) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode documentAttributeHasStickers#9801d2f7 as nil")
	}
	b.PutID(DocumentAttributeHasStickersTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DocumentAttributeHasStickers) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode documentAttributeHasStickers#9801d2f7 to nil")
	}
	if err := b.ConsumeID(DocumentAttributeHasStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode documentAttributeHasStickers#9801d2f7: %w", err)
	}
	return nil
}

// construct implements constructor of DocumentAttributeClass.
func (d DocumentAttributeHasStickers) construct() DocumentAttributeClass { return &d }

// Ensuring interfaces in compile-time for DocumentAttributeHasStickers.
var (
	_ bin.Encoder = &DocumentAttributeHasStickers{}
	_ bin.Decoder = &DocumentAttributeHasStickers{}

	_ DocumentAttributeClass = &DocumentAttributeHasStickers{}
)

// DocumentAttributeClass represents DocumentAttribute generic type.
//
// See https://core.telegram.org/type/DocumentAttribute for reference.
//
// Example:
//  g, err := DecodeDocumentAttribute(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *DocumentAttributeImageSize: // documentAttributeImageSize#6c37c15c
//  case *DocumentAttributeAnimated: // documentAttributeAnimated#11b58939
//  case *DocumentAttributeSticker: // documentAttributeSticker#6319d612
//  case *DocumentAttributeVideo: // documentAttributeVideo#ef02ce6
//  case *DocumentAttributeAudio: // documentAttributeAudio#9852f9c6
//  case *DocumentAttributeFilename: // documentAttributeFilename#15590068
//  case *DocumentAttributeHasStickers: // documentAttributeHasStickers#9801d2f7
//  default: panic(v)
//  }
type DocumentAttributeClass interface {
	bin.Encoder
	bin.Decoder
	construct() DocumentAttributeClass
	fmt.Stringer
}

// DecodeDocumentAttribute implements binary de-serialization for DocumentAttributeClass.
func DecodeDocumentAttribute(buf *bin.Buffer) (DocumentAttributeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DocumentAttributeImageSizeTypeID:
		// Decoding documentAttributeImageSize#6c37c15c.
		v := DocumentAttributeImageSize{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeAnimatedTypeID:
		// Decoding documentAttributeAnimated#11b58939.
		v := DocumentAttributeAnimated{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeStickerTypeID:
		// Decoding documentAttributeSticker#6319d612.
		v := DocumentAttributeSticker{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeVideoTypeID:
		// Decoding documentAttributeVideo#ef02ce6.
		v := DocumentAttributeVideo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeAudioTypeID:
		// Decoding documentAttributeAudio#9852f9c6.
		v := DocumentAttributeAudio{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeFilenameTypeID:
		// Decoding documentAttributeFilename#15590068.
		v := DocumentAttributeFilename{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	case DocumentAttributeHasStickersTypeID:
		// Decoding documentAttributeHasStickers#9801d2f7.
		v := DocumentAttributeHasStickers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DocumentAttributeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DocumentAttribute boxes the DocumentAttributeClass providing a helper.
type DocumentAttributeBox struct {
	DocumentAttribute DocumentAttributeClass
}

// Decode implements bin.Decoder for DocumentAttributeBox.
func (b *DocumentAttributeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode DocumentAttributeBox to nil")
	}
	v, err := DecodeDocumentAttribute(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DocumentAttribute = v
	return nil
}

// Encode implements bin.Encode for DocumentAttributeBox.
func (b *DocumentAttributeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DocumentAttribute == nil {
		return fmt.Errorf("unable to encode DocumentAttributeClass as nil")
	}
	return b.DocumentAttribute.Encode(buf)
}
