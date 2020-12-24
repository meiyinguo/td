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

// WallPaper represents TL type `wallPaper#a437c3ed`.
// Wallpaper settings.
//
// See https://core.telegram.org/constructor/wallPaper for reference.
type WallPaper struct {
	// Identifier
	ID int64
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Creator of the wallpaper
	Creator bool
	// Whether this is the default wallpaper
	Default bool
	// Pattern
	Pattern bool
	// Dark mode
	Dark bool
	// Access hash
	AccessHash int64
	// Unique wallpaper ID
	Slug string
	// The actual wallpaper
	Document DocumentClass
	// Wallpaper settings
	//
	// Use SetSettings and GetSettings helpers.
	Settings WallPaperSettings
}

// WallPaperTypeID is TL type id of WallPaper.
const WallPaperTypeID = 0xa437c3ed

// String implements fmt.Stringer.
func (w *WallPaper) String() string {
	if w == nil {
		return "WallPaper(nil)"
	}
	var sb strings.Builder
	sb.WriteString("WallPaper")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(w.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(w.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(w.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("\tSlug: ")
	sb.WriteString(fmt.Sprint(w.Slug))
	sb.WriteString(",\n")
	sb.WriteString("\tDocument: ")
	sb.WriteString(w.Document.String())
	sb.WriteString(",\n")
	if w.Flags.Has(2) {
		sb.WriteString("\tSettings: ")
		sb.WriteString(w.Settings.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (w *WallPaper) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode wallPaper#a437c3ed as nil")
	}
	b.PutID(WallPaperTypeID)
	b.PutLong(w.ID)
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode wallPaper#a437c3ed: field flags: %w", err)
	}
	b.PutLong(w.AccessHash)
	b.PutString(w.Slug)
	if w.Document == nil {
		return fmt.Errorf("unable to encode wallPaper#a437c3ed: field document is nil")
	}
	if err := w.Document.Encode(b); err != nil {
		return fmt.Errorf("unable to encode wallPaper#a437c3ed: field document: %w", err)
	}
	if w.Flags.Has(2) {
		if err := w.Settings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode wallPaper#a437c3ed: field settings: %w", err)
		}
	}
	return nil
}

// SetCreator sets value of Creator conditional field.
func (w *WallPaper) SetCreator(value bool) {
	if value {
		w.Flags.Set(0)
		w.Creator = true
	} else {
		w.Flags.Unset(0)
		w.Creator = false
	}
}

// SetDefault sets value of Default conditional field.
func (w *WallPaper) SetDefault(value bool) {
	if value {
		w.Flags.Set(1)
		w.Default = true
	} else {
		w.Flags.Unset(1)
		w.Default = false
	}
}

// SetPattern sets value of Pattern conditional field.
func (w *WallPaper) SetPattern(value bool) {
	if value {
		w.Flags.Set(3)
		w.Pattern = true
	} else {
		w.Flags.Unset(3)
		w.Pattern = false
	}
}

// SetDark sets value of Dark conditional field.
func (w *WallPaper) SetDark(value bool) {
	if value {
		w.Flags.Set(4)
		w.Dark = true
	} else {
		w.Flags.Unset(4)
		w.Dark = false
	}
}

// SetSettings sets value of Settings conditional field.
func (w *WallPaper) SetSettings(value WallPaperSettings) {
	w.Flags.Set(2)
	w.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (w *WallPaper) GetSettings() (value WallPaperSettings, ok bool) {
	if !w.Flags.Has(2) {
		return value, false
	}
	return w.Settings, true
}

// Decode implements bin.Decoder.
func (w *WallPaper) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode wallPaper#a437c3ed to nil")
	}
	if err := b.ConsumeID(WallPaperTypeID); err != nil {
		return fmt.Errorf("unable to decode wallPaper#a437c3ed: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field id: %w", err)
		}
		w.ID = value
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field flags: %w", err)
		}
	}
	w.Creator = w.Flags.Has(0)
	w.Default = w.Flags.Has(1)
	w.Pattern = w.Flags.Has(3)
	w.Dark = w.Flags.Has(4)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field access_hash: %w", err)
		}
		w.AccessHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field slug: %w", err)
		}
		w.Slug = value
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field document: %w", err)
		}
		w.Document = value
	}
	if w.Flags.Has(2) {
		if err := w.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field settings: %w", err)
		}
	}
	return nil
}

// construct implements constructor of WallPaperClass.
func (w WallPaper) construct() WallPaperClass { return &w }

// Ensuring interfaces in compile-time for WallPaper.
var (
	_ bin.Encoder = &WallPaper{}
	_ bin.Decoder = &WallPaper{}

	_ WallPaperClass = &WallPaper{}
)

