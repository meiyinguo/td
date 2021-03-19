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

// Invoker can invoke raw MTProto rpc calls.
type Invoker interface {
	InvokeRaw(ctx context.Context, input bin.Encoder, output bin.Decoder) error
}

// Client implement methods for calling functions from TL schema via Invoker.
type Client struct {
	rpc Invoker
}

func NewClient(invoker Invoker) *Client {
	return &Client{
		rpc: invoker,
	}
}