// WallPaperNoFile represents TL type `wallPaperNoFile#8af40b25`.
// No file wallpaper
//
// See https://core.telegram.org/constructor/wallPaperNoFile for reference.
type WallPaperNoFile struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is the default wallpaper
	Default bool
	// Dark mode
	Dark bool
	// Wallpaper settings
	//
	// Use SetSettings and GetSettings helpers.
	Settings WallPaperSettings
}

// WallPaperNoFileTypeID is TL type id of WallPaperNoFile.
const WallPaperNoFileTypeID = 0x8af40b25

// String implements fmt.Stringer.
func (w *WallPaperNoFile) String() string {
	if w == nil {
		return "WallPaperNoFile(nil)"
	}
	var sb strings.Builder
	sb.WriteString("WallPaperNoFile")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(w.Flags.String())
	sb.WriteString(",\n")
	if w.Flags.Has(2) {
		sb.WriteString("\tSettings: ")
		sb.WriteString(w.Settings.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (w *WallPaperNoFile) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode wallPaperNoFile#8af40b25 as nil")
	}
	b.PutID(WallPaperNoFileTypeID)
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode wallPaperNoFile#8af40b25: field flags: %w", err)
	}
	if w.Flags.Has(2) {
		if err := w.Settings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode wallPaperNoFile#8af40b25: field settings: %w", err)
		}
	}
	return nil
}

// SetDefault sets value of Default conditional field.
func (w *WallPaperNoFile) SetDefault(value bool) {
	if value {
		w.Flags.Set(1)
		w.Default = true
	} else {
		w.Flags.Unset(1)
		w.Default = false
	}
}

// SetDark sets value of Dark conditional field.
func (w *WallPaperNoFile) SetDark(value bool) {
	if value {
		w.Flags.Set(4)
		w.Dark = true
	} else {
		w.Flags.Unset(4)
		w.Dark = false
	}
}

// SetSettings sets value of Settings conditional field.
func (w *WallPaperNoFile) SetSettings(value WallPaperSettings) {
	w.Flags.Set(2)
	w.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (w *WallPaperNoFile) GetSettings() (value WallPaperSettings, ok bool) {
	if !w.Flags.Has(2) {
		return value, false
	}
	return w.Settings, true
}

// Decode implements bin.Decoder.
func (w *WallPaperNoFile) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode wallPaperNoFile#8af40b25 to nil")
	}
	if err := b.ConsumeID(WallPaperNoFileTypeID); err != nil {
		return fmt.Errorf("unable to decode wallPaperNoFile#8af40b25: %w", err)
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode wallPaperNoFile#8af40b25: field flags: %w", err)
		}
	}
	w.Default = w.Flags.Has(1)
	w.Dark = w.Flags.Has(4)
	if w.Flags.Has(2) {
		if err := w.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode wallPaperNoFile#8af40b25: field settings: %w", err)
		}
	}
	return nil
}

// construct implements constructor of WallPaperClass.
func (w WallPaperNoFile) construct() WallPaperClass { return &w }

// Ensuring interfaces in compile-time for WallPaperNoFile.
var (
	_ bin.Encoder = &WallPaperNoFile{}
	_ bin.Decoder = &WallPaperNoFile{}

	_ WallPaperClass = &WallPaperNoFile{}
)

// WallPaperClass represents WallPaper generic type.
//
// See https://core.telegram.org/type/WallPaper for reference.
//
// Example:
//  g, err := DecodeWallPaper(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *WallPaper: // wallPaper#a437c3ed
//  case *WallPaperNoFile: // wallPaperNoFile#8af40b25
//  default: panic(v)
//  }
type WallPaperClass interface {
	bin.Encoder
	bin.Decoder
	construct() WallPaperClass
	fmt.Stringer
}

// DecodeWallPaper implements binary de-serialization for WallPaperClass.
func DecodeWallPaper(buf *bin.Buffer) (WallPaperClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case WallPaperTypeID:
		// Decoding wallPaper#a437c3ed.
		v := WallPaper{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WallPaperClass: %w", err)
		}
		return &v, nil
	case WallPaperNoFileTypeID:
		// Decoding wallPaperNoFile#8af40b25.
		v := WallPaperNoFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WallPaperClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode WallPaperClass: %w", bin.NewUnexpectedID(id))
	}
}

// WallPaper boxes the WallPaperClass providing a helper.
type WallPaperBox struct {
	WallPaper WallPaperClass
}

// Decode implements bin.Decoder for WallPaperBox.
func (b *WallPaperBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode WallPaperBox to nil")
	}
	v, err := DecodeWallPaper(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.WallPaper = v
	return nil
}

// Encode implements bin.Encode for WallPaperBox.
func (b *WallPaperBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.WallPaper == nil {
		return fmt.Errorf("unable to encode WallPaperClass as nil")
	}
	return b.WallPaper.Encode(buf)
}
